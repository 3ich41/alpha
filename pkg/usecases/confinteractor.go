package usecases

import (
	"m15.io/alpha/pkg/domain"
)

type ConfInteractor struct {
	ConfRepository domain.ConfRepository
}

func (interactor *ConfInteractor) ConfigureClientApp(confRequest *domain.ConfRequest) (*domain.Conf, error) {
	conf, err := interactor.ConfRepository.GetConf(confRequest)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
