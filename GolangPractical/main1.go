package main

import (
	"fmt"
	"time"
)

type patient struct {
	name    string
	age     int
	gender string
	contact int64
}

type appointment struct {
	date string
	time string
	medicalPractitioner string
	patientId int
}

var patients map[int]patient
var appointments map[int]appointment
var patientId int = 1
var appointmentId int = 1

func main() {
	patients = make(map[int]patient)
	appointments = make(map[int]appointment)
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("Patient Management System")
		fmt.Println("1. Add new patient")
		fmt.Println("2. Update patient details")
		fmt.Println("3. Schedule appointment")
		fmt.Println("4. Generate report")
		fmt.Println("5. Exit")
		var choice int
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addPatient()
		case 2:
			updatePatient()
		case 3:
			scheduleAppointment()
		case 4:
			generateReport()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addPatient() {
	var p patient
	fmt.Println("Enter name:")
	fmt.Scanln(&p.name)
	fmt.Println("Enter age:")
	fmt.Scanln(&p.age)
	fmt.Println("Enter gender:")
	fmt.Scanln(&p.gender)
	fmt.Println("Enter contact:")
	fmt.Scanln(&p.contact)

	//check if contact length is 10 or not
	if p.contact < 1000000000 || p.contact > 9999999999 {
		fmt.Println("Invalid contact number")
		return
	}
	fmt.Println("Patient added")
	fmt.Println("-----------------------------------")
	patients[patientId] = p
	patientId++
}

func updatePatient() {
	var id int
	fmt.Println("Enter patient id:")
	fmt.Scanln(&id)
	p, ok := patients[id]
	if !ok {
		fmt.Println("Patient not found")
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Which Patient details You want to update")
	fmt.Println("1. Name")
	fmt.Println("2. Age")
	fmt.Println("3.Gender")
	fmt.Println("4. Contact")
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter name:")
		fmt.Scanln(&p.name)
	case 2:
		fmt.Println("Enter age:")
		fmt.Scanln(&p.age)
	case 3:
		fmt.Println("Enter gender:")
		fmt.Scanln(&p.gender)
	case 4:
		fmt.Println("Enter contact:")
		fmt.Scanln(&p.contact)
	default:
		fmt.Println("Invalid choice")
		return
	}
	fmt.Println("Patient details updated")
	fmt.Println("-----------------------------------")
	patients[id] = p

}


func scheduleAppointment() {
	var a appointment
	fmt.Println("Enter date:")
	fmt.Scanln(&a.date)
	fmt.Println("Enter time:")
	fmt.Scanln(&a.time)
	fmt.Println("Enter medical practitioner:")
	fmt.Scanln(&a.medicalPractitioner)
	fmt.Println("Enter patient id:")
	fmt.Scanln(&a.patientId)
	//check patientid is avialable or not in patient , otherwise return patient not found
	_, ok := patients[a.patientId]
	if !ok {
		fmt.Println("Patient not found")
		fmt.Println("-----------------------------------")
		return
	}
	appointments[appointmentId] = a
	fmt.Println("Appointment scheduled")
	fmt.Println("-----------------------------------")
	appointmentId++
}

func generateReport() {
	fmt.Println("Patients Details")
	for id, p := range patients {
		fmt.Println("-----------------------------------")
		fmt.Println("ID:", id)
		fmt.Println("Name:", p.name)
		fmt.Println("Age:", p.age)
		fmt.Println("Gender",p.gender)
		fmt.Println("Contact:", p.contact)
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Appointments Details")
	for id, a := range appointments {
		fmt.Println("-----------------------------------")
		fmt.Println("ID:", id)
		fmt.Println("Date:", a.date)
		fmt.Println("Time:", a.time)
		fmt.Println("Medical practitioner:", a.medicalPractitioner)
		fmt.Println("Patient ID:", a.patientId)
	}
	fmt.Println("Current date and time:", time.Now())
}






