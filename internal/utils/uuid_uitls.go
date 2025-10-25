package utils

import "github.com/google/uuid"

type UUID = uuid.UUID

func generateUUID() UUID {
	return uuid.New()
}
