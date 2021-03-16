package types

type RecentTracksType struct {
	Recenttracks struct {
		Attr struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`
		Track []struct {
			Artist struct {
				Mbid string `json:"mbid"`
				Text string `json:"#text"`
			} `json:"artist"`
			Attr struct {
				Nowplaying string `json:"nowplaying"`
			} `json:"@attr,omitempty"`
			Mbid  string `json:"mbid"`
			Album struct {
				Mbid string `json:"mbid"`
				Text string `json:"#text"`
			} `json:"album"`
			Streamable string `json:"streamable"`
			URL        string `json:"url"`
			Name       string `json:"name"`
			Image      []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Date struct {
				Uts  string `json:"uts"`
				Text string `json:"#text"`
			} `json:"date,omitempty"`
		} `json:"track"`
	} `json:"recenttracks"`
}
