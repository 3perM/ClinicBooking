package main

import (
    "fmt"
    "net"
    "google.golang.org/grpc"
)

func main() {
    fmt.Println("Starting gRPC server on :50051")

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()

    if err := s.Serve(lis); err != nil {
        panic(err)
    }
}
