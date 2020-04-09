package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Gandharv",
		Price: 1,
		SKU:   "abs-acd-erd",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
