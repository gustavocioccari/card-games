package deck

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gustavocioccari/toggl-cards/pkg/models"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/card"
	"github.com/gustavocioccari/toggl-cards/pkg/repositories/postgres/deck"
	utils "github.com/gustavocioccari/toggl-cards/pkg/utils/defaultDeck"
)

type service struct {
	deckRepository deck.DeckRepository
	cardRepository card.CardRepository
}

type DeckService interface {
	CreateDefault() (*models.Deck, error)
	Create(cards, shuffled string) (*models.Deck, error)
	GetByID(id string) (*models.Deck, error)
}

func NewDeckService(deckRepository deck.DeckRepository, cardRespository card.CardRepository) DeckService {
	return &service{
		deckRepository: deckRepository,
		cardRepository: cardRespository,
	}
}

func (s *service) CreateDefault() (*models.Deck, error) {
	var deck *models.Deck

	defaultDeckUtil := utils.NewDefaultDeckUtil(s.cardRepository)

	deck, err := defaultDeckUtil.GetDefaultDeck()
	if err != nil {
		return nil, err
	}

	err = s.deckRepository.Create(deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (s *service) Create(cards, shuffled string) (*models.Deck, error) {
	var deck *models.Deck

	defaultDeckUtil := utils.NewDefaultDeckUtil(s.cardRepository)

	deck, err := defaultDeckUtil.GetDefaultDeck()
	if err != nil {
		return nil, err
	}

	if len(cards) > 0 {
		cardsSlice := strings.Split(cards, ",")

		cards, err := s.cardRepository.FindByCode(cardsSlice)
		if err != nil {
			return nil, err
		}

		deck.Cards = cards
		deck.Remaining = len(deck.Cards)
	}
	log.Println("deck", deck)

	if shuffled == "true" {
		deck.Seed = time.Now().UnixNano()

		rand.Seed(deck.Seed)
		rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
		deck.Shuffled = true
	}

	err = s.deckRepository.Create(deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (s *service) GetByID(id string) (*models.Deck, error) {
	deck, err := s.deckRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if deck.Shuffled {
		rand.Seed(deck.Seed)
		rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
	}

	return deck, nil
}
