package main

import (
	"testing"

	"github.com/asaskevich/govalidator"

	"github.com/onsi/gomega"
)

func TestNegativeCode(t *testing.T) {
	g := gomega.NewWithT(t)

	u := Employees{
		Name:          "Test",
		Salary:        15000,
		EmployeesCode: "",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))

}
