package response

// PressReleases represents the response structure for press releases data.
type PressReleases struct {
	PressReleases []PressRelease `json:"press_releases"`
	Status        string         `json:"status"`
}

// PressRelease represents a single press release item.
type PressRelease struct {
	ID       string   `json:"id"`
	Datetime string   `json:"datetime"`
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Style    string   `json:"style"`
	Language []string `json:"language"`
}
