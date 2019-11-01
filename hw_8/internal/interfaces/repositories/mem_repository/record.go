package mem_repository

import "github.com/snarskliveshere/otus_golang/hw_8/entity"

func (r *RecordRepo) FindById(id uint64) (entity.Record, error) {
	r.handler.Execute("find by id")
	rec := entity.Record{
		Id:          id,
		Title:       "Title1",
		Description: "Desc",
	}
	return rec, nil
}

func (r *RecordRepo) Delete(record entity.Record) error {
	r.handler.Execute("delete")
	return nil
}

func (r *RecordRepo) Edit(record entity.Record) error {
	r.handler.Execute("edit")
	return nil
}

func (r *RecordRepo) Show() []entity.Record {
	r.handler.Execute("show")
	return []entity.Record{}
}

func (r *RecordRepo) Save(record entity.Record) error {
	r.handler.Execute("save")
	return nil
}
