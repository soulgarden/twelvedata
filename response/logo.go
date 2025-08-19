package response

// Logo represents company logo information with metadata and URL.
type Logo struct {
	Meta      LogoMeta `json:"meta"`
	URL       string   `json:"url"`
	LogoBase  string   `json:"logo_base,omitempty"`
	LogoQuote string   `json:"logo_quote,omitempty"`
}

// LogoMeta contains metadata for logo data.
type LogoMeta struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
}
