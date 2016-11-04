package route

import (
	"testing"
)

func TestGetDirections(t *testing.T) {
	options := []string{
		ROUTING_OPTION_AVOID_TRAILS,
		ROUTING_OPTION_ALLOW_UTURNS,
	}

	routeDirections, err := GetDirections(WAZE_ROW_ROUTING_MANAGER_SERVER, "-23.0065936", "-43.438938", "-22.8272883", "-43.06376460000001", options)

	if err != nil {
		t.Errorf("Error while get route directions: %v", err)
	}

	if routeDirections == nil {
		t.Error("Invalid params for route directions")
	}

	if len(routeDirections.Alternatives) != 3 {
		t.Errorf("Invalid route directions alternatives: %d", len(routeDirections.Alternatives))
	}

	if routeDirections.Alternatives[0].Response.RouteName != "Linha Amarela, BR-101 Ponte Pres. Costa e Silva Rio de Janeiro" {
		t.Errorf("Invalid route directions alternative 1: %v", routeDirections.Alternatives[0].Response.RouteName)
	}

	if routeDirections.Alternatives[1].Response.RouteName != "Av. Brasil Rio de Janeiro; BR-040 Rod. Washington Luiz Duque de Caxias" {
		t.Errorf("Invalid route directions alternative 2: %v", routeDirections.Alternatives[1].Response.RouteName)
	}

	if routeDirections.Alternatives[2].Response.RouteName != "Linha Amarela Rio de Janeiro; BR-040 Rod. Washington Luiz Duque de Caxias; BR-493 Est. do Contorno" {
		t.Errorf("Invalid route directions alternative 3: %v", routeDirections.Alternatives[2].Response.RouteName)
	}
}