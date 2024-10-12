package pages

type PageData struct {
	Title string
	Links map[string]string
}

func newPageData(title string) PageData {
	return PageData{
		Title: title,
		Links: make(map[string]string),
	}
}
