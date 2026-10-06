package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
    set := make(map[int]struct{})
    total := 0 
	for _, divisor := range divisors {
        for i := 1 ; i < limit ; i++{
            var t = divisor*i
            if t < limit {
              set[divisor*i] = struct{}{}
            }
        }
    }
    for key, _ := range set{
        total = total + key
    }
    return total 
}
