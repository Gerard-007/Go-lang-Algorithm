package main

import "fmt"

/*
3. Collect the scores for every students and every subjects.
4. Constraint: Score must be between 0 and 100.
5. Display class summary after input collection.
*/

func main() {
	//Ask the teacher for the number of students.
	var numberOfStudents int
	var numberOfSubjects int

	fmt.Print("Enter number of students: ")
	fmt.Scan(&numberOfStudents)

	fmt.Print("Enter number of subjects: ")
	fmt.Scan(&numberOfSubjects)

	fmt.Println("Saving >>>>>>>>>>>>>>>>>")

	students := make([]string, numberOfStudents)
	grades := make([][]int, numberOfStudents)
	for i := range grades {
		grades[i] = make([]int, numberOfSubjects)
	}

	var studentName string
	var subjectGrade int
	for i := 0; i < numberOfStudents; i++ {
		fmt.Printf("Enter name of student->%d: ", i+1)
		fmt.Scan(&studentName)
		for j := 0; j < numberOfSubjects; j++ {
			fmt.Printf("What did %s scored in subject->%d: ", studentName, j+1)
			fmt.Scan(&subjectGrade)
			if subjectGrade >= 0 && subjectGrade <= 100 {
				grades[i][j] = subjectGrade
			} else {
				fmt.Print("Invalid score entered must be between 0-100")
			}
		}
		students[i] = studentName
	}

	fmt.Println("Students:", students)
	fmt.Println("Grades:", grades)
}
