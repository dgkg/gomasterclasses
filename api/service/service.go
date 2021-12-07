package service

import "github.com/gin-gonic/gin"

type Service struct {
	router   *gin.Engine
	services []Handler
}

type Handler interface {
	InitRoutes(r *gin.Engine)
}

func New() *Service {
	return &Service{
		router: gin.Default(),
	}
}

func (s *Service) AddService(h Handler) {
	s.services = append(s.services, h)
}

func (s *Service) InitRoutes() *Service {
	// Handler user
	var uh *UserHandlers
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
