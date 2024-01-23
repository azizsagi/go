package main

import "fmt"

//Interface

type Financials interface {
	calculate_vat() float32
}

// Employee Struct
type Employee struct {
	employeeName string
	salary       float32
}

// Customer Struct
type Customer struct {
	customerName string
	salesInvoice float32
}

// Supplier Struct
type Supplier struct {
	supplierName    string
	purchaseInvoice float32
}

//implement VAT calculation for Employee
func (e Employee) calculate_vat() float32 {
	//20%
	return ((e.salary * 20) / 100)
}

//implement VAT calculation for Customer
func (c Customer) calculate_vat() float32 {
	//15%
	return ((c.salesInvoice * 15) / 100)
}

//implement VAT calculation for Customer
func (s Supplier) calculate_vat() float32 {
	//10%
	return ((s.purchaseInvoice * 10) / 100)
}

func print_vat(f Financials) {
	fmt.Println(f.calculate_vat())
}

func main() {
	emp := Employee{employeeName: "Muhammed", salary: 5000}
	cust := Customer{customerName: "Customer Lorem", salesInvoice: 100.50}
	sup := Supplier{supplierName: "Supplier Lipsum", purchaseInvoice: 150.59}

	fin := []Financials{emp, cust, sup}

	for _, v := range fin {
		print_vat(v)
	}

}
