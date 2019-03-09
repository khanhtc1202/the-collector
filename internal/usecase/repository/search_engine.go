package repository

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
)

type SearchEnginesRepository interface {
	FetchData(keyword domain.Keyword) (*domain.SearchEnginePool, error)
}