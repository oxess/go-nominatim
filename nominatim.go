package nominatim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Nominatim struct {
	Url string
}

type SearchResponseItem struct {
	PlaceId     int      `json:"place_id"`
	OSMId       int      `json:"osm_id"`
	OSMType     string   `json:"osm_type"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	BoundingBox []string `json:"boundingbox"`
	Licence     string   `json:"licence"`
	Icon        string   `json:"icon"`
}

func NewPublicApi() Nominatim {
	return Nominatim{Url: "https://nominatim.openstreetmap.org"}
}

func New(url string) Nominatim {
	return Nominatim{Url: url}
}

func (n Nominatim) Search(query string, limit int8) ([]SearchResponseItem, error) {
	var searchResponseItems []SearchResponseItem
	var limitString string
	var err error

	if limit != -1 {
		limitString = fmt.Sprintf("limit=%d&", limit)
	}

	requestUrl := fmt.Sprintf("%s/search?%sformat=json&q=%s", n.Url, limitString, query)
	response, err := http.Get(requestUrl)
	if err != nil {
		return searchResponseItems, err
	}

	err = json.NewDecoder(response.Body).Decode(&searchResponseItems)
	if err != nil {
		return searchResponseItems, err
	}

	return searchResponseItems, nil
}
