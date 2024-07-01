/*
 3. У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде

в таблицу вида
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
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
	results := make([]TableLine, 0)

	for _, val := range data.ListOfResults {

		currentLine := TableLine{Name: students[val.Student_id].Name,
			Grade:  students[val.Student_id].Grade,
			Object: objects[val.Obj_id].Name,
			Result: val.Result,
		}
		results = append(results, currentLine)
	}
	fmt.Println("______________________________________________")
	fmt.Println("|Student name  | Grade | Object    |   Result|")
	fmt.Println("_____________________________________________")

	for _, line := range results {
		fmt.Printf("|       %s     |    %d |      %s  |     %d  |\n",
			line.Name, line.Grade, line.Object, line.Result)
	}

}
