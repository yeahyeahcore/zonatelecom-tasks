package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/client"
)

const (
	votingStateURL = "/voting-stats"
)

type GammaClientDeps struct {
	Logger        *logrus.Logger
	Configuration *core.GammaServiceConfiguration
}

type GammaClient struct {
	logger *logrus.Logger
	client *resty.Client
}

func NewGammaClient(deps *GammaClientDeps) *GammaClient {
	return &GammaClient{
		logger: deps.Logger,
		client: resty.New().
			EnableTrace().
			SetBaseURL(deps.Configuration.BaseURL),
	}
}

func (receiver *GammaClient) SendVotingState(request *core.PreviousVotingState) error {
	if _, err := client.Post(votingStateURL, &client.RequestSettings[interface{}]{
		Driver: receiver.client,
		Body:   request,
	}); err != nil {
		requestError := fmt.Errorf("send voting state request error: %s", err.Error())
		receiver.logger.Errorln(requestError)
		return requestError
	}

	return nil
}
