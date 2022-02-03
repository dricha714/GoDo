package resolvers

import (
	"github.com/CodingProjects/Go/GoDo/models"
)

type VegetableResolver struct {
	V *models.Vegetable
}

func (r *VegetableResolver) Name() string   { return r.V.Name }
func (r *VegetableResolver) Price() int32   { return int32(r.V.Price) }
func (r *VegetableResolver) Image() *string { return r.V.Image }
