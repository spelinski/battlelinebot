package parser

import (
	"board"
	"player"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	player              player.Player
	lastCommandWasKnown bool
	pBoard              board.Board
}

func (p *Parser) ParseString(command string) {
	nameRegex := regexp.MustCompile("player\\s(.*)\\sname")
	colorsRegex := regexp.MustCompile("colors\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)")
	handRegex := regexp.MustCompile("player.*hand\\s(.*,.*)*")
	flagClaimRegex := regexp.MustCompile("flag claim-status\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)\\s(.*)")
	flagCardRegex := regexp.MustCompile("flag ([1-9]) cards (north|south) (.*)")
	emptyFlagCardRegex := regexp.MustCompile("flag ([1-9]) cards (north|south)")
	opponentPlayRegex := regexp.MustCompile("opponent play ([1-9]) (.*)")

	nameMatch := nameRegex.FindStringSubmatch(command)
	colorsMatch := colorsRegex.FindStringSubmatch(command)
	handMatch := handRegex.FindStringSubmatch(command)
	flagClaimMatch := flagClaimRegex.FindStringSubmatch(command)
	flagCardMatch := flagCardRegex.FindStringSubmatch(command)
	emptyFlagCardMatch := emptyFlagCardRegex.FindStringSubmatch(command)
	opponentPlayMatch := opponentPlayRegex.FindStringSubmatch(command)
	goPlayMatch, _ := regexp.MatchString("go play-card", command)
	p.lastCommandWasKnown = true
	if len(nameMatch) > 0 {
		p.player.HandleRespondingToName(nameMatch[1])
	} else if len(colorsMatch) > 0 {
		//Not doing anything with this right now
	} else if len(handMatch) > 0 {
		handMatch = strings.Split(handMatch[1], " ")
		p.player.HandleHandUpdate(handMatch)
	} else if len(flagClaimMatch) > 0 {
		flagClaimMatch = append(flagClaimMatch[:0], flagClaimMatch[1:]...)
		p.pBoard.HandleFlagClaimCommand(flagClaimMatch)
	} else if len(flagCardMatch) > 0 {
		flagIndex, flagDirection, cards := getFlagCardMatchInfo(flagCardMatch)
		p.pBoard.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	} else if len(emptyFlagCardMatch) > 0 {
		//Not doing anything with this right now
	} else if len(opponentPlayMatch) > 0 {
		//Not doing anything with this right now
	} else if goPlayMatch {
		//Not doing anything with this right now
	} else {
		p.lastCommandWasKnown = false
	}
}

func getFlagCardMatchInfo(command []string) (int, string, []string) {
	flagIndex, _ := strconv.Atoi(command[1])
	flagDirection := command[2]
	cards := strings.Split(command[3], " ")
	return flagIndex, flagDirection, cards
}
