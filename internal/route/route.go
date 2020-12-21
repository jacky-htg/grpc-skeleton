package route

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"waresix.com/cities/internal/pkg/db/redis"
	"waresix.com/cities/internal/service"
	"waresix.com/cities/waresix.com/cities"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *logrus.Entry, cache *redis.Cache) {
	cityServer := service.CityServer{Db: db, Log: log, Cache: cache}
	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)
}
