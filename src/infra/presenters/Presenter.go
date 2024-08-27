package presenters

type Presenter interface {
	Format(data interface{}) interface{}
}
