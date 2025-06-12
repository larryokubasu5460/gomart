package client

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type UserClient interface {
	UserExists(userID uuid.UUID) (bool, error)
}

type userClient struct {
	baseURL string
	client  *http.Client
}

func NewUserClient(baseURL string) UserClient {
	return &userClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (u *userClient) UserExists(userID uuid.UUID) (bool, error) {
	url := fmt.Sprintf("%s/users/%s", u.baseURL, userID.String())
	resp, err := u.client.Get(url)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("user-service error: %d", resp.StatusCode)
	}
	return true, nil
}
