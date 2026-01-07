package main

import (
	"testing"

	"github.com/asaskevich/govalidator"

	"github.com/onsi/gomega"
)

func TestNegativeSalary(t *testing.T) {
	g := gomega.NewWithT(t)

	u := Employees{
		Name:          "Test",
		Salary:        10,
		EmployeesCode: "AA-23",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Salary must be between 15000 and 200000"))

}
