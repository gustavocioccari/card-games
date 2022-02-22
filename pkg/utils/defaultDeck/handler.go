package utils

import (
	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/card"
)

type util struct {
	cardRepository card.CardRepository
}

type DefaultDeckUtil interface {
	GetDefaultDeck() (*models.Deck, error)
}

func NewDefaultDeckUtil(cardRepository card.CardRepository) DefaultDeckUtil {
	return &util{
		cardRepository: cardRepository,
	}
}

func (u *util) GetDefaultDeck() (*models.Deck, error) {
	codes := []string{
		"AS", "AD", "AC", "AH", "2S", "2D", "2C", "2H",
		"3S", "3D", "3C", "3H", "4S", "4D", "4C", "4H",
		"5S", "5D", "5C", "5H", "6S", "6D", "6C", "6H",
		"7S", "7D", "7C", "7H", "8S", "8D", "8C", "8H",
		"9S", "9D", "9C", "9H", "10S", "10D", "10C", "10H",
		"JS", "JD", "JC", "JH", "KS", "KD", "KC", "KH", "QS", "QD", "QC", "QH",
	}

	cards, err := u.cardRepository.FindByCode(codes)
	if err != nil {
		return nil, err
	}

	deck := &models.Deck{
		Shuffled:  false,
		Remaining: len(cards),
		Cards:     cards,
	}

	return deck, nil
}
