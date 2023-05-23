package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSearchRepository_SearchMembers(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "type", "role"}).
		AddRow(1, "John Doe", "CONTRACTOR", "Engineer").
		AddRow(2, "Jane Doe", "CONTRACTOR", "Engineer")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT `members`.* FROM `members` JOIN member_tags ON member_tags.member_id = members.id JOIN tags ON tags.id = member_tags.tag_id WHERE `members`.`deleted_at` IS NULL AND ((tags.name IN (?)) AND (members.type = ?)) GROUP BY members.id HAVING (COUNT(DISTINCT tags.id) = ?)")).
		WithArgs("Java", "CONTRACTOR", 1).
		WillReturnRows(rows)

	repo := NewSearchRepository(db)
	tags := []string{"Java"}
	members, err := repo.SearchMembers(tags, "CONTRACTOR")
	assert.NoError(t, err)
	assert.Len(t, members, 2)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
