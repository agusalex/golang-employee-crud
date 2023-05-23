package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagRepository_GetAll(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "tag1").
		AddRow(2, "tag2")

	mock.ExpectQuery("^SELECT (.+) FROM `tags`").WillReturnRows(rows)

	repo := NewTagRepository(db)
	tags, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
}

func TestTagRepository_GetByName(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tag1")
	mock.ExpectQuery("^SELECT (.+) FROM `tags` WHERE (.+)").WithArgs("tag1").WillReturnRows(rows)

	repo := NewTagRepository(db)
	tag, err := repo.GetByName("tag1")
	assert.NoError(t, err)
	assert.Equal(t, "tag1", tag.Name)

	mock.ExpectQuery("^SELECT (.+) FROM `tags` WHERE (.+)").WithArgs("nonexistent").WillReturnRows(sqlmock.NewRows([]string{}))
	tag, err = repo.GetByName("nonexistent")
	assert.NoError(t, err)
	assert.Nil(t, tag)
}
