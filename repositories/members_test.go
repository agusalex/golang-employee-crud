package repositories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberRepository_GetAll(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "type", "role"}).
		AddRow(1, "John Doe", "EMPLOYEE", "Engineer").
		AddRow(2, "Jane Doe", "EMPLOYEE", "Engineer")
	mock.ExpectQuery("^SELECT (.+) FROM `members` WHERE (.+)$").
		WillReturnRows(rows)

	tagsRows := sqlmock.NewRows([]string{"id", "name"})
	mock.ExpectQuery("^SELECT (.+) FROM `tags` INNER JOIN `member_tags` ON `member_tags`.`tag_id` = `tags`.`id` WHERE (.+)$").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(tagsRows)

	repo := NewMemberRepository(db)
	members, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, members, 2)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestMemberRepository_GetByID(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "type", "role"}).
		AddRow(1, "John Doe", "EMPLOYEE", "Engineer")
	mock.ExpectQuery("^SELECT (.+) FROM `members` WHERE (.+)$").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	tagsRows := sqlmock.NewRows([]string{"id", "name"})
	mock.ExpectQuery("^SELECT (.+) FROM `tags` INNER JOIN `member_tags` ON `member_tags`.`tag_id` = `tags`.`id` WHERE (.+)$").
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(tagsRows)

	repo := NewMemberRepository(db)
	member, err := repo.GetByID(fmt.Sprint(1))

	assert.NoError(t, err)
	assert.NotNil(t, member)
	assert.Equal(t, "John Doe", member.Name)
	assert.Equal(t, "EMPLOYEE", member.Type)
	assert.Equal(t, "Engineer", member.Role)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMemberRepository_Save(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		Tags:             nil,
		ContractDuration: 1,
	}

	mock.ExpectBegin()

	mock.ExpectExec("^INSERT INTO `members` (.+)$").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), member.Name, member.Type, sqlmock.AnyArg(), member.ContractDuration).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := NewMemberRepository(db)
	savedMember, err := repo.Save(member)

	assert.NoError(t, err)
	assert.NotNil(t, savedMember)
	assert.Equal(t, "John Doe", savedMember.Name)
	assert.Equal(t, "CONTRACTOR", savedMember.Type)
	assert.Equal(t, 1, savedMember.ContractDuration)
	assert.NoError(t, mock.ExpectationsWereMet())
}
