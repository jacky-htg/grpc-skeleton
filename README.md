# Golang gRPC Skeleton
gRPC skeleton using golang.

## How to Use
- create your repository,  ex : git@bitbucket.org:waresix/cities.git 
- open your terminal and clone grpc-skeleton: git clone git@bitbucket.org:waresix/grpc-skeleton.git
- git fetch origin
- git remote set-url origin git@bitbucket.org:waresix/cities.git
- go mod init waresix.com/cities
- find replace waresix.com/grpc-skeleton with waresix.com/cities
- cp .env.example .env 
- add sql script to internal/schema/migrate.go
- create file proto/cities.proto
- run protoc --go_out=plugins=grpc:. proto/cities.proto
- create file internal/service/cities.go
- create file internal/model/city.go
- add grpc route at internal/route/route.go
- test the grpc service with make a grpc client at client.go

## Example Schema Migrate
```
{
	Version:     1,
	Description: "Create cities Table",
	Script: `
		CREATE TABLE cities (
			id serial PRIMARY KEY,
			name varchar NOT NULL
		);
	`,
},
```

## Example proto/cities.proto
```
syntax = "proto3";
option go_package="waresix.com/cities";

message EmptyInput {}

message GetCityInput {
  int32 id = 1;
}

message NewCityInput {
  string name = 1;
}

message Cities {
  repeated City cities = 1;
}

message City {
  int32 id = 1;
  string name = 2;
}

service CitiesService {
 rpc GetCity(GetCityInput) returns (City) {}
 rpc GetCities(EmptyInput) returns (Cities) {}
 rpc CreateCity(NewCityInput) returns (City) {}
}

```

## Example internal/service/cities.go
```
package service

import (
	context "context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"waresix.com/cities/internal/model"
	"waresix.com/cities/internal/pkg/db/redis"
)

// CityServer struct for price agreement
type CityServer struct {
	Db    *sql.DB
	Log   *logrus.Entry
	Cache *redis.Cache
}

// GetCity function
func (s *CityServer) GetCity(ctx context.Context, in *city.GetCityInput) (*city.City, error) {
	var cityModel model.City
	cityModel.ID = in.Id
	err := cityModel.Get(ctx, s.Db)
	if err != nil {
		return &city.City{}, err
	}

	return &city.City{Id: cityModel.ID, Name: cityModel.Name}, nil
}

// GetCities function
func (s *CityServer) GetCities(ctx context.Context, in *city.EmptyInput) (*city.Cities, error) {
	var list city.Cities
	var cityModel model.City
	cities, err := cityModel.List(ctx, s.Db)

	if err != nil {
		return &list, err
	}

	for _, city := range cities {
		list = append(list, city.City{Id: city.ID, Name: city.Name})
	}

	return &list, nil
}

// CreateCity func
func (s *CityServer) CreateCity(ctx context.Context, in *city.NewCityInput) (*city.City, error) {

	var cityModel model.City
	err := cityModel.Create(ctx, s.Db, in)

	if err != nil {
		return &city.City{}, err
	}

	return &city.City{Id: cityModel.ID, Name: cityModel.Name}, nil
}

```

## Example internal/model/city.go
```
package model

import (
	"context"
	"database/sql"
)

// City struct
type City struct {
	ID   int32
	Name string
}

// List of City
func (u *City) List(ctx context.Context, db *sql.DB) ([]City, error) {
	var list []City
	var err error
	query := `SELECT id, name FROM cities`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var city City
		err = rows.Scan(&city.ID, &city.Name)
		if err != nil {
			return list, err
		}

		list = append(list, city)
	}

	return list, rows.Err()
}

// Get City
func (u *City) Get(ctx context.Context, db *sql.DB) error {
	return db.QueryRowContext(ctx, `SELECT id, name FROM cities WHERE id = $1`, u.ID).Scan(&u.ID, &u.Name)
}

// Create New City
func (u *City) Create(ctx context.Context, db *sql.DB, in *city.NewCityInput) error {

	_, err = tx.ExecContext(ctx, `INSERT INTO cities (name) VALUES ($1)`, in.Name)

	return err
}

```

## Example gRPC Routing
add this code at func GrpcRoute() in internal/route/route.go 
```
oncallServer := service.OnCallServer{Db: db, Log: log, Cache: cache}
oncall.RegisterOnCallPriceAgreementServiceServer(grpcServer, &oncallServer)
```

## Example test as grpc Client
create file client.go
```
package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"grpc-client/waresix.com/cities"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := cities.NewCitiesServiceClient(conn)

	response, err := c.CreateCity(context.Background(), &cities.NewCityInput{Name: "jakarta"})
	if err != nil {
		log.Fatalf("Error when calling CreateCity: %s", err)
	}
	log.Printf("Response from server: %v", response)

	response2, err := c.GetCity(context.Background(), &cities.GetCityInput{})
	if err != nil {
		log.Fatalf("Error when calling GetCity: %s", err)
	}
	log.Printf("Response from server: %v", response2)

  response3, err := c.GetCities(context.Background(), &cities.EmptyInput{})
	if err != nil {
		log.Fatalf("Error when calling GetCity: %s", err)
	}
	log.Printf("Response from server: %v", response3)
}

```