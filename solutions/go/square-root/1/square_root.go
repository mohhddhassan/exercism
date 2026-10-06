package squareroot
import (
	"math"
    "errors"
)

func SquareRoot(number int) (int, error) {
    if number < 0{
        return 0, errors.New("cannot calculate square root for negative number")
    }
	var sq = math.Sqrt(float64(number))
    return int(sq),nil
}
