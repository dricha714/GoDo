package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/CodingProjects/Go/GoDo/common"
	"github.com/CodingProjects/Go/GoDo/inputs"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
)

type Query struct{}

var vegetables = map[string]models.Vegetable{
	"tomato": {Name: "Tomato", Price: 100, Image: common.StrPtr("https://picsum.photos/id/152/100/100")},
	"potato": {Name: "Potato", Price: 50, Image: common.StrPtr("https://picsum.photos/id/159/100/100")},
	"corn":   {Name: "Corn", Price: 200},
}

func (q *Query) Vegetable(ctx context.Context, args struct{ Name string }) *resolvers.VegetableResolver {
	v, ok := vegetables[strings.ToLower(args.Name)]
	if ok {
		return &resolvers.VegetableResolver{V: &v}
	}
	return nil
}

func (q *Query) Vegetables() *[]*resolvers.VegetableResolver {
	var values []*resolvers.VegetableResolver
	for _, value := range vegetables {
		fmt.Println(value)
		tmp := value
		values = append(values, &resolvers.VegetableResolver{&tmp})
	}
	return &values
}

func (m *Query) CreateVegetable(args struct{ Vegetable *inputs.VegetableInput }) *resolvers.VegetableResolver {
	value := &models.Vegetable{
		Name:  args.Vegetable.Name,
		Price: int32(args.Vegetable.Price),
		Image: args.Vegetable.Image,
	}
	if vegetables == nil {
		vegetables[""] = models.Vegetable{}
	}
	vegetables[value.Name] = *value

	return &resolvers.VegetableResolver{value}
}
