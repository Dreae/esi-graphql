package resolvers

import (
	"context"
)

type ContextKey string

type Resolver struct{}

func (r *Resolver) DogmaAttribute(args *struct{ AttributeID int32 }) (*DogmaAttributeResolver, error) {
	return GetDogmaAttributeResolver(args.AttributeID)
}

func (r *Resolver) DogmaAttributes() (*[]*DogmaAttributeNodeResolver, error) {
	return GetDogmaList()
}

func (r *Resolver) Type(args *struct{ TypeID int32 }) (*EVETypeResolver, error) {
	return GetEVEType(args.TypeID)
}

func (r *Resolver) Character(args *struct{ CharacterID int32 }) (*CharacterResolver, error) {
	return GetCharacterByID(args.CharacterID)
}

func (r *Resolver) Corporation(args *struct{ CorporationID int32 }) (*CorporationResolver, error) {
	return GetCorpByID(args.CorporationID)
}

func (r *Resolver) Skills(ctx context.Context, args *struct{ CharacterID int32 }) (*CharacterSkillsResolver, error) {
	return GetSkillsForCharID(ctx.Value(ContextKey("auth")).(string), args.CharacterID)
}

func (r *Resolver) SkillQueue(ctx context.Context, args *struct{ CharacterID int32 }) (*[]*SkillQueueResolver, error) {
	return GetSkillQueueForCharID(ctx.Value(ContextKey("auth")).(string), args.CharacterID)
}

func (r *Resolver) Alliance(args *struct{ AllianceID int32 }) (*AllianceResolver, error) {
	return GetAllianceByID(args.AllianceID)
}

func (r *Resolver) Search(args *struct {
	SearchTypes []*string
	Keyword     string
}) (*SearchResultsResolver, error) {
	return DoSearch(&args.SearchTypes, args.Keyword)
}
