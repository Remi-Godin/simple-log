package validation

type Validator interface {
	Validate(fieldData string) error
}

type ValidationError struct {
	Message string
}

func (ve ValidationError) Error() string {
	return ve.Message
}

type ValidatedInputField interface {
	GetValidators() []Validator
	GetFieldValue() string
}

func Validate(data ValidatedInputField) error {
	for _, val := range data.GetValidators() {
		err := val.Validate(data.GetFieldValue())
		if err != nil {
			return err
		}
	}
	return nil
}
