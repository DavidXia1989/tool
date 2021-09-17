package template

var Proto = `syntax = "proto3"; //使用的proto3语法版本 执行命令protoc -I. --go_out=. --micro_out=. proto/example/example.proto
package example;// 包名
option go_package="/proto/example;example";

service Example {
    rpc Login(LoginMsgReq) returns(LoginMsgRes) {}
}

message LoginMsgReq{
    repeated Msg msg = 1;
}
message Msg{
    string name = 1;
    string password = 2;
}

message LoginMsgRes{
    string result = 1;
}
`