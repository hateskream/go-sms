package handlers

import (
	"space-management-system/app"
)

type Handlers struct {
	Storage  app.StorageInterface
	Hardware app.HardwareInterface
}
