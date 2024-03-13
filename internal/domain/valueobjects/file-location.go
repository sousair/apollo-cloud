package valueobjects

type FileLocation struct {
	URL       string `json:"url"`
	Provider  string `json:"provider"`
	Extension string `json:"extension"`
}
