package inmemorycache

import (
	"fmt"

	"github.com/google/uuid"

	bidder "ako.com/internal/bidder"
)

type InMemoryCache struct {
	ads       map[string]bidder.Ad
	sortedAds map[string](map[string]*bidder.Ad)
}

func (cache InMemoryCache) Upsert(newAd bidder.Ad) *bidder.Ad {
	if len(newAd.ID) == 0 {
		newAd.ID = uuid.NewString()
	}
	newAd.Key = fmt.Sprintf("%s:%d:%d", newAd.Filters.AdType, newAd.Filters.H, newAd.Filters.W)
	cache.ads[newAd.ID] = newAd
	cache.addToSortedAds(&newAd)
	return &newAd
}
func (cache InMemoryCache) addToSortedAds(ad *bidder.Ad) {
	if _, found := cache.sortedAds[ad.Key]; !found {
		cache.sortedAds[ad.Key] = make(map[string]*bidder.Ad)
	}
	cache.sortedAds[ad.Key][ad.ID] = ad
}
func (cache InMemoryCache) Get(id string) (bidder.Ad, bool) {
	ad, found := cache.ads[id]
	return ad, found
}
func (cache InMemoryCache) GetAll() map[string]bidder.Ad {
	return cache.ads
}
func (cache InMemoryCache) GetTree() map[string](map[string]*bidder.Ad) {
	return cache.sortedAds
}
func (cache InMemoryCache) GetByKey(key string) (map[string]*bidder.Ad, bool) {
	leaf, found := cache.sortedAds[key]
	return leaf, found
}
func (cache InMemoryCache) Delete(id string) bool {
	delete(cache.ads, id)
	delete(cache.sortedAds, id)
	return true
}
