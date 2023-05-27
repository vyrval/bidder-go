package ad

type AdDriver interface {
	upsert(ad Ad) Ad
	delete(id string) bool
	get(id string) Ad
	getAll() map[string]Ad
	getByKey(key string) map[string]Ad
}
