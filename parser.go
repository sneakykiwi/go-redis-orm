package redis_orm


type Parser interface {
	ParseModels(models ...interface{}) error
}