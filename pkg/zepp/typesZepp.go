package zp

type Update struct {
	Data Data `json:"data"`  
}

type Data struct {
	Summary []Summary `json:"summary"` 
	Code string `json:"code"`
}

type Summary struct {
	Distance string `json:"dis"` 
	Runtime string `json:"run_time"` 
	AvgPace string `json:"avg_pace"` 
}

type ResponseToken struct {
	TokenInfo TokenInfo `json:"token_info"` 
}

type TokenInfo struct {
	AppToken string `json:"app_token"` 
}