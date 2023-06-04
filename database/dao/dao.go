package dao

type DAO interface {
	FindAll() ([]interface{}, error)
	FindById(id int64) (interface{}, error)
	Create(entity interface{}) (interface{}, error)
	Update(entity interface{}) (interface{}, error)
	Delete(id int64) error
}
