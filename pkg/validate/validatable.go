package validate

import "github.com/amidgo/amiddocs/pkg/amiderrors"

type Validatable interface {
	Validate() *amiderrors.ErrorResponse
}

type ValidatableStruct interface {
	Validatable
	ValidatableVariables() []Validatable
}

func ValidateStructVariables(vars ...Validatable) *amiderrors.ErrorResponse {
	for _, v := range vars {
		err := v.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
