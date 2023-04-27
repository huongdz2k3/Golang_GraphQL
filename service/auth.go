package service

import (
	"context"
	"customer/ent"
	"customer/internal/utils"
	"customer/tool"
)

func Register(ctx context.Context, input ent.CreateCustomerInput, client ent.Client) (*ent.Jwt, error) {
	// check email exist

	_, err := GetCustomerByEmail(ctx, input.Email, client)
	if err == nil {
		return nil, utils.WrapGQLBadRequestError(ctx, "This email already exists")
	}
	input.Password = tool.HashPassword(input.Password)
	user, err := CreateCustomer(ctx, input, client)

	if err != nil {
		panic(err)
	}

	token, err := JwtGenerate(ctx, user.ID)

	if err != nil {
		return nil, utils.WrapGQLBadRequestError(ctx, "Failed to generate JWT")
	}
	var t ent.Jwt = ent.Jwt{
		token,
	}
	return &t, nil
}

func Login(ctx context.Context, input ent.LoginInput, client ent.Client) (*ent.Jwt, error) {
	cus, err := GetCustomerByEmail(ctx, input.Email, client)

	if err := tool.ComparePassword(cus.Password, input.Password); err != nil {
		return nil, err
	}
	token, err := JwtGenerate(ctx, cus.ID)
	if err != nil {
		panic(err)
	}
	return &ent.Jwt{
		token,
	}, nil
}
