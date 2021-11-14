package graph

import (
	"abdukhashimov/mybron.uz/logger"
	"abdukhashimov/mybron.uz/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services *services.Services
	Log      logger.Logger
}
