package bet

var openBets []Bet

func AddOpen(b Bet) {
	openBets = append(openBets, b)
}

func GetOpen() []Bet {
	return openBets
}

func Settle(b Bet) {
	for i, bet := range openBets {
		if b.Team == bet.Team && b.Prediction == bet.Prediction && b.Size == bet.Size {
			openBets[i] = openBets[len(openBets)-1]
			openBets = openBets[:len(openBets)-1]
			return
		}
	}
}

func ClearAll() {
	openBets = nil
}
