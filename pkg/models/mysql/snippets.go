package mysql

import (
	"database/sql"
	"errors"
	"snippitbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?,  UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() and id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
