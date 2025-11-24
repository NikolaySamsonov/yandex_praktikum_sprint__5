package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных: должно быть 2 поля")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	if steps <= 0 {
		return errors.New("значение шагов должно быть больше 0")
	}
	d, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = d
	if d <= 0 {
		return errors.New("длительность тренировки должна быть больше 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		calories,
	)

	return result, nil
}
