package tg

/*type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result Update `json:"result"`
}*/

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

var myChatId int = 450892706
var authLinkZepp string = "http://account.xiaomi.com/oauth2/authorize?skip_confirm=false&client_id=428135909242707968&pt=1&scope=1+6000+16001+16002+20000&redirect_uri=https%3A%2F%2Fapi-mifit-cn.huami.com%2Fhuami.health.loginview.do&_locale=en_US&response_type=code&userId=7034468780&nonce=bkGvbHWsdKwBnE7Q&confirmed=true&from_login=true&sign=vi0n5XQ0mQgIvyXvXXilf4HXQV8%3D"
var awaitSupportMsg bool