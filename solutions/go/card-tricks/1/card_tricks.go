package cards

import "slices"
// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	Fav := []int{2,6,9}
    return Fav
}

// GetItem retrieves an item from a slice at given position.

func GetItem(slice []int, index int) int {
    if index >= 0 && index <= len(slice) - 1{ 
        indexitem := slice[index]
        return indexitem
    }
    return -1
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index >= 0 && index <= len(slice) -1 {
        slice[index] = value
        return slice
    }
	slice = append(slice, value)
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var slice2 []int
    for _,v := range values { 
         slice2 = append(slice2,v)
    }
    slice3 := append(slice2, slice...)
    return slice3
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index >= 0 && index <= len(slice) -1 {
		slice = slices.Delete(slice, index, index+1) 
    return slice
    }
    return slice
}
