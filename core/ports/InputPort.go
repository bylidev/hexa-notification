package ports

import (
	"github.com/igloar96/hexa-notification/core/domain"
)

type InputPort interface {
	GetMessage() (*domain.Message, error)
}
