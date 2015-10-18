package api

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/endophage/godota2"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL           = "https://api.steampowered.com/IDOTA2Match_570"
	keyQueryParam     = "key"
	startIDQueryParam = "start_at_match_id"
	matchIDQueryParam = "match_id"
)

var (
	matchHistoryURL = baseURL + "/GetMatchHistory/V001/"
	matchDetailURL  = baseURL + "/GetMatchDetails/V001/"
)

type Result struct {
	Result json.RawMessage `json:"result"`
}

type dota2Api struct {
	key string
}

func NewDota2Api(key string) dota2Api {
	return dota2Api{
		key: key,
	}
}

func (d dota2Api) Matches(from int) (*dota2.Matches, error) {
	u, err := url.Parse(matchHistoryURL)
	if err != nil {
		return nil, err
	}
	p := url.Values{}
	if from != 0 {
		p.Add(startIDQueryParam, fmt.Sprint("%d", from))
	}
	p.Add(keyQueryParam, d.key)
	u.RawQuery = p.Encode()
	logrus.Debug(u.String())
	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	w := &Result{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, w)
	if err != nil {
		return nil, err
	}
	matches := &dota2.Matches{}
	err = json.Unmarshal(w.Result, matches)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (d dota2Api) MatchDetail(matchID int) (*dota2.MatchDetail, error) {
	u, err := url.Parse(matchDetailURL)
	if err != nil {
		return nil, err
	}
	p := url.Values{}
	p.Add(keyQueryParam, d.key)
	u.RawQuery = p.Encode()
	logrus.Debug(u.String())
	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	w := &Result{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, w)
	if err != nil {
		return nil, err
	}
	detail := &dota2.MatchDetail{}
	err = json.Unmarshal(w.Result, detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}
