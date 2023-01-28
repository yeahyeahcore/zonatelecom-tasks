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

type BetaClientDeps struct {
	Logger        *logrus.Logger
	Configuration *core.BetaServiceConfiguration
}

type BetaClient struct {
	logger *logrus.Logger
	client *resty.Client
}

func NewBetaClient(deps *BetaClientDeps) *BetaClient {
	return &BetaClient{
		logger: deps.Logger,
		client: resty.New().
			EnableTrace().
			SetBaseURL(deps.Configuration.BaseURL),
	}
}

func (receiver *BetaClient) SendVotingState(request *core.VotingState) error {
	if _, err := client.Post(votingStateURL, &client.RequestSettings[interface{}]{
		Driver:   receiver.client,
		Formdata: request,
	}); err != nil {
		requestError := fmt.Errorf("send voting state request error: %s", err.Error())
		receiver.logger.Errorln(requestError)
		return requestError
	}

	return nil
}
