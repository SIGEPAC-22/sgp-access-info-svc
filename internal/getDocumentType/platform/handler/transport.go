package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"sgp-access-info-svc/kit/constants"
)

func NewGetDocumentTypeHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle(path,
		httptransport.NewServer(endpoints,
			DecodeRequestGetDocumentType,
			EncodeRequestGetDocumentType,
		)).Methods(http.MethodGet)
	return r
}

func DecodeRequestGetDocumentType(ctx context.Context, r *http.Request) (interface{}, error) {
	processID, _ := uuid.NewUUID()
	ctx = context.WithValue(ctx, constants.UUID, processID.String())
	return GetDocumentTypeInternalRequest{ctx: ctx}, nil
}

func EncodeRequestGetDocumentType(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, _ := response.(GetDocumentTypeInternalResponse)
	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch resp.Err {
		case constants.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
			break
		case constants.ErrNoContent:
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return json.NewEncoder(w).Encode(resp.Err.Error())
	}
	return json.NewEncoder(w).Encode(resp.Response)
}
