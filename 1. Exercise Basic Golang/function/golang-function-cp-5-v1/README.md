# Exercise Function

## Sum Minimum and Maximum number

### Description

Terdapat sekumpulan angka yang tidak diketahui berapa banyaknya / dinamis. Buatlah sebuah _code_ yang dapat mencari angka terkecil dan sebuah _code_ yang dapat mencari angka terbesar dari sekumpulan angka tersebut. Kemudian jumlahkanlah kedua angka tersebut (terkecil dan terbesar).

Disediakan fungsi `FindMin` dengan parameter `nums` bertipe `...int` (data dinamis) untuk mencari angka terkecil dari sekumpulan angka. Kembalikan _output_ berupa `int` yang berisi angka terkecil.

Terdapat juga fungsi `FindMax` dengan parameter `nums` bertipe `...int` (data dinamis) untuk mencari angka terbesar dari sekumpulan angka. Kembalikan _output_ berupa `int` yang berisi angka terbesar.

Terakhir, ada fungsi `SumMinMax` dengan parameter `nums` bertipe `...int` (data dinamis) untuk menjumlahkan angka terkecil dan angka terbesar. Gunakan 2 fungsi sebelumnya (`FindMin` dan `FindMax`) untuk mencari angka terkecil dan angka terbesar.

Kembalikan _output_ berupa `int` yang berisi hasil penjumlahan angka terkecil dan angka terbesar.

Contoh, terdapat sekumpulan angka `1, 2, 3, 4, 5, 6, 7, 8, 9, 10`, angka terkecil adalah `1` dan angka terbesar adalah `10`. Jika dijumlahkan, maka hasilnya adalah `11`.

### Constraints

- Input `nums` akan selalu berupa sekumpulan angka yang terdiri dari angka 0-100.000, dengan jumlah data dinamis maksimal 250 data.

### Test Case Examples

#### Test Case 1

**Input**:

```txt
nums = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
```

**Expected Output / Behavior**:

```txt
11
```

**Explanation**:

Output yang diharapkan adalah `11` dengan rincian sebagai berikut:

- Kita menggunakan fungsi `FindMin` untuk mencari angka terkecil dari sekumpulan angka. Angka terkecil yang di dapat adalah `1`.
- Kita menggunakan fungsi `FindMax` untuk mencari angka terbesar dari sekumpulan angka. Angka terbesar yang di dapat adalah `10`.

Kemudian kita jumlahkan kedua angka tersebut di fungsi `SumMinMax`, yaitu `1 + 10 = 11`.

#### Test Case 2

**Input**:

```txt
nums = 333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111
```

**Expected Output / Behavior**:

```txt
3111
```

**Explanation**:

Output yang diharapkan adalah `3111` dengan rincian sebagai berikut:

- Kita menggunakan fungsi `FindMin` untuk mencari angka terkecil dari sekumpulan angka. Angka terkecil yang di dapat adalah `111`.
- Kita menggunakan fungsi `FindMax` untuk mencari angka terbesar dari sekumpulan angka. Angka terbesar yang di dapat adalah `3000`.

Kemudian kita jumlahkan kedua angka tersebut di fungsi `SumMinMax`, yaitu `111 + 3000 = 3111`.
