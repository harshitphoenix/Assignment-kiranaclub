package main

import (
	"database/sql"
	"server/controllers"
	"server/models"
	"testing"

	"github.com/stretchr/testify/require"

	"net/http/httptest"

	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	// utilHandler *UtilHandler
	repository controllers.ControllerInterface
	job        *models.Job
	metadata   *models.Metadata
	store      *models.Store
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)
	s.repository = controllers.NewUtilHandler(s.DB)
}
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
func (s *Suite) Test_GetJobStatus() {
	s.mock.ExpectQuery("[SELECT * FROM \"jobs\" WHERE id = $1]").WillReturnRows(sqlmock.NewRows([]string{"id", "store_id", "status"}).AddRow(1, "1", "completed"))
	req := httptest.NewRequest("GET", "https://example.com?job_id=1", nil)
	w := httptest.NewRecorder()
	s.repository.GetJobStatus(w, req)
	require.Equal(s.T(), 200, w.Code)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
