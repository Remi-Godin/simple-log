package forms

type FormData struct {
	FormName             string
	FormDesc             string
	FormFields           []string
	FormSubmissionLink   string
	ResourceDeletionLink string
	Patch                bool
}

func NewFormData(formName string) FormData {
	return FormData{
		FormName: formName,
	}
}
