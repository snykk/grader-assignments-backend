
# Phase 1

Kita akan mengimplementasikan _interface_ untuk masing-masing _struct_ yang mewakili format invoice dari tiap departement.

Tedapat _package_ invoice yang berisi semua data yang diperlukan terkait _interface_ dan _struct_ dari setiap departement. Terdapat file `invoice.go` yang berisi _interface_ `Invoice` sebagai kontrak untuk masing-masing _struct_ invoice. Terdapat juga _struct_ `InvoiceData` yang berisi data yang akan dihasilkan dari setiap invoice dengan properti yang sudah ditentukan.

```go
type Invoice interface {
    RecordInvoice() (InvoiceData, error)
}
```

Buatlah _method_ yang dibutuhkan untuk _Struct_ dari departemen finance di file `finance.go` agar dapat mengimplementasikan _interface_ `Invoice` di atas. Kembalikan `InvoiceData` dengan format sebagai berikut:

```go
type InvoiceData struct {
    Date         string
    TotalInvoice float64
    departemen   DepartmentName
}

type DepartmentName string

const (
    Finance   DepartmentName = "finance"
    Warehouse DepartmentName = "warehouse"
    Marketing DepartmentName = "marketing"
)
```

Struct `InvoiceData` akan digunakan untuk menyimpan data invoice yang formatnya sudah bisa digunakan untuk melakukan rekapitulasi data. Terdapat data berupa tanggal, total harga, dan nama departement.

- `Date` bertipe `string` berformat `"DD-month-YYYY"` (contoh: "01-January-2022")
- `TotalInvoice` bertipe `float64` yang akan berisi total harga dari invoice finance di data `Details`
- `DepartmentName` akan berisi informasi departemen mana yang menghasilkan invoice tersebut ("finance", "warehouse", dan "marketing").

Return ke 2 adalah `error`, lakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong maka return `error` dengan pesan `"invoice date is empty"`
- Jika isi dari `Details` kosong maka return `error` dengan pesan `"invoice details is empty"`
- Jika `Status` di invoice finance berisi string kosong atau selain "paid" dan "unpaid", maka return `error` dengan pesan `"invoice status is not valid"`
- Jika terdapat `Total` di `Details` yang bernilai 0 atau negatif, maka return `error` dengan pesan `"total price is not valid"`

## Test Case

**Input**:

```txt
FinanceInvoice{
    Date: "01/01/2022",
    Status: "paid",
    Approved: "true"
    Details: []FinanceInvoiceDetail{
        {
            Name: "Pembalian nota",
            Total: 4000,
        },
        {
            Name: "Pembelian alat tulis",
            Total: 4000,
        },
    },
}
```

**Output**:

```txt
InvoiceData{
    Date: "01-January-2022",
    TotalInvoice: 8000,
    Department: "finance",
}
```

Format date akan berubah menjadi `DD-month-YYYY` seperti contoh di atas, dan hasil `TotalInvoice` akan diakumulasikan dari `Total` yang ada di `Details` yaitu `4000` + `4000`.  Dan terakhir `Department` akan bernilai `"finance"`.
