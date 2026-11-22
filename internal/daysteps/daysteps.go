package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print()
}

/*
	Метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps
*/
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataArr := strings.Split(datastring, ",")

	if len(dataArr) != 2 {
		return fmt.Errorf("некорректный формат входного параметра")
	}

	steps, err := strconv.Atoi(dataArr[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(dataArr[1])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("Кол-во шагов должно быть больше 0")
	}

	if duration <= 0 {
		return fmt.Errorf("длительность должна быть больше 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

/*
	Метод формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
	Формат вывода:
		Количество шагов: 792.
		Дистанция составила 0.51 км.
		Вы сожгли 221.33 ккал.
*/
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		cal,
	), nil
}
