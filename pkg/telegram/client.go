package tg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	er "github.com/rottenahorta/tgbotsche/pkg/int"
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

	/*q := url.Values{}
	q.Add("url", c.host+"/"+c.path)
	_, err := c.doRequest("setWebhook", q)
	if err != nil {
		er.Log("cant send msg", err)
	}*/

	handler := func(w http.ResponseWriter, r *http.Request) {
		var res Update

		defer func() { _ = r.Body.Close() }()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("method of handler WH: %s", r.Method)
		if err := json.Unmarshal(body, &res); err != nil {
			log.Fatal(err)
			return
		}

		c.Fetch(res)
		
		log.Printf("inside handler: " + res.Msg.Text)
		
	}

	go http.ListenAndServe(c.listenPort, http.HandlerFunc(handler))
}

func (c *Client) Send(chatId int, m string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", m)
	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		return er.Log("cant send msg", err)
	}
	return nil
}

func (c *Client) doRequest(method string, q url.Values) (d []byte, err error) {
	defer func() { err = er.Log("cant do req", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   c.tghost,
		Path:   path.Join(c.path, method),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
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
