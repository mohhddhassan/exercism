package isbnverifier

import "strings"
func IsValidISBN(isbn string) bool {
    if(isbn ==""){
        return false
    }
    hyp := strings.ReplaceAll(isbn,"-","")
    if(len(hyp)!=10){
        return false
    }
    digi :=0
    for i:=0;i<10;i++{
        char := hyp[i]
		w:=0
        multi := 10 -i
        if i == 9 && (char == 'X' || char == 'x') {
			w = 10
		} else if char >= '0' && char <= '9' {
			w = int(char - '0')
		} else {
			return false
		}
        digi+= w * multi
    }
    if digi%11!=0{
        return false
    }
    return true
}