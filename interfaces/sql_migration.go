package interfaces

type Migrater interface {
	Migrate() error
}