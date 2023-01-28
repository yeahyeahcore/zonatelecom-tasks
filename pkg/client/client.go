package client

import (
	"bytes"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/convert"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/utils"
)

var (
	defaultHeaders = map[string]string{
		"Content-Type": "application/json",
	}
)

type RequestSettings[T interface{}] struct {
	Driver   *resty.Client
	Headers  map[string]string
	Body     interface{}
	Formdata interface{}
	result   T
}

func Post[T interface{}](url string, settings *RequestSettings[T]) (*T, error) {
	request := setSettings(settings.Driver.R(), settings)

	return makeRequest[T](url, request.Post)
}

func Get[T interface{}](url string, settings *RequestSettings[T]) (*T, error) {
	request := setSettings(settings.Driver.R(), settings)

	return makeRequest[T](url, request.Get)
}

func setSettings[T interface{}](request *resty.Request, settings *RequestSettings[T]) *resty.Request {
	if settings == nil {
		return request
	}
	if settings.Body != nil {
		request.SetBody(settings.Body)
	}
	if settings.Formdata != nil {
		formdata, _ := convert.ObjectToStringMap(settings.Formdata, "formdata")
		request.SetFormData(formdata)
	}

	headers := utils.MergeMaps(defaultHeaders, settings.Headers)

	request.SetHeaders(headers)
	request.SetResult(settings.result)

	return request
}

func makeRequest[T interface{}](url string, requestMethod func(url string) (*resty.Response, error)) (*T, error) {
	responseInstance, err := requestMethod(url)
	if err != nil {
		return nil, err
	}
	if responseInstance.IsError() {
		return nil, fmt.Errorf("%s request error with status %d: %s",
			responseInstance.Request.Method,
			responseInstance.StatusCode(),
			responseInstance.Body(),
		)
	}

	responseBody, err := json.Parse[T](bytes.NewReader(responseInstance.Body()))
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
