package bet

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var openBets = make(map[string]Bet)

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
	for k := range openBets {
		delete(openBets, k)
	}
}

func SaveOpen() error {
	f, err := os.Create("open.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	for k, v := range openBets {
		s := fmt.Sprintf("%s:%v\n", k, v)
		f.WriteString(s)
	}
	return nil
}

func LoadOpen() error {
	f, err := os.Open("open.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		text := strings.SplitN(fileScanner.Text(), ":", 2)
		if len(text) == 2 {
			b, err := Decouple(text[1], "")
			if err != nil {
				return err
			}
			openBets[text[0]] = b
		}
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
