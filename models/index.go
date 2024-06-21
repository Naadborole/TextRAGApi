package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Index struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NDocuments int    `json:"nDocuments"`
}

func NewIndex(nm string) Index {
	return Index{
		Name:       nm,
		Id:         uuid.NewString(),
		NDocuments: 0,
	}
}

func (i Index) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, NDocs: %v", i.Id, i.Name, i.NDocuments)
}
