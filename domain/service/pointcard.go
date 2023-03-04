package service

import (
	"fmt"
	"time"
)

type PointCard struct{}

func (p *PointCard) IsCard() bool {
	return true
}
func (p *PointCard) HasBarcode() bool {
	return true
}

func (p *PointCard) ReadCard() int {
	fmt.Println("カードを読み取る💳")
	point := p.addPoint()
	return point
}
func (p *PointCard) addPoint() int {
	point := 3
	time.Sleep(1 * time.Second)
	fmt.Println("........読み込み中")
	time.Sleep(1 * time.Second)
	fmt.Println("........ DONE!")
	return point
}
