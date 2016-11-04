package domain

type WazeAddress struct {
	Bounds       struct {
					 Bottom float64 `json:"bottom"`
					 Left   float64 `json:"left"`
					 Right  float64 `json:"right"`
					 Top    float64 `json:"top"`
				 } `json:"bounds"`
	BusinessName string `json:"businessName"`
	City         string `json:"city"`
	Location     struct {
					 Lat float64 `json:"lat"`
					 Lon float64 `json:"lon"`
				 } `json:"location"`
	Name         string `json:"name"`
	Number       string `json:"number"`
	Provider     string `json:"provider"`
	SegmentId    int `json:"segmentId"`
	State        string `json:"state"`
	Street       string `json:"street"`
	StreetId     int `json:"streetId"`
}