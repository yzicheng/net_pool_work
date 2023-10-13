package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

func TestReq(t *testing.T) {
	GetReq()
}

func GetReq() {
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get("http://localhost:5536/Api/GetTrueResult?random=" + fmt.Sprintf("%d", i))
			if err != nil {
				fmt.Println("请求失败:", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取响应内容失败:", err)
				return
			}

			fmt.Println(string(body))
		}(i)
	}

	wg.Wait()

}
