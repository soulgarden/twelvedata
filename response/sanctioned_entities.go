package response

// SanctionedEntities represents the response structure for sanctioned entities data.
type SanctionedEntities struct {
	Sanctions []SanctionedEntity `json:"sanctions"`
	Count     int                `json:"count"`
	Status    string             `json:"status"`
}

// SanctionedEntity represents a single sanctioned entity record.
type SanctionedEntity struct {
	Symbol   string   `json:"symbol"`
	Name     string   `json:"name"`
	MicCode  string   `json:"mic_code"`
	Country  string   `json:"country"`
	Sanction Sanction `json:"sanction"`
}

// Sanction represents a sanction record for an entity.
type Sanction struct {
	Source  string         `json:"source"`
	Program string         `json:"program"`
	Notes   string         `json:"notes"`
	Lists   []SanctionList `json:"lists"`
}

// SanctionList represents a sanctions list entry.
type SanctionList struct {
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
}
