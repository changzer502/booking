package response

type PageData struct {
	Total    int64 `json:"total"`
	PageData any   `json:"pageData"`
}
