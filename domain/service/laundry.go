package service

import (
	"fmt"
	"time"
)

type Laundry struct{}

func (d *Laundry) HasTwoPlugs() bool {
	return true
}
func (d *Laundry) LengthOfPlag() bool {
	return true
}
func (d *Laundry) DistanceOfPlugs() bool {
	return true
}
func (d *Laundry) SwitchOn() bool {
	fmt.Println("お洗濯スタート💇‍♀️")
	d.washing()
	return true
}
func (d *Laundry) washing() bool {
	time.Sleep(2 * time.Second)
	fmt.Println("........洗濯中")
	time.Sleep(2 * time.Second)
	fmt.Println("........洗濯終わった！")
	return true
}
