# Exercise file 1

## Calculate profit or loss

Terdapat perusahaan A yang menjual barang dagangan berupa makanan dan minuman. Perusahaan tersebut melakukan pencatatan transaksi sederhana secara manual di dalam sebuah file.

Setiap file akan berisi seluruh transaksi dalam satu bulan. Transaksi yang dicacat hanya dikategorikan sebagai pemasukan atau "income" dan pengeluaran atau "expense". Format transaksinya berupa `[dd/mm/yyyy];[type];[amount]`.

Contoh data transaksi:

- `transactions.txt`

```txt
1/1/2021;income;100000
1/1/2021;expense;50000
2/1/2021;income;200000
2/1/2021;expense;100000
3/1/2021;income;300000
3/1/2021;expense;150000
4/1/2021;income;400000
4/1/2021;expense;200000
...
```

Buatlah sebuah fungsi yang dapat digunakan untuk mendapatkan data dari file tersebut, kemudian menghitung total keuntungan ("_profit_") atau kerungian ("loss") dari perhitungan transaksi dalam satu bulan. Kembalikan nilai dalam format yang sama dengan tanggal terakhir dari transaksi tersebut.

Contoh, jika transaksi dalam bulan januari dan mengalami keuntungan, maka:

```txt
"4/1/2021;income;100000"
```

Berikut _template_ yang diperlukan:

```go

// gunakan fungsi ini untuk mendapatkan data file
// kembalikan data berupa slice of string
func Readfile(path string) ([]string, error) {
    // kerjakan di sini
}

func CalculateProfitLoss(data string) string {
    // kerjakan di sini
}
```

> Hint : kalian dapat memisah setiap baris data di `transactions.txt` dengan pemisah `"\n"` (_new line_)
