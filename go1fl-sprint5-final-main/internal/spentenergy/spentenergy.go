package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчётов.
const (
	mInKm                      = 1000 // количество метров в километре
	minInH                     = 60   // количество минут в часе
	stepLengthCoefficient      = 0.45 // коэффициент для расчёта длины шага на основе роста
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчёта калорий при ходьбе
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("длительность должна быть больше 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationMinutes) / 60.0

	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть больше 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationMinutes) / 60.0

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if steps < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distKm := Distance(steps, height)

	durationHours := duration.Hours()

	return distKm / durationHours
}

func Distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}
