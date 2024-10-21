package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Value struct {
	s    string
	i    int
	kind string
}

type Storage struct {
	storage map[string]Value
	Logger  *zap.Logger
}

const (
	KindDigit   = "D"
	KindString  = "S"
	KindUnknown = ""
)

func InitStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}
	defer logger.Sync()
	// logger.Info("storage initialized")
	return Storage{
		storage: make(map[string]Value),
		Logger:  logger,
	}, nil
}

func (storage Storage) Set(key, value string) {
	var val Value
	val.s = value
	if i, err := strconv.Atoi(value); err == nil {
		val.i = i
		val.kind = KindDigit
	} else {
		val.kind = KindString
	}

	storage.storage[key] = val
	// storage.Logger.Info("Set key value")
}

func (storage Storage) Get(key string) *string {
	out, ok := storage.storage[key]
	// storage.Logger.Info("Get key value")
	if ok {
		return &out.s
	}
	return nil
}

func (storage Storage) GetKind(key string) string {

	out, ok := storage.storage[key]
	// storage.Logger.Info("Get kind key value")
	if !ok {
		return KindUnknown
	}
	return out.kind
}

func (storage Storage) getTest(key string) Value {
	out, ok := storage.storage[key]
	if !ok {
		return Value{}
	}
	return out
}
