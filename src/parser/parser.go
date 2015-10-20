package parser

import (
	"board"
	"errors"
	"player"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	Player              player.Player
	Board               board.Board
}

const (
	UNKNOWN_COMMAND = iota
	NAME_COMMAND
	COLORS_COMMAND
	HAND_COMMAND
	FLAG_CLAIM_COMMAND
	FLAG_CARD_COMMAND
	EMPTY_FLAG_CARD_COMMAND
	OPPONENT_PLAY_COMMAND
	GO_PLAY_COMMAND
)

func (p *Parser) ParseString(command string) error {
	commandType, parsedCommand := getCommandMatch(command)
	if commandType == NAME_COMMAND {
		p.Player.HandleRespondingToName(parsedCommand[1])
	} else if commandType == COLORS_COMMAND {
		//Not doing anything with this right now
	} else if commandType == HAND_COMMAND {
		parsedCommand = strings.Split(parsedCommand[1], " ")
		p.Player.HandleHandUpdate(parsedCommand)
	} else if commandType == FLAG_CLAIM_COMMAND {
		parsedCommand = append(parsedCommand[:0], parsedCommand[1:]...)
		p.Board.HandleFlagClaimCommand(parsedCommand)
	} else if commandType == FLAG_CARD_COMMAND {
		flagIndex, flagDirection, cards := getFlagCardMatchInfo(parsedCommand)
		p.Board.HandleFlagAddCardCommand(flagIndex, flagDirection, cards)
	} else if commandType == EMPTY_FLAG_CARD_COMMAND {
		//Not doing anything with this right now
	} else if commandType == OPPONENT_PLAY_COMMAND {
		//Not doing anything with this right now
	} else if commandType == GO_PLAY_COMMAND {
		//Not doing anything with this right now
	} else {
		return errors.New("Unkown command")
	}
	return nil
}

func getCommandMatch(command string) (int, []string) {
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

	if len(nameMatch) > 0 {
		return NAME_COMMAND, nameMatch
	} else if len(colorsMatch) > 0 {
		return COLORS_COMMAND, colorsMatch
	} else if len(handMatch) > 0 {
		return HAND_COMMAND, handMatch
	} else if len(flagClaimMatch) > 0 {
		return FLAG_CLAIM_COMMAND, flagClaimMatch
	} else if len(flagCardMatch) > 0 {
		return FLAG_CARD_COMMAND, flagCardMatch
	} else if len(emptyFlagCardMatch) > 0 {
		return EMPTY_FLAG_CARD_COMMAND, emptyFlagCardMatch
	} else if len(opponentPlayMatch) > 0 {
		return OPPONENT_PLAY_COMMAND, opponentPlayMatch
	} else if goPlayMatch {
		return GO_PLAY_COMMAND, nil
	} else {
		return UNKNOWN_COMMAND, nil
	}
}

func getFlagCardMatchInfo(command []string) (int, string, []string) {
	flagIndex, _ := strconv.Atoi(command[1])
	flagDirection := command[2]
	cards := strings.Split(command[3], " ")
	return flagIndex, flagDirection, cards
}
