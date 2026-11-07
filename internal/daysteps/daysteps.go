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
	// TODO: добавить поля
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata,
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	if datastring == "" {
		return errors.New("datastring Пустой")
	}

	parts := strings.Split(datastring, ",") // Разбивает строку s по разделителю

	if len(parts) != 2 {
		return errors.New("length(parts) != 2")
	}

	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	if ds.Steps <= 0 {
		return errors.New("steps <= 0")
	}
	if ds.Duration <= 0 {
		return errors.New("duration <= 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps <= 0 {
		return "", errors.New("steps <= 0")
	}
	if ds.Duration <= 0 {
		return "", errors.New("duration <= 0")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight <= 0")
	}
	if ds.Height <= 0 {
		return "", errors.New("height <= 0")
	}

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	template := `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`

	return fmt.Sprintf(template, ds.Steps, distance, calories), nil
}
