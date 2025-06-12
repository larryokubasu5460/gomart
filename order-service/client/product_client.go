package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ProductDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

type ProductClient interface {
	GetProduct(productID uuid.UUID) (*ProductDTO, error)
}

type productClient struct {
	baseURL string
	client *http.Client
}

func NewProductClient(baseURL string) ProductClient {
	return &productClient{
		baseURL: baseURL,
		client: &http.Client{},
	}
}

func (p *productClient) GetProduct(productID uuid.UUID) (*ProductDTO, error) {
	url := fmt.Sprintf("%s/products/%s",p.baseURL,productID.String())
	resp, err := p.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("product not found")
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product-service error: %d", resp.StatusCode)
	}

	var product ProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}