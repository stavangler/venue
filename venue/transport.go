package venue

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/log"

	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("PUT").Path("/create").Handler(createVenueHandler(e, options))
	r.Methods("PUT").Path("/name").Handler(createSetNameHandler(e, options))
	r.Methods("GET").Path("/query").Handler(createQueryHandler(e, options))

	return r
}

func createVenueHandler(e Endpoints, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		e.CreateVenue,
		decodeCreateVenue,
		encodeResponse,
		options...,
	)
}

func createSetNameHandler(e Endpoints, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		e.SetName,
		decodeSetName,
		encodeResponse,
		options...,
	)
}

func createQueryHandler(e Endpoints, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		e.Query,
		decodeQuery,
		encodeResponse,
		options...,
	)
}

func decodeQuery(_ context.Context, r *http.Request) (request interface{}, err error) {
	query := r.URL.Query()
	qm := &QueryModel{}
	qm.Id = query.Get("id")
	qm.Name = query.Get("name")

	return qm, err
}

func decodeSetName(_ context.Context, r *http.Request) (request interface{}, err error) {
	var command SetName

	err = json.NewDecoder(r.Body).Decode(&command)

	return command, err
}

func decodeCreateVenue(_ context.Context, r *http.Request) (request interface{}, err error) {
	var create Create

	err = json.NewDecoder(r.Body).Decode(&create)

	return create, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrMissingID:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
