package main

import (
	"fmt"
	"sync"
)

type URLShortener interface {
	Shorten(longURL string) string
	Expand(shortURL string) (string, bool)
	GetStats(shortURL string) int
}

type URLData struct {
	LongURL  string
	ShortURL string
	Visits   int
}

type URLService struct {
	mu           sync.RWMutex
	shortToLong  map[string]*URLData  // 短链接 -> URL数据
	longToShort  map[string]string    // 长链接 -> 短链接 (避免重复)
}

func NewURLService() URLShortener {
	return &URLService{
		shortToLong: make(map[string]*URLData),
		longToShort: make(map[string]string),
	}
}

func (s *URLService) Shorten(longURL string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// 检查是否已经存在
	if shortURL, exists := s.longToShort[longURL]; exists {
		return shortURL
	}
	
	// 生成新的短链接
	shortURL := Encode(longURL)
	
	// 处理哈希冲突（简单处理：如果冲突就加后缀）
	originalShort := shortURL
	counter := 1
	for {
		if _, exists := s.shortToLong[shortURL]; !exists {
			break
		}
		shortURL = fmt.Sprintf("%s%d", originalShort, counter)
		counter++
	}
	
	// 存储映射关系
	urlData := &URLData{
		LongURL:  longURL,
		ShortURL: shortURL,
		Visits:   0,
	}
	
	s.shortToLong[shortURL] = urlData
	s.longToShort[longURL] = shortURL
	
	return shortURL
}

func (s *URLService) Expand(shortURL string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	urlData, exists := s.shortToLong[shortURL]
	if !exists {
		return "", false
	}
	
	// 增加访问计数
	urlData.Visits++
	
	return urlData.LongURL, true
}

func (s *URLService) GetStats(shortURL string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	urlData, exists := s.shortToLong[shortURL]
	if !exists {
		return 0
	}
	
	return urlData.Visits
}

func main() {
	// 创建URL服务
	service := NewURLService()
	
	// 测试基本功能
	fmt.Println("=== 基本功能测试 ===")
	longURL1 := "https://www.google.com"
	longURL2 := "https://www.github.com"
	
	// 生成短链接
	shortURL1 := service.Shorten(longURL1)
	shortURL2 := service.Shorten(longURL2)
	
	fmt.Printf("长链接: %s\n", longURL1)
	fmt.Printf("短链接: %s\n", shortURL1)
	fmt.Printf("长链接: %s\n", longURL2)
	fmt.Printf("短链接: %s\n", shortURL2)
	
	// 测试重复URL
	fmt.Println("\n=== 重复URL测试 ===")
	shortURL1Again := service.Shorten(longURL1)
	fmt.Printf("相同长链接再次生成: %s\n", shortURL1Again)
	fmt.Printf("是否相同: %t\n", shortURL1 == shortURL1Again)
	
	// 测试还原
	fmt.Println("\n=== 还原测试 ===")
	expandedURL, ok := service.Expand(shortURL1)
	fmt.Printf("短链接: %s\n", shortURL1)
	fmt.Printf("还原结果: %s, 成功: %t\n", expandedURL, ok)
	
	// 测试不存在的短链接
	_, ok = service.Expand("notexist")
	fmt.Printf("不存在的短链接: 成功: %t\n", ok)
	
	// 测试访问统计
	fmt.Println("\n=== 访问统计测试 ===")
	fmt.Printf("初始访问次数: %d\n", service.GetStats(shortURL1))
	
	// 多次访问
	service.Expand(shortURL1)
	service.Expand(shortURL1)
	service.Expand(shortURL1)
	
	fmt.Printf("访问3次后的次数: %d\n", service.GetStats(shortURL1))
	
	// 测试并发安全
	fmt.Println("\n=== 并发安全测试 ===")
	testConcurrency(service)
}

func testConcurrency(service URLShortener) {
	var wg sync.WaitGroup
	testURL := "https://www.example.com"
	
	// 10个goroutine同时访问
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			shortURL := service.Shorten(testURL)
			service.Expand(shortURL)
		}()
	}
	
	wg.Wait()
	
	shortURL := service.Shorten(testURL)
	visits := service.GetStats(shortURL)
	fmt.Printf("并发测试URL: %s\n", testURL)
	fmt.Printf("短链接: %s\n", shortURL)
	fmt.Printf("总访问次数: %d\n", visits)
}