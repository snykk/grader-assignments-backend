# Exercise JSON

## Loan Report

### Description

Perusahaan memberlakukan pinjaman uang kepada karyawan dengan harapan ingin meringankan beban biaya karyawan tanpa harus memberikan bayaran tambahan dari perusahaan. Perusahaan selalu menyediakan saldo awal untuk peminjaman di awal bulan. Setiap peminjaman dari karyawan akan diberikan uang sebesar 50.000 per hari. Jika saldo dari perusahan tidak cukup untuk membayar peminjaman, maka karyawan tidak akan mendapatkan uang.

Terdapat 3 informasi yang dikumpulkan dari divisi keuangan perusahaan, yaitu:

- saldo awal uang peminjaman
- Data peminjaman karyawan
- Data informasi karyawan

Karena data tersebut disimpan ke dalam program, maka informasi tersebut akan disimpan ke dalam _struct_ `LoanData`.

```go
type LoanData struct {
    StartBalance int
    Data         []Loan
    Employees    []Employee
}

type Loan struct {
    Date        string
    EmployeeIDs []string
}

type Employee struct {
    ID       string
    Name     string
    Position string
}
```

Dalam struct `LoanData` terdapat 3 properti yaitu:

- `StartBalance` bertipe `int` yang menyimpan saldo awal uang peminjaman
- `Data` bertipe `[]Loan` yang menyimpan data peminjaman karyawan, dalam struct `Loan` akan menyimpan tanggal peminjaman dan `EmployeeIDs` yang menyimpan kumpulan ID karyawan yang melakukan peminjaman di tanggal tersebut
  - data tanggal akan berformat `"DD-month-YY"`, seperti `"01-January-2021"`
- `Employees` bertipe `[]Employee` yang menyimpan data informasi karyawan, dalam struct `Employee` akan menyimpan ID, nama, dan posisi karyawan
  - `ID` berisi data unik, tidak sama dengan `ID` karyawan lain, seperti `EMP001`

Buatlah sebuah program yang dapat menghasilkan laporan peminjaman (rekap) selama sebulan dari beberapa data tersebut ke _struct_ `LoanRecord` dengan struktur sebagai berikut:

- `MonthDate` bertipe `string` yang menyimpan selama bulan peminjaman dalam format `"month YYYY"`, seperti `January 2021`
- `StartBalance` bertipe `int` yang menyimpan saldo awal uang peminjaman
- `EndBalance` bertipe `int` yang menyimpan saldo akhir uang peminjaman
- `Borrowers` bertipe _struct_ yang akan menyimpan data peminjaman karyawan, dalam _struct_ tersebut akan menyimpan employee ID dan total peminjaman yang dilakukan oleh karyawan tersebut

Gunakan fungsi `LoanReport` untuk menghasilkan laporan peminjaman. Fungsi tersebut memiliki parameter `data` bertipe `LoanData` yang berisi data peminjaman yang sudah dijelaskan di atas. Kembalikan nilai berupa _struct_ `LoanRecord` yang berisi rekap data peminjaman.

Kemudian, gunakan fungsi `RecordJSON` untuk menyimpan laporan ke dalam _file_ JSON. Fungsi `RecordJSON` memiliki 2 parameter yaitu:

- `filename` bertipe `string` yang menyimpan nama _file_ yang akan disimpan
- `record` bertipe _struct_ `LoanRecord` yang berisi rekap data peminjaman yang akan disimpan ke dalam _file_.

Contoh:

```json
{
    "month": "January 2021",
    "start_balance": 1000000,
    "end_balance": 850000,
    "borrowers": [
        {
            "id": "EMP001",
            "total_loan": 50000
        },
        {
            "id": "EMP002",
            "total_loan": 100000
        }
    ]
}
```

> hint: Gunakan `json.Marshal` untuk mengubah object menjadi JSON byte lalu gunakan simpan JSON byte tersebut ke file

### Constraints

- `StartBalance` bertipe `int` yang akan selalu bernilai lebih dari `0`
- `Data` bertipe `[]Loan` yang akan selalu memiliki data (minimal 1 data) dan sudah terurut berdasarkan tanggal peminjaman (tidak perlu di-sort)
- `Employees` bertipe `[]Employee` yang akan selalu memiliki data karyawan (minimal 1 data)
- Seluruh item dalam `Data` merupakan daftar pinjaman dalam bulan yang sama

### Test Case Examples

#### Test Case 1

**Input**:

