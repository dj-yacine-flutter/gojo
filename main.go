package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	rt "runtime"

	"github.com/dj-yacine-flutter/gojo/api"
	"github.com/dj-yacine-flutter/gojo/api/shared"
	v1 "github.com/dj-yacine-flutter/gojo/api/v1"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	_ "github.com/dj-yacine-flutter/gojo/doc/statik"
	"github.com/dj-yacine-flutter/gojo/mail"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	sk "github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// Set GOMAXPROCS to the number of available CPU cores
	rt.GOMAXPROCS(rt.NumCPU())
}

func main() {
	var err error

	config, err := utils.LoadConfig(".", "gojo")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the DB")
	}

	gojo := db.NewGojo(conn)

	taskDistributor := worker.NewRedisTaskDistributor(asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	})

	ping := ping.NewPingSystem(config)

	go gateway(config, gojo, taskDistributor, ping)

	server := grpc.NewServer()

	err = v1.StartGRPCApi(server, config, gojo, taskDistributor, ping)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	reflection.Register(server)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gRPC listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START gRPC server -AT- %s\u001b[0m\n", listener.Addr().String())

	err = server.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}

	//		fmt.Printf("\u001b[38;5;125m\u001b[48;5;0m%s\u001b[0m\n", fmt.Sprintln(`
	//
	//	    ██████╗  ██████╗      ██╗ ██████╗
	//	   ██╔════╝ ██╔═══██╗     ██║██╔═══██╗
	//	   ██║  ███╗██║   ██║     ██║██║   ██║
	//	   ██║   ██║██║   ██║██   ██║██║   ██║
	//	   ╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝
	//	    ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝
	//	                                           `))
	//
	//		cache := runRedisCache(config.RedisCacheAddress)
	//
	//		go runRedisQueue(config, redisOpt, gojo)
	//		go runGatewayServer(config, gojo, taskDistributot, cache)
	//
	//		runGRPCServer(config, gojo, taskDistributot, cache)
}

func gateway(config utils.Config, gojo db.Gojo, taskDistrinbutor worker.TaskDistributor, ping *ping.PingSystem) {
	var err error

	httpMux := http.NewServeMux()
	err = v1.StartGatewayApi(httpMux, config, gojo, taskDistrinbutor, ping)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START HTTP server -AT- %s\u001b[0m\n", listener.Addr().String())

	err = http.Serve(listener, httpMux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the Gateway server")
	}
}

func runRedisQueue(config utils.Config, redisOpt asynq.RedisClientOpt, gojo db.Gojo) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, gojo, mailer)

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START REDIS TASK PROCESSOR -AT- %s\u001b[0m\n", config.RedisQueueAddress)

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGRPCServer(config utils.Config, gojo db.Gojo, taskDistrinbutor worker.TaskDistributor) {
	server, err := api.NewServer(config, gojo, taskDistrinbutor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create GRPC server")
	}

	gprcLogger := grpc.UnaryInterceptor(shared.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterGojoServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gRPC listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START gRPC SERVER -AT- %s\u001b[0m\n", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the gRPC server")
	}
}

func runGatewayServer(config utils.Config, gojo db.Gojo, taskDistrinbutor worker.TaskDistributor) {
	server, err := api.NewServer(config, gojo, taskDistrinbutor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway server")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()

	err = pb.RegisterGojoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := sk.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik")
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START HTTP SERVER -AT- %s\u001b[0m\n", listener.Addr().String())

	handler := shared.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the Gateway server")
	}
}
