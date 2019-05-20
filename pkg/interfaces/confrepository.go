package interfaces

import (
	"m15.io/alpha/pkg/delivery/grpc/conf_grpc"
	"m15.io/alpha/pkg/domain"
)

type ConfHandler interface {
	GetConf(fetchRequest *conf_grpc.FetchRequest) (*conf_grpc.Conf, error)
}

type GrpcConfRepository struct {
	handler ConfHandler
}

func NewGrpcConfRepository(confHandler ConfHandler) *GrpcConfRepository {
	repository := new(GrpcConfRepository)
	repository.handler = confHandler

	return repository
}

func (repo *GrpcConfRepository) GetConf(confRequest *domain.ConfRequest) (*domain.Conf, error) {
	fetchRequest := &conf_grpc.FetchRequest{
		Username:  confRequest.Username,
		Ipaddr:    confRequest.IPAddr,
		Mac:       confRequest.Mac,
		Timestamp: confRequest.Timestamp,
	}

	grpcConf, err := repo.handler.GetConf(fetchRequest)
	if err != nil {
		return nil, err
	}

	conf := repo.transformGrpcDomain(grpcConf)

	return conf, nil
}

func (repo *GrpcConfRepository) transformGrpcDomain(grpcConf *conf_grpc.Conf) *domain.Conf {
	conf := new(domain.Conf)
	conf.Username = grpcConf.Username

	buttons := make([]domain.Button, 0, 25)
	for _, grpcButton := range grpcConf.Button {
		button := domain.Button{
			Text:  grpcButton.Text,
			Value: grpcButton.Value,
		}
		buttons = append(buttons, button)
	}
	conf.Buttons = buttons

	return conf
}
