package lib

import "github.com/google/uuid"

func GenUuid() (id string) {
	id = uuid.New()
	return
}
