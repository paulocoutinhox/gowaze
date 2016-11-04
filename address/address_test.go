package address

import (
	"testing"
	"fmt"
)

func TestGetSingleAddress(t *testing.T) {
	address, err := GetSingle("São Gonçalo")

	if err != nil {
		t.Errorf("Error while get single address: %v", err)
	}

	if address == nil {
		t.Error("Address is invalid")
	}

	if address.Name != "São Gonçalo, RJ" {
		t.Errorf("Address name is wrong: %v", address.Name)
	}
}

func TestGetAddressList(t *testing.T) {
	addressList, err := GetList("São Gonçalo")

	fmt.Println(addressList)

	if err != nil {
		t.Errorf("Error while get address list: %v", err)
	}

	if addressList == nil {
		t.Error("Empty address list")
	}

	if len(addressList) != 4 {
		t.Errorf("Address list with wrong size: %d", len(addressList))
	}
}