package main

type Score = int

type Card rune

func (card Card) Score() Score {
	scores := map[Card]Score{
		'*': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	if score, ok := scores[card]; ok {
		return score
	} else {
		panic(card)
	}
}

var AllCards = []Card{
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'T',
	'J',
	'Q',
	'K',
	'A',
}

type Hand [5]Card

const (
	highCard Score = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func NewHand(s string) Hand {
	if len(s) != 5 {
		panic(s)
	}

	cards := [5]Card{}
	for i, c := range s {
		cards[i] = Card(c)
	}
	return cards
}

func NewWildHand(s string) Hand {
	if len(s) != 5 {
		panic(s)
	}

	cards := [5]Card{}
	for i, c := range s {
		if c == 'J' {
			c = '*'
		}
		cards[i] = Card(c)
	}
	return cards
}

var scoreMemo = map[Hand]Score{}

func (hand Hand) Score() Score {
	if score, found := scoreMemo[hand]; found {
		return score
	} else {
		score = hand.calcScore()
		scoreMemo[hand] = score
		return score
	}
}

func (hand Hand) calcScore() Score {
	cardCount := map[Card]int{}

	for i, card := range hand {
		if card == '*' {
			bestScore := -1
			for _, wild := range AllCards {
				next := hand
				next[i] = wild
				score := next.Score()
				if score > bestScore {
					bestScore = score
				}
			}
			return bestScore
		}

		cardCount[card] += 1
	}

	hasThree := false
	hasTwo := false

	for _, count := range cardCount {
		switch count {
		case 5:
			return fiveOfAKind
		case 4:
			return fourOfAKind
		case 3:
			hasThree = true
		case 2:
			if hasTwo {
				return twoPair
			}
			hasTwo = true
		}
	}

	if hasThree && hasTwo {
		return fullHouse
	}
	if hasThree {
		return threeOfAKind
	}
	if hasTwo {
		return onePair
	}

	return highCard
}

func (h Hand) String() string {
	s := ""
	for _, c := range h {
		s += string(c)
	}
	return s
}

func HandCompare(a, b Hand) int {
	diff := a.Score() - b.Score()

	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	}

	for i := range (Hand{}) {
		diff = a[i].Score() - b[i].Score()

		if diff < 0 {
			return -1
		} else if diff > 0 {
			return 1
		}
	}

	return 0
}
