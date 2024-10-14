package pages

type PageData struct {
	Title string
	Data  map[string]string
	Links map[string]string
}

func NewPageData(title string) PageData {
	return PageData{
		Title: title,
		Data:  make(map[string]string),
		Links: make(map[string]string),
	}
}
