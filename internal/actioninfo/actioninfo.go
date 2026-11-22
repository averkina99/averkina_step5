package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error) // метод парсинга строки и зааполнение объекта
	ActionInfo() (string, error)         // формирование строки о состоянии тренировки
}

/*
	Функция для выводи информации о любом виде тренировок
*/
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}

		nf, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(nf)
	}
}
