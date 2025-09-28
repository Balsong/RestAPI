package saveinfo

import "github.com/rs/zerolog"

type ISaveService interface {
}

type service struct {
	logger *zerolog.Logger
}

func NewSaveService(logger *zerolog.Logger) ISaveService {
	return &service{
		logger: logger,
	}
}
