package documentService

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllDocumentsEndpoint(srv DocumentService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req := request.(GetAllDocumentsRequest)
		documents, err := srv.GetAllDocuments(ctx, req.Skip, req.Take)
		if err != nil {
			return GetAllDocumentsResponse{Documents: documents, Err: "no data found"}, err
		}

		return GetAllDocumentsResponse{Documents: documents, Err: ""}, nil
	}
}

func MakeGetDocumentByIdEndpoint(srv DocumentService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req := request.(DocumentByIdRequest)
		document, err := srv.GetDocumentById(ctx, req.Id)
		if err != nil {
			return GetDocumentByIdResponse{
				Document: document,
				Err:      "Id not found",
			}, nil
		}
		return GetDocumentByIdResponse{Document: document, Err: ""}, nil
	}
}

func MakeDeleteDocumentByIdEndpoint(srv DocumentService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req := request.(DocumentByIdRequest)
		status, err := srv.DeleteDocumentById(ctx, req.Id)
		return DeleteDocumentByIdResponse{
			Status: status,
			Err:    err,
		}, nil
	}
}

func MakeAddDocumentEndpoint(srv DocumentService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req := request.(DocumentRequest)
		msg, err := srv.AddDocument(ctx, req.document)
		return SetDocumentResponse{Msg: msg, Err: err}, nil
	}
}

func MakeUpdateDocumentEndpoint(srv DocumentService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {
		req := request.(DocumentRequest)
		msg, err := srv.UpdateDocument(ctx, req.document)
		return SetDocumentResponse{Msg: msg, Err: err}, nil
	}
}
