package tg

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	er "github.com/rottenahorta/tgbotsche/pkg/int"
	repo "github.com/rottenahorta/tgbotsche/pkg/repo"
	zp "github.com/rottenahorta/tgbotsche/pkg/zepp"
)

type Client struct {
	client     http.Client
	host       string
	botPath    string
	listenPort string
	tghost     string
	repo 	   *repo.Repo
}

func NewClient(h, t, lp, r string) *Client {
	return &Client{
		client:     http.Client{},
		botPath:    makePath(t),
		host:       h,
		tghost:     "api.telegram.org",
		listenPort: lp,
		repo:		repo.NewRepo(r)}
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
	q.Set("chat_id", strconv.Itoa(chatId))
	q.Set("text", m)
	_, err := c.doRequest(c.tghost, c.botPath+"/sendMessage", "", "", "GET", q)
	if err != nil {
		return er.Log("cant send msg", err)
	}
	return nil
}

func (c *Client) GetZeppData(chatId int) (zp.Update, error) {
	zpToken, err := repo.GetZeppToken(chatId, c.repo.DBPostgres)
	if err != nil {
		return zp.Update{}, er.Log("cant retrieve zpToken from db", err)
	}
	var res zp.Update
	b, err := c.doRequest("api-mifit-de2.huami.com", "v1/sport/run/history.json", "apptoken", zpToken, "GET", nil)
	if err != nil {
		return zp.Update{}, er.Log("cant get zepp data", err)
	}
	//log.Print("GetZeppData: " + string(b))
	if err := json.Unmarshal(b, &res); err != nil {
		return zp.Update{}, er.Log("cant unmarshal zepp data", err)
	}
	if res.Data.Code == "0102" {
		c.Send(chatId, msgErrorToken+"\n"+msgUpdateToken+"\n"+authLinkZepp+"\n"+msgSupport)
		return res, er.Log("GetZeppData: ", errors.New("invalid zp token"))
	}
	log.Printf("zepp req: %v", res.Data)
	return res, nil
}

func (c *Client) GetZeppLastDistByToken(zpToken string) (string, error) {
	var res zp.Update
	b, err := c.doRequest("api-mifit-de2.huami.com", "v1/sport/run/history.json", "apptoken", zpToken, "GET", nil)
	if err != nil {
		return "", er.Log("cant get zepp data", err)
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return "", er.Log("cant unmarshal zepp data", err)
	}
	if res.Data.Code == "0102" {
		return "", er.Log("GetZeppLastDistByToken: ", errors.New("invalid zp token"))
	}
	return res.Data.Summary[0].Distance, nil
}

func (c *Client) GetZeppTokenFromUser(code string, chatId int) (error) {
	var res zp.ResponseToken
	q := url.Values{}
	q.Set("code",code)
	q.Set("grant_type","request_token")
	q.Set("country_code","RU")
	q.Set("device_id","w")
	q.Set("third_name","xiaomi-hm-mifit")
	q.Set("app_version","w")
	q.Set("device_model","w")
	q.Set("app_name","com.xiaomi.hm.health")
	b, err := c.doRequest("account.huami.com", "v2/client/login", "", "", "POST", q)
	if err != nil {
		return er.Log("cant get zepp apptoken", err)
	}
	log.Print("getzepptokenfromuser: " + string(b))
	if err := json.Unmarshal(b, &res); err != nil {
		return er.Log("cant unmarshal zepp apptoken", err)
	}

	var id int
	row := c.repo.DBPostgres.QueryRow("INSERT INTO users (chatid,zptoken) values ($1,$2) ON CONFLICT (chatid) DO UPDATE SET zptoken = $2 RETURNING id", chatId, res.TokenInfo.AppToken)
	if err := row.Scan(&id); err != nil {
		return er.Log("cant retrieve zptoken", err)
	}
	log.Printf("new zptoken retrieved w id: %d", id)
	return nil
}

func (c *Client) doRequest(host, path, headerName, headerValue, method string, q url.Values) (d []byte, err error) {
	defer func() { err = er.Log("cant do req", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	req, err := http.NewRequest(method, u.String(), strings.NewReader(q.Encode())) 
	if err != nil {
		return nil, err
	}
	if headerName != "" {
		req.Header.Set(headerName, headerValue)
	}
	req.Header.Set("Content-Type","application/json")
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
