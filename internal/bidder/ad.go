package bidder

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

type AdDriver interface {
	Upsert(ad Ad) *Ad
	Delete(id string) bool
	Get(id string) (Ad, bool)
	GetAll() map[string]Ad
	GetTree() map[string](map[string]*Ad)
	GetByKey(key string) (map[string]*Ad, bool)
}
