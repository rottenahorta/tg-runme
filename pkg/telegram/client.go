package tg

import (
	"encoding/json"
	"io"
	"log"
	"net"

	//"net"
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

func (c *Client) Update() (chan *Update, error) {
	/*q := url.Values{} // addin params
	q.Add("offset", strconv.Itoa(o))
	q.Add("limit", strconv.Itoa(l))
	d, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, err
	}*/
	/*
	d, err := c.doRequest("", nil)
	if err != nil {
		return nil, err
	}*/

	updates := make(chan *Update)

	//var res Update
	handler := func(w http.ResponseWriter, r *http.Request) {

		res := &Update{}

		defer func() { _ = r.Body.Close() }()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf(string(body))
		if err := json.Unmarshal(body, &res); err != nil {
			log.Fatal(err)
			return
		}

		updates <- res
		
		log.Printf(res.Msg.Text)
		
	}
	
	l, err := net.Listen("tcp", c.listenPort)
	if err != nil {
		return nil, err
	}
	go http.Serve(l, http.HandlerFunc(handler))

	/*go http.HandleFunc("/"+c.path, handler)
	if err := http.ListenAndServe(c.listenPort, nil); err != nil {
		log.Fatal(err)
	}*/
	log.Print("debuggin Update() after Serve()")
	return updates, nil

	/*var res UpdateResponse
	if err := json.Unmarshal(d, &res); err != nil {
		return nil, err
	}
	return res.Result, nil*/
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
