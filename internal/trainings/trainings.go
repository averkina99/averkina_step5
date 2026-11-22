package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку.
	TrainingType          string        // тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print()
}

/*
	Метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
*/

func (t *Training) Parse(datastring string) (err error) {
	dataArr := strings.Split(datastring, ",")

	if len(dataArr) != 3 {
		return fmt.Errorf("Некоректный формат входного параметра")
	}

	steps, err := strconv.Atoi(dataArr[0])
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(dataArr[2])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("Кол-во шагов должно быть больше 0")
	}

	if duration <= 0 {
		return fmt.Errorf("Дистанция должна быть больше 0")
	}

	t.Steps = steps
	t.TrainingType = dataArr[1]
	t.Duration = duration

	return nil
}

/*
	Метод формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
	Формат вывода:
		Тип тренировки: Бег
		Длительность: 0.75 ч.
		Дистанция: 10.00 км.
		Скорость: 13.34 км/ч
		Сожгли калорий: 18621.75
*/
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var cal float64
	var err error

	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = fmt.Errorf("Неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		avgSpeed,
		cal,
	), nil
}
