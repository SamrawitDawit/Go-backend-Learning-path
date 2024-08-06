// Student grade calculator
package main

import "fmt"

//defining data type to store student data
type student struct {
	name              string
	subject_and_grade map[string]float64
}

//defining a method to calculate average grade
func (s student) averageGrade() float64 {
	var sum float64
	for _, grade := range s.subject_and_grade {
		sum += grade
	}
	return sum / float64(len(s.subject_and_grade))
}

func main() {
	test()
	//Implementing a loop to ask for student data
	var studentName string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&studentName)
	studentData := student{
		name:              studentName,
		subject_and_grade: make(map[string]float64),
	}
	var numSubjects int
	fmt.Println("Enter number of subjects: ")
	fmt.Scanln(&numSubjects)
	for j := 0; j < numSubjects; j++ {
		var subject string
		fmt.Println("Enter subject: ")
		fmt.Scanln(&subject)
		var grade float64
		fmt.Println("Enter grade: ")
		fmt.Scanln(&grade)
		//validating the grade
		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade.")
			return
		}
		//storing the student data
		studentData.subject_and_grade[subject] = grade
	}
	//Displaying the student's name, individual grades and average grade
	fmt.Println("Student Name:", studentData.name)
	fmt.Println("Grades:")
	for subject, grade := range studentData.subject_and_grade {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", studentData.averageGrade())
}

//testing the code
func test() {
	//Test case 1
	studentData := student{
		name:              "Samri",
		subject_and_grade: map[string]float64{"Maths": 95, "Physics": 92, "English": 95},
	}
	if studentData.averageGrade() != 94 {
		fmt.Println("Test case 1 failed.")
		return
	}
	//Test case 2
	studentData = student{
		name:              "Redi",
		subject_and_grade: map[string]float64{"Maths": 100, "Physics": 100, "English": 100},
	}
	if studentData.averageGrade() != 100 {
		fmt.Println("Test case 2 failed.")
		return
	}
	fmt.Println("All test cases passed.")
}