```txt
data = LoanData{
    StartBalance: 500000,
    Data: []Loan{
        {
            Date:        "01-January-2021",
            EmployeeIDs: []string{"EMP001", "EMP002"},
        },
        {
            Date:        "02-January-2021",
            EmployeeIDs: []string{"EMP001", "EMP003"},
        },
    },
    Employees: []Employee{
        {
            ID:       "EMP001",
            Name:     "Eddy Assidiqi",
            Position: "Data Engineer",
        },
        {
            ID:       "EMP002",
            Name:     "Imam Permana",
            Position: "Frontend Engineer",
        },
        {
            ID:       "EMP003",
            Name:     "Rizky Ramadhan",
            Position: "Data Engineer",
        }
    }
}
```

**Expected Output / Behavior**:

```json
{
    "month": "January 2021",
    "start_balance": 500000,
    "end_balance": 300000,
    "borrowers": [
        {
            "id": "EMP001",
            "total_loan": 100000
        },
        {
            "id": "EMP002",
            "total_loan": 50000
        },
        {
            "id": "EMP003",
            "total_loan": 50000
        }
    ]
}
```

**Explanation**:

Peminjaman terjadi di bulan January 2021, dengan saldo awal 500.000. Terdapat peminjaman dari tanggal 1 January - 2 January 2021, dan peminjamnya adalah 3 karyawan. Tanggal 1 January ada 2 karyawan yang melakukan peminjaman, yaitu Eddy dan Imam. Masing-masing karyawan melakukan peminjaman sebesar 50.000. Tanggal 2 January ada 2 karyawan yang melakukan peminjaman, yaitu Eddy dan Rizky. Masing-masing karyawan melakukan peminjaman sebesar 50.000.

Jadi total peminjaman yang dilakukan adalah 200.000. Jadi saldo akhir adalah 500.000 - 200.000 = 300.000. Dan data peminjamannya adalah:

- Eddy Assidiqi melakukan peminjaman sebesar 100.000
- Imam Permana melakukan peminjaman sebesar 50.000
- Rizky Ramadhan melakukan peminjaman sebesar 50.000

#### Test Case 2

**Input**:

```txt
data = []LoanData{
    StartBalance: 200000,
    Data: []Loan{
        {
            Date:        "01-March-2021",
            EmployeeIDs: []string{"EMP001", "EMP002", "EMP003"},
        },
        {
            Date:        "04-March-2021",
            EmployeeIDs: []string{"EMP001", "EMP002", "EMP003"},
        },
    },
    Employees: []Employee{
        {
            ID:       "EMP001",
            Name:     "Eddy Assidiqi",
            Position: "Data Engineer",
        },
        {
            ID:       "EMP002",
            Name:     "Imam Permana",
            Position: "Frontend Engineer",
        },
        {
            ID:       "EMP003",
            Name:     "Rizky Ramadhan",
            Position: "Data Engineer",
        }
    }
}
```

**Expected Output / Behavior**:

```json
{
    "month": "March 2021",
    "start_balance": 200000,
    "end_balance": 0,
    "borrowers": [
        {
            "id": "EMP001",
            "total_loan": 100000
        },
        {
            "id": "EMP002",
            "total_loan": 50000
        },
        {
            "id": "EMP003",
            "total_loan": 50000
        }
    ]
}
```

**Explanation**:

Peminjaman terjadi di bulan Maret 2021, dengan saldo awal 200.000. Terdapat peminjaman dari tanggal 1 Maret - 2 Maret 2021, dan peminjamnya adalah 3 karyawan. Tanggal 1 Maret ada 3 karyawan yang melakukan peminjaman, yaitu Eddy, Imam, dan Rizky. Masing-masing karyawan melakukan peminjaman sebesar 50.000.

Tanggal 4 Maret ada 3 karyawan yang melakukan peminjaman, yaitu Eddy, Imam, dan Rizky. Karena sisa saldo di tanggal 1 Maret sudah terpakai 150.000, maka sisa saldo di tanggal 4 Maret hanya 50.000. Karena peminjam ada 3 orang dan peminjam pertama adalah Eddy, maka Eddy mendapat pinjaman 50.000, namun Imam dan Rizky tidak dapat uang pinjaman karena saldo sudah habis.

Jadi saldo sudah habis `0`. Dan data peminjamannya adalah:

- Eddy Assidiqi melakukan peminjaman sebesar 100.000
- Imam Permana melakukan peminjaman sebesar 50.000
- Rizky Ramadhan melakukan peminjaman sebesar 50.000

### Template

Berikut _template_ pengerjaannya :

```go

// buat property yang dibutuhkan untuk menampung data yang akan diubah ke JSON
type LoanRecord struct {}

// gunakan fungsi ini untuk menghitung total peminjaman
// kembalian berupa struct 'LoanRecord'
func LoanReport(data LoanData) LoanRecord {
    // kerjakan di sini
}

func RecordJSON(record LoanRecord, path string) error {
    // kerjakan di sini
}
```
