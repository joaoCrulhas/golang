package main

import "fmt"

//Payment describe how to calculate employee paycheck
type Payment interface {
	calculatePayCheck() float32
}

//PersonalData employee's age and names
type PersonalData struct {
	Name         string
	Age          int
	BasePaycheck float32
}

//FullTimeEmployee the employee that work as a full-time employee
type FullTimeEmployee struct {
	PersonalData PersonalData
}

//FreelancerEmployee the employee that work as a full-time employee
type FreelancerEmployee struct {
	PersonalData PersonalData
}

//ManagerEmployee the employee that work as a full-time employee
type ManagerEmployee struct {
	PersonalData PersonalData
}

func calculatePayCheck(employee Payment) float32 {
	return employee.calculatePayCheck()
}

func (fr FreelancerEmployee) calculatePayCheck() float32 {
	return fr.PersonalData.BasePaycheck + (0.1 * fr.PersonalData.BasePaycheck)
}

func (f FullTimeEmployee) calculatePayCheck() float32 {
	return f.PersonalData.BasePaycheck + (0.4 * f.PersonalData.BasePaycheck)
}

func (m ManagerEmployee) calculatePayCheck() float32 {
	return m.PersonalData.BasePaycheck + (0.5 * m.PersonalData.BasePaycheck)
}

func main() {
	manager := ManagerEmployee{PersonalData: PersonalData{Name: "Joao Crulhas", Age: 29, BasePaycheck: 2300.00}}
	freelancerEmployee := FreelancerEmployee{PersonalData: PersonalData{Name: "Joao Crulhas", Age: 29, BasePaycheck: 2300.00}}
	fullTimeEmployee := FullTimeEmployee{PersonalData: PersonalData{Name: "Joao Crulhas", Age: 29, BasePaycheck: 2300.00}}

	managerPaycheck := calculatePayCheck(manager)
	freelancerPaycheck := calculatePayCheck(freelancerEmployee)
	fullTimePaycheck := calculatePayCheck(fullTimeEmployee)

	fmt.Printf("Employee %v => Salary => $%.2f \n", manager.PersonalData.Name, managerPaycheck)
	fmt.Printf("Employee %v => Salary => $%.2f \n", freelancerEmployee.PersonalData.Name, freelancerPaycheck)
	fmt.Printf("Employee %v => Salary => $%.2f \n", fullTimeEmployee.PersonalData.Name, fullTimePaycheck)

}
