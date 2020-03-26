package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 4.99,
		SKU:   "abs-asc-asf",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
