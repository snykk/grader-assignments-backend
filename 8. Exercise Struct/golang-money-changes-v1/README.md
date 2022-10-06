# Exercise struct 2

## Money Changes

Di sebuah pusat perbelanjaan terdapat sebuah mesin yang dapat mengeluarkan uang kembalian, mesin ini akan menghitung total belanjaan yang dibeli oleh pelanggan dan mengeluarkan uang kembalian sesuai dengan jumlah uang yang dibeli oleh pelanggan. Dan mesin ini akan mengembalikan uang dengan beberapa pecahan uang yaitu 1000, 500, 200, 100, 50, 20, 10, 5, 1.

Kita ingin membuat fungsi sejenis dari sistem kerja kasir tersebut dengan mencoba mengimplementasikannya. Nama fungsi tersebut adalah `MoneyChanges`. Fungsi ini akan menerima dua parameter :

- `amount` bertipe `int`: jumlah uang yang dibeli oleh pelanggan
- _slice of struct_ `Product` yang memiliki beberapa properti yaitu `Name` bertipe `string`, `Price` bertipe `int` dan `Tax` bertipe `int`

Kita akan menghitung semua harga dari semua produk (`price` + `tax`) yang dibeli oleh pelanggan dan kemudian menghitung kembalian. Kembalian akan disimpan dalam sebuah _slice of int_ yang berisi pecahan uang yang akan dikembalikan dari terbesar ke terkecil.

Berikut _template Pengerjaannya_ :

```go
type Product struct {
    Name  string
    Price int
    Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
    // kerjakan di sini
}
```

Contoh _test case_ :

```txt
input: 10000, [{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}]
output: [1000, 200]

input: 30000, [{Name: "baju 1", Price: 10000, Tax: 1000}, {Name: "Sepatu", Price: 15550, Tax: 1555}]
Output: [1000, 500, 200, 100, 50, 20, 20, 5]

input: 5500, [{Name: "Baju", Price: 5000, Tax: 500}]
output: []

...
```
