package resolvers

import (
	"context"
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

// Character holds all information about an EVE character
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

// AllianceID returns the alliance ID of the character's alliance,
// if they're in one
func (c *CharacterResolver) AllianceID() *int32 {
	return &c.character.AllianceID
}

// AncestryID returns the ID of the character's ancestry
func (c *CharacterResolver) AncestryID() *int32 {
	return &c.character.AncestryID
}

// Birthday returns the string date the character was created
func (c *CharacterResolver) Birthday() string {
	return c.character.Birthday
}

// BloodlineID returns the ID of the character's blood line
func (c *CharacterResolver) BloodlineID() int32 {
	return c.character.BloodlineID
}

// CorporationID returns the ID of the corporation the character is a member of
func (c *CharacterResolver) CorporationID() int32 {
	return c.character.CorporationID
}

// Description returns the character's bio
func (c *CharacterResolver) Description() *string {
	return &c.character.Description
}

// Gender returns the character's gender as a string
func (c *CharacterResolver) Gender() string {
	return c.character.Gender
}

// Name returns the character's name
func (c *CharacterResolver) Name() string {
	return c.character.Name
}

// RaceID returns the ID of the character's race
func (c *CharacterResolver) RaceID() int32 {
	return c.character.RaceID
}

// SecurityStatus returns the character's security status
func (c *CharacterResolver) SecurityStatus() *float64 {
	return &c.character.SecurityStatus
}

// Corporation serves as a pointer to fetch the the details of the corporation
// the character is a member of
func (c *CharacterResolver) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(c.character.CorporationID)
}

// Alliance serves as a pointer to fetch the alliance the character is a member of
func (c *CharacterResolver) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(c.character.AllianceID)
}

// Skills fetches an object representing the character's skills.
// Using this method requires an auth token
func (c *CharacterResolver) Skills(ctx context.Context) (*CharacterSkillsResolver, error) {
	return GetSkillsForCharID(ctx.Value(ContextKey("auth")).(string), c.character.CharacterID)
}

// SkillQueue fetches an object represeting the character's current skill queue.
// Using this method requires an auth token
func (c *CharacterResolver) SkillQueue(ctx context.Context) (*[]*SkillQueueResolver, error) {
	return GetSkillQueueForCharID(ctx.Value(ContextKey("auth")).(string), c.character.CharacterID)
}

// RecentKillmails fetches the character's recent killmails.
// Using this method requires an auth token
func (c *CharacterResolver) RecentKillmails(ctx context.Context, args *struct{ Before *int32 }) (*[]*KillmailNodeResolver, error) {
	return GetRecentKillsByCharacter(c.character.CharacterID, ctx.Value(ContextKey("auth")).(string), args.Before)
}

// GetCharacterByID fetches an EVE character by a given ID
func GetCharacterByID(charID int32) (*CharacterResolver, error) {
	var char Character
	resp, err := http.MakeRequest("characters/%d/", charID)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&char); err != nil {
		return nil, err
	}

	char.CharacterID = charID

	return &CharacterResolver{&char}, nil
}
