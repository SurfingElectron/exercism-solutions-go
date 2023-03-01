package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValue := map[string] int {
        "ace": 11,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "ten": 10,
        "jack": 10,
        "queen": 10,
        "king": 10, 	
    }
	return cardValue[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
        case isWin(card1, card2, dealerCard):
    		return "W"
        case isSplit(card1, card2):
    		return "P"
        case isStand(card1, card2, dealerCard):
    		return "S"
        default:
    		return "H"
    }
}

// sumCards sums the player's two cards and returns that value
func sumCards(card1, card2 string) int {
    return ParseCard(card1) + ParseCard(card2)
}

// isBlackjack returns whether the first pair is a Blackjack
func isBlackjack(card1, card2 string) bool {
	return ParseCard(card1) + ParseCard(card2) == 21
}

// isWin returns whether the first hand meets the win conditions
func isWin(card1, card2, dealerCard string) bool {
    return isBlackjack(card1, card2) && ParseCard(dealerCard) < 10
}

// isSplit returns whether the player should split after their first hand
func isSplit(card1, card2 string) bool {
    return card1 == "ace" && card2 == "ace"
}

// isStand returns whether the player should stand after their first hand
func isStand(card1, card2, dealerCard string) bool {
    switch {
    case isBlackjack(card1, card2) && ParseCard(dealerCard) >= 10:
    	return true
    case sumCards(card1, card2) >= 17 && sumCards(card1, card2) <= 20:
    	return true
    case sumCards(card1, card2) >= 12 && sumCards(card1, card2) <= 16 && ParseCard(dealerCard) < 7:
    	return true
    default:
    	return false
    }  
}