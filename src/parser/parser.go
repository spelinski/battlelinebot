package parser

import (
	"regexp"
	"strconv"
	"strings"
)

var BotVisualName string = "SynergyBot"

type Card struct {
	color  string
	number int
}

type Flag struct {
	claimer string
	north   []Card
	south   []Card
}

type Board struct {
	flags [9]Flag
}

type Parser struct {
	visualName          string
	direction           string
	lastCommandWasKnown bool
	colors              []string
	hand                []Card
	board               Board
}

func (p *Parser) ParseString(command string) {
	NameRegex := regexp.MustCompile("player\\s(.*)\\sname")
	colorsRegex := regexp.MustCompile("colors\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)")
	handRegex := regexp.MustCompile("player.*hand\\s(.*,.*)*")
	flagClaimRegex := regexp.MustCompile("flag claim-status\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)")
	flagCardRegex := regexp.MustCompile("flag ([1-9]) cards (north|south) (.*)")
	opponentPlayRegex := regexp.MustCompile("opponent play ([1-9]) (.*)")

	nameMatch := NameRegex.FindStringSubmatch(command)
	colorsMatch := colorsRegex.FindStringSubmatch(command)
	handMatch := handRegex.FindStringSubmatch(command)
	flagClaimMatch := flagClaimRegex.FindStringSubmatch(command)
	flagCardMatch := flagCardRegex.FindStringSubmatch(command)
	opponentPlayMatch := opponentPlayRegex.FindStringSubmatch(command)
	goPlayMatch, _ := regexp.MatchString("go play-card", command)
	p.lastCommandWasKnown = true
	if len(nameMatch) > 0 {
		p.visualName = BotVisualName
		p.direction = nameMatch[1]
	} else if len(colorsMatch) > 0 {
		colorsMatch = append(colorsMatch[:0], colorsMatch[1:]...)
		for _, color := range colorsMatch {
			p.colors = append(p.colors, color)
		}
	} else if len(handMatch) > 0 {
		handMatch = strings.Split(handMatch[1], " ")
		for _, card := range handMatch {
			cardDetails := strings.Split(card, ",")
			cardNumber, _ := strconv.Atoi(cardDetails[1])
			nextCard := Card{cardDetails[0], cardNumber}
			p.hand = append(p.hand, nextCard)
		}
	} else if len(flagClaimMatch) > 0 {
		flagClaimMatch = append(flagClaimMatch[:0], flagClaimMatch[1:]...)
		for i, claimer := range flagClaimMatch {
			p.board.flags[i].claimer = claimer
		}
	} else if len(flagCardMatch) > 0 {
		flagIndex, _ := strconv.Atoi(flagCardMatch[1])
		flagDirection := flagCardMatch[2]
		flagCardMatch = strings.Split(flagCardMatch[3], " ")
		tempFlagCardsList := []Card{}
		for _, card := range flagCardMatch {
			cardDetails := strings.Split(card, ",")
			cardNumber, _ := strconv.Atoi(cardDetails[1])
			nextCard := Card{cardDetails[0], cardNumber}
			tempFlagCardsList = append(tempFlagCardsList, nextCard)
		}
		if flagDirection == "north" {
			p.board.flags[flagIndex-1].north = tempFlagCardsList
		} else {
			p.board.flags[flagIndex-1].south = tempFlagCardsList
		}
	} else if len(opponentPlayMatch) > 0 {
		//Not doing anything with this right now
	} else if goPlayMatch {
		//Not doing anything with this right now
	} else {
		p.lastCommandWasKnown = false
	}
}
