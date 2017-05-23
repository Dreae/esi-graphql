package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/cache"
	"github.com/dreae/esi-graphql/resolvers/http"
)

type EVEType struct {
	TypeID          int32                `json:"type_id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Published       bool                 `json:"published"`
	GroupID         int32                `json:"group_id"`
	Radius          float64              `json:"radius"`
	Volume          float64              `json:"volume"`
	Capacity        float64              `json:"capacity"`
	PortionSize     int32                `json:"portion_size"`
	Mass            float64              `json:"mass"`
	GraphicID       int32                `json:"graphic_id"`
	DogmaAttributes []DogmaAttributeNode `json:"dogma_attributes"`
}

type EVETypeResolver struct {
	EveType *EVEType
}

func (t *EVETypeResolver) TypeID() *int32 {
	return &t.EveType.TypeID
}

func (t *EVETypeResolver) Name() *string {
	return &t.EveType.Name
}

func (t *EVETypeResolver) Description() *string {
	return &t.EveType.Description
}

func (t *EVETypeResolver) Published() *bool {
	return &t.EveType.Published
}

func (t *EVETypeResolver) GroupID() *int32 {
	return &t.EveType.GroupID
}

func (t *EVETypeResolver) Radius() *float64 {
	return &t.EveType.Radius
}

func (t *EVETypeResolver) Volume() *float64 {
	return &t.EveType.Volume
}

func (t *EVETypeResolver) Capacity() *float64 {
	return &t.EveType.Capacity
}

func (t *EVETypeResolver) PortionSize() *int32 {
	return &t.EveType.PortionSize
}

func (t *EVETypeResolver) Mass() *float64 {
	return &t.EveType.Mass
}

func (t *EVETypeResolver) GraphicID() *int32 {
	return &t.EveType.GraphicID
}

func (t *EVETypeResolver) DogmaAttributes() *[]*DogmaAttributeNodeResolver {
	var nodes []*DogmaAttributeNodeResolver
	for _, attributeNode := range t.EveType.DogmaAttributes {
		nodes = append(nodes, &DogmaAttributeNodeResolver{&DogmaAttributeNode{
			attributeNode.AttributeID,
			attributeNode.Value,
		}})
	}

	return &nodes
}

var typeCache = cache.New(3600)

func GetEVEType(typeID int32) (*EVETypeResolver, error) {
	if item, ok := typeCache.Get(typeID); ok {
		eveType := item.(EVEType)
		return &EVETypeResolver{&eveType}, nil
	}

	var type_ EVEType
	resp, err := http.MakeRequest("universe/types/%d/", typeID)
	if err != nil {
		return &EVETypeResolver{&type_}, err
	}

	json.NewDecoder(resp.Body).Decode(&type_)

	typeCache.Set(typeID, type_)

	return &EVETypeResolver{&type_}, nil
}
