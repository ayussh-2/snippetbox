package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	title   string
	content string
	created time.Time
	expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// (m *SnippetModel) method reciver attaches the below fn to snippetModel to be used as SnippetModel.Insert()....

func (m *SnippetModel) Insert(title string, content string, created time.Time, expires time.Time) (int, *error) {
	return 0, nil
}

func (m *SnippetModel) Get(ID int) (*SnippetModel, *error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*SnippetModel, error) {
	return nil, nil
}
