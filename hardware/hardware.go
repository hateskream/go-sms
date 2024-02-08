package hardware

import (
	"log"
	"space-management-system/app"
)

type Methods struct {
	Storage app.StorageInterface
}

func (m *Methods) SetLocker(id int) error {
	log.Println("Locker set", id)
	return nil
}
