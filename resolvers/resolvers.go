package resolvers

import (
	"github.com/dreae/esi-graphql/resolvers/dogma"
	"github.com/dreae/esi-graphql/resolvers/universe"
)

type Resolver struct{}

func (r *Resolver) DogmaAttribute(args *struct{ AttributeID int32 }) (*dogma.DogmaAttributeResolver, error) {
	return dogma.GetDogmaAttributeResolver(args.AttributeID)
}

func (r *Resolver) DogmaAttributes() (*[]*dogma.DogmaAttributeNodeResolver, error) {
	return dogma.GetDogmaList()
}

func (r *Resolver) Type(args *struct{ TypeID int32 }) (*universe.EVETypeResolver, error) {
	return universe.GetEVEType(args.TypeID)
}
