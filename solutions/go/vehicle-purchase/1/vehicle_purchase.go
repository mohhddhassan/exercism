package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
import "fmt"
func NeedsLicense(kind string) bool {
	if kind=="car" || kind =="truck" {
        return true
    } 
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var r string 
    if option1 < option2 {
        r =fmt.Sprintf("%s is clearly the better choice.",option1)
    }else{
    	r = fmt.Sprintf("%s is clearly the better choice.",option2)
    }
    return r
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var resellprice float64
	if age < 3 {
        resellprice = originalPrice * 80/100
    }else if age >= 10 {
        resellprice = originalPrice * 50/100
    }else if age >= 3 && age < 10{
        resellprice = originalPrice * 70/100
    }
    return resellprice
}
