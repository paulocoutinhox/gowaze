package address

import (
	"github.com/parnurzeal/gorequest"
	"github.com/prsolucoes/gowaze/domain"
	"encoding/json"
	"errors"
)

func GetList(address string) ([]*domain.WazeAddress, error) {
	endpoint := "https://www.waze.com/SearchServer/mozi?q=" + address + "&lang=eng&lon=0&lat=0&origin=livemap"

	request := gorequest.New()
	_, body, errs := request.Get(endpoint).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	var addressList []*domain.WazeAddress
	err := json.Unmarshal(body, &addressList)

	if err != nil {
		return nil, err
	}

	return addressList, nil
}

func GetSingle(address string) (*domain.WazeAddress, error) {
	addressList, err := GetList(address)

	if err != nil {
		return nil, err
	}

	if len(addressList) == 0 {
		return nil, errors.New("Address not found")
	}

	return addressList[0], nil
}
