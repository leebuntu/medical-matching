package sets

import (
	"database/sql"
	"math"
)

type PriorityInjection struct {
	db *sql.DB
}

func NewPriorityInjection(db *sql.DB) *PriorityInjection {
	return &PriorityInjection{db: db}
}

func (p *PriorityInjection) countSelected(combination []int) int {
	count := 0
	for _, v := range combination {
		if v == 1 {
			count++
		}
	}
	return count
}

func (p *PriorityInjection) generateValidCombinations(n int) [][]int {
	var combinations [][]int

	combinations = append(combinations, make([]int, n))

	for i := 1; i < int(math.Pow(2, float64(n))); i++ {
		combination := make([]int, n)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				combination[j] = 1
			}
		}

		if count := p.countSelected(combination); count >= 1 && count <= 3 {
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

func (p *PriorityInjection) alreadyInjected(finalCount int) (bool, error) {
	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM priority_set").Scan(&count)
	if err != nil {
		return false, err
	}

	return count == finalCount, nil
}

func (p *PriorityInjection) InjectPriority(count int, finalCount int) error {
	alreadyInjected, err := p.alreadyInjected(finalCount)
	if err != nil {
		return err
	}
	if alreadyInjected {
		return nil
	}

	combinations := p.generateValidCombinations(count)

	stmt, err := p.db.Prepare("INSERT INTO priority_set (priority_1, priority_2, priority_3, priority_4, priority_5)VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, combination := range combinations {
		_, err := stmt.Exec(combination[0], combination[1], combination[2], combination[3], combination[4])
		if err != nil {
			return err
		}
	}

	return nil
}
