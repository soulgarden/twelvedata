package response

// PressReleases represents the response structure for press releases data.
type PressReleases struct {
	Pagination    PressReleasesPagination `json:"pagination"`
	PressReleases []PressRelease          `json:"press_releases"`
	Status        string                  `json:"status"`
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

// PressReleasesPagination identifies the current page and its requested size.
type PressReleasesPagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}
