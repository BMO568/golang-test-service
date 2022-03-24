package documentService

import (
	"context"
	"encoding/json"
	"myTest/pkg/models"
	"net/http"
)

type DocumentByIdRequest struct {
	Id int32 `json:"documentId"`
}

type GetAllDocumentsRequest struct {
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
}

type DocumentRequest struct {
	document models.Document
}

type GetDocumentByIdResponse struct {
	Document *models.Document `json:"document"`
	Err      string           `json:"error,omitempty"`
}

type GetAllDocumentsResponse struct {
	Documents []*models.Document `json:"documents,omitempty"`
	Err       string             `json:"error,omitempty"`
}

type SetDocumentResponse struct {
	Msg string `json:"msg"`
	Err error  `json:"error,omitempty"`
}

type DeleteDocumentByIdResponse struct {
	Status string `json:"status"`
	Err    error  `json:"error,omitempty"`
}

func DecodeGetDocumentByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DocumentByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeDeleteDocumentByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DocumentByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetAllDocumentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllDocumentsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeAddDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req.document); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req.document); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
