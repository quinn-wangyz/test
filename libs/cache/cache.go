package cache

var items map[string]string

func init() {
	items = make(map[string]string)
}

func Set(key string, value string) {
	items[key] = value
}

func Get(key string) string {
	return items[key]
}
