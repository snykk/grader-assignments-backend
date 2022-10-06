# Assignment Map 5

## Delivery order day

### Description

Terdapat perusahaan yang menyediakan jasa pengiriman barang. Perusahaan ini menyediakan jasa pengiriman barang dengan harga berbeda-beda berdasarkan kota tujuan dan hari pengiriman. Jam operasi perusahaan adalah setiap hari senin - sabtu pukul 08.00 - 17.00. Saat ini, perusahaan hanya memiliki 4 kota tujuan yaitu JKT, BDG, BKS dan DPK.

```txt
Kode Lokasi
JKT = "JKT"
BDG = "BDG"
BKS = "BKS"
DPK = "DPK"

Biaya admin pengiriman berdasarkan hari
senin, rabu, jumat = 10%
selasa, kamis, sabtu = 5%

Lokasi yang dapat di kirim sesuai hari
JKT = setiap hari ( selain minggu )
BDG = rabu, kamis, sabtu
BKS = selasa, kamis, jumat
DPK = senin, selasa
```

Saat ini perusahaan tersebut menyimpan informasi pesanan dalam format `[first-name]:[last-name]:[price]:[kode-lokasi]`. Buatlah sebuah fungsi `deliveryOrder` yang dapat melakukan _filtering_ berdasarkan hari dan menghitung total biaya pengiriman. Terdapat 2 parameter yang harus diinputkan yaitu `data` dan `day`. Parameter `data` merupakan _slice of string_ yang berisi data-data pesanan. Parameter `day` merupakan hari pengiriman barang. Fungsi `deliveryOrder` akan mengembalikan nilai berupa _map_ yang berisi data-data pesanan yang telah di _filter_ dan perhitungan total biaya pengiriman.

### Test Case

#### Test Case 1

**Input**:

```txt
data = ["Budi:Gunawan:10000:JKT", "Andi:Sukirman:20000:JKT", "Budi:Sukirman:30000:BDG", "Andi:Gunawan:40000:BKS", "Budi:Gunawan:50000:DPK"]

day = "sabtu"
```

**Expected Output / Behavior**:

```txt
map[
    Andi-Sukirman: 21000, 
    Budi-Gunawan: 10500, 
    Budi-Sukirman: 31500,
]
```

**Explanation**:

Terdapat 5 data yang harus dikirim. Karena hari ini adalah `"sabtu"`, maka data tersebut akan dipilih hanya untuk lokasi yang jadwal pengirimannya pada hari `"sabtu"`, yaitu "JKT" dan "BDG". Terdapat 3 data pengiriman yang sesuai dengan jadwal pengiriman hari `"sabtu"`, yaitu:
`"Budi:Gunawan:10000:JKT", "Andi:Sukirman:20000:JKT","Budi:Sukirman:30000:BDG"`. Kemudian kita tambahkan dengan total biaya pengiriman untuk hari sabtu yaitu 5% dari total harga barang. Hasilnya adalah:

```txt
map[
    Andi-Sukirman: 21000, 
    Budi-Gunawan: 10500, 
    Budi-Sukirman: 31500,
]
```

#### Test Case 2

**Input**:

```txt
data = ["Anggi:Anggraini:20000:DPK", "Andi:Sukirman:15000:JKT"]

day = "selasa"
```

**Expected Output / Behavior**:

```txt
map[
    Anggi-Anggraini: 21000, 
    Andi-Sukirman: 15750
]
```

**Explanation**:

Terdapat 2 data yang harus dikirim. Karena hari ini adalah `"selasa"`, maka data tersebut akan dipilih hanya untuk lokasi yang jadwal pengirimannya pada hari `"selasa"`, yaitu "JKT", "BKS" dan "DPK".

Semua data yang ada sesuai dengan jadwal pengiriman hari `"selasa"`, yaitu:
`"Anggi:Anggraini:20000:DPK", "Andi:Sukirman:15000:JKT"`. Kemudian kita tambahkan dengan total biaya pengiriman untuk hari selasa yaitu 5% dari total harga barang. Hasilnya adalah:

```txt
map[
    Anggi-Anggraini: 21000, 
    Andi-Sukirman: 15750
]
```

### Template

Berikut _template_ pengerjaannya:

```go
func deliveryOrder(data []string, day string) map[string]float32 {
    
    // kerjakan di sini
}
```
