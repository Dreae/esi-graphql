package resolvers

import (
	"context"
)

// ContextKey is used to access values in the context object
type ContextKey string

// Resolver is a placeholder type representing the root query
type Resolver struct{}

// DogmaAttribute queries ESI for a dogma attribute by ID
func (r *Resolver) DogmaAttribute(args *struct{ AttributeID int32 }) (*DogmaAttributeResolver, error) {
	return GetDogmaAttributeResolver(args.AttributeID)
}

// DogmaAttributes returns all dogma attributes reported by ESI
func (r *Resolver) DogmaAttributes() (*[]*DogmaAttributeNodeResolver, error) {
	return GetDogmaList()
}

// Type queries ESI for an EVEType by ID
func (r *Resolver) Type(args *struct{ TypeID int32 }) (*EVETypeResolver, error) {
	return GetEVEType(args.TypeID)
}

// Character queries ESI for a character by character ID
func (r *Resolver) Character(args *struct{ CharacterID int32 }) (*CharacterResolver, error) {
	return GetCharacterByID(args.CharacterID)
}

// Corporation queries ESI for a corporation by corporation ID
func (r *Resolver) Corporation(args *struct{ CorporationID int32 }) (*CorporationResolver, error) {
	return GetCorpByID(args.CorporationID)
}

// Skills queries ESI for a this skills of the character represented by the provided character ID
func (r *Resolver) Skills(ctx context.Context, args *struct{ CharacterID int32 }) (*CharacterSkillsResolver, error) {
	return GetSkillsForCharID(ctx.Value(ContextKey("auth")).(string), args.CharacterID)
}

// SkillQueue queries ESI for the skill queue of the character represented by the provided character ID
func (r *Resolver) SkillQueue(ctx context.Context, args *struct{ CharacterID int32 }) (*[]*SkillQueueResolver, error) {
	return GetSkillQueueForCharID(ctx.Value(ContextKey("auth")).(string), args.CharacterID)
}

// Alliance queries ESI for an alliance by alliance ID
func (r *Resolver) Alliance(args *struct{ AllianceID int32 }) (*AllianceResolver, error) {
	return GetAllianceByID(args.AllianceID)
}

// Search is used to search ESI for various types
func (r *Resolver) Search(args *struct {
	SearchTypes []*string
	Keyword     string
}) (*SearchResultsResolver, error) {
	return DoSearch(&args.SearchTypes, args.Keyword)
}

// Structures is used to get a list of all structures as reported by ESI
func (r *Resolver) Structures() (*[]*StructureNode, error) {
	return GetAllStructures()
}

// Structure queries ESI for details of a structure by structure ID
func (r *Resolver) Structure(ctx context.Context, args *struct{ StructureID float64 }) (*StructureResolver, error) {
	return GetStructureByID(args.StructureID, ctx.Value(ContextKey("auth")).(string))
}

// SolarSystem queries ESI for the details of a solar system by ID
func (r *Resolver) SolarSystem(args *struct{ SystemID int32 }) (*SolarSystemResolver, error) {
	return GetSolarSystemByID(args.SystemID)
}

// Killmail queries ESI to get details of a killmail
func (r *Resolver) Killmail(args *struct {
	KillmailID   int32
	KillmailHash string
}) (*KillmailResolver, error) {
	return GetKillmailDetails(args.KillmailID, args.KillmailHash)
}
