package raindrops
import (
    "strings"
    "fmt"
)
func Convert(number int) string {
    var result []string
    var res string
	if number%3 == 0{
       res = "Pling"
       result = append(result, res)
    }
    if number%5 == 0{
       res = "Plang"
       result = append(result, res)
     }
    if number%7 == 0{
       res = "Plong"
       result = append(result, res)
    }
    if len(result)>0{
        return strings.Join(result,"")
    }
	return fmt.Sprintf("%d",number)
}
