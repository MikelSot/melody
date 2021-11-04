package interfaces

type CreateTabler interface {
	CreateTable() error
}


type InsertDataer interface {
	InsertData() error
}
