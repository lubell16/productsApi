package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Sam",
		Price: 1.00,
		SKU:   "aaa-aaa-aaa",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
