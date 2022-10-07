# Exercise file 2

## Record transactions

### Description

Perusahaan ingin mencatat semua transaksi dalam bentuk harian ke dalam file `txt`. Data transaksi saat ini masih tersimpan di dalam code program berbentuk `struct` dengan data berupa `date_transaction`,`transaction_type` dan  `amount`.

Buatlah sebuah program yang dapat mencatat transaksi ke dalam file `txt` dengan format `date;transaction_type;amount` dalam bentuk harian. Transaksi akan diakumulasi (total income - total expense) dalam satu hari. Jika total _income_ lebih besar dari _expense_, maka _transaction_type_ akan diisi `income`. Jika total _expense_ lebih besar dari _income_, maka _transaction_type_ akan diisi `expense`.

### Format Pengerjaan

Disediakan fungsi `RecordTransactions` dengan signature

```go
func RecordTransactions(path string, transactions []Transaction) error
```

Implementasikanlah fungsi tersebut sehingga fungsi tersebut akan mengolah `transactions` sesuai dengan `Description` di atas lalu menyimpan hasilnya ke file yang ditunjuk oleh parameter `path`.

### Constraints

Terdapat beberapa batasan yang perlu diperhatikan dalam mengerjakan soal ini:

1. Penulisan ke file tidak diakhiri dengan empty new line (baris terakhir di file tersebut haruslah berisi data transaksi)
2. Penulisan transaksi ke file harus urut berdasarkan tanggal. Contoh: rekap transaksi di tanggal `01/01/2021` harus berada di atas rekap transaksi pada tanggal `02/01/2021`

### Test Case Examples

#### Test Case 1

**Input**:

```txt
01/01/2021;income;100000
01/01/2021;expense;50000
01/01/2021;expense;30000
01/01/2021;income;20000
02/01/2021;expense;10000
02/01/2021;expense;10000
```

**Expected Output / Behavior**:

```txt
01/01/2021;income;40000
02/01/2021;expense;20000
```

**Explanation**:

Pada tanggal `01/01/2021` terdapat 4 transaksi (2 income , 2 expense) dengan total income 120000, dan total expense 80000, maka akumulasinya menjadi `40000`. Karena total profit lebih besar dari expense maka akan dicatat menjadi `01/01/2021;income;40000`.

Pada tanggal `02/01/2021` terdapat 2 transaksi (0 income , 2 expense) dengan tidak ada income, dan total expense 20000, maka akumulasinya menjadi `-20000`. Karena total profit lebih kecil dari expense maka akan dicatat menjadi `02/01/2021;expense;20000`.

di file `txt` akan terlihat seperti berikut:

```txt
01/01/2021;income;40000
02/01/2021;expense;20000
```

#### Test Case 2

**Input**:

```txt
01/01/2021;income;100000
01/01/2021;expense;50000
01/01/2021;expense;30000
01/01/2021;income;20000
01/01/2021;expense;50000
```

**Expected Output / Behavior**:

```txt
01/01/2021;expense;10000
```

**Explanation**:

Pada tanggal `01/01/2021` terdapat 5 transaksi (2 income , 3 expense) dengan total income 120000, dan total expense 130000, maka akumulasinya menjadi `-10000`. Karena total profit lebih kecil dari expense maka akan dicatat menjadi `01/01/2021;expense;10000`.

### Template

Berikut _template_ pengerjaannya :

```go
type Transaction struct {
    Date   string
    Type   string
    Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
    // kerjakan di sini
}
```
