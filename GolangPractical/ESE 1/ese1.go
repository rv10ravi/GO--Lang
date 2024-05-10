//create a user-friendly student management systemusing go . The system will have following features:
//1. Add student to the system
//2. Delete student from the system
//3. Display all students
//4. Search student by ID
//5.Add multiple students at a time 

package main

import (
	"fmt"
)


// struct for student
type Student struct {
	ID     int
	Name   string
	Course string
	Marks  int
}

// interface for student manager
type StudentManager interface {
	AddStudent(s Student) error
	DeleteStudent(ID int) error
	DisplayAllStudents() ([]Student, error)
	SearchStudent(ID int) (Student, error)
	AddMultipleStudents(s []Student) error
}

// error for student not found
type StudentNotFoundError struct {
	ID int
}


func (e StudentNotFoundError) Error() string {
	return fmt.Sprintf("Student with ID %d not found", e.ID)
}

// error for duplicate student
type DuplicateStudentError struct {
	ID int
}

func (e DuplicateStudentError) Error() string {
	return fmt.Sprintf("Student with ID %d already exists", e.ID)
}

// student manager implementation
type StudentManagerImpl struct {
	students map[int]Student
}

// create new student manager
func NewStudentManager() StudentManager {
	return &StudentManagerImpl{
		students: make(map[int]Student),
	}
}

// add student to student manager
func (sm *StudentManagerImpl) AddStudent(s Student) error {
	if _, ok := sm.students[s.ID]; ok {
		return DuplicateStudentError{s.ID}
	}
	sm.students[s.ID] = s
	return nil
}

// delete student from student manager
func (sm *StudentManagerImpl) DeleteStudent(ID int) error {
	if _, ok := sm.students[ID]; !ok {
		return StudentNotFoundError{ID}
	}
	delete(sm.students, ID)
	return nil
}

// display all students in student manager
func (sm *StudentManagerImpl) DisplayAllStudents() ([]Student, error) {
	var students []Student
	for _, s := range sm.students {
		students = append(students, s)
	}
	return students, nil
}

// search student by ID in student manager
func (sm *StudentManagerImpl) SearchStudent(ID int) (Student, error) {
	s, ok := sm.students[ID]
	if !ok {
		return Student{}, StudentNotFoundError{ID}
	}
	return s, nil
}

// add multiple students to student manager
func (sm *StudentManagerImpl) AddMultipleStudents(s []Student) error {
	for _, student := range s {
		if _, ok := sm.students[student.ID]; ok {
			return DuplicateStudentError{student.ID}
		}
		sm.students[student.ID] = student
	}
	return nil
}

// display menu
func displayMenu() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Welcome to Student Management System")
	fmt.Println("---------------------------------------------")
	fmt.Println("Menu:")
	fmt.Println("1. Add Student")
	fmt.Println("2. Delete Student")
	fmt.Println("3. Display All Students")
	fmt.Println("4. Search Student")
	fmt.Println("5. Add Multiple Students")
	fmt.Println("6. Exit")
	fmt.Println("---------------------------------------------")
}

func main() {
	sm := NewStudentManager()
	var choice int

	for {
		displayMenu()
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		

		switch choice {
		case 1:
			var id, marks int
			var name, course string
			fmt.Print("Enter student ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter student name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter student course: ")
			fmt.Scanln(&course)
			fmt.Print("Enter student marks: ")
			fmt.Scanln(&marks)
			if err := sm.AddStudent(Student{ID: id, Name: name, Course: course, Marks: marks}); err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Student added successfully")
            }
		case 2:
			var id int
			fmt.Print("Enter ID of student to delete: ")
			fmt.Scanln(&id)
			if err := sm.DeleteStudent(id); err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Student Deleted Successfully")
            }
		
		case 3:
			students, err := sm.DisplayAllStudents()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("All Students:")
				for _, student := range students {
					fmt.Println("---------------------------------------------")
					
					fmt.Printf("ID:%d", student.ID)
					fmt.Printf("\nName:%s", student.Name)
					fmt.Printf("\nCourse:%s", student.Course)
					fmt.Printf("\nMarks:%d", student.Marks)

					fmt.Println("\n---------------------------------------------")
				}

			}
		case 4:
			var id int
			fmt.Print("Enter ID of student to search: ")
			fmt.Scanln(&id)
			student, err := sm.SearchStudent(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Found Student:")
				fmt.Printf("ID:%d", student.ID)
				fmt.Printf("\nName:%s", student.Name)
				fmt.Printf("\nCourse:%s", student.Course)
				fmt.Printf("\nMarks:%d", student.Marks)
				fmt.Println("\n---------------------------------------------")
			}
		case 5:
			var count int
			fmt.Print("Enter number of students to add: ")
			fmt.Scanln(&count)
			var students []Student
			for i := 0; i < count; i++ {
				var id, marks int
				var name, course string
				fmt.Printf("Enter ID for student %d: ", i+1)
				fmt.Scanln(&id)
				fmt.Printf("Enter name for student %d: ", i+1)
				fmt.Scanln(&name)
				fmt.Printf("Enter course for student %d: ", i+1)
				fmt.Scanln(&course)
				fmt.Printf("Enter marks for student %d: ", i+1)
				fmt.Scanln(&marks)
				students = append(students, Student{ID: id, Name: name, Course: course, Marks: marks})
			}
		
			if err := sm.AddMultipleStudents(students); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Student added successfully")
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}






