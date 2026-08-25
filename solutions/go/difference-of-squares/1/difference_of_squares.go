package differenceofsquares

func SquareOfSum(n int) int {
    count:=0
	for j:=1;j<=n;j++{
        count +=j
    }
    count = count*count
    return count
}

func SumOfSquares(n int) int {
    count:=0
	for j:=1;j<=n;j++{
        count +=j*j
    }
    return count
}

func Difference(n int) int {
	squareos :=SquareOfSum(n)
    sumos :=SumOfSquares(n)
    dif := squareos - sumos
    return dif
}
