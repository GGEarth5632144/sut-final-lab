package main

import (
	"testing"

	"github.com/asaskevich/govalidator"

	"github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := gomega.NewWithT(t)

	u := Employees{
		Name:          "Test",
		Salary:        15001,
		EmployeesCode: "AA-23",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}
