package hardware

import (
	"log"
	"space-management-system/app"
)

type Methods struct {
	App *app.App
}

func (m *Methods) SetLocker(id int) error {
	log.Println("Locker set", id)
	return nil
}
