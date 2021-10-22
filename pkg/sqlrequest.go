package pkg

import (
	"context"
	"fmt"
	"helper/models"
	"io"
	"log"
	pb "pkg/gen"

	"google.golang.org/grpc"
)

const address = "vazman.ru:50051"

func SqlRequest(number string) []models.Price_st {
	req := []models.Price_st{}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("is connected")
	defer conn.Close()

	ctx := context.Background()
	c := pb.NewSqlRequestClient(conn)

	timeStream, _ := c.StreamSql(ctx, &pb.Request{Number: number})
	for {
		stream, err := timeStream.Recv()
		if err == io.EOF {
			break
		}
		tmp := models.Price_st{Number: number, FirmSql: stream.Caption, PresencePrice: stream.PresencePrice, SalesPrice: stream.SalesPrice, Oid: stream.Oid}
		req = append(req, tmp)
		fmt.Printf("Получено :%v\n", stream)
	}
	return req

}
