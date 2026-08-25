package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
   id = strings.ReplaceAll(id," ","")
   if(len(id)<=1){
       return false
   }    
   var double bool
   slice:=[]rune(id) 
   sum :=0
   for i:=len(slice)-1;i>=0;i--{
       	if unicode.IsDigit(slice[i]) == false{
			return false
		}
       d := int(slice[i] -'0')
       if double{
           d*=2
           if d>9{
           d = d-9
       }
       }
       sum+=d
       double = !double
   }
   return sum%10==0
}

