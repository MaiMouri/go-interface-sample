package main

import "fmt"

// Dryer構造体の持つメソッドをラップするように実装
type IBathroomOutlet interface {
	HasTwoPlugs() bool
	LengthOfPlag() bool
	DistanceOfPlugs() bool
	SwitchOn() bool
}

type BathroomOutlet struct {
	// dr service.Dryer
	firstOutlet  IBathroomOutlet // interfaceを実装する
	secondOutlet IBathroomOutlet // interfaceを実装する
}

// func NewBathroomOutlet(ib IBathroomOutlet) *BathroomOutlet {
// 	return &BathroomOutlet{
// 		ib: ib,
// 	}
// }

func NewBathroomOutlet(one IBathroomOutlet, two IBathroomOutlet) *BathroomOutlet {
	return &BathroomOutlet{
		firstOutlet:  one,
		secondOutlet: two,
	}
}

func (b *BathroomOutlet) HairDrying() {
	fmt.Println("First outlet gets connected.")
	b.firstOutlet.SwitchOn()
}

func (b *BathroomOutlet) Laundry() {
	fmt.Println("Second outlet gets connected.")
	b.secondOutlet.SwitchOn()
}
