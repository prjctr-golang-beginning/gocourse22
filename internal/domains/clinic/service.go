package clinic

import (
	"github.com/samber/do"
	"gocourse22/pkg/extend"
	"sync"
)

const visits int = 110
const weeks int = 3

func ProvideService(_ *do.Injector) (*Service, error) {
	return NewService(), nil
}
func NewService() *Service {
	return &Service{}
}

type Service struct {
}

type GroupedVisits struct {
	Week  int
	Count int
}

func (s *Service) GroupPatientsVisits() []GroupedVisits {
	visitsCount := make(map[int]int)
	workers := 7

	chanStrategy(workers, visitsCount)

	var res []GroupedVisits
	for week, count := range visitsCount {
		res = append(res, GroupedVisits{week + 1, count})
	}

	return res
}

func (s *Service) DeleteClinic() error {
	return extend.NewFormattedError(1, `Clinic deletion is impossible`, nil)
}

func muxStrategy(workers int, visitsResult map[int]int) {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(week int) {
			defer wg.Done()
			for j := 0; j < (visits / weeks); j++ {
				mutex.Lock()
				visitsResult[week]++
				mutex.Unlock()
			}
		}(i % weeks)
	}

	wg.Wait()
}

func chanStrategy(workers int, visitsResult map[int]int) {
	results := make(chan map[int]int)

	for i := 0; i < workers; i++ {
		go func(week int) {
			result := make(map[int]int)
			for j := 0; j < (visits / weeks); j++ {
				result[week]++
			}
			results <- result
		}(i % weeks)
	}

	for i := 0; i < workers; i++ {
		for week, count := range <-results {
			visitsResult[week] += count
		}
	}
}
