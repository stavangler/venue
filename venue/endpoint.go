package venue

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateVenue endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateVenue: MakeCreateVenueEndpoint(s),
	}
}

func MakeCreateVenueEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Create)
		err = s.Create(req.ID, req.Name)
		return apiResponse{Err: err}, err
	}
}

type apiResponse struct {
	Err error `json:"err,omitempty"`
}
