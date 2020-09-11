# requirement
### gRPCサーバの実装に必要
```$ go get -u google.golang.org/grpc```
### Protocol Buffersの定義からGolangのソースコードを生成に必要
```$ go get -u github.com/golang/protobuf/protoc-gen-go```


# 起動
## フォルダ構成
grpc-sample
├ proto
│ └ ProtocolBuffers定義
├ server
│ ├ Railsアプリ
│ └ lib
│ 　 └ protoから自動生成
└ pinger
　 ├ Golangアプリ
　 └ lib
　 　 └ protoから自動生成
## サーバー側
go で実装
```
protoc -I ./proto pinger.proto --go_out=plugins=grpc:./pinger/lib
go run ./pinger/server.go
```
## クライアント側
rails で実装
```
# cd grpc-study/server
$ bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=lib --grpc_out=lib ../proto/pinger.proto
$ bundle exec rails server
$ bundle exec rails s
```
