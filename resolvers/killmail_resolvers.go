package resolvers

import "github.com/dreae/esi-graphql/resolvers/http"
import "encoding/json"
import "strconv"

// KillmailNode holds references to a recent killmail
type KillmailNode struct {
	KillmailHash string `json:"killmail_hash"`
	KillmailID   int32  `json:"killmail_id"`
}

type KillmailNodeResolver struct {
	node *KillmailNode
}

// KillmailID returns the ID of the killmail
func (r *KillmailNodeResolver) KillmailID() int32 {
	return r.node.KillmailID
}

// KillmailHash returns the hash of the killmail
func (r *KillmailNodeResolver) KillmailHash() string {
	return r.node.KillmailHash
}

// Killmail serves as a pointer to fetch the details of this killmail
func (r *KillmailNodeResolver) Killmail() (*KillmailResolver, error) {
	return GetKillmailDetails(r.node.KillmailID, r.node.KillmailHash)
}

// Killmail holds all of the details for a kill in EVE
type Killmail struct {
	Attackers     []KillmailAttacker `json:"attackers"`
	KillmailID    int32              `json:"killmail_id"`
	KillmailTime  string             `json:"killmail_time"`
	MoonID        int32              `json:"moon_id"`
	SolarSystemID int32              `json:"solar_system_id"`
	Victim        KillmailVictim     `json:"victim"`
	WarID         int32              `json:"war_id"`
}

// KillmailAttacker holds details about an attacker
type KillmailAttacker struct {
	AllianceID     int32   `json:"alliance_id"`
	CharacterID    int32   `json:"character_id"`
	CorporationID  int32   `json:"corporation_id"`
	DamageDone     int32   `json:"damage_done"`
	FactionID      int32   `json:"faction_id"`
	FinalBlow      bool    `json:"final_blow"`
	SecurityStatus float64 `json:"security_status"`
	ShipTypeID     int32   `json:"ship_type_id"`
	WeaponTypeID   int32   `json:"weapon_type_id"`
}

// KillmailVictim holds details about a victim
type KillmailVictim struct {
	AllianceID    int32          `json:"alliance_id"`
	CharacterID   int32          `json:"character_id"`
	CorporationID int32          `json:"corporation_id"`
	DamageTaken   int32          `json:"damage_taken"`
	FactionID     int32          `json:"faction_id"`
	Items         []KillmailItem `json:"items"`
	Position      EVEPosition    `json:"position"`
	ShipTypeID    int32          `json:"ship_type_id"`
}

// KillmailItem holds details about an item in a killmail
type KillmailItem struct {
	Flag              int32          `json:"flag"`
	ItemTypeID        int32          `json:"item_type_id"`
	Items             []KillmailItem `json:"items"`
	QuantityDestroyed int32          `json:"quantity_destroyed"`
	QuantityDropped   int32          `json:"quantity_dropped"`
	Singleton         int32          `json:"singleton"`
}

type KillmailResolver struct {
	killmail *Killmail
}

type KillmailAttackerResolver struct {
	attacker *KillmailAttacker
}

type KillmailVictimResolver struct {
	victim *KillmailVictim
}

type KillmailItemResolver struct {
	item *KillmailItem
}

// Attackers returns the array of attackers for this killmail
func (r *KillmailResolver) Attackers() *[]*KillmailAttackerResolver {
	var resolvers []*KillmailAttackerResolver
	for idx := range r.killmail.Attackers {
		resolvers = append(resolvers, &KillmailAttackerResolver{&r.killmail.Attackers[idx]})
	}

	return &resolvers
}

// KillmailID returns the ID of this killmail
func (r *KillmailResolver) KillmailID() int32 {
	return r.killmail.KillmailID
}

// KillmailTime returns the time of the EVE kill
func (r *KillmailResolver) KillmailTime() string {
	return r.killmail.KillmailTime
}

// MoonID returns the ID of the moon where this killmail took place
func (r *KillmailResolver) MoonID() *int32 {
	return &r.killmail.MoonID
}

// SolarSystemID returns the ID of the solar system where this kill took place
func (r *KillmailResolver) SolarSystemID() int32 {
	return r.killmail.SolarSystemID
}

// Victim returns the victim of this killmail
func (r *KillmailResolver) Victim() *KillmailVictimResolver {
	return &KillmailVictimResolver{&r.killmail.Victim}
}

// WarID returns the ID of any associated wars
func (r *KillmailResolver) WarID() *int32 {
	return &r.killmail.WarID
}

// AllianceID returns the ID of the attacker's alliance
func (r *KillmailAttackerResolver) AllianceID() *int32 {
	return &r.attacker.AllianceID
}

// CorporationID returns the ID of the attacker's corporaion
func (r *KillmailAttackerResolver) CorporationID() *int32 {
	return &r.attacker.CorporationID
}

// CharacterID returns the ID of the attacking character
func (r *KillmailAttackerResolver) CharacterID() *int32 {
	return &r.attacker.CharacterID
}

