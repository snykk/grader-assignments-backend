package internal

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var CalculatorTestCase = func() {
	Describe("Add", func() {
		When("add 10 to 10", func() {
			It("should return 20", func() {
				c := NewCalculator(10)
				c.Add(10)
				Expect(c.Result()).To(Equal(float32(20)))
			})
		})

		When("add 10 to 10 and then add 10", func() {
			It("should return 30", func() {
				c := NewCalculator(10)
				c.Add(10)
				c.Add(10)
				Expect(c.Result()).To(Equal(float32(30)))
			})
		})

		When("add 10 to 10 and then add 10 and then add 10", func() {
			It("should return 40", func() {
				c := NewCalculator(10)
				c.Add(10)
				c.Add(10)
				c.Add(10)
				Expect(c.Result()).To(Equal(float32(40)))
			})
		})
	})

	Describe("Subtract", func() {
		When("subtract 10 from 10", func() {
			It("should return 2", func() {
				c := NewCalculator(10)
				c.Subtract(8)
				Expect(c.Result()).To(Equal(float32(2)))
			})
		})

		When("subtract 20 from 10 and then subtract again 5", func() {
			It("should return 5", func() {
				c := NewCalculator(20)
				c.Subtract(10)
				c.Subtract(5)
				Expect(c.Result()).To(Equal(float32(5)))
			})
		})

		When("subtract 10 from 10 and subtract again 10", func() {
			It("should return -10", func() {
				c := NewCalculator(10)
				c.Subtract(10)
				c.Subtract(10)
				Expect(c.Result()).To(Equal(float32(-10)))
			})
		})
	})

	Describe("Multiply", func() {
		When("multiply 10 by 10", func() {
			It("should return 100", func() {
				c := NewCalculator(10)
				c.Multiply(10)
				Expect(c.Result()).To(Equal(float32(100)))
			})
		})

		When("multiply 10 by 10 and then multiply again by 10", func() {
			It("should return 1000", func() {
				c := NewCalculator(10)
				c.Multiply(10)
				c.Multiply(10)
				Expect(c.Result()).To(Equal(float32(1000)))
			})
		})

		When("multiply 10 by 10 and then multiply again by 10 and then multiply again by 10", func() {
			It("should return 10000", func() {
				c := NewCalculator(10)
				c.Multiply(10)
				c.Multiply(10)
				c.Multiply(10)
				Expect(c.Result()).To(Equal(float32(10000)))
			})
		})
	})

	Describe("Divide", func() {
		When("divide 10 by 10", func() {
			It("should return 1", func() {
				c := NewCalculator(10)
				c.Divide(10)
				Expect(c.Result()).To(Equal(float32(1)))
			})
		})

		When("divide 100 by 10 and then divide again by 10", func() {
			It("should return 1", func() {
				c := NewCalculator(100)
				c.Divide(10)
				c.Divide(10)
				Expect(c.Result()).To(Equal(float32(1)))
			})
		})

		When("divide 100 by 10 and then divide again by 10 and then divide again by 10", func() {
			It("should return 0.1", func() {
				c := NewCalculator(100)
				c.Divide(10)
				c.Divide(10)
				c.Divide(10)
				Expect(c.Result()).To(Equal(float32(0.1)))
			})
		})
	})
}
