package resolvers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func (c *CorporationResolver) AllianceID() *int32 {
	return &c.corp.AllianceID
}

func (c *CorporationResolver) CEOID() *int32 {
	return &c.corp.CEOID
}

func (c *CorporationResolver) Description() *string {
	return &c.corp.Description
}

func (c *CorporationResolver) Name() *string {
	return &c.corp.Name
}

func (c *CorporationResolver) CreationDate() *string {
	return &c.corp.CreationDate
}

func (c *CorporationResolver) CreatorID() *int32 {
	return &c.corp.CreatorID
}

func (c *CorporationResolver) Faction() *string {
	return &c.corp.Faction
}

func (c *CorporationResolver) MemberCount() *int32 {
	return &c.corp.MemberCount
}

func (c *CorporationResolver) TaxRate() *float64 {
	return &c.corp.TaxRate
}

func (c *CorporationResolver) Ticker() *string {
	return &c.corp.Ticker
}

func (c *CorporationResolver) URL() *string {
	return &c.corp.URL
}

func (c *CorporationResolver) CEO() (*CharacterResolver, error) {
	return GetCharacterByID(c.corp.CEOID)
}

func (c *CorporationResolver) Creator() (*CharacterResolver, error) {
	return GetCharacterByID(c.corp.CreatorID)
}

func (c *CorporationResolver) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(c.corp.AllianceID)
}

func GetCorpByID(corpID int32) (*CorporationResolver, error) {
	var corp Corporation
	resp, err := http.Get(fmt.Sprintf("https://esi.tech.ccp.is/latest/corporations/%d/?datasource=tranquility", corpID))
	if err != nil {
		return &CorporationResolver{&corp}, err
	}

	json.NewDecoder(resp.Body).Decode(&corp)

	return &CorporationResolver{&corp}, nil
}
