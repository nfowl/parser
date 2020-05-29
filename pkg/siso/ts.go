package siso

import (
	"fmt"
)

func init() {
	AvailableWriters.addWriter("ts", &TSWriter{})
}

type TSWriter struct {
	Ayo string
}

func (t TSWriter) WriteEnum(d *DisEnum) error {
	// fmt.Println(d)
	return nil
}

func (t TSWriter) WriteCet(c *Cet) error {
	fmt.Println(c.Name)
	return nil
}
