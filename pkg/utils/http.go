package utils

import (
	"github.com/go-resty/resty/v2"
)

func Get(url string, params map[string]string, headers map[string]string) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetDebugBodyLimit(1000000)
	resp, err := client.R().
		SetQueryParams(params).
		SetHeaders(headers).
		Get(url)

	//OutputRequestResponse(resp, err)
	return resp, err
}

func Post(url string, data interface{}, headers map[string]string) (*resty.Response, error) {
	//logs.Infof("post data =%+v", data)
	client := resty.New()
	client.SetDebug(true)
	client.SetDebugBodyLimit(1000000)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeaders(headers).
		SetBody(data).
		Post(url)
	//OutputRequestResponse(resp, err)
	return resp, err
}

func Put(url string, data interface{}, headers map[string]string) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetDebugBodyLimit(1000000)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeaders(headers).
		SetBody(data).
		Put(url)
	//OutputRequestResponse(resp, err)
	return resp, err
}

func Delete(url string, data interface{}, headers map[string]string) (*resty.Response, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetDebugBodyLimit(1000000)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeaders(headers).
		SetBody(data).
		Delete(url)
	//OutputRequestResponse(resp, err)
	return resp, err
}
