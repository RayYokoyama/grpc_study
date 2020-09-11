package main

import (
  "fmt"
  "context"
  "log"
  "net"

  "./lib"
  "google.golang.org/grpc"
)

func main() {
  listener, err := net.Listen("tcp", ":5300")
  if err != nil {
    log.Fatalf("failed to listen: %v\n", err)
    return
  }
  grpcSrv := grpc.NewServer()
  pinger.RegisterPingerServer(grpcSrv, &server{
    Counter: map[string]int32{},
  })
  log.Printf("Pinger service is running!")
  grpcSrv.Serve(listener)
}

type server struct{
  Counter map[string]int32
}

type counter struct{
  msg string
  count int32
}

func (s *server) Ping(ctx context.Context, req *pinger.Reqest) (*pinger.Response, error) {
  text := req.Text
  s.count_msg(text)

  pong := &pinger.Response{
    Text: text,
    Count: s.Counter[text],
  }
  return pong, nil
}

func (s *server) count_msg(msg string) {
  fmt.Println(s.Counter)
  s.Counter[msg] ++
  fmt.Println(msg)
}
