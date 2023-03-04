package service

import (
	"fmt"
	"time"
)

type Dryer struct{}

func (d *Dryer) HasTwoPlugs() bool {
	return true
}
func (d *Dryer) LengthOfPlag() bool {
	return true
}
func (d *Dryer) DistanceOfPlugs() bool {
	return true
}
func (d *Dryer) SwitchOn() bool {
	fmt.Println("ドライヤースタート💇‍♀️")
	d.drying()
	return true
}
func (d *Dryer) drying() bool {
	time.Sleep(1 * time.Second)
	fmt.Println("........乾かし中")
	time.Sleep(1 * time.Second)
	fmt.Println("........乾いた！")
	return true
}
