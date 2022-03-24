package models

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
