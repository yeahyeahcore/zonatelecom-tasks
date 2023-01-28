package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/client"
)

const (
	checkDigestURL = "/check"
)

type DigestClientDeps struct {
	Logger        *logrus.Logger
	Configuration *core.DigestServiceConfiguration
}

type DigestClient struct {
	logger *logrus.Logger
	client *resty.Client
}

func NewDigestClient(deps *DigestClientDeps) *DigestClient {
	return &DigestClient{
		logger: deps.Logger,
		client: resty.New().
			EnableTrace().
			SetBaseURL(deps.Configuration.BaseURL),
	}
}

func (receiver *DigestClient) Check(digest string) error {
	if _, err := client.Post(checkDigestURL, &client.RequestSettings[interface{}]{
		Driver: receiver.client,
		Body:   map[string]string{"digest": digest},
	}); err != nil {
		requestError := fmt.Errorf("check digest request error: %s", err.Error())
		receiver.logger.Errorln(requestError)
		return ErrWrongDigest
	}

	return nil
}
