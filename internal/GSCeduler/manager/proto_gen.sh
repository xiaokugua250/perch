#protoc  .\proto\client_manager.proto --go_out=.
# ref https://juejin.cn/post/6875963835342127111
#protoc -I .\proto\ .\proto\client_manager.proto --go_out=.
protoc -I . --go_out=plugins=grpc:. .\proto\client_manager.proto
protoc -I . --go_out=plugins=grpc:. .\proto\*.proto
 protoc -I . --go_out=plugins=grpc:. .\proto\pb_stream\*.proto
 protoc -I . --go_out=plugins=grpc:. .\proto\pb_normal\*.proto
