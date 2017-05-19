package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Character struct {
	CharacterID    int32
	AllianceID     int32   `json:"alliance_id"`
	AncestryID     int32   `json:"ancestry_id"`
	Birthday       string  `json:"birthday"`
	BloodlineID    int32   `json:"bloodline_id"`
	CorporationID  int32   `json:"corporation_id"`
	Description    string  `json:"description"`
	Gender         string  `json:"gender"`
	Name           string  `json:"name"`
	RaceID         int32   `json:"race_id"`
	SecurityStatus float64 `json:"security_status"`
}

type CharacterResolver struct {
	character *Character
}

func (c *CharacterResolver) AllianceID() *int32 {
	return &c.character.AllianceID
}

func (c *CharacterResolver) AncestryID() *int32 {
	return &c.character.AncestryID
}

func (c *CharacterResolver) Birthday() *string {
	return &c.character.Birthday
}

func (c *CharacterResolver) BloodlineID() *int32 {
	return &c.character.BloodlineID
}

func (c *CharacterResolver) CorporationID() *int32 {
	return &c.character.CorporationID
}

func (c *CharacterResolver) Description() *string {
	return &c.character.Description
}

func (c *CharacterResolver) Gender() *string {
	return &c.character.Gender
}

func (c *CharacterResolver) Name() *string {
	return &c.character.Name
}

func (c *CharacterResolver) RaceID() *int32 {
	return &c.character.RaceID
}

func (c *CharacterResolver) SecurityStatus() *float64 {
	return &c.character.SecurityStatus
}

func (c *CharacterResolver) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(c.character.CorporationID)
}

func (c *CharacterResolver) Skills(ctx context.Context) (*CharacterSkillsResolver, error) {
	return GetSkillsForCharID(ctx.Value(ContextKey("auth")).(string), c.character.CharacterID)
}

func GetCharacterByID(charID int32) (*CharacterResolver, error) {
	var char Character
	resp, err := http.Get(fmt.Sprintf("https://esi.tech.ccp.is/latest/characters/%d/?datasource=tranquility", charID))
	if err != nil {
		return &CharacterResolver{&char}, err
	}

	json.NewDecoder(resp.Body).Decode(&char)
	char.CharacterID = charID

	return &CharacterResolver{&char}, nil
}
