package resolvers

import (
	"context"

	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

// StructureNode is a wrapper type to support fetching
// details from a list of structures
type StructureNode float64

// Structure contains the details of a strucutre in EVE
type Structure struct {
	Name          string      `json:"name"`
	SolarSystemID int32       `json:"solar_system_id"`
	TypeID        int32       `json:"type_id"`
	Position      EVEPosition `json:"position"`
}

type StructureResolver struct {
	structure *Structure
}

// Name returns the name of the structure
func (r *StructureResolver) Name() *string {
	return &r.structure.Name
}

// SolarSystemID returns the ID of the solar system where the
// structure is located
func (r *StructureResolver) SolarSystemID() *int32 {
	return &r.structure.SolarSystemID
}

// TypeID returns the ID of the structure's type
func (r *StructureResolver) TypeID() *int32 {
	return &r.structure.TypeID
}

// Position returns the structure's position in space
func (r *StructureResolver) Position() *PositionResolver {
	return &PositionResolver{&r.structure.Position}
}

// StructureID returns the structure's ID
func (n *StructureNode) StructureID() *float64 {
	return (*float64)(n)
}

// Structure serves as a pointer to get the details of a structure.
// This endpoint requires an authorization token
func (n *StructureNode) Structure(ctx context.Context) (*StructureResolver, error) {
	return GetStructureByID(*(*float64)(n), ctx.Value(ContextKey("auth")).(string))
}

// GetStructureByID returns the structure represented by the provided ID.
// This endpoint requires an authorization token
func GetStructureByID(id float64, auth string) (*StructureResolver, error) {
	resp, err := http.MakeAuthorizedRequest(auth, "universe/structures/%.00f/", id)
	if err != nil {
		return nil, err
	}

	var structure Structure
	if err := json.NewDecoder(resp.Body).Decode(&structure); err != nil {
		return nil, err
	}

	return &StructureResolver{&structure}, nil
}

// GetAllStructures returns all structures reported by ESI
func GetAllStructures() (*[]*StructureNode, error) {
	resp, err := http.MakeRequest("universe/structures/")
	if err != nil {
		return nil, err
	}

	var ids []StructureNode
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	var structures []*StructureNode
	for idx := range ids {
		structures = append(structures, &ids[idx])
	}

	return &structures, nil
}
