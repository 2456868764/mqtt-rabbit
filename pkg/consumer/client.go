package consumer

import (
	"fmt"

	"bifromq_engine/pkg/utils"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c Client) RulesetImport(req string) error {
	importUrl := fmt.Sprintf("%s%s", c.url, "/ruleset/import")
	resp, err := utils.Post(importUrl, req, map[string]string{})
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("response error status %s from %s", resp.Status(), importUrl)
	}
	return nil
}

func (c Client) RuleStatus(name string) (string, error) {
	statusUrl := fmt.Sprintf("%s/rules/%s/status", c.url, name)
	resp, err := utils.Get(statusUrl, map[string]string{}, map[string]string{})
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", fmt.Errorf("response error status %s from %s", resp.Status(), statusUrl)
	}
	return resp.String(), nil
}
