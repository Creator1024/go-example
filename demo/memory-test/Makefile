
build-image:
	GOOS=linux go build -o main
	docker build -t registry.cn-hangzhou.aliyuncs.com/creator1024/demo:memory_test_v1 .

clean:
	rm -f main