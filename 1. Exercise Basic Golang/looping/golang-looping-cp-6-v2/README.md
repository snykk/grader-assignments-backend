# Exercise Looping

## Biggest Pair Number

### Description

Buatlah sebuah _code_ yang dapat mencari pasangan angka yang paling besar (penjumlahan 2 angka tersebut) dari deret angka yang diberikan. Pasangan angka adalah 2 angka dari setiap urutan angka mulai dari urutan paling kiri. Contoh deretan angka `1234`, maka pasangan pertama adalah `12` (angka pertama dan angka ke-2) dan pasangan kedua adalah `23` (angka ke-2 dan angke ke-3), dan pasangan terakhir adalah `34` (angka ke-3 dan angka ke-4).

Terdapat fungsi bernama `BiggestPairNumber` dengan parameter `numbers` bertipe `int` yang berisi deret angka _input_. Terakhir kembalikan _output_ berupa `int` yang berisi pasangan angka yang paling besar.

Contoh, terdapat deret angka `11223344`, pasangan angka pertama adalah `11` diambil dari urutan paling kiri (`[11]223344`) jika dijumlahkan akan menjadi `2`, pasangan angka kedua adalah `12` dari deretan di posisi selanjutnya (`1[12]23344`) jika dijumlahkan menjadi `3`, dan seterusnya. Nah, jika dilihat dari beberapa pasangan yang terbentuk, maka pasangan angka yang paling besar adalah `44` dengan jumlah `8`.

### Constraints

- Input `number` akan selalu berupa deret angka yang terdiri dari angka 0-9, dengan panjang karakter 2 <= `number` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
11223344
```

**Expected Output / Behavior**:

```txt
44
```

**Explanation**:

Output yang diharapkan adalah `44` karena pasangan angka yang paling besar adalah `44` dengan jumlah `8`.

#### Test Case 2

**Input**:

```txt
89083278
```

**Expected Output / Behavior**:

```txt
89
```

**Explanation**:

Output yang diharapkan adalah `89` karena pasangan angka yang paling besar adalah `89` dengan jumlah `17`.
