package usecases

import (
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/entity"
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
