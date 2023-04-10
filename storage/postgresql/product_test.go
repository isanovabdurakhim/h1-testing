package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateProduct
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateProduct{
				ProductName: "Iphone",
				BrandId:     1,
				CategoryId:  1,
				ModelYear:   2022,
				ListPrice:   10000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := productTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  *models.Product
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				ProductId: 22,
			},
			Output: &models.Product{
				ProductName: "Iphone",
				BrandId:     1,
				CategoryId:  1,
				ModelYear:   2022,
				ListPrice:   10000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			product, err := productTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if product.ProductId != test.Output.ProductId || product.ProductName != test.Output.ProductName {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *product, *test.Output)
				return
			}

		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateProduct
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateProduct{
				ProductName: "Iphone",
				BrandId:  1,
				CategoryId:     1,
				ModelYear:     2022,
				ListPrice:    10000,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				ProductId: 22,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Delete(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}
