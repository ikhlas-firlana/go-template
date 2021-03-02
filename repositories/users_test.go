package repositories

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ikhlas-firlana/go-template/interfaces"
	"github.com/ikhlas-firlana/go-template/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB    *gorm.DB
	mock  sqlmock.Sqlmock
	mocky mock.Mock

	repo interfaces.IUsers
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repo = Provide(s.DB)

}

func TestInitUserRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetName() {
	expected := models.Users{
		Name:     "abc",
		Username: "def",
	}
	row := sqlmock.
		NewRows([]string{"username", "name"}).
		AddRow(expected.Username, expected.Name)

	sqlContent := "SELECT * FROM `users` WHERE  ((username = ?))"
	s.mock.
		ExpectQuery(regexp.QuoteMeta(sqlContent)).
		WithArgs(expected.Username).
		WillReturnRows(row)

	res, err := s.repo.GetUserByUsername(expected.Username)

	s.T().Logf("Expected! %v", res)

	require.Error(s.T(), err, nil)
	assert.Equal(s.T(), expected, res)
}
