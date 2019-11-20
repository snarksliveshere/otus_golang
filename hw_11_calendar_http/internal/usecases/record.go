package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"
)

func (act *Actions) AddRecord(title, description string) (rec entity.Record, err error) {
	id := helpers.MakeTimestampId()
	rec = entity.Record{
		Id:          id,
		Title:       title,
		Description: description,
	}

	err = act.RecordRepository.Save(rec)
	if err != nil {
		act.Logger.Info("An error occurred while record added")
		return rec, err
	}
	act.Logger.Info("Record added successfully")

	return rec, nil
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
