package service

import (
	"context"
	"customer/ent"
	"customer/ent/role"
	"customer/internal/utils"
)

func CreateRole(ctx context.Context, input ent.CreateRoleInput, client ent.Client) (*ent.Role, error) {
	return client.Role.Create().SetInput(input).Save(ctx)
}

func GetRole(ctx context.Context, input int, client ent.Client) (*ent.Role, error) {
	role, err := client.Role.Query().Where(role.ID(input)).Only(ctx)
	if err != nil {
		return nil, utils.WrapGQLNotFoundError(ctx)
	}
	return role, nil
}
