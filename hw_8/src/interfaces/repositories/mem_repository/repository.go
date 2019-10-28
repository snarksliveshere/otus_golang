package mem_repository

type MemHandler interface {
	Execute(i interface{})
}

type DayRepo struct {
	Repo
}

type RecordRepo struct {
	Repo
}

type Repo struct {
	handler MemHandler
	//DayRepo entity.DayRepository
	//RecordRepo entity.RecordRepository
}

func (r *Repo) createRepo() *Repo {
	return new(Repo)
}
