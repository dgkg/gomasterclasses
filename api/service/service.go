package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db"
)

type Service struct {
	router   *gin.Engine
	services []Handler
	db       db.Storage
	log      *zap.Logger
}

// Handler is an interface for creating new services
// that can be added with AddService function.
type Handler interface {
	InitRoutes(r *gin.Engine)
}

func New(db db.Storage, log *zap.Logger) *Service {
	s := &Service{
		router: gin.Default(),
		db:     db,
		log:    log,
	}
	return s
}

// AddService is adding a service to the main Service.
func (s *Service) AddService(h Handler) {
	s.services = append(s.services, h)
}

func (s *Service) InitRoutes() *Service {
	// Handler user
	var uh *UserHandlers = &UserHandlers{
		db:  s.db,
		log: s.log,
	}

	uh.InitRoutes(s.router)

	// Handlers added
	for k := range s.services {
		s.services[k].InitRoutes(s.router)
	}
	return s
}

func (s *Service) Run(port string) {
	if len(port) == 0 {
		panic("please register a port.")
	}

	s.router.Run(":" + port)
}
