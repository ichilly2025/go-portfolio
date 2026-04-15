package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ichilly2025/go-portfolio/network/grpc/pb"
)

type BenchmarkResult struct {
	ServerType      string
	TotalRequests   int
	Concurrency     int
	TotalTime       time.Duration
	TPS             float64 // Transactions Per Second
	AvgResponseTime time.Duration
	MinResponseTime time.Duration
	MaxResponseTime time.Duration
	SuccessCount    int
	FailCount       int
}

// gRPC server implementation for benchmark
type grpcCalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *grpcCalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	fmt.Println("=== HTTP Server vs RESTful vs Gin vs Echo vs Go-Zero vs gRPC Benchmark ===\n")

	concurrencyLevels := []int{10, 50, 100, 500}
	requestsPerLevel := 1000

	for _, concurrency := range concurrencyLevels {
		fmt.Printf("Testing with %d concurrent requests (%d total requests):\n", concurrency, requestsPerLevel)
		fmt.Println(strings.Repeat("-", 80))

		// Benchmark standard HTTP server
		httpResult := benchmarkHTTPServer(concurrency, requestsPerLevel)
		printResult(httpResult)

		// Benchmark RESTful server
		restfulResult := benchmarkRESTfulServer(concurrency, requestsPerLevel)
		printResult(restfulResult)

		// Benchmark Gin server
		ginResult := benchmarkGinServer(concurrency, requestsPerLevel)
		printResult(ginResult)

		// Benchmark Echo server
		echoResult := benchmarkEchoServer(concurrency, requestsPerLevel)
		printResult(echoResult)

		// Benchmark Go-Zero server
		zeroResult := benchmarkGoZeroServer(concurrency, requestsPerLevel)
		printResult(zeroResult)

		// Benchmark gRPC server
		grpcResult := benchmarkGRPCServer(concurrency, requestsPerLevel)
		printResult(grpcResult)

		// Compare results
		compareResults(httpResult, restfulResult, ginResult, echoResult, zeroResult, grpcResult)
		fmt.Println()
	}
}

func benchmarkHTTPServer(concurrency, totalRequests int) BenchmarkResult {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	return runBenchmark("Standard HTTP", server.URL, concurrency, totalRequests)
}

func benchmarkRESTfulServer(concurrency, totalRequests int) BenchmarkResult {
	ws := new(restful.WebService)
	ws.Path("/api")
	ws.Route(ws.GET("/hello").To(func(req *restful.Request, resp *restful.Response) {
		resp.WriteAsJson(map[string]string{"message": "Hello, World!"})
	}))

	container := restful.NewContainer()
	container.Add(ws)

	server := httptest.NewServer(container)
	defer server.Close()

	return runBenchmark("RESTful", server.URL+"/api/hello", concurrency, totalRequests)
}

func benchmarkGinServer(concurrency, totalRequests int) BenchmarkResult {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	return runBenchmark("Gin", server.URL+"/hello", concurrency, totalRequests)
}

func benchmarkEchoServer(concurrency, totalRequests int) BenchmarkResult {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	server := httptest.NewServer(e)
	defer server.Close()

	return runBenchmark("Echo", server.URL+"/hello", concurrency, totalRequests)
}

func benchmarkGoZeroServer(concurrency, totalRequests int) BenchmarkResult {
	r := router.NewRouter()

	r.Handle(http.MethodGet, "/hello", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		httpx.OkJson(w, map[string]string{
			"message": "Hello, World!",
		})
	}))

	server := httptest.NewServer(r)
	defer server.Close()

	return runBenchmark("Go-Zero", server.URL+"/hello", concurrency, totalRequests)
}

func benchmarkGRPCServer(concurrency, totalRequests int) BenchmarkResult {
	// Start gRPC server
	lis, err := net.Listen("tcp", "127.0.0.1:0") // Use random available port
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &grpcCalculatorServer{})

	// Start server in background
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(fmt.Sprintf("failed to serve: %v", err))
		}
	}()
	defer s.Stop()

	// Get the actual address
	addr := lis.Addr().String()

	return runGRPCBenchmark("gRPC", addr, concurrency, totalRequests)
}

