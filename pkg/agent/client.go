package agent

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
	url := fmt.Sprintf("%s%s", c.url, "/ruleset/import_replace")
	resp, err := utils.Post(url, []byte(req), map[string]string{})
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("response error status %s from %s, error massage:%s", resp.Status(), url, resp.String())
	}
	return nil
}

func (c Client) RuleStatus(name string) (string, error) {
	url := fmt.Sprintf("%s/rules/%s/status", c.url, name)
	resp, err := utils.Get(url, map[string]string{}, map[string]string{})
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", fmt.Errorf("response error status %s from %s, error massage:%s", resp.Status(), url, resp.String())
	}
	return resp.String(), nil
}
