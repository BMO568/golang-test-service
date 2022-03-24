package repositories

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"myTest/pkg/models"

	"github.com/restream/reindexer"
)

type DocumentRepository interface {
	CreateDocument(ctx context.Context, document models.Document) (int, error)
	GetAllDocuments(ctx context.Context, offset *int32, limit *int32) ([]*models.Document, error)
	GetDocumentById(ctx context.Context, id int32) (*models.Document, error)
	UpdateDocument(ctx context.Context, document models.Document) (int, error)
	DeleteDocumentById(ctx context.Context, id int32) error
}

type reindexerRepo struct {
	reindexerDb        *reindexer.Reindexer
	documentsNamespace string
}

func NewReindexerDocumentRepo(reindexerDb *reindexer.Reindexer) (DocumentRepository, error) {
	const documentsNamespace = "documents"

	// Temp
	reindexerDb.DropNamespace(documentsNamespace)

	reindexerDb.OpenNamespace(documentsNamespace, reindexer.DefaultNamespaceOptions(), models.Document{})

	// Default items generation
	documents := generateDocuments()
	for _, v := range documents {
		err := reindexerDb.Upsert(documentsNamespace, &v)
		if err != nil {
			panic(err)
		}
	}

	return &reindexerRepo{
		reindexerDb:        reindexerDb,
		documentsNamespace: documentsNamespace,
	}, nil
}

func (rr *reindexerRepo) CreateDocument(ctx context.Context, document models.Document) (int, error) {
	count, err := rr.reindexerDb.Insert(rr.documentsNamespace, &document)
	return count, err
}

func (rr *reindexerRepo) GetDocumentById(ctx context.Context, id int32) (*models.Document, error) {
	document, found := rr.reindexerDb.Query(rr.documentsNamespace).
		Where("id", reindexer.EQ, id).
		Get()

	if found {
		return document.(*models.Document), nil
	}
	return nil, errors.New("not found")
}

func (rr *reindexerRepo) DeleteDocumentById(ctx context.Context, id int32) error {
	err := rr.reindexerDb.Delete(rr.documentsNamespace, &models.Document{Id: id})
	return err
}

func (rr *reindexerRepo) UpdateDocument(ctx context.Context, document models.Document) (int, error) {
	count, err := rr.reindexerDb.Update(rr.documentsNamespace, &document)
	return count, err
}

func (rr *reindexerRepo) GetAllDocuments(ctx context.Context, offset *int32, limit *int32) ([]*models.Document, error) {
	var documentsIt *reindexer.Iterator
	if offset != nil || limit != nil {
		var offsetV = 0
		if offset != nil {
			offsetV = int(*offset)
		}

		var limitV = 5
		if limit != nil {
			limitV = int(*limit)
		}

		documentsIt = rr.reindexerDb.Query(rr.documentsNamespace).
			Offset(offsetV).
			Limit(limitV).
			Exec()
	} else {
		documentsIt = rr.reindexerDb.Query(rr.documentsNamespace).Exec()
	}

	data, err := documentsIt.FetchAll()
	if err != nil {
		fmt.Println(err)
	}

	documents := make([]*models.Document, len(data))
	for i, arg := range data {
		documents[i] = arg.(*models.Document)
	}

	return documents, err
}

func generateDocuments() []models.Document {
	var generateDVal = func() models.DValue {
		return models.DValue{
			Id:    int32(rand.Intn(100)),
			Title: "TitleV",
		}
	}

	var generateDOpt = func() models.DOption {
		var dValues []models.DValue = make([]models.DValue, 5)
		for i := range dValues {
			dValues[i] = generateDVal()
		}

		return models.DOption{
			Id:          int32(rand.Intn(200)),
			Sort:        int32(rand.Intn(200)),
			Description: "DescriptionO",
			Values:      dValues,
		}
	}

	var generateDocument = func(indx int) models.Document {
		var dOpts = make([]models.DOption, 5)
		for i := range dOpts {
			dOpts[i] = generateDOpt()
		}

		return models.Document{
			Id:          int32(indx),
			Description: "DocumentDescription",
			Options:     dOpts,
		}
	}

	var documents []models.Document = make([]models.Document, 5)

	for i := range documents {
		documents[i] = generateDocument(i + 1)
	}

	return documents
}
