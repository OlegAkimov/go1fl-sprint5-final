package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps                 int           // количество шагов, проделанных за тренировку.
	TrainingType          string        // тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata,
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	if datastring == "" {
		return errors.New("datastring Пустой")
	}

	parts := strings.Split(datastring, ",") // Разбивает строку s по разделителю

	if len(parts) != 3 {
		return errors.New("length(parts) != 3")
	}

	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	t.TrainingType = parts[1]
	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	if t.Steps <= 0 {
		return errors.New("steps <= 0")
	}
	if t.Duration <= 0 {
		return errors.New("duration <= 0")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	template := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	return fmt.Sprintf(template, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil

}
