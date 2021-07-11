package main

import (
	"fmt"
)

/*
interface 的一个基础demo
*/

type People interface {
	name(name string)
	age(age int)
}

type Student struct {
	score uint
}

type Teacher struct {
}

func (s *Student) name(nm string) {
	fmt.Printf("student name is %s\n", nm)
}

func (s *Student) age(age int) {
	fmt.Println("student age is :", age)
}

func (t *Teacher) name(nm string) {
	fmt.Println("Teacher name is  ", nm)
}

func (t *Teacher) age(age int) {
	fmt.Println("Teacher age is ", age)
}

var conn People

func GetStudentConn() People {
	conn = &Student{
		score: 100,
	}
	return conn
}

func GetTeacherConn() People {
	conn = &Teacher{}
	return conn
}

func main() {
	stuConn := GetStudentConn()
	stuConn.age(100)

	teaConn := GetTeacherConn()
	teaConn.name("tea")
}
