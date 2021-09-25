package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "server/gen"
	"server/pkg"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang/protobuf/ptypes/wrappers"

	//"gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/github.com/golang/protobuf@v1.5.0/ptypes/wrappers"
	"google.golang.org/grpc"
)

const port = ":50051"

type Server struct {
	pb.UnimplementedSqlRequestServer
}

func (s *Server) StreamSql(number *pb.Request, stream pb.SqlRequest_StreamSqlServer) error {
	fmt.Printf("Number = %v\n", number.Number)
	a := pkg.QuerySQL(number.Number)
	for _, val := range a {
		tmp := pb.Answer{Oid: val.Oid,
			FirmSql:       val.Firm,
			PresencePrice: val.PresencePrice,
			SalesPrice:    val.SalesPrice,
			Caption:       val.Caption}
		err := stream.Send(&tmp)
		if err != nil {
			return fmt.Errorf("error sending message to stream : %v", err)
		}
	}
	return nil

}
func (s *Server) Change(stream pb.SqlRequest_ChangeServer) error {
	for {
		order, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&wrappers.BoolValue{Value: true})
		}
		//fmt.Printf("order = %v\n\n", order)
		pkg.Change(order.SalesPrice, order.Number, order.Oid, order.PresencePrice)

	}
	//fmt.Printf("stream = %v", stream)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSqlRequestServer(s, &Server{})
	log.Printf("Starting gRPC listen on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
