package route

import (
	"github.com/prsolucoes/gowaze/domain"
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

const (
	WAZE_ROW_ROUTING_MANAGER_SERVER = "row-RoutingManager"
	WAZE_ROUTING_MANAGER_SERVER = "RoutingManager"
	WAZE_IL_ROUTING_MANAGER_SERVER = "il-RoutingManager"

	ROUTING_OPTION_AVOID_TRAILS = "AVOID_TRAILS%3At"
	ROUTING_OPTION_ALLOW_UTURNS = "ALLOW_UTURNS%3At"
)

func GetDirections(server string, startLat string, startLong string, endLat string, endLong string, options []string) (*domain.WazeRouteDirections, error) {
	// https://www.waze.com/row-RoutingManager/routingRequest?
	// from=x:-43.438938 y:-23.0065936&
	// to=x:-43.06376460000001 y:-22.8272883&
	// at=0&
	// returnJSON=true&
	// returnGeometries=true&
	// returnInstructions=true&
	// timeout=60000&
	// nPaths=3&
	// clientVersion=4.0.0&
	// options=AVOID_TRAILS:t,ALLOW_UTURNS:t

	optionsParam := "options="

	if len(options) > 0 {
		for index, option := range options {
			if index == 0 {
				optionsParam += option
			} else {
				optionsParam += "%2C" + option
			}
		}
	}

	endpoint := "https://www.waze.com/" + server + "/routingRequest?from=x%3A" + startLong + "+y%3A" + startLat + "&to=x%3A" + endLong + "+y%3A" + endLat + "&at=0&returnJSON=true&returnGeometries=true&returnInstructions=true&timeout=60000&nPaths=3&" + optionsParam

	request := gorequest.New()
	_, body, errs := request.Get(endpoint).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	var routeDirections *domain.WazeRouteDirections
	err := json.Unmarshal(body, &routeDirections)

	if err != nil {
		return nil, err
	}

	return routeDirections, nil
}
