package api

import (
	"context"

	"github.com/vektah/gqlparser/gqlerror"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Test1(ctx context.Context) (*Test1, error) {
	if !hasRole(ctx.Value("roles").([]string), "test_role_1") {
		return nil, gqlerror.Errorf("Not authorized")
	}

	return &Test1{ID: "hoge"}, nil
}

func (r *queryResolver) Test2(ctx context.Context) (*Test2, error) {
	if !hasRole(ctx.Value("roles").([]string), "test_role_10") {
		return nil, gqlerror.Errorf("Not authorized")
	}
	return &Test2{ID: "fuga"}, nil
}

func hasRole(userRole []string, role string) bool {
	for _, value := range userRole {
		if value == role {
			return true
		}
	}

	return false
}
