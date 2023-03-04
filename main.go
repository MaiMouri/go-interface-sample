package main

import (
	"time"

	"github.com/MaiMouri/go-interface-sample/domain/service"
)

func main() {

	d := &service.Dryer{}
	l := &service.Laundry{}
	c := &service.PointCard{}

	b := NewBathroomOutlet(d, l)

	// ドライヤーメソッドを呼ぶ
	b.HairDrying()

	time.Sleep(2 * time.Second)

	// 洗濯メソッドを呼ぶ
	b.Laundry()

	time.Sleep(2 * time.Second)

	s := NewShopping(c)

	// ショッピングメソッドを呼ぶ
	s.ShowCard()

}
