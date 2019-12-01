package usecases

import (
	"errors"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"time"
)

func (act *Actions) CreateEvent(title, description string, t time.Time) (uint64, error) {
	date := t.Format(config.TimeLayout)
	day, err := act.DateRepository.FindByDay(date)
	if err != nil {
		return 0, err
	}
	var dateId uint32
	if day.Id == 0 {
		newDateId, err := act.CreateEmptyDate(date)
		if err != nil {
			return 0, err
		}
		dateId = newDateId
	} else {
		dateId = day.Id
	}
	rec, err := act.AddRecord(title, description, dateId, t)

	return rec.Id, nil
}

func (act *Actions) AddRecord(title, description string, dateFk uint32, t time.Time) (rec entity.Record, err error) {
	rec = entity.Record{
		Title:       title,
		Description: description,
		Time:        t,
		DateFk:      dateFk,
	}

	recId, err = act.RecordRepository.Save(rec)
	if err != nil {
		act.Logger.Info("An error occurred while record added")
		return rec, err
	}
	act.Logger.Info("Record added successfully")

	return rec, nil
}

func (act *Actions) GetEventsByDay(date string) ([]entity.Record, error) {
	day, err := act.DateRepository.FindByDay(date)
	if err != nil {
		act.Logger.Info("An error occurred while record added")
		return []entity.Record{}, err
	}
	if day.Id == 0 {
		act.Logger.Info("An empty ady")
		return []entity.Record{}, errors.New("i cant find this day")
	}
	records, err := act.RecordRepository.GetEventsByDay(day.Id)
	if err != nil {
		act.Logger.Info("An error occurred while record added")
		return []entity.Record{}, err
	}

	return records, nil
}

func (act *Actions) EditRecord(id uint64) error {
	rec, err := act.getRecordById(id)
	if err != nil {
		act.Logger.Info(err.Error())
		return err
	}
	err = act.RecordRepository.Edit(rec)
	if err != nil {
		act.Logger.Info("An error occurred while record updating")
		return err
	}
	return nil
}

func (act *Actions) DeleteRecord(id uint64) error {
	rec, err := act.getRecordById(id)
	if err != nil {
		act.Logger.Info(err.Error())
		return err
	}
	err = act.RecordRepository.Delete(rec)
	if err != nil {
		act.Logger.Info("An error occurred while record deleting")
		return err
	}
	return nil
}

func (act *Actions) getRecordById(id uint64) (entity.Record, error) {
	record, err := act.RecordRepository.FindById(id)
	if err != nil {
		act.Logger.Info("An error occurred while get record")
		return entity.Record{}, err
	}
	return record, nil
}
