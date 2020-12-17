package route

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"waresix.com/grpc-skeleton/internal/pkg/db/redis"
)

// GrpcRoute func
func GrpcRoute(grpcServer *grpc.Server, db *sql.DB, log *logrus.Entry, cache *redis.Cache) {

}
