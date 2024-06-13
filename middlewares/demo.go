package middlewares

import (
	"fmt" //输出到控制台
	"github.com/gin-gonic/gin"
	"time"
)

func PrintOne(c *gin.Context) {
	fmt.Println(1)
	c.Set("name", "何风顺")
	//定义goroutine 线程 统计日志
	// 复制上下文
	cCopy := c.Copy()
	//拷贝解决的问题
	//1.请求已结束:
	//	如果 HTTP 请求已经结束，而 goroutine 还在运行，gin.Context 可能已经被清理或重用，这会导致 goroutine 中的访问出现不可预测的行为或崩溃。
	//2.数据竞争:
	//	如果同时有多个 goroutine 访问或修改 gin.Context 的数据，会引发数据竞争，导致数据不一致或崩溃。
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + cCopy.Request.URL.Path)
	}()
}
