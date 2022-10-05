# Exercise Function

## Money Change

### Description

Terdapat toko yang ingin menghitung kembalian uang yang diberikan kepada pembeli. Buatlah sebuah _code_ yang dapat menghitung kembalian uang yang diberikan kepada pembeli. Terdapat fungsi bernama `MoneyChange` dengan parameter `money` bertipe `int` yang berisi uang yang diberikan oleh pembeli. Parameter ke 2 `listPrice` bertipe `...int` yang berisi harga barang yang dibeli oleh pembeli. Terakhir kembalikan _output_ berupa `string` yang berisi kembalian uang yang diberikan kepada pembeli dengan format mata uang rupiah `"Rp. xxx.xxx.xxx"`, contoh seperti `"Rp. 100.000"`. Jika uang yang diberikan kurang dari harga barang yang dibeli, maka kembalikan `"Uang tidak cukup"`.

Terdapat fungsi `ChangeToCurrency` berisi parameter `change` bertipe `int` yang digunakan untuk mengubah angka menjadi format mata uang rupiah `"Rp. xxx.xxx.xxx"`. Gunakan fungsi ini di fungsi `MoneyChange` untuk mengubah kembalian uang menjadi format mata uang rupiah.

Contoh, terdapat pembeli yang membayar dengan uang `100000` dan membeli barang dengan _list_ harga `50000`, `10000` , `10000`, `5000`, `5000`. Maka kembalian yang didapat adalah `20000`, terakhir kita konversi ke format mata uang rupiah menjadi `"Rp. 20.000"`.

### Constraints

- Input `money` akan selalu berupa angka bulat positif, dengan panjang karakter 1 <= `money` <= 100.000.000
- Input `listPrice` akan selalu berupa angka bulat positif dengan data dinamis, dengan panjang karakter 1 <= `listPrice` <= 100.000.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
money = 100000, listPrice = 50000, 10000, 10000, 5000, 5000
```

**Expected Output / Behavior**:

```txt
"Rp. 20.000"
```

**Explanation**:

Output yang diharapkan adalah `"Rp. 20.000"` karena kembalian uang yang diberikan kepada pembeli adalah `20000` (100000 dikurangi total harga = 80000). Terakhir tinggal kita konversi ke format mata uang rupiah menjadi `"Rp. 20.000"`.

#### Test Case 2

**Input**:

```txt
money = 150000, listPrice = 50000, 30000, 30000, 30000, 20000, 50000
```

**Expected Output / Behavior**:

```txt
"Uang tidak cukup"
```

**Explanation**:

Output yang diharapkan adalah `"Uang tidak cukup"` karena uang yang diberikan oleh pembeli kurang dari total harga barang yang dibeli (150000 dikurangi total harga = 210000).
