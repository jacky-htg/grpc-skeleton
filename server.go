package main

import (
	"context"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"waresix.com/grpc-skeleton/internal/config"
	"waresix.com/grpc-skeleton/internal/pkg/db/postgres"
	"waresix.com/grpc-skeleton/internal/pkg/db/redis"
	"waresix.com/grpc-skeleton/internal/pkg/log/logruslog"
	"waresix.com/grpc-skeleton/internal/route"
)

const defaultPort = "9001"

func main() {
	// lookup and setup env
	if _, ok := os.LookupEnv("PORT"); !ok {
		_, err := os.Stat(".env.prod")
		if os.IsNotExist(err) {
			config.Setup(".env")
		} else {
			config.Setup(".env.prod")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// init log
	log := logruslog.Init()

	// create postgres database connection
	db, err := postgres.Open()
	if err != nil {
		log.Errorf("connecting to db: %v", err)
		return
	}
	log.Info("connecting to postgresql database")

	defer db.Close()

	// create redis cache connection
	cache, err := redis.NewCache(context.Background(), os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"), 24*time.Hour)
	if err != nil {
		log.Errorf("cannot create redis connection: %v", err)
		return
	}
	log.Info("connecting to redis cache")

	// listen tcp port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// routing grpc services
	route.GrpcRoute(grpcServer, db, log, cache)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	log.Info("serve grpc on port: " + port)
}
