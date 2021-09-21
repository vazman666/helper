package main

import (
	pb "client/gen"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("is connected")
	defer conn.Close()

	ctx := context.Background()
	c := pb.NewSqlRequestClient(conn)

	timeStream, _ := c.StreamSql(ctx, &pb.Request{Number: "6PK1000"})
	for {
		stream, err := timeStream.Recv()
		if err == io.EOF {
			break
		}

		fmt.Printf("Получено :%v\n", stream)
	}

}
