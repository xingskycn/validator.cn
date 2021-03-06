package validator_test

import (
	"fmt"
	"testing"

	"github.com/liuhengloveyou/validator.cn"
)

type ValidateExample struct {
	IdCard      string `validate:"idcard"`
	Id2         int64  `validate:"idcard"`
	Name        string `validate:"nonzero"`
	Description string
	Age         int    `validate:"min=18"`
	Email       string `validate:"nonzero,email"`
	Address     struct {
		Street string `validate:"nonzero"`
		City   string `validate:"nonzero"`
	}
}

func TestParseTags(t *testing.T) {

	ve := ValidateExample{
		Name:        "Joe Doe", // valid as it's nonzero
		Description: "",        // valid no validation tag exists
		Age:         19,        // invalid as age is less than required 18
	}

	ve.Email = "@not.a.valid.email"
	ve.Address.City = "Some City" // valid
	ve.Address.Street = ""        // invalid
	ve.IdCard = "41289719870908876x"
	ve.Id2 = 123

	err := validator.Validate(ve)
	fmt.Println(err)
}
