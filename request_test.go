package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
	"verif_net_work/pool"
)

var reqPool *pool.ConnectionPool

func TestSlice(t *testing.T) {
	reqPool = pool.NewConnectionPool()
	for i := 0; i < 1000; i++ {
		go func() {
			GetReq()
		}()
	}
	time.Sleep(60 * time.Second)
}

func GetReq() {
	random := rand.Int()
	url := fmt.Sprintf("http://127.0.0.1:5536/Api/GetTrueResult?random=%d", random)
	p := pools.GetConnection()
	defer pools.PutConnection(p)
	resp, err := p.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	fmt.Println("请求头：", random, "响应体:", responseString)

}
