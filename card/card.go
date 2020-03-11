package card

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/emanpicar/cc-api/logger"
	"github.com/emanpicar/cc-api/luhnalg"
)

type (
	card struct {
		luhnManager luhnalg.Manager
	}

	CardDetails struct {
		Result     string `json:"result"`
		Name       string `json:"name"`
		CardNumber string `json:"card_number"`
		IsValid    bool   `json:"is_valid"`
	}

	Manager interface {
		ValidateCardList(body io.ReadCloser) (*[]CardDetails, error)
		GetCardListDetails(data []string) *[]CardDetails
		GetCardDetails(data string) CardDetails
	}
)

func New(luhnManager luhnalg.Manager) Manager {
	return &card{luhnManager}
}

func (c *card) ValidateCardList(body io.ReadCloser) (*[]CardDetails, error) {
	var user []string
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		return nil, errors.New("Invalid data format. Requires string array")
	}

	logger.Log.Infoln("Card list validation complete")
	return c.GetCardListDetails(user), nil
}

func (c *card) GetCardListDetails(data []string) *[]CardDetails {
	cardList := []CardDetails{}

	for _, cardData := range data {
		cardList = append(cardList, c.GetCardDetails(cardData))
	}

	return &cardList
}

func (c *card) GetCardDetails(data string) CardDetails {
	cardName := c.getCardName(data)
	isValid := c.luhnManager.Validate(data)
	isValidValue := "invalid"
	if isValid {
		isValidValue = "valid"
	}

	return CardDetails{
		Result:     fmt.Sprintf("%v: %v (%v)", cardName, data, isValidValue),
		Name:       cardName,
		CardNumber: data,
		IsValid:    isValid,
	}
}

func (c *card) getCardName(data string) string {
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

func (c *card) isCardAMEX(data string) bool {
	newData := c.stripStringSpaces(data)

	return (strings.HasPrefix(newData, "34") || strings.HasPrefix(newData, "37")) && len(newData) == 15
}

func (c *card) isCardDiscover(data string) bool {
	newData := c.stripStringSpaces(data)

	return strings.HasPrefix(newData, "6011") && len(newData) == 16
}

func (c *card) isCardMastercard(data string) bool {
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

func (c *card) isCardVisa(data string) bool {
	newData := c.stripStringSpaces(data)

	return strings.HasPrefix(newData, "4") && (len(newData) == 13 || len(newData) == 16)
}

func (c *card) stripStringSpaces(data string) string {
	return strings.Join(strings.Fields(data), "")
}
