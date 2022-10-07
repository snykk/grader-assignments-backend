package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	service "a21hc3NpZ25tZW50/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var serviceTest service.ServiceInterface

var _ = Describe("CashierApp", func() {
	BeforeEach(func() {
		dbTest := database.NewDatabase()
		serviceTest = main.CashierApp(dbTest)
	})

	Context("Add valid product to cart", func() {
		It("should add product to cart in database", func() {
			errAddCart1 := serviceTest.AddCart("Kaos Polos", 2)
			errAddCart2 := serviceTest.AddCart("Kaos sablon", 1)

			act, err := serviceTest.ShowCart()

			Expect(act).To(HaveLen(2))
			Expect(act[0].ProductName).To(Equal("Kaos Polos"))
			Expect(act[0].Quantity).To(Equal(2))
			Expect(act[1].ProductName).To(Equal("Kaos sablon"))
			Expect(act[1].Quantity).To(Equal(1))
			Expect(errAddCart1).To(BeNil())
			Expect(errAddCart2).To(BeNil())
			Expect(err).To(BeNil())
		})
	})

	Context("Add valid product to cart with invalid quantity", func() {
		It("should return error invalid quantity", func() {
			errAddCart1 := serviceTest.AddCart("Topi", 3)
			errAddCart2 := serviceTest.AddCart("Kaos Polos", 0)
			errAddCart3 := serviceTest.AddCart("Kaos sablon", -3)

			act, _ := serviceTest.ShowCart()

			Expect(act).To(HaveLen(1))
			Expect(act[0].ProductName).To(Equal("Topi"))
			Expect(act[0].Quantity).To(Equal(3))
			Expect(errAddCart1).To(BeNil())
			Expect(errAddCart2).To(Not(BeNil()))
			Expect(errAddCart2.Error()).To(Equal("invalid quantity"))
			Expect(errAddCart3).To(Not(BeNil()))
			Expect(errAddCart3.Error()).To(Equal("invalid quantity"))
		})
	})

	Context("Add non-existing product to cart", func() {
		It("should return product not found error", func() {
			err := serviceTest.AddCart("RandomName", 99)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("product not found"))
		})
	})

	Context("Remove valid product from cart", func() {
		It("should remove product from cart in database", func() {
			serviceTest.AddCart("Kaos Polos", 2)
			serviceTest.AddCart("Kaos sablon", 1)
			serviceTest.AddCart("Celana hitam", 3)

			errRemoveCart := serviceTest.RemoveCart("Kaos sablon")

			act, err := serviceTest.ShowCart()

			Expect(act).To(HaveLen(2))
			Expect(act[0].ProductName).To(Equal("Kaos Polos"))
			Expect(act[0].Quantity).To(Equal(2))
			Expect(act[1].ProductName).To(Equal("Celana hitam"))
			Expect(act[1].Quantity).To(Equal(3))
			Expect(errRemoveCart).To(BeNil())
			Expect(err).To(BeNil())
		})
	})

	Context("Remove valid product but not exists in cart", func() {
		It("should return product not found error", func() {
			err := serviceTest.RemoveCart("Kaos Polos")
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("product not found"))
		})
	})

	Context("Remove non-existing product from cart", func() {
		It("should return product not found error", func() {
			err := serviceTest.RemoveCart("RandomName")
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("product not found"))
		})
	})

	Context("Reset cart", func() {
		It("should reset cart in database", func() {
			dbReset := database.NewDatabase()

			dbReset.Save([]entity.CartItem{
				{ProductName: "Kaos Polos", Quantity: 2, Price: 50000},
			})

			cashier := main.CashierApp(dbReset)

			cashier.AddCart("Kaos Polos", 2)
			cashier.AddCart("Kaos sablon", 1)

			cashier.ResetCart()

			act, err := cashier.ShowCart()

			Expect(act).To(HaveLen(0))
			Expect(err).To(BeNil())
		})
	})

	Context("Get all product", func() {
		It("should get all product in database", func() {
			act, err := serviceTest.GetAllProduct()

			Expect(act).To(HaveLen(9))
			Expect(err).To(BeNil())
		})
	})

	Context("Paid", func() {
		It("should return payment information and reset cart", func() {
			serviceTest.AddCart("Kaos Polos", 2)
			serviceTest.AddCart("Kaos sablon", 1)

			act, err := serviceTest.Paid(500000)

			Expect(act).To(Equal(entity.PaymentInformation{
				MoneyPaid:  500000,
				Change:     330000,
				TotalPrice: 170000,
				ListProduct: []entity.CartItem{
					{
						ProductName: "Kaos Polos",
						Price:       50000,
						Quantity:    2,
					},
					{
						ProductName: "Kaos sablon",
						Price:       70000,
						Quantity:    1,
					},
				},
			}))

			Expect(err).To(BeNil())

			currentCart, _ := serviceTest.ShowCart()
			Expect(currentCart).To(HaveLen(0))
		})
	})

	Context("Paid with insufficient amount of money", func() {
		It("should return money is not enough error", func() {
			serviceTest.AddCart("Kaos Polos", 2)
			serviceTest.AddCart("Kaos sablon", 1)

			act, err := serviceTest.Paid(1000)

			Expect(act).To(Equal(entity.PaymentInformation{}))
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("money is not enough"))
		})
	})
})
