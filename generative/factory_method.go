package generative

import "io"

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

/*func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( ... )
	case DiskStorage:
		return newDiskStorage( ... )
	default:
		return newTempStorage( ... )
	}
}*/
