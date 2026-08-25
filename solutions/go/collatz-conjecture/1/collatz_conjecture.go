package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if n<=0{
        return 0, errors.New("The integer value cannot be less than 0")
    }
    count:=0
    for n>1{
    if(n%2==0){
        n=n/2     
        count ++
    }else{
        n*=3
        n+=1
        count ++
    }
    }
    return count,nil
}
