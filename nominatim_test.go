package nominatim

import (
	"testing"
)

func TestNominatim_PublicSearch(t *testing.T) {
	geocoder := NewPublicApi()
	_, err := geocoder.Search("Warszawa, plac zbawiciela 1", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestNominatim_Search(t *testing.T) {
	geocoder := New("https://nominatim.openstreetmap.org")
	_, err := geocoder.Search("Kalisz, ratusz", 10)
	if err != nil {
		t.Error(err)
	}
}
