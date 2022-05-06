package services

import (
	"llevapp/src/middlewares"
	"llevapp/src/routes"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//APIService APIService
type APIService struct {
	port              string
	logger            *logrus.Entry
	Engine            *gin.Engine
	PostgreSQLService *PostgreSQLService
}

//NewAPIService returns a service instance.
func NewAPIService(port string) *APIService {
	return &APIService{
		port: port,
	}
}

//Health Health
func (service *APIService) Health() bool {
	return true
}

//InjectServices InjectServices
func (service *APIService) InjectServices(logger *logrus.Entry, services []Service) {

	service.logger = logger

	for _, otherService := range services {
		if PostgreSQLService, ok := otherService.(*PostgreSQLService); ok {

			if PostgreSQLService.connectionString != "" {
				service.PostgreSQLService = PostgreSQLService
			} else {
				service.PostgreSQLService = nil
			}

		}
	}

}

//Init Init this service
func (service *APIService) Init() error {
	service.logger.Info("[APIService] Initializing...")
	service.logger.Info("[APIService] Using port: " + service.port)

	service.Engine = gin.Default()
	service.Engine.Use(middlewares.CORSMiddleware())
	service.Engine.Use(gin.Recovery())
	return nil
}

//Execute Execute this service
func (service *APIService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[APIService] Executing...")

	err := routes.EndpointGroup(service.Engine, service.PostgreSQLService.db)
	err = service.Engine.Run(":" + service.port)
	if err != nil {
		service.logger.Fatal("[APIService] Failed running api server: " + err.Error())
		return err
	}

	return nil
}

// Endpoint
