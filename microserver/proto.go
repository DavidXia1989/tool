package template

var Proto = `syntax = "proto3"; //使用的proto3语法版本 执行命令protoc -I. --go_out=. --micro_out=. proto/example.proto
package example;// 包名
option go_package="/proto/example;example";

service Example {
    rpc Hello(HelloReq) returns(HelloRep) {}
}

message HelloReq{
    string name = 1;
}

message HelloRep{
    string msg = 1;
}
`