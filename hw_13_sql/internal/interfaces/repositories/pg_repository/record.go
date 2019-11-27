package pg_repository

import (
	"errors"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/pg_repository/pg_models"
)

func (r *EventRepo) FindById(id uint64) (entity.Event, error) {
	err := r.db.Model(r.row).
		Column("time", "title", "description", "time", "id", "date_fk").
		Where("id = ?", id).
		Select()
	if err != nil {
		return entity.Event{}, err
	}

	rec := entity.Event{
		Id:          r.row.Id,
		Title:       r.row.Title,
		Description: r.row.Description,
		Time:        r.row.Time,
		DateFk:      r.row.DateFk,
	}

	return rec, nil
}

func (r *EventRepo) Delete(event entity.Event) error {
	_, err := r.db.Model(r.row).
		Where("id = ?", event.Id).
		Delete()

	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepo) Edit(event entity.Event) error {
	_, err := r.db.Model(r.row).
		Set("title = ?", event.Title).
		Set("description = ?", event.Description).
		Where("id = ?", event.Id).
		Update()

	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepo) GetEventsByDay(dateFk uint32) ([]entity.Event, error) {
	err := r.db.Model(&r.rows).
		Column("time", "title", "description", "time", "id", "date_fk").
		Where("date_fk = ?", dateFk).
		Select()
	if err != nil {
		return nil, err
	}

	var events []entity.Event

	for _, v := range r.rows {
		rec := entity.Event{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Time:        v.Time,
		}
		events = append(events, rec)
	}

	if len(events) == 0 {
		return []entity.Event{}, errors.New("there are no events in this day")
	}

	return events, nil
}

func (r *EventRepo) GetEventsByDateInterval(from, till string) ([]entity.Event, error) {
	err := r.db.Model(&r.rows).
		Column("event.time", "event.title", "event.description", "event.time", "event.id", "event.date_fk").
		Join("JOIN calendar.calendar ON event.date_fk = calendar.id").
		Where("calendar.date >= ?", from).
		Where("calendar.date <= ?", till).
		Select()
	if err != nil {
		return nil, err
	}

	var events []entity.Event

	for _, v := range r.rows {
		rec := entity.Event{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Time:        v.Time,
		}
		events = append(events, rec)
	}

	if len(events) == 0 {
		return []entity.Event{}, errors.New("there are no events in this day")
	}

	return events, nil
}

func (r *EventRepo) Save(rec entity.Event) (uint64, error) {
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

func (r *EventRepo) Show() []entity.Event {

	return []entity.Event{}
}
