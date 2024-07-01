/*
 2. Подсчет голосов.

Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]
*/
package main

import (
	"fmt"
	"sort"
)

type Candidate struct {
	Name  string
	Votes int
}

func main() {

	pool := []string{
		"Jane",
		"John",
		"Ann",
		"Kate",
		"Peter",
		"Kate",
		"Ann",
		"Ann",
		"Helen",
	}

	votes := Votes(pool)

	candy := ConvertVotesToResult(votes)

	printList(candy)

	sort.SliceStable(candy, func(i, j int) bool {
		return candy[i].Votes > candy[j].Votes
	})

	printList(candy)

}

func Votes(poles []string) map[string]int {

	resultMap := make(map[string]int)

	for _, name := range poles {

		if _, ok := resultMap[name]; !ok {
			resultMap[name] = 1
			continue
		}

		val := resultMap[name] + 1
		resultMap[name] = val

	}

	return resultMap

}

func ConvertVotesToResult(votes map[string]int) []Candidate {
	result := make([]Candidate, 0)

	for name, vote := range votes {
		result = append(result, Candidate{Name: name, Votes: vote})

	}

	return result
}

func printList(c []Candidate) {

	for i, val := range c {
		fmt.Printf("№ %v, Name: %s , Votes: %v \n", i+1, val.Name, val.Votes)
	}

}
