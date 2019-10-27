package usecases

import "github.com/snarskliveshere/otus_golang/hw_8/src/entity"

func (act *Actions) AddRecordToDay(recId, dayId int) error {
	day, err := act.getDayById(dayId)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	rec, err := act.getRecordById(recId)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.DayRepository.AddRecordToDay(rec, day)
	if err != nil {
		act.Logger.Log("An error occurred while add day records")
		return err
	}

	return nil
}

func (act *Actions) AddDay(id int) error {
	day, err := act.getDayById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.DayRepository.Save(day)
	if err != nil {
		act.Logger.Log("An error occurred while day added")
		return err
	}
	act.Logger.Log("Day added successfully")

	return nil
}

func (act *Actions) ShowDayRecords(id int) ([]entity.Record, error) {
	day, err := act.getDayById(id)
	records, err := act.DayRepository.ShowDayRecords(day)
	if err != nil {
		act.Logger.Log("An error occurred while show day records")
		return []entity.Record{}, err
	}
	return records, nil
}

func (act *Actions) getDayById(id int) (entity.Day, error) {
	day, err := act.DayRepository.FindById(uint64(id))
	if err != nil {
		act.Logger.Log("An error occurred while get day")
		return entity.Day{}, err
	}
	return day, nil
}
