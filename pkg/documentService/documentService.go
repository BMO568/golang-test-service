package documentService

import (
	"context"
	"fmt"
	"github.com/jellydator/ttlcache/v2"
	"myTest/pkg/models"
	"myTest/pkg/repositories"
	"sort"
	"strconv"
	"sync"
	"time"
)

type DocumentService interface {
	AddDocument(ctx context.Context, document models.Document) (string, error)
	UpdateDocument(ctx context.Context, document models.Document) (string, error)
	GetAllDocuments(ctx context.Context, offset *int32, limit *int32) ([]*models.Document, error)
	GetDocumentById(ctx context.Context, id int32) (*models.Document, error)
	DeleteDocumentById(ctx context.Context, id int32) (string, error)
}

type documentService struct {
	repository repositories.DocumentRepository
	cache      *ttlcache.Cache
}

func NewReindexerService(repo repositories.DocumentRepository) DocumentService {
	return &documentService{
		repository: repo,
		cache:      ttlcache.NewCache(),
	}
}

func (s documentService) AddDocument(ctx context.Context, document models.Document) (string, error) {
	if _, err := s.repository.CreateDocument(ctx, document); err != nil {
		fmt.Println(err)
		return "fail", err
	}
	return "ok", nil
}

func (s documentService) GetDocumentById(ctx context.Context, id int32) (*models.Document, error) {

	var document *models.Document
	documentI, exists := s.cache.Get(strconv.Itoa(int(id)))

	if exists != nil {
		var err error
		document, err = s.repository.GetDocumentById(ctx, id)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		s.cache.SetWithTTL(strconv.Itoa(int(id)), document, 15*time.Minute)
	} else {
		document = documentI.(*models.Document)
	}

	sortOptionsInDocs(document)

	return document, nil
}

func (s documentService) DeleteDocumentById(ctx context.Context, id int32) (string, error) {
	err := s.repository.DeleteDocumentById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return "fail", err
	}
	return "ok", nil
}

func (s documentService) UpdateDocument(ctx context.Context, document models.Document) (string, error) {
	if _, err := s.repository.UpdateDocument(ctx, document); err != nil {
		fmt.Println(err)
		return "fail", err
	}
	return "ok", nil
}

func (s documentService) GetAllDocuments(ctx context.Context, offset *int32, limit *int32) ([]*models.Document, error) {
	documents, err := s.repository.GetAllDocuments(ctx, offset, limit)
	if err != nil {
		fmt.Println(err)
		var empty []*models.Document
		return empty, err
	}

	sortOptionsInDocs(documents...)

	return documents, nil
}

func sortOptionsInDocs(documents ...*models.Document) {
	var wg sync.WaitGroup
	wg.Add(len(documents))

	var sortOptionsInDoc = func(document *models.Document) {
		fmt.Println(document.Id)
		defer wg.Done()
		sort.SliceStable(document.Options, func(i, j int) bool {
			return document.Options[i].Sort > document.Options[j].Sort
		})
	}

	for _, v := range documents {
		go sortOptionsInDoc(v)
	}

	wg.Wait()
	fmt.Println("stop")
}
