package bet

var openBets map[string]Bet

func AddOpen(messageID string, b Bet) {
	openBets[messageID] = b
}

func GetOpen() map[string]Bet {
	return openBets
}

func Settle(messageID string) {
	delete(openBets, messageID)
}

func ClearAll() {
	openBets = nil
}

func FormatOpenBets() string {
	betFormats := make([]string, len(openBets))
	for _, bet := range openBets {
		betFormats = append(betFormats, bet.Format())
	}
	var result string
	for i := range betFormats {
		result = result + betFormats[i]
	}
	return result
}