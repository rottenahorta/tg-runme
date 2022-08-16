package tg

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	Id  int    `json:"update_id"`
	Msg *MessageInc `json:"message"` // ptr cos if nil it parses right (tg api optional tag)
}

type MessageInc struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Uname string `json:"username"`
}

type Chat struct {
	Id int `json:"id"`
}
