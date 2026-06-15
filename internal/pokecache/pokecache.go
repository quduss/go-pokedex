


type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}
