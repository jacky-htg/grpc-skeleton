package service

import (
	context "context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"waresix.com/cities/internal/model"
	"waresix.com/cities/internal/pkg/db/redis"
	"waresix.com/cities/waresix.com/cities"
)

// CityServer struct for price agreement
type CityServer struct {
	Db    *sql.DB
	Log   *logrus.Entry
	Cache *redis.Cache
}

// GetCity function
func (s *CityServer) GetCity(ctx context.Context, in *cities.GetCityInput) (*cities.City, error) {
	var cityModel model.City
	cityModel.ID = in.Id
	err := cityModel.Get(ctx, s.Db)
	if err != nil {
		s.Log.Error(err)
		return &cities.City{}, err
	}

	return &cities.City{Id: cityModel.ID, Name: cityModel.Name}, nil
}

// GetCities function
func (s *CityServer) GetCities(ctx context.Context, in *cities.EmptyInput) (*cities.Cities, error) {
	var list cities.Cities
	var cityModel model.City
	result, err := cityModel.List(ctx, s.Db)

	if err != nil {
		s.Log.Error(err)
		return &list, err
	}

	for _, city := range result {
		list.Cities = append(list.Cities, &cities.City{Id: city.ID, Name: city.Name})
	}

	return &list, nil
}

// CreateCity func
func (s *CityServer) CreateCity(ctx context.Context, in *cities.NewCityInput) (*cities.City, error) {

	var cityModel model.City
	cityModel.Name = in.Name
	err := cityModel.Create(ctx, s.Db)

	if err != nil {
		s.Log.Error(err)
		return &cities.City{}, err
	}

	return &cities.City{Id: cityModel.ID, Name: cityModel.Name}, nil
}
