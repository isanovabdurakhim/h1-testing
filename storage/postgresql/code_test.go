package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateCode(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCode
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCode{
				CodeName:        "Test Code",
				Discount:        20000,
				DiscountType:    "fixed",
				OrderLimitPrice: 40000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := codeTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdCode(t *testing.T) {
	tests := []struct {
		CodeName string
		Input    *models.CodePrimaryKey
		Output   *models.Code
		WantErr  bool
	}{
		{
			CodeName: "Case 1",
			Input: &models.CodePrimaryKey{
				Code_Id: 2,
			},
			Output: &models.Code{
				Code_Id:         2,
				CodeName:        "Test Code",
				Discount:        20000,
				DiscountType:    "fixed",
				OrderLimitPrice: 40000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.CodeName, func(t *testing.T) {

			code, err := codeTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.CodeName, err)
				return
			}

			if code.CodeName != test.Output.CodeName {
				t.Errorf("%s: got: %v, expected: %v", test.CodeName, *code, *test.Output)
				return
			}

		})
	}
}

func TestUpdateActor(t *testing.T) {
	tests := []struct {
		CodeName string
		Input    *models.UpdateCode
		Output   int64
		WantErr  bool
	}{
		{
			CodeName: "Case 1",
			Input: &models.UpdateCode{
				CodeName:        "Test Code",
				Discount:        20000,
				DiscountType:    "fixed",
				OrderLimitPrice: 40000,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.CodeName, func(t *testing.T) {
			rows, err := codeTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.CodeName, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.CodeName, rows, test.Output)
				return
			}
		})
	}
}

func TestDeleteCodeode(t *testing.T) {
	tests := []struct {
		CodeName string
		Input    *models.CodePrimaryKey
		Output   int64
		WantErr  bool
	}{
		{
			CodeName: "Case 1",
			Input: &models.CodePrimaryKey{
				Code_Id: 2,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.CodeName, func(t *testing.T) {

			rows, err := codeTestRepo.Delete(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.CodeName, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.CodeName, rows, test.Output)
				return
			}

		})
	}
}
