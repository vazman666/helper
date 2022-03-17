package main

import (
	"context"
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

func (s *Server) StreamSql(ctx context.Context, in *pb.Request) (*pb.Answer, error) {
	fmt.Printf("Запрос на номер %v   фирмы %v\n", in.Number, in.Firm)
	a := pkg.QuerySQL(in.Number, in.Firm)

	ans := pb.Answer{Oid: a.Oid,
		FirmSql:       a.Firm,
		PresencePrice: a.PresencePrice,
		SalesPrice:    a.SalesPrice,
		Caption:       a.Caption,
		Cellm:         a.Cellm,
		Cellt:         a.Cellt,
		Name:          a.Name,
	}
	/*err := stream.Send(&tmp)
	if err != nil {
		return ans, fmt.Errorf("error sending message to stream : %v", err)
	}*/

	return &ans, nil

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
func (s *Server) Analogs(in *pb.Request, stream pb.SqlRequest_AnalogsServer) error {
	fmt.Printf("Запрос на номер %v   фирмы %v\n", in.Number, in.Firm)
	pkg.Analogue(pkg.Analog{in.Firm, in.Number})
	//fmt.Printf("нашли %v\n", pkg.Analogs)
	for _, feature := range pkg.Analogs {
		resp := pb.Request{Number: feature.Number, Firm: feature.Firm}
		if err := stream.Send(&resp); err != nil {
			return err

		}
	}

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
