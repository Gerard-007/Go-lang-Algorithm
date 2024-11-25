package phaseGate

import "fmt"

/*
1. Ask the teacher for the number of students.
2. Ask the teacher for the number of subjects.
3. Collect the scores for every students and every subjects.
4. Constraint: Score must be between 0 and 100.
5. Display class summary after input collection.
*/

func main() {
	//Ask the teacher for the number of students.
	var numberOfStudents int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&numberOfStudents)
	fmt.Printf("Number of students %d", numberOfStudents)
}
