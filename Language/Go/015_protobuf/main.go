package main

import (
	"example/go_pkg_name"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	rqst := &go_pkg_name.ExampleServiceRqst{}
	rsp := &go_pkg_name.ExampleServiceRsp{}
	filed1 := "aaaa"
	result := int32(0)
	resInfo := "ok"
	rqst.Filed1 = &filed1
	rsp.Result = &result
	rsp.ResInfo = &resInfo
	// 序列化成字节流
	rqstBinData, _ := proto.Marshal(rqst)
	fmt.Printf("%v\n", rqstBinData)
	// 序列化成字节流
	rspBinData, _ := proto.Marshal(rsp)
	fmt.Printf("%v\n", rspBinData)

	rsp2 := &go_pkg_name.ExampleServiceRsp{}
	// 反序列化成结构
	_ = proto.Unmarshal(rspBinData, rsp2)
	fmt.Printf("%v\n", *rsp2.ResInfo)
}
