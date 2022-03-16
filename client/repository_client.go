package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository struct {
	Owner           *string `json:"owner,omitempty"`
	Name            *string `json:"name,omitempty"`
	Private         *bool   `json:"private,omitempty"`
	HasIssues       *bool   `json:"has_issues,omitempty"`
	AutoInit        *bool   `json:"auto_init,omitempty"`
	LicenseTemplate *string `json:"license_template,omitempty"`
}

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}

func (c *Client) Create(ctx context.Context, repo *Repository) (*Repository, error) {
	body, _ := json.Marshal(repo)
	req, err := http.NewRequest(http.MethodPost, "https://api.github.com/user/repos", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return newRepository(resp), nil
}

func (c *Client) Delete(ctx context.Context, owner, name string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v", owner, name)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func newRepository(r *http.Response) *Repository {
	repo := &Repository{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, repo)
	return repo
}
