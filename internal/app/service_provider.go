package app

import (
	"basic-backend/internal/api/user"
	"basic-backend/internal/config"
	"basic-backend/internal/repository"
	userRepository "basic-backend/internal/repository/user"
	"basic-backend/internal/service"
	userService "basic-backend/internal/service/user"
	"log"
)

type serviceProvider struct {
	httpConfig     config.HTTPConfig
	userAPI        *user.API
	userService    service.UserService
	userRepository repository.UserRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("app: %q", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.New()
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.New(s.UserRepository())
	}

	return s.userService
}

func (s *serviceProvider) UserAPI() *user.API {
	if s.userAPI == nil {
		s.userAPI = user.New(s.UserService())
	}

	return s.userAPI
}
