package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

// EVEType holds all of the details of an EVE type
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

// EVEPosition contains a representation of a point in space
type EVEPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type PositionResolver struct {
	position *EVEPosition
}

// X returns the X position of this point
func (r *PositionResolver) X() *float64 {
	return &r.position.X
}

// Y returns the Y position of this point
func (r *PositionResolver) Y() *float64 {
	return &r.position.Y
}

// Z returns the Z position of this point
func (r *PositionResolver) Z() *float64 {
	return &r.position.Z
}

type EVETypeResolver struct {
	EveType *EVEType
}

// TypeID returns the type's ID
func (t *EVETypeResolver) TypeID() *int32 {
	return &t.EveType.TypeID
}

// Name returns the type's name
func (t *EVETypeResolver) Name() *string {
	return &t.EveType.Name
}

// Description returns the type's string description
func (t *EVETypeResolver) Description() *string {
	return &t.EveType.Description
}

// Published returns whether or not a type is published
func (t *EVETypeResolver) Published() *bool {
	return &t.EveType.Published
}

// GroupID returns the type's group ID
func (t *EVETypeResolver) GroupID() *int32 {
	return &t.EveType.GroupID
}

// Radius returns the type's radius
func (t *EVETypeResolver) Radius() *float64 {
	return &t.EveType.Radius
}

// Volume returns the type's volume
func (t *EVETypeResolver) Volume() *float64 {
	return &t.EveType.Volume
}

// Capacity returns the capacity of the type
func (t *EVETypeResolver) Capacity() *float64 {
	return &t.EveType.Capacity
}

// PortionSize returns the type's portion size
func (t *EVETypeResolver) PortionSize() *int32 {
	return &t.EveType.PortionSize
}

// Mass returns the type's mass
func (t *EVETypeResolver) Mass() *float64 {
	return &t.EveType.Mass
}

// GraphicID returns the graphic ID provided on a type
func (t *EVETypeResolver) GraphicID() *int32 {
	return &t.EveType.GraphicID
}

// DogmaAttributes returns an array represeting all of the attributes
// on a type
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

// GetEVEType fetches an EVEType with the provided ID
func GetEVEType(typeID int32) (*EVETypeResolver, error) {
	resp, err := http.MakeRequest("universe/types/%d/", typeID)
	if err != nil {
		return nil, err
	}

	var eveType EVEType
	if err := json.NewDecoder(resp.Body).Decode(&eveType); err != nil {
		return nil, err
	}

	return &EVETypeResolver{&eveType}, nil
}
