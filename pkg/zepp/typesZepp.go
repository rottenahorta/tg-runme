package zp

type Update struct {
	Data Data `json:"data"` 
}

type Data struct {
	Summary []Summary `json:"summary"` 
}

type Summary struct {
	Distance string `json:"dis"` 
	Runtime string `json:"run_time"` 
	AvgPace string `json:"avg_pace"` 
}