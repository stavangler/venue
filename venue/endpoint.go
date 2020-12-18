package venue

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateVenue endpoint.Endpoint
	SetName     endpoint.Endpoint
	Query       endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateVenue: MakeCreateVenueEndpoint(s),
		SetName:     MakeSetNameEndpoint(s),
		Query:       MakeQueryVenueEndpoint(s),
	}
}

func MakeCreateVenueEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Create)
		err = s.Create(req.ID)
		return apiResponse{Err: err}, err
	}
}

func MakeSetNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetName)
		err = s.SetName(req.ID, req.Name)
		return apiResponse{Err: err}, err
	}
}

func MakeQueryVenueEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		qm := request.(*QueryModel)
		result := s.Query(qm)
		return apiResponse{Data: result, Err: err}, err
	}
}

type apiResponse struct {
	Err  error       `json:"err,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
