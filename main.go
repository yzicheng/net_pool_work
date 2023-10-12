package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"verif_net_work/pool"
)

var pools *pool.ConnectionPool
var reqChan chan int

func init() {
	pools = pool.NewConnectionPool()
	reqChan = make(chan int, 50)
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	router := gin.Default()
	group := router.Group("Api")
	{
		group.GET("/GetTrueResult", getTrueResult)
	}
	router.Run(":5536")
}

func getTrueResult(ctx *gin.Context) {

	respChan := make(chan string)
	go func() {
		// 同步等待请求
		<-reqChan

		// 异步调用第三方接口
		time.Sleep(20 * time.Second) // 模拟接口调用耗时
		resp := ctx.DefaultQuery("random", "0") + "调用成功"

		// 将结果发送到通道
		respChan <- resp
	}()

	// 发送请求到通道
	reqChan <- 1
	ctxTimeOut, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	for {
		select {
		// 等待接收结果
		case resp := <-respChan:
			// 返回结果给请求
			ctx.JSON(http.StatusOK, gin.H{
				"message": ctx.DefaultQuery("random", "0") + "|" + resp,
			})
			return
		case <-ctxTimeOut.Done():
			// 返回超时
			ctx.JSON(http.StatusOK, gin.H{
				"message": "已超时",
			})
			return
		}
	}

}
