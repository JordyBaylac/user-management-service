package utils

import (
	uuid "github.com/satori/go.uuid"
)

type UniqueIDGenerator interface {
	GenerateID() string
}

type UUIDGenerator struct {
}

func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}

func (generator UUIDGenerator) GenerateID() string {
	return uuid.NewV4().String()
}
