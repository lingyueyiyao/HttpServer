package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	conf clientv3.Config
	client *clientv3.Client
	err error
)



func main() {
	conf = clientv3.Config{
		Endpoints:   []string{"localhost:2479", "localhost:2579", "localhost:2679"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(conf); err != nil {
		fmt.Println("连接失败，Error：", err)
		return
	}

	fmt.Println("连接成功")

	kv := clientv3.NewKV(client)
	fmt.Println("正在读取")
	getResp, _ := kv.Get(context.TODO(), "msg")
	fmt.Println("读取成功")
	fmt.Println(string(getResp.Kvs[0].Key), ":", string(getResp.Kvs[0].Value))

	defer client.Close()


	//engine := http.NewEngine()
	//err := engine.AddRoute("/hello", func(ctx *http.Context){
	//	ctx.Writer.Write([]byte("hello"))
	//})
	//
	//if err != nil {
	//	log.Fatalf("fail to addroute /hello : %v", err)
	//}
	//
	//if err = engine.ListenAndServe(":8080"); err != nil {
	//	log.Fatalf("fail to Listen :8080 : %v", err)
	//}
}

