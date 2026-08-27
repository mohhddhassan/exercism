package differenceofsquares

func SquareOfSum(n int) int {
    count :=0
    for i := range n{ 
        j := i+1
        count +=j
    }
    count *= count
    return count
}

func SumOfSquares(n int) int {
    count :=0
    for i := range n{
        j := i+1 
        count +=j*j
    }
    return count
}

func Difference(n int) int {
	squareos :=SquareOfSum(n)
    sumos :=SumOfSquares(n)
    return squareos - sumos
}
