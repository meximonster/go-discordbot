package bet

var openBets = make(map[string]openBet)

type openBet struct {
	message_id string
	Bet
}

func AddOpen(messageID string, b Bet) {
	openBets[messageID] = openBet{
		message_id: messageID,
		Bet:        b,
	}
}

func GetOpen() map[string]openBet {
	return openBets
}

func Settle(messageID string) {
	delete(openBets, messageID)
}

func ClearAll() {
	for k := range openBets {
		delete(openBets, k)
	}
}

func SaveOpen() error {
	for k, v := range openBets {
		q := `INSERT INTO open_bets (message_id,team,prediction,size,odds) VALUES ($1,$2,$3,$4,$5)`
		_, err := dbC.Exec(q, k, v.Team, v.Prediction, v.Size, v.Odds)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadOpen() error {
	bets := []openBet{}
	err := dbC.Select(&bets, `SELECT message_id,team,prediction,size,odds FROM open_bets`)
	if err != nil {
		return err
	}
	for _, bet := range bets {
		openBets[bet.message_id] = bet
	}
	_, err = dbC.Exec(`DELETE FROM open_bets`)
	if err != nil {
		return err
	}
	return nil
}

func FormatOpenBets() string {
	betFormats := make([]string, 0, len(openBets))
	for _, bet := range openBets {
		betFormats = append(betFormats, bet.Format())
	}
	var result string
	for i := range betFormats {
		result = result + betFormats[i]
	}
	return result
}
