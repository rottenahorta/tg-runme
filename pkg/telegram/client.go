package tg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	er "github.com/rottenahorta/tgbotsche/pkg/int"
	zp "github.com/rottenahorta/tgbotsche/pkg/zepp"
)

type Client struct {
	client http.Client
	host   string
	path   string
	listenPort string
	tghost string
}

func NewClient(h, t, lp string) *Client {
	return &Client{
		client: http.Client{},
		path:   makePath(t),
		host:   h,
		tghost: "api.telegram.org",
		listenPort: lp}
}

func (c *Client) Update() (){
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
	_, err := c.doRequest("sendMessage", c.tghost, "", "", q)
	if err != nil {
		return er.Log("cant send msg", err)
	}
	return nil
}

func (c *Client) GetZeppData() (zp.Update, error) {
	var res zp.Update
	b, err := c.doRequest("", "api-mifit-de2.huami.com", "apptoken", os.Getenv("ZPTOKEN"), nil)
	if err != nil {
		return  zp.Update{}, er.Log("cant get zepp data", err)
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return  zp.Update{}, er.Log("cant unmarshal zepp data", err)
	}
	log.Printf("zepp req summary: %v", res.Data.Summary)
	return res, nil
}

func (c *Client) doRequest(method, host, headerName, headerValue string, q url.Values) (d []byte, err error) {
	defer func() { err = er.Log("cant do req", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   func() string {
			if headerName == "" {
				return path.Join(c.path, method) 
			} else {
				return "v1/sport/run/history.json"
			}}(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if headerName != ""{
		req.Header.Set(headerName, headerValue)
	} 
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
