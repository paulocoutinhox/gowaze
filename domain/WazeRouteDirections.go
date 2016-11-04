package domain

type WazeRouteDirections struct {
	Alternatives []struct {
		Response  struct {
					  Results                 []struct {
						  Path                     struct {
													   SegmentID int `json:"segmentId"`
													   NodeID    int `json:"nodeId"`
													   X         float64 `json:"x"`
													   Y         float64 `json:"y"`
												   } `json:"path"`
						  Street                   int `json:"street"`
						  AltStreets               interface{} `json:"altStreets"`
						  Distance                 int `json:"distance"`
						  Length                   int `json:"length"`
						  CrossTime                int `json:"crossTime"`
						  CrossTimeWithoutRealTime int `json:"crossTimeWithoutRealTime"`
						  Tiles                    interface{} `json:"tiles"`
						  ClientIds                interface{} `json:"clientIds"`
						  Instruction              struct {
													   Opcode          string `json:"opcode"`
													   Arg             int `json:"arg"`
													   InstructionText interface{} `json:"instructionText"`
													   Name            interface{} `json:"name"`
													   Tts             interface{} `json:"tts"`
												   } `json:"instruction"`
						  KnownDirection           bool `json:"knownDirection"`
						  Penalty                  int `json:"penalty"`
						  RoadType                 int `json:"roadType"`
						  AdditionalInstruction    interface{} `json:"additionalInstruction"`
						  IsToll                   bool `json:"isToll"`
						  NaiveRoute               interface{} `json:"naiveRoute"`
						  DetourSavings            int `json:"detourSavings"`
						  DetourSavingsNoRT        int `json:"detourSavingsNoRT"`
						  UseHovLane               bool `json:"useHovLane"`
					  } `json:"results"`
					  StreetNames             []interface{} `json:"streetNames"`
					  TileIds                 []interface{} `json:"tileIds"`
					  TileUpdateTimes         []interface{} `json:"tileUpdateTimes"`
					  Geom                    interface{} `json:"geom"`
					  FromFraction            float64 `json:"fromFraction"`
					  ToFraction              float64 `json:"toFraction"`
					  SameFromSegment         bool `json:"sameFromSegment"`
					  SameToSegment           bool `json:"sameToSegment"`
					  AstarPoints             interface{} `json:"astarPoints"`
					  WayPointIndexes         interface{} `json:"wayPointIndexes"`
					  WayPointFractions       interface{} `json:"wayPointFractions"`
					  TollMeters              int `json:"tollMeters"`
					  PreferedRouteID         int `json:"preferedRouteId"`
					  IsInvalid               bool `json:"isInvalid"`
					  IsBlocked               bool `json:"isBlocked"`
					  ServerUniqueID          string `json:"serverUniqueId"`
					  DisplayRoute            bool `json:"displayRoute"`
					  AstarVisited            int `json:"astarVisited"`
					  AstarResult             string `json:"astarResult"`
					  AstarData               interface{} `json:"astarData"`
					  IsRestricted            bool `json:"isRestricted"`
					  AvoidStatus             string `json:"avoidStatus"`
					  DueToOverride           interface{} `json:"dueToOverride"`
					  PassesThroughDangerArea bool `json:"passesThroughDangerArea"`
					  DistanceFromSource      int `json:"distanceFromSource"`
					  DistanceFromTarget      int `json:"distanceFromTarget"`
					  RouteType               []string `json:"routeType"`
					  RouteAttr               []string `json:"routeAttr"`
					  AstarCost               int `json:"astarCost"`
					  SegGeoms                interface{} `json:"segGeoms"`
					  RouteName               string `json:"routeName"`
					  RouteNameStreetIds      []int `json:"routeNameStreetIds"`
				  } `json:"response"`
		Coords    []struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z string `json:"z"`
		} `json:"coords"`
		SegCoords interface{} `json:"segCoords"`
	} `json:"alternatives"`
}