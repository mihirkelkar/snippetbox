package mysql

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("Models: mo matching record found")

type SnippetDB interface {
	Insert(string, string, string) (int, error)
	Get(int) (*Snippet, error)
	Latest() ([]*Snippet, error)
}

type snippetDB struct {
	DB *sql.DB
}

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

func NewSnippetDB(db *sql.DB) SnippetDB {
	return &snippetDB{DB: db}
}

func (s *snippetDB) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
           VALUES(?, ?, UTC_TIMESTASMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *snippetDB) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
           WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := s.DB.QueryRow(stmt, id)
	snippet := &Snippet{}

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	}
	return snippet, nil
}

func (s *snippetDB) Latest() ([]*Snippet, error) {
	stmt := `
           SELECT id, title, content, created, expires FROM snippets
           WHERE expires > UTC_TIMESTAMP ORDER BY created DESC LIMIT 10
          `
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)

		//Check if any errors happened during the time of execution
		if err = rows.Err(); err != nil {
			return nil, err
		}

	}

	return snippets, nil
}
