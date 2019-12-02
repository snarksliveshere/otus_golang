package pg_repository

import (
	"errors"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/pg_repository/pg_models"
)

func (r *RecordRepo) FindById(id uint64) (entity.Record, error) {
	err := r.db.Model(r.row).
		Column("time", "title", "description", "time", "id", "date_fk").
		Where("id = ?", id).
		Select()
	if err != nil {
		return entity.Record{}, err
	}

	rec := entity.Record{
		Id:          r.row.Id,
		Title:       r.row.Title,
		Description: r.row.Description,
		Time:        r.row.Time,
		DateFk:      r.row.DateFk,
	}

	return rec, nil
}

func (r *RecordRepo) Delete(record entity.Record) error {
	_, err := r.db.Model(r.row).
		Where("id = ?", record.Id).
		Delete()

	if err != nil {
		return err
	}
	return nil
}

func (r *RecordRepo) GetEventsByDay(dateFk uint32) ([]entity.Record, error) {
	err := r.db.Model(&r.rows).
		Column("time", "title", "description", "time", "id", "date_fk").
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

func (r *RecordRepo) Save(rec entity.Record) (uint64, error) {
	m := pg_models.Event{
		Title:       rec.Title,
		Description: rec.Description,
		Time:        rec.Time,
		DateFk:      rec.DateFk,
	}
	_, err := r.db.Model(&m).
		OnConflict("(date_fk, time) DO UPDATE").
		Set("title = EXCLUDED.title").
		Set("description = EXCLUDED.description").
		Insert()

	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

func (r *RecordRepo) Edit(record entity.Record) error {

	return nil
}

func (r *RecordRepo) Show() []entity.Record {

	return []entity.Record{}
}
