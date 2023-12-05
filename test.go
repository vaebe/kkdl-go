package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	// 反转字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func generateRandomNumber() int {
	// 设置随机数种子，确保每次运行程序时都得到不同的结果
	rand.NewSource(time.Now().UnixNano())
	max := 9999999999
	min := 1
	return rand.Intn(max-min+1) + min
}

func topNum(idObj map[string]int) {
	// 将键值对提取到结构体切片中
	type KeyValue struct {
		Key   string
		Value int
	}

	var keyValueSlice []KeyValue
	for key, value := range idObj {
		keyValueSlice = append(keyValueSlice, KeyValue{Key: key, Value: value})
	}

	// 定义排序规则
	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value < keyValueSlice[j].Value
	})

	// 获取前十个键值对
	topTen := keyValueSlice
	if len(keyValueSlice) > 10 {
		topTen = keyValueSlice[:10]
	}

	// 输出前十个键值对
	for _, kv := range topTen {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}

func getId() {
	idObj := make(map[string]int)

	for a := 0; a < 10000; a++ {
		base62Str := IntToBase62(generateRandomNumber())

		println(base62Str)

		if _, ok := idObj[base62Str]; ok {
			idObj[base62Str] = idObj[base62Str] + 1
		} else {
			idObj[base62Str] = 1
		}
	}

	topNum(idObj)
}

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// 执行重定向到新的 URL
		http.Redirect(w, r, "https://gitee.com/go-nunu/nunu-layout-advanced/blob/main/pkg/helper/convert/convert.go", http.StatusFound)
	})

	// 启动 HTTP 服务器，监听端口 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
