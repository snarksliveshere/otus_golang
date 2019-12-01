package pg_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/pg_repository/pg_models"
	"sync"
	"time"
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
		Records:       nil,
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
	//err := d.db.Insert(&m)

	if err != nil {
		return 0, err
	}

	return m.Id, nil
}

var (
	calendar     *entity.Calendar
	calendarOnce sync.Once
)

func createCalendar() *entity.Calendar {
	calendarOnce.Do(func() {
		calendar = &entity.Calendar{}
	})
	return calendar
}

func (d *DateRepo) GetCalendar() *entity.Calendar {
	return createCalendar()
}

func (d *DateRepo) AddDateToCalendar(day entity.Date) error {
	calendar := d.GetCalendar()
	calendar.Dates = append(calendar.Dates, &day)

	return nil
}

func (d *DateRepo) AddRecordToDate(record entity.Record, day *entity.Date) error {
	day.Records = append(day.Records, record)

	return nil
}

func (d *DateRepo) ShowDayRecords(day *entity.Date) ([]entity.Record, error) {

	return nil, nil
}

func (d *DateRepo) GetDateFromString(date string) (time.Time, error) {
	t, err := time.Parse(config.TimeLayout, date)
	if err != nil {
		d.logger.Info("Wrong incoming day pattern")
		return t, err
	}
	return t, nil
}
