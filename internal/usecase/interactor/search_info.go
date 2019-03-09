package interactor

import (
	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/usecase/presenter"
	"github.com/khanhtc1202/boogeyman/internal/usecase/repository"
	"github.com/pkg/errors"
)

type InfoSearch struct {
	ranker    *domain.Ranker
	poolRepo  repository.SearchEnginesRepository
	presenter presenter.TextPresenter
}

func NewInfoSearch(
	presenter presenter.TextPresenter,
	poolRepo repository.SearchEnginesRepository,
) *InfoSearch {
	return &InfoSearch{
		ranker:    domain.NewRanker(),
		poolRepo:  poolRepo,
		presenter: presenter,
	}
}

func (i *InfoSearch) Search(
	queryString string,
	strategy domain.FilterStrategyType,
) error {
	// fetch data from search engines
	resultPool, err := i.poolRepo.FetchData(domain.NewKeyword(queryString))
	if err != nil {
		return errors.Wrap(err, "Error on fetch data from pool!\n")
	}

	// merge by strategy
	var queryResult *domain.QueryResults
	switch strategy {
	case domain.TOP:
		queryResult, err = i.ranker.Top(resultPool)
		break
	case domain.CROSS:
		queryResult, err = i.ranker.CrossMatch(resultPool)
		break
	case domain.ALL:
		queryResult, err = i.ranker.All(resultPool,
			config.GetConfig().RankerConf.MaxReturnItems)
		break
	default:
		queryResult, err = i.ranker.CrossMatch(resultPool)
		break
	}
	if err != nil {
		return err
	}

	// printout
	if err = i.presenter.PrintList(queryResult); err != nil {
		return err
	}

	return nil
}
