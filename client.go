package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"waresix.com/cities/waresix.com/cities"
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

	response2, err := c.GetCity(context.Background(), &cities.GetCityInput{Id: 1})
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
