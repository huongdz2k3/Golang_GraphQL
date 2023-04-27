package service

import (
	"context"
	"customer/ent"
	"customer/ent/customer"
	"customer/internal/utils"
	"log"
	"strconv"
)

func GetCustomerById(ctx context.Context, input string, client ent.Client) (*ent.Customer, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("\"Failed to convert string to int:\", err")
		return nil, err
	}
	return client.Customer.Query().Where(customer.ID(id)).Only(ctx)
}

func GetCustomerByEmail(ctx context.Context, email string, client ent.Client) (*ent.Customer, error) {
	u, err := client.Customer.Query().Where(customer.Email(email)).Only(ctx)
	if err != nil {
		return nil, utils.WrapGQLNotFoundError(ctx)
	}
	return u, nil
}

func CreateCustomer(ctx context.Context, input ent.CreateCustomerInput, client ent.Client) (*ent.Customer, error) {
	return client.Customer.Create().SetInput(input).Save(ctx)
}
