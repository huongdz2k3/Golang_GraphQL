package directives

import (
	"context"
	"customer/internal/utils"
	"customer/middleware"
	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	if tokenData == nil {
		return nil, utils.WrapGQLUnauthorizedError(ctx)
	}

	return next(ctx)
}
