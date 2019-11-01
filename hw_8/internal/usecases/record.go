package usecases

import "github.com/snarskliveshere/otus_golang/hw_8/entity"

func (act *Actions) AddRecord(id int) error {
	rec, err := act.getRecordById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.RecordRepository.Save(rec)
	if err != nil {
		act.Logger.Log("An error occurred while record added")
		return err
	}
	act.Logger.Log("Record added successfully")

	return nil
}

func (act *Actions) EditRecord(id int) error {
	rec, err := act.getRecordById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.RecordRepository.Edit(rec)
	if err != nil {
		act.Logger.Log("An error occurred while record updating")
		return err
	}
	return nil
}

func (act *Actions) DeleteRecord(id int) error {
	rec, err := act.getRecordById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.RecordRepository.Delete(rec)
	if err != nil {
		act.Logger.Log("An error occurred while record deleting")
		return err
	}
	return nil
}

func (act *Actions) getRecordById(id int) (entity.Record, error) {
	record, err := act.RecordRepository.FindById(uint64(id))
	if err != nil {
		act.Logger.Log("An error occurred while get record")
		return entity.Record{}, err
	}
	return record, nil
}
