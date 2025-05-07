package requests

type CalculateAreaRequest struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}
