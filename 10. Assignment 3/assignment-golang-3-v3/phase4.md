# Phase 4

Kita akan gunakan _interface_ ke dalam fungsi `RecapDataInvoince` untuk menghitung total invoice dari setiap departemen di file `main.go` yang ada di luar _package_ invoice.

Buatlah fungsi `RecapDataInvoince` yang menerima parameter `invoices` bertipe `[]Invoice` (_interface_) dari _package_ `invoice`.

Fungsi ini akan mengembalikan _array_ `InvoiceData` dan `error`. Kita dapat menggunakan method dari `RecordInvoice` untuk mengubahnya menjadi `InvoiceData`. Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai dari invoice ke 3 departement yang sudah dibuat sebelumnya (**phase 1** sampai **phase 3**).

```go
func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
    // kerjakan di sini
}
```

FUngsi ini akan mengumpulkan semua invoice dengan ketentuan sebagai berikut:

- Hanya akan mengumpulkan invoice yang `Status` nya adalah `"paid"` (khusus untuk invoice finance)
- Hanya akan mengumpulkan `InvoiceType` nya adalah `"purchsae"` (khusus untuk invoice Warehouse)
- `Approved` nya bernilai `true` atau sudah ditandatangani
- Menggabungkan invoice dengan tanggal yang sama dan departemen yang sama menjadi satu data.

Contoh, jika terdapat 3 invoice sebagai berikut:

```txt
data 1 : invoice finance 
[
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 40000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 20000
            }
        ]
        Approved: true
        Status: "paid"
    },
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota 2",
                "Total": 50000
            },
            {
                "Description": "pembelian peralatan",
                "Total": 100000
            }
        ]
        Approved: true
        Status: "paid"
    },
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota 3",
                "Total": 50000
            },
            {
                "Description": "pembelian peralatan",
                "Total": 100000
            }
        ]
        Approved: false
        Status: "unpaid"
    }
]
```

Maka Invoice yang akan dihasilkan adalah sebagai berikut:

```txt
[
    {
        Date: "01-January-2022",
        Department: "finance",
        Total: 210000
    }
]
```

Karena ketiga invoice tersebut memiliki tanggal yang sama dan departemen yang sama, maka akan dijumlahkan menjadi satu data. Dan yang dihitung hanya berstatus `"paid"` dan `Approved` yang sudah ditandatangani `true`.
