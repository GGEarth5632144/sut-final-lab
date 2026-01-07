package main

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name string `valid:"required~You Name is required"`
	Salary float64 `valid:"range(15000|20000)~Salary must be between 15000 and 200000"`
	EmployeesCode string `valid:"required~EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"`
}