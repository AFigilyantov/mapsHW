/*
 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.

Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	first := []int{1, 2, 3, 2}
	second := []int{3, 2}
	third := []int{}

	firstMap := JoinSlice(first, second, third)

	res := CleanMap(firstMap, first, second, third)

	fmt.Println(res)

}

func JoinSlice(args ...[]int) map[int]struct{} {

	resultMap := make(map[int]struct{})

	for _, arg := range args {

		for _, i := range arg {

			if _, ok := resultMap[i]; ok {
				continue
			}
			resultMap[i] = struct{}{}

		}

	}

	return resultMap
}

func CleanMap(m map[int]struct{}, args ...[]int) []int {
	currentMap := m
	for _, arg := range args {
		tempMap := make(map[int]struct{})

		for _, value := range arg {
			if _, ok := currentMap[value]; ok {
				tempMap[value] = struct{}{}
			}
		}
		currentMap = tempMap
	}

	resultSlice := make([]int, 0)

	for index := range currentMap {
		resultSlice = append(resultSlice, index)
	}

	sort.Ints(resultSlice)

	return resultSlice
}
