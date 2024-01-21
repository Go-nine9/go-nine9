package services

import "github.com/google/uuid"

func GenerateUUID() (uuid.UUID, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
