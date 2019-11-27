package pg_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/pg_repository/pg_models"
)

func (d *DateRepo) FindByDay(date string) (entity.Date, error) {
	err := d.db.Model(d.row).
		Column("id", "date", "description", "is_celebration").
		Where("date = ?", date).
		Select()
	if err != nil {
		return entity.Date{}, err
	}

	day := entity.Date{
		Id:            d.row.Id,
		Day:           d.row.Date,
		Description:   d.row.Description,
		IsCelebration: d.row.IsCelebration,
		Events:        nil,
	}
	return day, nil
}

func (d *DateRepo) Save(date entity.Date) (uint32, error) {
	m := pg_models.Calendar{
		Date:          date.Day,
		Description:   date.Description,
		IsCelebration: date.IsCelebration,
	}
	_, err := d.db.Model(&m).
		OnConflict("(date) DO UPDATE").
		Set("description = EXCLUDED.description").
		Set("is_celebration = EXCLUDED.is_celebration").
		Insert()

	if err != nil {
		return 0, err
	}

	return m.Id, nil
}
