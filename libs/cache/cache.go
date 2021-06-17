package cache

import "fmt"

var items map[string]string

func init() {
	items = make(map[string]string)
}

func Set(key string, value string) {
	items[key] = value
}

func Get(key string) string {
	fmt.Println("99999999999999")
	fmt.Println(items)
	return items[key]
}
