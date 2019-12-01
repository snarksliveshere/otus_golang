package pg_repository

import (
	"errors"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
)

func (r *RecordRepo) FindById(id uint64) (entity.Record, error) {

	rec := entity.Record{
		Id:          id,
		Title:       "Title1",
		Description: "Desc",
	}
	return rec, nil
}

func (r *RecordRepo) GetEventsByDay(dateFk uint32) ([]entity.Record, error) {
	err := r.db.Model(&r.rows).
		Column("time", "title", "description", "time", "id").
		Where("date_fk = ?", dateFk).
		Select()
	if err != nil {
		return nil, err
	}

	var records []entity.Record

	for _, v := range r.rows {
		rec := entity.Record{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Time:        v.Time,
		}
		records = append(records, rec)
	}

	if len(records) == 0 {
		return []entity.Record{}, errors.New("there are no records in this day")
	}

	return records, nil
}

func (r *RecordRepo) Delete(record entity.Record) error {

	return nil
}

func (r *RecordRepo) Edit(record entity.Record) error {

	return nil
}

func (r *RecordRepo) Show() []entity.Record {

	return []entity.Record{}
}

func (r *RecordRepo) Save(record entity.Record) error {

	return nil
}
