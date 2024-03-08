package uuidv4

import (
	"github.com/google/uuid"
	"github.com/sousair/apollo-cloud/internal/domain/providers"
)

type UuidV4Provider struct{}

var _ providers.UuidProvider = (*UuidV4Provider)(nil)

func NewUuidV4Provider() *UuidV4Provider {
	return &UuidV4Provider{}
}

func (p UuidV4Provider) Generate() string {
	return uuid.New().String()
}
