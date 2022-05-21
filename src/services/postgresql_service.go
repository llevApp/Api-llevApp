package services

import (
	"database/sql"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//PostgreSQLService PostgreSQLService
type PostgreSQLService struct {
	connectionString string
	logger           *logrus.Entry
	db               *sql.DB
}

//NewPostgreSQLService returns a service instance.
func NewPostgreSQLService(connectionString string) *PostgreSQLService {
	return &PostgreSQLService{
		connectionString: connectionString,
	}
}

//Health Health
func (service *PostgreSQLService) Health() bool {
	return true
}

//InjectServices InjectServices
func (service *PostgreSQLService) InjectServices(logger *logrus.Entry, otherServices []Service) {
	service.logger = logger
}

//Init Init this service
func (service *PostgreSQLService) Init() error {
	service.logger.Info("[PostgreSQLService] Initializing...")
	service.logger.Info("[PostgreSQLService] Using connection string: " + service.connectionString)

	var err error

	service.db, err = sql.Open("postgres", service.connectionString)
	if err != nil {
		service.logger.Fatal("[PostgreSQLService] Failed connecting to database: " + err.Error())
		return err
	}
	//defer service.db.Close()
	service.logger.Info("[PostgreSQLService] Connected to database successfully")
	return nil
}

//Execute Execute this service
func (service *PostgreSQLService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[PostgreSQLService] Executing...")

	defer service.db.Close()

	waitGroup.Wait()

	return nil
}
