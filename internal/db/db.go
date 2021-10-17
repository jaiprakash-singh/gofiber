package db

type Connection interface {
	Connect() string
}

type Database interface {
	List() string
}

type Collection interface {
	Init()
	Insert()
	Update()
	Find()
	Delete()
}
