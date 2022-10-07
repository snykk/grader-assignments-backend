# Assignment 3

## Invoice Reporting

### Warning

Pastikan saat mengerjakan ini, kamu sudah mengerti tentang:

- Struct
- Interface
- Error Handling
- Package dan Import

### Description

Terdapat perusahaan "kitaSemangat" yang memiliki banyak departemen dan bergerak di penjualan produk (retail). Perusahaan ini memiliki kendala dalam melakukan pelaporan faktur / _invoice_ setiap akhir bulan karena ada beberapa format _faktur_ yang berbeda-beda di setiap departemen. Khususnya terdapat 3 departement yang memiliki format faktur yang sangat berbeda, yaitu Finance, Warehouse, dan Marketing.

Buatlah program yang dapat melakukan rekap menjadi format invoice yang sudah disepakati bersama. Karena kita sudah belajar `struct`, maka format invoice dari tiap departemen akan di konversi menjadi sebagai berikut:

1. Finance

    Invoice finance umumnya berisi pembelian barang-barang. Terdapat data berupa tanggal, status, approval status dan detail pembelian.

    - Tanggal (Date) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Status bertipe `string` adalah informasi khusus di departemen keuangan yang menunjukkan apakah invoice tersebut sudah terbayar atau tidak
    - Approval status bertipe `bool` yang akan berisi informasi yang menandakan invoice tersebut sudah ditandatangani oleh atasan, jika bernilai `true` maka invoice sudah ditandatangani, dan akan `false` jika sebaliknya
    - Detail pembelian akan berisi deskripsi (bertipe `string`) dan total pembelian (bertipe `float64`), di setiap detail ini bisa lebih dari 1 data

    ```go
    type FinanceInvoice struct {
        Date       string
        Status     InvoiceStatus // status: "paid", "unpaid"
        Approved   bool
        Details   []Detail
    }

    type InvoiceStatus string

    const (
        PAID   InvoiceStatus = "paid"
        UNPAID InvoiceStatus = "unpaid"
    )

    type Detail struct {
        Description string
        Total       float64
    }
    ```

2. Warehouse

    Invoice departemen warehouse umumnya digunakan untuk mencatat barang masuk dan keluar termasuk harga dan diskonnya. Terdapat data berupa tanggal, tipe invoice, approval status dan detail produk.

    - Tanggal (Date) bertipe `string` berformat `"DD-month-YYYY"` (contoh: "01-January-2022")
    - Tipe invoice bertipe `InvoiceTypeName` yang akan berisi informasi apakah invoice tersebut merupakan invoice pembelian atau penjualan, jika bernilai `purchase` maka invoice tersebut merupakan invoice pembelian, dan akan `sales` jika sebaliknya
    - Approval status bertipe `bool` yang sama seperti invoice dari departemen Finance
    - Detail product (Products) bertipe _struct_ yang akan berisi informasi product (nama, total, harga per product, dan diskon yang diberikan)

    ```go
    type WarehouseInvoice struct {
        Date     string
        InvoiceType InvoiceTypeName
        Approved bool
        Products []Product
    }

    type InvoiceTypeName string

    const (
        PURCHASE InvoiceTypeName = "purchase"
        SALES    InvoiceTypeName = "sales"
    )


    type Product struct {
        Name     string
        Unit     int
        Price    float64
        Discount int
    }
    ```

3. Marketing

    Invoice marketing umumnya digunakan untuk keluar kota. Terdapat data berupa tanggal invoice, tanggal mulai perjalanan sampai selesai (untuk kebutuhan biaya penginapan), biaya lain yang harus dihitung juga, dan approval status.

    - Tanggal (Date) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Tanggal mulai perjalanan (StartDate) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Tanggal selesai perjalanan (EndDate) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Biaya per hari (PricePerDay) bertipe `float64` yang akan berisi informasi biaya penginapan per hari berdasarkan jarak tanggal mulai perjalanan sampai selesai
    - Biaya lain (AnotherFee) bertipe `int` yang akan berisi informasi biaya lain yang harus dibayarkan
    - Approval status bertipe `bool` yang sama seperti invoice dari departemen Finance dan Warehouse

    ```go
    type MarketingInvoice struct {
        Date        string
        StartDate   string
        EndDate     string
        PricePerDay float64
        AnotherFee  int
        Approved    bool
    }
    ```

### Instructions

- Pertama, baca soal phase 1 di file `phase1.md` dan kerjakan di file yang sudah ditentukan.
- Kedua, baca soal phase 2 di file `phase2.md` dan kerjaakan di file yang sudah ditentukan.
- Ketiga, baca soal phase 3 di file `phase3.md` dan kerjaakan di file yang sudah ditentukan.
- Terakhir, baca kerjakan baca soal phase 4 di file `phase4.md` dan kerjaakan di file yang sudah ditentukan.

### Test Case Examples

#### Test Case 1

**Input**:

Kita menggunakan fungsi `RecapDataInvoice` di file `main.go` untuk melakukan pengujian ini

```txt
data = [
    // invoice Finance
    {
        Date: "01/02/2020", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },
    {
        Date: "01/02/2020", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },

    // invoice warehouse
    {
        Date: "01-February-2020",
        Products: [
            {
                Name: "Beras",
                Unit: 10,
                Price: 10000,
                Quantity: 0.1,
            },
            {
                Name: "Gula",
                Unit: 5,
                Price: 15000,
                Discout: 0.2
            },
            }

        ],
        InvoiceType: "purchase",
        Approved:    true,
    },

    // invoice marketing
    {
        Date: "01/02/2020",
        StartDate: 20/01/2020",
        EndDate: "25/01/2020",
        PricePerDay: 10000,
        AnotherFree: 5000,
        Approved: true,
    },
]
```

**Expected Output / Behavior**:

```txt
InvoiceData = [
    {
        Date: "01-February-2020",
        Department: "finance",
        Total: 16000
    },
    {
        Date: "01-February-2020",
        Department: "warehouse",
        Total: 150000
    },
    {
        Date: "01-February-2020",
        Department: "marketing",
        Total: 65000
    }
]

error = nil
```

**Explanation**:

Seperti yang dijelaskan sebelumnya bahwa fungsi ini akan mengumpulkan semua invoice dengan tanggal yang sama dan departemen yang sama menjadi satu data. Dan yang dihitung hanya berstatus `"paid"` dan `Approved` yang sudah ditandatangani `true`. Selama tidak terjadi error maka akan mengembalikan `InvoiceData` seperti contoh di atas

#### Test Case 2

**Input**:

Kita menggunakan fungsi `RecapDataInvoice` untuk melakukan pengujian ini

```txt
data = [{
    // invoice Finance
    {
        Date: "", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },
}]
```

**Expected Output / Behavior**:

```txt
InvoiceData = nil

error = "invoice date is empty"
```

**Explanation**:

Karena invoice date kosong maka akan mengembalikan _error_ `"invoice date is empty"`.
