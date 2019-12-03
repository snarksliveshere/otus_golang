package mem_repository

import "github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"

func (r *EventRepo) FindById(id uint64) (entity.Event, error) {
	r.handler.Execute("find by id")
	rec := entity.Event{
		Id:          id,
		Title:       "Title1",
		Description: "Desc",
	}
	return rec, nil
}

func (r *EventRepo) Delete(event entity.Event) error {
	r.handler.Execute("delete")
	return nil
}

func (r *EventRepo) Edit(event entity.Event) error {
	r.handler.Execute("edit")
	return nil
}

func (r *EventRepo) Show() []entity.Event {
	r.handler.Execute("show")
	return []entity.Event{}
}

func (r *EventRepo) Save(event entity.Event) error {

	r.handler.Execute("save")
	return nil
}
