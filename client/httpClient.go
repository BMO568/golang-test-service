package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type DValue struct {
	Id    int32  `json:"id" reindex:"valueId"`
	Title string `json:"title" reindex:"title"`
}

type DOption struct {
	Id          int32    `json:"id" reindex:"optionId"`
	Sort        int32    `json:"sort" reindex:"sort,tree"`
	Description string   `json:"description,omitempty" reindex:"description"`
	Values      []DValue `json:"values,omitempty" reindex:"values"`
}

type Document struct {
	Id          int32     `json:"id" reindex:"id,,pk"`
	Description string    `json:"description,omitempty" reindex:"description"`
	Options     []DOption `json:"options,omitempty" reindex:"options"`
}

type SetDocumentResponse struct {
	Msg string `json:"msg"`
	Err error  `json:"error,omitempty"`
}

func generateDocuments() []Document {
	var generateDVal = func() DValue {
		return DValue{
			Id:    int32(rand.Intn(100)),
			Title: "TitleNew " + time.Now().Format("2006-01-02T15:04:05"),
		}
	}

	var generateDOpt = func() DOption {
		var dValues []DValue = make([]DValue, 5)
		for i := range dValues {
			dValues[i] = generateDVal()
		}

		return DOption{
			Id:          int32(rand.Intn(200)),
			Sort:        int32(rand.Intn(200)),
			Description: "DescriptionONew " + time.Now().Format("2006-01-02T15:04:05"),
			Values:      dValues,
		}
	}

	var generateDocument = func(indx int) Document {
		var dOpts = make([]DOption, 5)
		for i := range dOpts {
			dOpts[i] = generateDOpt()
		}

		return Document{
			Id:          int32(indx),
			Description: "DocumentDescription " + time.Now().Format("2006-01-02T15:04:05"),
			Options:     dOpts,
		}
	}

	var documents []Document = make([]Document, 1)

	for i := range documents {
		documents[i] = generateDocument(i + 2)
	}

	return documents
}

func main() {
	fmt.Println("httpClient start")
	addNewDoc()
	// updateDoc()
}

func updateDoc() {
	document := generateDocuments()[0]

	json_data, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://localhost:8000/updateDocument",
		"application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		panic(err)
	}

	var documentResp SetDocumentResponse

	json.NewDecoder(resp.Body).Decode(&documentResp)

	fmt.Println(documentResp)
}

func addNewDoc() {
	document := generateDocuments()[0]

	json_data, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://localhost:8000/addDocument",
		"application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		panic(err)
	}

	var documentResp SetDocumentResponse

	json.NewDecoder(resp.Body).Decode(&documentResp)

	fmt.Println(documentResp)
}
