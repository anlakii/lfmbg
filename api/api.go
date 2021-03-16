package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lakiluki1/lfmbg/api/types"
	"github.com/lakiluki1/lfmbg/config"
)

const apiUrl = "https://ws.audioscrobbler.com/2.0/?"

func GetRecentTracks(limit uint) (*types.RecentTracksType, error) {
	config, err := config.GetConfig()

	if err != nil {
		return nil, err
	}

	values := url.Values{
		"method":  []string{"user.getRecentTracks"},
		"user":    []string{config.Username},
		"api_key": []string{config.ApiKey},
		"format":  []string{"json"},
		"limit":   []string{fmt.Sprint(limit)},
	}

	resp, err := http.Get(apiUrl + values.Encode())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("Request was not OK: " + resp.Status)
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var tracks types.RecentTracksType

	err = dec.Decode(&tracks)
	if err != nil {
		return nil, err
	}

	return &tracks, nil
}
