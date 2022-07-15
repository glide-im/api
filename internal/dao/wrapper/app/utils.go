package app

import (
	"github.com/google/uuid"
)

func GenerateAppId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
