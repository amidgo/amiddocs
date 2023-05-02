package validate

type Validatable interface {
	Validate() error
}

type ValidatableStruct interface {
	ValidatableVariables() []Validatable
}

func ValidateFields(v ValidatableStruct) error {
	for _, v := range v.ValidatableVariables() {
		err := v.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
