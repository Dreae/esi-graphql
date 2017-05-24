package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

// Corporation holds the details of a corporation in EVE
type Corporation struct {
	AllianceID   int32   `json:"alliance_id"`
	CEOID        int32   `json:"ceo_id"`
	Description  string  `json:"corporation_description"`
	Name         string  `json:"corporation_name"`
	CreationDate string  `json:"creation_date"`
	CreatorID    int32   `json:"creator_id"`
	Faction      string  `json:"faction"`
	MemberCount  int32   `json:"member_count"`
	TaxRate      float64 `json:"tax_rate"`
	Ticker       string  `json:"ticker"`
	URL          string  `json:"url"`
}

type CorporationResolver struct {
	corp *Corporation
}

// AllianceID returns the ID of the alliance the corporation is in, if
// it's a member of an alliance
func (c *CorporationResolver) AllianceID() *int32 {
	return &c.corp.AllianceID
}

// CEOID returns the ID of the corporation's CEO
func (c *CorporationResolver) CEOID() *int32 {
	return &c.corp.CEOID
}

// Description returns the corporation's description
func (c *CorporationResolver) Description() *string {
	return &c.corp.Description
}

// Name returns the name of the corporation
func (c *CorporationResolver) Name() *string {
	return &c.corp.Name
}

// CreationDate returns the date string that the corporation was created
func (c *CorporationResolver) CreationDate() *string {
	return &c.corp.CreationDate
}

// CreatorID returns the ID of the corporation's founder
func (c *CorporationResolver) CreatorID() *int32 {
	return &c.corp.CreatorID
}

// Faction returns the ID of the corporation's faction, if it's a member of one
func (c *CorporationResolver) Faction() *string {
	return &c.corp.Faction
}

// MemberCount returns the number of memebers in a corporation
func (c *CorporationResolver) MemberCount() *int32 {
	return &c.corp.MemberCount
}

// TaxRate returns the tax rate of the corporation
func (c *CorporationResolver) TaxRate() *float64 {
	return &c.corp.TaxRate
}

// Ticker returns the corporation's ticker string
func (c *CorporationResolver) Ticker() *string {
	return &c.corp.Ticker
}

// URL returns the corporation's URL
func (c *CorporationResolver) URL() *string {
	return &c.corp.URL
}

// CEO serves as a pointer to fetch more details of the corporation's CEO
func (c *CorporationResolver) CEO() (*CharacterResolver, error) {
	return GetCharacterByID(c.corp.CEOID)
}

// Creator serves as a pointer to fetch more details of the corporation's founder
func (c *CorporationResolver) Creator() (*CharacterResolver, error) {
	return GetCharacterByID(c.corp.CreatorID)
}

// Alliance serves as a pointer to fetch more details of the corporation's alliance
func (c *CorporationResolver) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(c.corp.AllianceID)
}

// GetCorpByID fetches the details of a corporation by corporation ID
func GetCorpByID(corpID int32) (*CorporationResolver, error) {
	var corp Corporation
	resp, err := http.MakeRequest("corporations/%d/", corpID)
	if err != nil {
		return &CorporationResolver{&corp}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&corp); err != nil {
		return nil, err
	}

	return &CorporationResolver{&corp}, nil
}
