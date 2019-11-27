package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
)

func (act *Actions) FindByDay(day string) (entity.Date, error) {
	date, err := act.DateRepository.FindByDay(day)
	if err != nil {
		return entity.Date{}, err
	}
	return date, nil
}

func (act *Actions) CreateEmptyDate(day string) (uint32, error) {
	date := entity.Date{
		Day:           day,
		Description:   "",
		IsCelebration: false,
		Events:        nil,
	}
	dayId, err := act.DateRepository.Save(date)
	if err != nil {
		act.Logger.Info("An error occurred while create day")
		return 0, err
	}
	return dayId, nil
}
