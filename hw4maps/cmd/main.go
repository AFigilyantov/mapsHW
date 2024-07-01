/* 4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:
________________
Math	 | Mean
________________
 9 grade | 4.5
10 grade | 5
11 grade | 3.5
________________
mean     | 4		- среднее значение среди всех учеников
________________
________________
Biology	 | Mean
________________
...
Вводные данные представлены в файле dz3.json
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../dz3.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	data := Data{}

	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	students := make(map[int]Student)

	for _, val := range data.ListOfStudents {
		students[val.Id] = Student{Id: val.Id,
			Name:  val.Name,
			Grade: val.Grade}
	}

	objects := make(map[int]Object)

	for _, val := range data.ListOfObjects {
		objects[val.Id] = Object{Id: val.Id,
			Name: val.Name}

	}

	results := make(map[string]map[int]float64)

	for _, val := range data.ListOfResults {

		if res, ok := results[objects[val.Obj_id].Name]; ok {
			currentVal := res[students[val.Student_id].Grade]
			currentVal = (float64(val.Result) + currentVal) / 2
			res[students[val.Student_id].Grade] = currentVal
			results[objects[val.Obj_id].Name] = res
			continue
		}

		if _, ok := results[objects[val.Obj_id].Name]; !ok {
			valueToResult := make(map[int]float64)
			valueToResult[students[val.Student_id].Grade] = float64(val.Result)
			results[objects[val.Obj_id].Name] = valueToResult
		}

	}

	for name, val := range results {
		fmt.Println("______________")
		fmt.Printf("%s| Mean \n", name)
		fmt.Println("______________")

		averageMean := 0.0
		var index int

		for grade, mean := range val {

			fmt.Printf("%d grade | %.1f \n", grade, mean)
			averageMean += mean
			index++

		}
		fmt.Println("______________")

		fmt.Printf("mean  |  %.1f \n", (averageMean / float64(index)))

		fmt.Println("______________")
	}

}
