package main

import (
	"fmt"
)

func main() {
	userID := "123"
	user, address, err := searchUserAndAddressByUserID(userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Address: %+v\n", address)
}

type (
	User struct {
		ID     string
		Name   string
		Status string
	}
	Address struct {
		Country       string
		City          string
		StreetAddress string
	}
)

func searchUserAndAddressByUserID(userID string) (*User, *Address, error) {
	user, err := getUser(userID)
	if err != nil {
		return nil, nil, err
	}
	if user.Status != "active" {
		// This should be return nil, nil, ErrNotFound.
		return nil, &Address{
			Country:       "Japan",
			City:          "Tokyo",
			StreetAddress: "1-1",
		}, nil
	}

	address, err := getAddress()
	if err != nil {
		return nil, nil, err
	}

	return user, address, nil
}

func getUser(userID string) (*User, error) {
	return &User{
		ID:     userID,
		Name:   "john",
		Status: "active",
	}, nil
}

func getAddress() (*Address, error) {
	return &Address{
		Country:       "Japan",
		City:          "Tokyo",
		StreetAddress: "1-1",
	}, nil
}
