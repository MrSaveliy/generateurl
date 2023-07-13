package store

import (
	"log"
	"name/generator/internal/pkg/repository"
	"name/generator/internal/pkg/store/database"
	in_memory "name/generator/internal/pkg/store/in-memory"
	"os"
	"sync"
)

const (
	dbType       = "database"
	inMemoryType = "in-memory"
)

var (
	storeImpl repository.Repository
	once      = &sync.Once{}
)

func GetStore() repository.Repository {
	once.Do(func() {
		switch os.Getenv("DB_TYPE") {
		case dbType:
			storeImpl = database.GetStore()
		case inMemoryType:
			storeImpl = in_memory.GetStore()
		default:
			log.Fatal("invalid memory type")
		}
	})
	return storeImpl
}
