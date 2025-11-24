package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, data := range dataset {

		if err := dp.Parse(data); err != nil {
			log.Println("Ошибка при разборе данных:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка при формировании информации:", err)
			continue
		}

		log.Println(info)
	}
}
