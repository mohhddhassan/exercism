package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var val int
	switch {
        case card == "ace":
        	val = 11
        case card == "two":
        	val = 2
        case card == "three":
        	val = 3
        case card == "four":
        	val = 4
        case card == "five":
        	val = 5
        case card == "six":
        	val = 6
        case card == "seven":
        	val = 7
        case card == "eight":
        	val = 8
        case card == "nine":
        	val = 9
        case card == "ten" || card == "jack" || card == "queen" || card == "king" :
        	val = 10
        default:
        	val = 0
    }
    return val
}
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    a := ParseCard(card1)
    b := ParseCard(card2)
    c := ParseCard(dealerCard)
    if card1 == "ace" && card2 == "ace"{
        return "P"
    }else if a + b == 21 && c != 11 && c !=10{ 
        return "W"        
    }else if a+b >= 17 && a+b <= 20{
        return "S"
    }else if a+b >= 12 && a+b <= 16 && c >=7{
        return "H"
    }else if a+b >= 12 && a+b <= 16{
        return "S"
    }else if a+b <= 11{
        return "H"
    }
	return "S"
}
