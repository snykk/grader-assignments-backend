package invoice

type InvoiceData struct {
	Date         string
	TotalInvoice float64
	Departemen   DepartmentName
}

type DepartmentName string

const (
	Finance   DepartmentName = "finance"
	Warehouse DepartmentName = "warehouse"
	Marketing DepartmentName = "marketing"
)

type Invoice interface {
	RecordInvoice() (InvoiceData, error)
}
