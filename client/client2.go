package main

import (
	pb "client/gen"
	"context"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/tealeg/xlsx/v3"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	/*condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}*/

	wb, err := xlsx.OpenFile("a.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["test"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}

	/*tmp := pb.Answer{Oid: "17291",
	FirmSql:       "RIBROWER",
	PresencePrice: "313.00",
	SalesPrice:    "600.00",
	Caption:       "228",
	Number:        "6PK1000"}*/
	var a []pb.Answer

	for i := 1; i < sh.MaxRow; i++ {

		theCell, err := sh.Cell(i, 6)
		if err != nil {
			panic(err)
		}
		presencePrice, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 1)
		if err != nil {
			panic(err)
		}
		number, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 5)
		if err != nil {
			panic(err)
		}
		salesPrice, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 11)
		if err != nil {
			panic(err)
		}
		oid, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		a = append(a, pb.Answer{Oid: oid,
			FirmSql:       "-----",
			PresencePrice: presencePrice,
			SalesPrice:    salesPrice,
			Caption:       "-----",
			Number:        number})
	}

	fmt.Printf("a=%v\n", a)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("is connected")
	defer conn.Close()

	ctx := context.Background()
	client := pb.NewSqlRequestClient(conn)
	updateStream, err := client.Change(ctx)
	if err != nil {
		log.Fatalf("%v.Change(_) = _,%v", client, err)
	}
	for _, val := range a {
		if err := updateStream.Send(&val); err != nil {
			log.Fatalf("%v.Send(%v)=%v", updateStream, val, err)
		}

	}
	_, err = updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloaseAndRecv() got error %v, want %v",
			updateStream, err, nil)
	}
}
