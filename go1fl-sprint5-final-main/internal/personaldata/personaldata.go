package personaldata

import (
	"fmt"
	"log"
)

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	msg := fmt.Sprintf(
		"Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n",
		p.Name,
		p.Weight,
		p.Height,
	)

	log.Println("PRINT:", msg)

	fmt.Print(msg)
}
