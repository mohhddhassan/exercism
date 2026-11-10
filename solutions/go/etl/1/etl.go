package etl

import "strings"
func Transform(in map[int][]string) map[string]int {
	oto := make(map[string]int)
    for key,value := range in {
        v := key
        for _,value := range value{
            key := strings.ToLower(value)
            oto[key] = v
        }
    }
    return oto
}
