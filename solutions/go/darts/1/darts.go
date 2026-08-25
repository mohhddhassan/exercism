package darts
import "math"

func Score(x, y float64) int {
	sol :=math.Sqrt((x*x)+(y*y))
    if (sol<=1){
        res := 10
        return res
    } else if(sol>=1 && sol<=5){
        res:= 5
        return res
    }else if(sol>=5 && sol<=10 ){
        res:=1
        return res
    }else{
        return 0
    }
}
