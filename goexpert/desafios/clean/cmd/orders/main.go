package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/configs"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/internal/event/handler"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/internal/infra/graph"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/internal/infra/grpc/pb"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/internal/infra/grpc/service"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/internal/infra/web/webserver"
	"github.com/osvaldotcf/pgfcycle/goexpert/desafios/cleanarch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// ----- CONFIGS
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// ----- DATABASE
	dbcon := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", configs.DBUser, configs.DBPassword,
		configs.DBHost, configs.DBPort, configs.DBName)
	fmt.Println("DBCON: " + dbcon)
	db, err := sql.Open(configs.DBDriver, dbcon)

	//sql.Open("postgres", "postgres://pgclean:1010aa@pgclean-host/pgcleandb?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		if err = db.Ping(); err != nil {
			panic(err)
		} else {

			fmt.Println("POSTGRESQL OK!")

		}
	}
	defer db.Close()

	// ----- RABBITMQ
	rabbitMQChannel := getRabbitMQChannel()
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUsecase := NewListOrderUseCase(db, eventDispatcher)

	// ----- WEBSERVER
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("POST /order", webOrderHandler.Create)
	webserver.AddHandler("GET /order", webOrderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// ----- GRPC_SERVER
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrderUsecase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// ----- GRAPHQL_SERVER
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUsecase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-host")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
