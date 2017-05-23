package resolvers

import (
	"context"

	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

type StructureNode struct {
	id float64
}

type Structure struct {
	Name          string      `json:"name"`
	SolarSystemID int32       `json:"solar_system_id"`
	TypeID        int32       `json:"type_id"`
	Position      EVEPosition `json:"position"`
}

type StructureResolver struct {
	structure *Structure
}

func (r *StructureResolver) Name() *string {
	return &r.structure.Name
}

func (r *StructureResolver) SolarSystemID() *int32 {
	return &r.structure.SolarSystemID
}

func (r *StructureResolver) TypeID() *int32 {
	return &r.structure.TypeID
}

func (r *StructureResolver) Position() *PositionResolver {
	return &PositionResolver{&r.structure.Position}
}

func (n *StructureNode) StructureID() *float64 {
	return &n.id
}

func (n *StructureNode) Structure(ctx context.Context) (*StructureResolver, error) {
	return GetStructureByID(n.id, ctx.Value(ContextKey("auth")).(string))
}

func GetStructureByID(id float64, auth string) (*StructureResolver, error) {
	resp, err := http.MakeAuthorizedRequest(auth, "universe/structures/%.00f/", id)
	if err != nil {
		return nil, err
	}

	var structure Structure
	json.NewDecoder(resp.Body).Decode(&structure)

	return &StructureResolver{&structure}, nil
}

func GetAllStructures() (*[]*StructureNode, error) {
	resp, err := http.MakeRequest("universe/structures/")
	if err != nil {
		return nil, err
	}

	var ids []float64
	var structures []*StructureNode
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	for _, id := range ids {
		structures = append(structures, &StructureNode{id})
	}

	return &structures, nil
}
