package fields

type GenericFormField struct {
	FieldId       string
	FieldName     string
	FieldType     string
	FieldValue    string
	FieldNameText string
}

func NewGenericFormField() GenericFormField {
	return GenericFormField{}
}
