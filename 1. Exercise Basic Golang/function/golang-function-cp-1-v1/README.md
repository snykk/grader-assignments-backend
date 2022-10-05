# Exercise Function

## Humanize Date format

### Description

Terdapat sebuah data berupa angka yang merepresentasikan tanggal dengan 3 data yaitu `day`, `month` dan `year`. Kita ingin mengubahnya menjadi tanggal yang mudah dibaca dengan format `"[day]-[month]-[year]"` dengan `month` adalah nama bulan dalam bahasa Inggris. Jika tanggal tersebut kurang dari 10, maka tambahkan `0` di depannya.

Terdapat fungsi bernama `DateFormat` yang menerima 3 parameter bertipe `int` yaitu `day`, `month` dan `year`. Kembalikan _output_ berupa `string` yang berisi tanggal yang sudah diubah ke format yang diinginkan.

Contoh sebagai berikut, jika day = 1, month = 1, year = 2020, maka _output_ yang diharapkan adalah `"01-January-2020"`.

### Constraints

- Input `day`, `month`, `year` akan selalu berisi angka bulat positif dan tidak akan melebihi batas maksimal tanggal yang ada di bulan tersebut.

### Test Case Examples

#### Test Case 1

**Input**:

```txt
day = 1, month = 1, year = 2020
```

**Expected Output / Behavior**:

```txt
01-January-2020
```

**Explanation**:

Output yang diharapkan adalah `"01-January-2020"` karena :

- tanggal yang diinput adalah `1` maka akan diubah menjadi `01`
- bulan yang dimasukkan adalah `1` maka akan diubah menjadi `January`.

Sisanya tinggal digabung dengan format yang diinginkan menjadi `"01-January-2020"`.

#### Test Case 2

**Input**:

```txt
day = 31, month = 12, year = 2020
```

**Expected Output / Behavior**:

```txt
31-December-2020
```

**Explanation**:

Output yang diharapkan adalah `"31-December-2020"` karena :

- tanggal yang diinput adalah `31` maka akan diubah menjadi `31`
- bulan yang dimasukkan adalah `12` maka akan diubah menjadi `December`.

Sisanya tinggal digabung dengan format yang diinginkan menjadi `"31-December-2020"`.
