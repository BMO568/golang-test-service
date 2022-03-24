package pkg

import (
	"myTest/pkg/documentService"
	"myTest/pkg/repositories"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func CreateRouter(repository repositories.DocumentRepository) *http.ServeMux {
	var svc documentService.DocumentService = documentService.NewReindexerService(repository)

	router := http.NewServeMux()

	AddDocumentHandler := httptransport.NewServer(
		documentService.MakeAddDocumentEndpoint(svc),
		documentService.DecodeAddDocumentRequest,
		documentService.EncodeResponse,
	)

	GetAllDocumentsHandler := httptransport.NewServer(
		documentService.MakeGetAllDocumentsEndpoint(svc),
		documentService.DecodeGetAllDocumentsRequest,
		documentService.EncodeResponse,
	)

	GetDocumentByIdHandler := httptransport.NewServer(
		documentService.MakeGetDocumentByIdEndpoint(svc),
		documentService.DecodeGetDocumentByIdRequest,
		documentService.EncodeResponse,
	)

	DeleteDocumentByIdHandler := httptransport.NewServer(
		documentService.MakeDeleteDocumentByIdEndpoint(svc),
		documentService.DecodeDeleteDocumentByIdRequest,
		documentService.EncodeResponse,
	)

	UpdateDocumentHandler := httptransport.NewServer(
		documentService.MakeUpdateDocumentEndpoint(svc),
		documentService.DecodeUpdateDocumentRequest,
		documentService.EncodeResponse,
	)

	router.Handle("/addDocument", AddDocumentHandler)

	router.Handle("/getAllDocuments", GetAllDocumentsHandler)

	router.Handle("/getDocumentById", GetDocumentByIdHandler)

	router.Handle("/updateDocument", UpdateDocumentHandler)

	router.Handle("/deleteDocumentById", DeleteDocumentByIdHandler)

	return router
}