func runGRPCBenchmark(serverType, addr string, concurrency, totalRequests int) BenchmarkResult {
	result := BenchmarkResult{
		ServerType:    serverType,
		TotalRequests: totalRequests,
		Concurrency:   concurrency,
	}

	var wg sync.WaitGroup
	requestsPerWorker := totalRequests / concurrency

	responseTimes := make([]time.Duration, totalRequests)
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Create gRPC client for this worker
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				mu.Lock()
				failCount += requestsPerWorker
				mu.Unlock()
				return
			}
			defer conn.Close()

			client := pb.NewCalculatorClient(conn)

			for j := 0; j < requestsPerWorker; j++ {
				reqStart := time.Now()
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				_, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 20})
				cancel()
				reqDuration := time.Since(reqStart)

				mu.Lock()
				idx := workerID*requestsPerWorker + j
				responseTimes[idx] = reqDuration

				if err != nil {
					failCount++
				} else {
					successCount++
				}
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	// Calculate statistics
	var totalResponseTime time.Duration
	minResponseTime := responseTimes[0]
	maxResponseTime := responseTimes[0]

	for _, rt := range responseTimes {
		totalResponseTime += rt
		if rt < minResponseTime {
			minResponseTime = rt
		}
		if rt > maxResponseTime {
			maxResponseTime = rt
		}
	}
	avgResponseTime := totalResponseTime / time.Duration(totalRequests)

	result.TotalTime = totalTime
	result.TPS = float64(totalRequests) / totalTime.Seconds()
	result.AvgResponseTime = avgResponseTime
	result.MinResponseTime = minResponseTime
	result.MaxResponseTime = maxResponseTime
	result.SuccessCount = successCount
	result.FailCount = failCount

	return result
}

func runBenchmark(serverType, url string, concurrency, totalRequests int) BenchmarkResult {
	result := BenchmarkResult{
		ServerType:    serverType,
		TotalRequests: totalRequests,
		Concurrency:   concurrency,
	}

	var wg sync.WaitGroup
	requestsPerWorker := totalRequests / concurrency
	
	responseTimes := make([]time.Duration, totalRequests)
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			for j := 0; j < requestsPerWorker; j++ {
				reqStart := time.Now()
				resp, err := http.Get(url)
				reqDuration := time.Since(reqStart)

				mu.Lock()
				idx := workerID*requestsPerWorker + j
				responseTimes[idx] = reqDuration

				if err != nil || resp.StatusCode != http.StatusOK {
					failCount++
				} else {
					successCount++
					io.ReadAll(resp.Body)
					resp.Body.Close()
				}
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	// Calculate statistics
	var totalResponseTime time.Duration
	minResponseTime := responseTimes[0]
	maxResponseTime := responseTimes[0]
	
	for _, rt := range responseTimes {
		totalResponseTime += rt
		if rt < minResponseTime {
			minResponseTime = rt
		}
		if rt > maxResponseTime {
			maxResponseTime = rt
		}
	}
	avgResponseTime := totalResponseTime / time.Duration(totalRequests)

	result.TotalTime = totalTime
	result.TPS = float64(totalRequests) / totalTime.Seconds()
	result.AvgResponseTime = avgResponseTime
	result.MinResponseTime = minResponseTime
	result.MaxResponseTime = maxResponseTime
	result.SuccessCount = successCount
	result.FailCount = failCount

	return result
}

func printResult(result BenchmarkResult) {
	fmt.Printf("  %s Server:\n", result.ServerType)
	fmt.Printf("    Total Time:          %v\n", result.TotalTime)
	fmt.Printf("    TPS (Req/sec):       %.2f\n", result.TPS)
	fmt.Printf("    Avg Response Time:   %v\n", result.AvgResponseTime)
	fmt.Printf("    Min Response Time:   %v\n", result.MinResponseTime)
	fmt.Printf("    Max Response Time:   %v\n", result.MaxResponseTime)
	fmt.Printf("    Success:             %d\n", result.SuccessCount)
	fmt.Printf("    Failed:              %d\n", result.FailCount)
	fmt.Println()
}

func compareResults(http, restful, gin, echo, zero, grpc BenchmarkResult) {
	fmt.Println("  Comparison:")
	
	// Find the fastest
	results := []BenchmarkResult{http, restful, gin, echo, zero, grpc}
	names := []string{"Standard HTTP", "RESTful", "Gin", "Echo", "Go-Zero", "gRPC"}
	
	fastest := results[0]
	fastestName := names[0]
	for i, r := range results {
		if r.TPS > fastest.TPS {
			fastest = r
			fastestName = names[i]
		}
	}

	fmt.Printf("    Fastest: %s (%.2f TPS)\n", fastestName, fastest.TPS)
	fmt.Println()

	// Compare each to the fastest
	for i, r := range results {
		if r.TPS != fastest.TPS {
			diff := ((fastest.TPS - r.TPS) / r.TPS) * 100
			fmt.Printf("    %s is %.2f%% slower (%.2f TPS)\n", names[i], diff, r.TPS)
		}
	}
}
