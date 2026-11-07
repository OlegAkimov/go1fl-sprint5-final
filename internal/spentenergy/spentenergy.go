package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("Количество шагов 0 или меньше")
		return 0, err
	}

	if weight == 0 || height == 0 {
		err := errors.New("weight или height равны 0")
		return 0, err
	}

	if duration == 0 {
		err := errors.New("Продолжительность равна 0")
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("Количество шагов 0 или меньше")
		return 0, err
	}

	if weight == 0 || height == 0 {
		err := errors.New("weight или height равны 0")
		return 0, err
	}

	if duration == 0 {
		err := errors.New("Продолжительность равна 0")
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration == 0 || steps <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stideLength := height * stepLengthCoefficient
	return (stideLength * float64(steps)) / mInKm
}
