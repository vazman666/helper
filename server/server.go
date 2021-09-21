package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	pb "server/gen"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"

	"google.golang.org/grpc"
)

const port = ":50051"

type Server struct {
	pb.UnimplementedSqlRequestServer
}

func (s *Server) StreamSql(number *pb.Request, stream pb.SqlRequest_StreamSqlServer) error {
	fmt.Printf("Number = %v\n", number.Number)
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	a := pb.Answer{
		Oid:           "1232",
		Firm:          "Mits",
		PresencePrice: "20",
		SalesPrice:    "323",
		Caption:       "123",
	}
	rows, err := condb.Query("use basebasebase") //используем базу данных basebasebase
	if err != nil {
		log.Fatal(err)
	}
	//Заброс к бд с OriginalCode=number
	rows, err = condb.Query("SELECT TOP (1000) [Oid],[Producer] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? ", number.Number)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Oid, &a.Firm)
		if err != nil {
			log.Fatal(err)

		}
		rows2, err := condb.Query("SELECT  [Caption] FROM [basebasebase].[dbo].[Producer]  WHERE [OID]=? ", a.Firm)
		if err != nil {
			log.Fatal(err)
		}
		for rows2.Next() {

			err := rows2.Scan(&a.Caption) //фирма из базы
			if err != nil {
				log.Fatal(err)

			}

		}

		rows3, err := condb.Query("SELECT TOP (1000) [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", a.Oid) //используем базу данных tim
		if err != nil {
			log.Fatal(err)
		}

		for rows3.Next() {

			err := rows3.Scan(&a.PresencePrice, &a.SalesPrice)
			if err != nil {
				log.Fatal(err)
			}

		}
		if a.Oid == "" {
			a.PresencePrice = "---"
			a.SalesPrice = "---"
		} else {
			pp, _ := strconv.ParseFloat(a.PresencePrice, 2)
			sp, _ := strconv.ParseFloat(a.SalesPrice, 2)
			a.PresencePrice = fmt.Sprintf("%5.2f", pp)
			a.SalesPrice = fmt.Sprintf("%5.2f", sp)
		}

		err = stream.Send(&a)
		if err != nil {
			return fmt.Errorf("error sending message to stream : %v", err)
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
