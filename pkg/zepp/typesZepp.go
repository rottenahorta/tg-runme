package zp

type Update struct {
	Data Data `json:"data"` 
}

type Data struct {
	Summary []Summary `json:"summary"` 
}

type Summary struct {
	Distance int `json:"dis"` 
}