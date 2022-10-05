# Exercise slice multidimention 2

## Count profit selling

Terdapat perusahaan bernama "Ruangprofit" yang menjual barang-barang berupa komputer, laptop, dan printer. Perusahaan selalu melakukan rekap penjualan setiap bulan dengan mencatat total penjualan dan total pengeluaran. karena perusahaan ini memiliki banyak cabang, maka perusahaan ini memiliki data penjualan dan pengeluaran untuk setiap cabang setiap bulan.

Buatlah fungsi yang dapat menghitung total keuntungan (_profit_) dari penjualan dan pengeluaran untuk setiap cabang di setiap bulan. Data yang didapat saat ini masih dalam bentuk _slice_ multidimensi dengan struktur sebagai berikut:

- _Slice_ data terluar adalah berisi data dari setiap cabang
- _Slice_ di tiap cabang adalah berisi data dari setiap bulan
- _Slice_ di tiap bulan adalah berisi data dari penjualan dan pengeluaran (hanya 2 data)

```go
[][][2]int
```

Kembalikan sebuah data _slice_ yang berisi total keuntungan dari penjualan dan pengeluaran seluruh cabang di setiap bulan.
Contoh, sebagai berikut:

```txt
data = [[[1000, 500], [500, 200]],[[1200, 200],[1000, 800]],[[500, 100],[700, 100]]]
```

Data di atas berarti perusahaan mendapat informasi dari 3 cabang dengan informasi setiap cabang melaporkan penjualan dan pengeluaran selama 2 bulan.

- Cabang 1 di bulan 1 terdapat penjualan 1000 dan pengeluaran 500, maka keuntungannya adalah 1000 - 500 = 500. Di bulan 2 terdapat penjualan 500 dan pengeluaran 200, maka keuntungannya adalah 500 - 200 = 300.
- Cabang 2 di bulan 1 terdapat penjualan 1200 dan pengeluaran 200, maka keuntungannya adalah 1200 - 200 = 1000. Di bulan 2 terdapat penjualan 1000 dan pengeluaran 800, maka keuntungannya adalah 1000 - 800 = 200.
- Cabang 3 di bulan 1 terdapat penjualan 500 dan pengeluaran 100, maka keuntungannya adalah 500 - 100 = 400. Di bulan 2 terdapat penjualan 700 dan pengeluaran 100, maka keuntungannya adalah 700 - 100 = 600.

Total keuntungan di bulan 1 adalah 500 + 1000 + 400 = 1900. Total keuntungan di bulan 2 adalah 300 + 200 + 600 = 1100.

Maka hasilnya adalah:

```txt
[1900, 1100]
```

Berikut _template_ pengerjaannya :

```go
func CountProfit(data [][][2]int) []int {
    // kerjakan di sini
}
```

Contoh _test case_ :

```txt
input: data = [[[1000, 800],[700, 500]], [[1000, 800],[900, 200]]]
output: [400, 900]

input: data = [[[1000, 500], [500, 150], [600, 100], [800, 750]]]
output: [500, 350, 500, 50]

input: data = [[[1000, 200]], [[500, 100]], [[600, 100]], [[450, 150]],[[100, 50]]]
output: [2050]

input data = []
output: []
```
