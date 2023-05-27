package ad

import (
	"fmt"

	"github.com/google/uuid"
)

type Filter struct {
	AdType string `json:"type"`
	H      int    `json:"h"`
	W      int    `json:"w"`
}

type Ad struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Filters Filter  `json:"filters"`
	Price   float32 `json:"price"`
	Key     string
}

var ads = make(map[string]Ad)
var sortedAds = make(map[string](map[string]*Ad))

func CreateAd(newAd Ad) *Ad {
	if len(newAd.ID) == 0 {
		newAd.ID = uuid.NewString()
	}
	newAd.Key = fmt.Sprintf("%s:%d:%d", newAd.Filters.AdType, newAd.Filters.H, newAd.Filters.W)
	ads[newAd.ID] = newAd
	addToSortedAds(&newAd)
	return &newAd
}

func addToSortedAds(ad *Ad) {
	if _, found := sortedAds[ad.Key]; !found {
		sortedAds[ad.Key] = make(map[string]*Ad)
	}
	sortedAds[ad.Key][ad.ID] = ad
}

func GetAdById(id string) (Ad, bool) {
	ad, found := ads[id]
	return ad, found
}
func GetAllAds() map[string]Ad {
	return ads
}
func GetTree() map[string](map[string]*Ad) {
	return sortedAds
}
func GetLeafByKey(key string) (map[string]*Ad, bool) {
	leaf, found := sortedAds[key]
	return leaf, found
}
