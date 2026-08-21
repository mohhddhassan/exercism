// Package twofer returns the name for whom will the cookie given.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"
// ShareWith will give the name of the person the cookie will be given or just return one for you one for me

func ShareWith(name string) string {
    if name==""{
        return "One for you, one for me."
    }
	return fmt.Sprintf("One for %s, one for me.",name)
}
