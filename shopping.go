package main

import "fmt"

// Card構造体の持つメソッドをラップするように実装
type ICard interface {
	IsCard() bool
	HasBarcode() bool
	ReadCard() int
}

type Shopping struct {
	card ICard // interfaceを実装する
}

func NewShopping(card ICard) *Shopping {
	return &Shopping{
		card: card,
	}
}

func (s *Shopping) ShowCard() int {
	fmt.Println("Show the point card.")
	point := s.card.ReadCard()
	fmt.Printf("........ %d ポイント加算！", point)

	return point
}