// DamageDone returns the ammount of damage the attacker did
func (r *KillmailAttackerResolver) DamageDone() *int32 {
	return &r.attacker.DamageDone
}

// FactionID returns the ID of the attacker's faction
func (r *KillmailAttackerResolver) FactionID() *int32 {
	return &r.attacker.FactionID
}

// FinalBlow returns whether or no this attacker got the final blow
func (r *KillmailAttackerResolver) FinalBlow() *bool {
	return &r.attacker.FinalBlow
}

// SecurityStatus returns the security status of the attacker
func (r *KillmailAttackerResolver) SecurityStatus() *float64 {
	return &r.attacker.SecurityStatus
}

// WeaponTypeID returns the attacker's weapon type
func (r *KillmailAttackerResolver) WeaponTypeID() *int32 {
	return &r.attacker.WeaponTypeID
}

// ShipTypeID returns the attacker's ship type
func (r *KillmailAttackerResolver) ShipTypeID() *int32 {
	return &r.attacker.ShipTypeID
}

// Character is a pointer to fetch the details of the attacking character
func (r *KillmailAttackerResolver) Character() (*CharacterResolver, error) {
	return GetCharacterByID(r.attacker.CharacterID)
}

// CharacterID returns the ID of the victim's character
func (r *KillmailVictimResolver) CharacterID() *int32 {
	return &r.victim.CharacterID
}

// AllianceID returns the ID of the vicitim's alliance
func (r *KillmailVictimResolver) AllianceID() *int32 {
	return &r.victim.AllianceID
}

// CorporationID returns the ID of the victim's corporation
func (r *KillmailVictimResolver) CorporationID() *int32 {
	return &r.victim.CorporationID
}

// DamageTaken returns the ammount of damage the victim took
func (r *KillmailVictimResolver) DamageTaken() *int32 {
	return &r.victim.DamageTaken
}

// Position retuns the position where the victim was killed
func (r *KillmailVictimResolver) Position() *PositionResolver {
	return &PositionResolver{&r.victim.Position}
}

// Items returns an array of items the victim was carrying
func (r *KillmailVictimResolver) Items() *[]*KillmailItemResolver {
	var items []*KillmailItemResolver
	for idx := range r.victim.Items {
		items = append(items, &KillmailItemResolver{&r.victim.Items[idx]})
	}

	return &items
}

// ShipTypeID returns the ID of the type of the victim's ship
func (r *KillmailVictimResolver) ShipTypeID() *int32 {
	return &r.victim.ShipTypeID
}

// FactionID returns the ID of the victim's faction
func (r *KillmailVictimResolver) FactionID() *int32 {
	return &r.victim.FactionID
}

// Flag returns the flag of this item
func (r *KillmailItemResolver) Flag() *int32 {
	return &r.item.Flag
}

// ItemTypeID returns the ID of this item
func (r *KillmailItemResolver) ItemTypeID() *int32 {
	return &r.item.ItemTypeID
}

// QuantityDestroyed returns the number of items destroyed
func (r *KillmailItemResolver) QuantityDestroyed() *int32 {
	return &r.item.QuantityDestroyed
}

// QuantityDropped returns the number of items dropped
func (r *KillmailItemResolver) QuantityDropped() *int32 {
	return &r.item.QuantityDropped
}

// Singleton returns this item's singleton flag
func (r *KillmailItemResolver) Singleton() *int32 {
	return &r.item.Singleton
}

// Items return the items contained in this items
func (r *KillmailItemResolver) Items() *[]*KillmailItemResolver {
	var resolvers []*KillmailItemResolver
	for idx := range r.item.Items {
		resolvers = append(resolvers, &KillmailItemResolver{&r.item.Items[idx]})
	}

	return &resolvers
}

// GetRecentKillsByCharacter gets the recent killmails for given character ID
func GetRecentKillsByCharacter(characterID int32, auth string, before *int32) (*[]*KillmailNodeResolver, error) {
	params := make(map[string]string)
	if before != nil {
		params["max_kill_id"] = strconv.Itoa(int(*before))
	}

	resp, err := http.MakeAuthorizedQuery(auth, params, "characters/%d/killmails/recent/", characterID)
	if err != nil {
		return nil, err
	}

	var mails []KillmailNode
	if err := json.NewDecoder(resp.Body).Decode(&mails); err != nil {
		return nil, err
	}

	var resolvers []*KillmailNodeResolver
	for idx := range mails {
		resolvers = append(resolvers, &KillmailNodeResolver{&mails[idx]})
	}

	return &resolvers, nil
}

// GetKillmailDetails returns the details of a specific killmail
func GetKillmailDetails(killmailID int32, killmailHash string) (*KillmailResolver, error) {
	resp, err := http.MakeRequest("killmails/%d/%s/", killmailID, killmailHash)
	if err != nil {
		return nil, err
	}

	var killmail Killmail
	if err := json.NewDecoder(resp.Body).Decode(&killmail); err != nil {
		return nil, err
	}

	return &KillmailResolver{&killmail}, nil
}
