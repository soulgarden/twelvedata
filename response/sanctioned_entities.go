package response

import "github.com/guregu/null/v6"

// SanctionedEntities represents the response structure for sanctioned entities data.
type SanctionedEntities struct {
	Meta               SanctionedEntitiesMeta `json:"meta"`
	SanctionedEntities []SanctionedEntity     `json:"sanctioned_entities"`
	Pagination         null.Value[Pagination] `json:"pagination"`
}

// SanctionedEntitiesMeta contains metadata for sanctioned entities data.
type SanctionedEntitiesMeta struct {
	Source      string `json:"source"`
	LastUpdated string `json:"last_updated"`
	TotalCount  int    `json:"total_count"`
}

// SanctionedEntity represents a single sanctioned entity record.
type SanctionedEntity struct {
	EntityID        string      `json:"entity_id"`
	EntityName      string      `json:"entity_name"`
	EntityType      string      `json:"entity_type"`
	SanctionDate    string      `json:"sanction_date"`
	SanctionProgram string      `json:"sanction_program"`
	SanctionReason  null.String `json:"sanction_reason"`
	Country         null.String `json:"country"`
	Nationality     null.String `json:"nationality"`
	DateOfBirth     null.String `json:"date_of_birth"`
	PlaceOfBirth    null.String `json:"place_of_birth"`
	AlternateNames  []string    `json:"alternate_names"`
	Addresses       []Address   `json:"addresses"`
	SanctionStatus  string      `json:"sanction_status"`
	LastUpdated     string      `json:"last_updated"`
}

// Address represents an address associated with a sanctioned entity.
type Address struct {
	Address1   string      `json:"address1"`
	Address2   null.String `json:"address2"`
	City       null.String `json:"city"`
	State      null.String `json:"state"`
	Country    null.String `json:"country"`
	PostalCode null.String `json:"postal_code"`
}

// Pagination represents pagination information for sanctioned entities.
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalCount int `json:"total_count"`
}
