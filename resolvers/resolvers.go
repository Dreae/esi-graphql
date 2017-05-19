package resolvers

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
