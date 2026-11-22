package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

/*
	Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
*/
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0.0
	}

	return height * stepLengthCoefficient * float64(steps) / mInKm
}

/*
	Функция принимает количество шагов steps, рост пользователя height и продолжительность активности duration
	и возвращает среднюю скорость.
*/
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0.0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

/*
	Функция принимает:
		steps int — количество шагов.
		weight, height float64 — вес(кг.) и рост(м.) пользователя.
		duration time.Duration — продолжительность бега.

	И возвращает два значения:
		float64 — количество калорий, потраченных при беге.
		error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
*/
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0.0, fmt.Errorf("Кол-во шагов должно быть больше 0")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("Рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("длительность должна быть больше 0")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("Вес должен быть больше 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)

	return weight * avgSpeed * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	runCal, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0.0, err
	}

	return runCal * walkingCaloriesCoefficient, nil
}
