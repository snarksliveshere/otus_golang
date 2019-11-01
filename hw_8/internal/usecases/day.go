package usecases

import "github.com/snarskliveshere/otus_golang/hw_8/entity"

func (act *Actions) AddRecordToDate(recId, dayId int) error {
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
	err = act.DateRepository.AddRecordToDate(rec, day)
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
	err = act.DateRepository.Save(day)
	if err != nil {
		act.Logger.Log("An error occurred while day added")
		return err
	}
	act.Logger.Log("Day added successfully")

	return nil
}

func (act *Actions) ShowDayRecords(id int) ([]entity.Record, error) {
	day, err := act.getDayById(id)
	records, err := act.DateRepository.ShowDayRecords(day)
	if err != nil {
		act.Logger.Log("An error occurred while show day records")
		return []entity.Record{}, err
	}
	return records, nil
}

func (act *Actions) getDayById(id int) (entity.Date, error) {
	day, err := act.DateRepository.FindById(uint64(id))
	if err != nil {
		act.Logger.Log("An error occurred while get day")
		return entity.Date{}, err
	}
	return day, nil
}
