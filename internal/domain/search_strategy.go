package domain

/*
Filter Strategy Type
*/
type FilterStrategyType int

const (
	ALL FilterStrategyType = iota
	TOP
	CROSS
)

func (s FilterStrategyType) String() string {
	switch s {
	case ALL:
		return "ALL"
	case CROSS:
		return "CROSS"
	case TOP:
		return "TOP"
	}
	return "ALL"
}

/*
Filter result strategies
*/
type FilterSearch interface {
	Filter() (*QueryResults, error)
}

type topResultsByEngines struct {
	engines *SearchEnginePool
}

func TopResultsByEngines(engines *SearchEnginePool) *topResultsByEngines {
	return &topResultsByEngines{
		engines: engines,
	}
}

func (t *topResultsByEngines) Filter() (*QueryResults, error) {
	topResults := EmptyQueryResult()
	for _, searchEngine := range *t.engines {
		if searchEngine.GetQueryResults().Length() > 0 {
			topResults.Add(searchEngine.TopResult())
		}
	}
	return topResults, nil
}

type crossMatchByEngines struct {
	engines *SearchEnginePool
}

func CrossMatchByEngines(engines *SearchEnginePool) *crossMatchByEngines {
	return &crossMatchByEngines{
		engines: engines,
	}
}

func (c *crossMatchByEngines) Filter() (*QueryResults, error) {
	crossMatchedResults := EmptyQueryResult()
	for _, searchEngine := range *c.engines {
		crossMatchedResults.Concatenate(searchEngine.GetQueryResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

type mergeResultsByEngines struct {
	engines        *SearchEnginePool
	maxReturnItems int
}

func MergeResultsByEngines(engines *SearchEnginePool, limit int) *mergeResultsByEngines {
	return &mergeResultsByEngines{
		engines:        engines,
		maxReturnItems: limit,
	}
}

func (m *mergeResultsByEngines) Filter() (*QueryResults, error) {
	allResults := EmptyQueryResult()
	for _, searchEngine := range *m.engines {
		allResults.Concatenate(searchEngine.GetQueryResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(m.maxReturnItems), nil
}
