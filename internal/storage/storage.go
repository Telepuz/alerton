package storage

type Storage interface {
	IsCooldown(name string) bool
	ClearByTTL()
}
