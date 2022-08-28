package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func testRoutine(c *gin.Context)  {
	val, _ := c.Get("test")
	fmt.Println("在goroutine中使用Context的副本拿到数据：", val)
}
func h1(c *gin.Context){
	fmt.Println("h2 start...")
	// 设置一些参数，可以给后续的中间件使用
	c.Set("test", "111")

	go testRoutine(c.Copy())
	start := time.Now()
	// 不执行之后的中间件
	//c.Abort()
	// 先继续执行之后的中间件
	c.Next()
	timeCost := time.Since(start)
	fmt.Println(timeCost)
	fmt.Println("h2 end...")
}

func h2(c *gin.Context) {
	fmt.Println("h1 start")
	c.JSON(http.StatusOK, gin.H{
		"name": "ww",
	})
}

func h3(c *gin.Context) {
	testVal, _ := c.Get("test")
	fmt.Println("h3中拿到h1的数据", testVal)
	fmt.Println("h3 start")
}

func authMiddleware(doAuth bool) gin.HandlerFunc{
	return func(c *gin.Context) {
		if doAuth{
			c.Abort()
		}else{
			c.Next()
		}

	}
}


func main() {
	r := gin.Default()

	r.Use(h1, authMiddleware(false))

	r.GET("/index", h2, h3)
	r.GET("/index2",h2, h3)

	if err := r.Run(); err != nil{
		fmt.Println(err)
	}
}
