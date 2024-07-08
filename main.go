package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"salary_project/db"
	"salary_project/entity/proto"
	v1 "salary_project/handlers/http/v1"
	"salary_project/repository"

	grpchandler "salary_project/handlers/grpc_handler"
	"salary_project/service"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {

	db, err := db.DBconnection()

	fmt.Println(err)

	e := echo.New()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	v1.SetUpRouter(e, db)

	go grpcServer(s, db)

	e.Logger.Fatal(e.Start(":4001"))

}

func grpcServer(g *grpc.Server, db *sql.DB) {
	//fmt.Println("start")
	lis, err := net.Listen("tcp", ":2101")
	if err != nil {
		log.Fatalf("GRPC: failed to listen: %v", err)
	}
	//fmt.Println("start1")
	repo := repository.NewSalaryRepo(db)
	service := service.NewSalaryService(repo)

	salaryHandler := grpchandler.SalaryGrpcHandler{
		SalaryService: service,
	}
	//fmt.Println("start2")
	proto.RegisterEmployeeToSalaryServer(g, &salaryHandler)
	//fmt.Println("start3")
	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := g.Serve(lis); err != nil {
		log.Fatalf("GRPC: failed to serve: %v", err)
	}
}
