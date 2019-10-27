package mem_repository

type MemHandler interface {
	Execute(i interface{})
}

type Repo struct {
	handler MemHandler
}

func (r *Repo) createRepo() *Repo {
	r.handler.Execute()
}
