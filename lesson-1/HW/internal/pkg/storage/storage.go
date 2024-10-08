package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Value struct {
	s    string
	f    float64
	kind string
}

type Storage struct {
	storage map[string]Value
	logger  *zap.Logger
}

func InitStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}
	defer logger.Sync()
	logger.Info("storage initialized")
	return Storage{
		storage: make(map[string]Value),
		logger:  logger,
	}, nil
}

func (storage Storage) Set(key, value string) {
	defer storage.logger.Sync()
	var val Value
	val.s = value
	if f, err := strconv.ParseFloat(value, 64); err == nil {
		val.f = f
		val.kind = "D"
	} else {
		val.kind = "S"
	}

	storage.storage[key] = val
	storage.logger.Info("Set key value")
}

func (storage Storage) Get(key string) *string {
	defer storage.logger.Sync()
	out, ok := storage.storage[key]
	storage.logger.Info("Get key value")
	if ok {
		return &out.s
	}
	return nil
}

func (storage Storage) GetKind(key string) string {
	defer storage.logger.Sync()
	out, ok := storage.storage[key]
	storage.logger.Info("Get kind key value")
	if !ok {
		return "N"
	}
	return out.kind
}
