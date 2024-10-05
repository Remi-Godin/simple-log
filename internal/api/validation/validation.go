package api

type FieldValidationData struct {
	FieldData string
	Err       string
	Links     map[string]string
}

func NewFieldValidationData() FieldValidationData {
	return FieldValidationData{
		Links: make(map[string]string),
	}
}
