package tg

import (
	//"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	//"path"
	"strconv"

	er "github.com/rottenahorta/tgbotsche/pkg/int"
	zp "github.com/rottenahorta/tgbotsche/pkg/zepp"
)

type Client struct {
	client     http.Client
	host       string
	botPath    string
	listenPort string
	tghost     string
}

func NewClient(h, t, lp string) *Client {
	return &Client{
		client:     http.Client{},
		botPath:    makePath(t),
		host:       h,
		tghost:     "api.telegram.org",
		listenPort: lp}
}

func (c *Client) Update() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var res Update

		defer func() { _ = r.Body.Close() }()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := json.Unmarshal(body, &res); err != nil {
			log.Fatal(err)
			return
		}
		c.Fetch(res)
	}

	go http.ListenAndServe(c.listenPort, http.HandlerFunc(handler))
}

func (c *Client) Send(chatId int, m string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", m)
	_, err := c.doRequest(c.tghost, c.botPath+"/sendMessage", "", "", "GET", q)
	if err != nil {
		return er.Log("cant send msg", err)
	}
	return nil
}

func (c *Client) GetZeppData() (zp.Update, error) {
	var res zp.Update
	b, err := c.doRequest("api-mifit-de2.huami.com", "v1/sport/run/history.json", "apptoken", os.Getenv("HUAMITOKEN"), "GET", nil)
	if err != nil {
		return zp.Update{}, er.Log("cant get zepp data", err)
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return zp.Update{}, er.Log("cant unmarshal zepp data", err)
	}
	log.Printf("zepp req summary: %v", res.Data.Summary)
	return res, nil
}

func (c *Client) GetZeppToken(code string) (string, error) {
	var res zp.ResponseToken
	q := url.Values{}
	q.Add("code",code)
	q.Add("grant_type","request_token")
	q.Add("country_code","RU")
	q.Add("device_id","w")
	q.Add("third_name","xiaomi-hm-mifit")
	q.Add("app_version","w")
	q.Add("device_model","w")
	q.Add("app_name","com.xiaomi.hm.health")
	b, err := c.doRequest("account.huami.com", "v2/client/login", "", "", "POST", q)
	if err != nil {
		return "", er.Log("cant get zepp apptoken", err)
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return "", er.Log("cant unmarshal zepp apptoken", err)
	}
	log.Printf("zepp apptoken: %s", res.TokenInfo.AppToken)
	return res.TokenInfo.AppToken, nil
}

func (c *Client) doRequest(host, path, headerName, headerValue, method string, q url.Values) (d []byte, err error) {
	defer func() { err = er.Log("cant do req", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	req, err := http.NewRequest(method, u.String(), strings.NewReader(q.Encode())) // todo : do i need body?
	if err != nil {
		return nil, err
	}
	if headerName != "" {
		req.Header.Set(headerName, headerValue)
	}
	req.Header.Set("Content-Type","application/x-www-form-urlencoded") // todo : do i need it?
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func makePath(t string) string {
	return "bot" + t
}
