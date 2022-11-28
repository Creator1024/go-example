package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

// 执行结果：
//{"level":"info","ts":1669559449.8381262,"caller":"zap-logger/main.go:32","msg":"Success..","statusCode":"200 OK","url":"http://www.baidu.com"}

//Get "http://www.google.com": dial tcp 185.60.218.50:80: i/o timeout
//{"level":"error","ts":1669559479.8404,"caller":"zap-logger/main.go:27","msg":"Error fetching url..","url":"http://www.google.com","error":"Get \"http://www.google.com\": dial tcp 185.60.218.50:80: i/o timeout","stacktrace":"main.simpleHttpGet\n\t/Users/oneway/go/src/go-example/demo/zap-logger/main.go:27\nmain.main\n\t/Users/oneway/go/src/go-example/demo/zap-logger/main.go:16\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"}
