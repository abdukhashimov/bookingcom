package graph

import (
	"abdukhashimov/mybron.uz/pkg/logger"
	"abdukhashimov/mybron.uz/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services *services.Services
	log      logger.Logger
}

func NewResolver(log logger.Logger, services *services.Services) Resolver {
	return Resolver{
		log:      log,
		services: services,
	}
}
