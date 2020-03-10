package card

import (
	"strconv"
	"strings"

	"github.com/emanpicar/cc-api/luhnalg"
)

type (
	Card struct {
		luhnManager luhnalg.LuhnManager
	}

	CardDetails struct {
		Name       string
		CardNumber string
		IsValid    bool
	}

	CardManager interface {
		GetCardListDetails(data []string) []CardDetails
		GetCardDetails(data string) CardDetails
	}
)

func New(luhnManager luhnalg.LuhnManager) CardManager {
	return &Card{luhnManager}
}

func (c *Card) GetCardListDetails(data []string) []CardDetails {
	cardList := []CardDetails{}

	for _, cardData := range data {
		cardList = append(cardList, c.GetCardDetails(cardData))
	}

	return cardList
}

func (c *Card) GetCardDetails(data string) CardDetails {
	return CardDetails{
		Name:       c.getCardName(data),
		CardNumber: data,
		IsValid:    c.luhnManager.Validate(data),
	}
}

func (c *Card) getCardName(data string) string {
	switch {
	case c.isCardAMEX(data):
		return "AMEX"
	case c.isCardDiscover(data):
		return "Discover"
	case c.isCardMastercard(data):
		return "Mastercard"
	case c.isCardVisa(data):
		return "Visa"
	default:
		return "Unknown"
	}
}

func (c *Card) isCardAMEX(data string) bool {
	newData := c.stripStringSpaces(data)

	return (strings.HasPrefix(newData, "34") || strings.HasPrefix(newData, "37")) && len(newData) == 15
}

func (c *Card) isCardDiscover(data string) bool {
	newData := c.stripStringSpaces(data)

	return strings.HasPrefix(newData, "6011") && len(newData) == 16
}

func (c *Card) isCardMastercard(data string) bool {
	newData := c.stripStringSpaces(data)
	firstDigit, err := strconv.ParseInt(newData[0:1], 10, 8)
	if err != nil {
		return false
	}

	secondDigit, err := strconv.ParseInt(newData[1:2], 10, 8)
	if err != nil {
		return false
	}

	return (firstDigit == 5 && (secondDigit >= 1 && secondDigit <= 5)) && len(newData) == 16
}

func (c *Card) isCardVisa(data string) bool {
	newData := c.stripStringSpaces(data)

	return strings.HasPrefix(newData, "4") && (len(newData) == 13 || len(newData) == 16)
}

func (c *Card) stripStringSpaces(data string) string {
	return strings.Join(strings.Fields(data), "")
}
