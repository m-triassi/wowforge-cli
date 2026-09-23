package search

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// LoadAddons returns the tracked addons as a map of id -> installed folder
// names. Legacy config shapes ([]int, or id->name string) are migrated in
// place with empty folder lists that self-heal on the next add/update.
func LoadAddons() map[string][]string {
	m := map[string][]string{}
	for key, val := range viper.GetStringMap("addons") {
		if folders, ok := val.([]interface{}); ok {
			for _, f := range folders {
				m[key] = append(m[key], fmt.Sprint(f))
			}
		} else {
			m[key] = []string{}
		}
	}
	if len(m) > 0 {
		return m
	}

	for _, id := range viper.GetIntSlice("addons") {
		m[strconv.Itoa(id)] = []string{}
	}
	return m
}
