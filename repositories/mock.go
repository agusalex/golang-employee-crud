package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func newMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}
