package main

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Object struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	Obj_id     int `json:"object_id"`
	Student_id int `json:"student_id"`
	Result     int `json:"result"`
}

type Students struct {
	ListOfStudents []Student
}

type Objects struct {
	ListOfObjects []Object
}

type Results struct {
	ListOfResults []Result
}

type Data struct {
	ListOfStudents []Student `json:"students"`
	ListOfObjects  []Object  `json:"objects"`
	ListOfResults  []Result  `json:"results"`
}

type TableLine struct {
	Name   string
	Grade  int
	Object string
	Result int
}
