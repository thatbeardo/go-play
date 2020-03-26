package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product struct to denote a Product at the coffee shoppe
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products represents a list of products.
type Products []*Product

// FromJSON converts the payload from JSON format
func (p *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)
}

// ToJSON converts the payload to JSON format
func (p *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)
}

// Database and methods to modify underlying data
var productsList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       3.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong shot of coffee",
		Price:       4.45,
		SKU:         "xyz123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// AddProduct is called to update the existing list of products
func AddProduct(product *Product) {
	product.ID = getNextID()
	productsList = append(productsList, product)
}

// GetProducts is called when a GET call is made to retrieve all current Products
func GetProducts() Products {
	return productsList
}

// UpdateProduct is called when a PATCH call is made to change fields of a product.
func UpdateProduct(id int, p *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}
	productsList[index] = p
	return nil
}

// DeleteProduct will remove a product from the database
func DeleteProduct(id int) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}
	productsList[len(productsList)-1], productsList[index] = productsList[index], productsList[len(productsList)-1]
	productsList = productsList[:len(productsList)-1]
	return nil
}

func getNextID() int {
	currentID := productsList[len(productsList)-1].ID
	return currentID + 1
}

func findProduct(id int) (*Product, int, error) {
	for index, product := range productsList {
		if product.ID == id {
			return product, index, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// Validate determines if a Product payload is valid
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}

func validateSku(fl validator.FieldLevel) bool {
	// Sku is of format abc-asf-asw
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

// ErrProductNotFound is thrown when a an attempt is made to update non-existent product
var ErrProductNotFound = fmt.Errorf("Product not found")
