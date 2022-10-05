# Exercise Looping

## Counting Letter

### Description

Terapat orang cadel yang ingin membaca sebuah kalimat panjang. Namun, dia penasaran berapa banyak huruf yang dia tidak bisa baca dengan benar. Huruf yang tidak bisa dibaca adalah `R`, `S`, `T` dan `Z` (baik itu huruf besar atau huruf kecil). Buatlah sebuah _code_ yang digunakan untuk menghitung banyaknya huruf yang tidak bisa dibaca dengan benar oleh orang cadel tersebut. Terdapat fungsi `CountingLetter` dengan parameter `text` bertipe `string` yang berisi kalimat yang akan dibaca oleh orang cadel tersebut. Terakhir kembalikan _output_ berupa `int` yang berisi banyaknya huruf yang tidak bisa dibaca oleh orang cadel tersebut.

Contoh sebagai berikut, terdapat text `"Remaja Semarak"`, maka _output_ yang diharapkan adalah `3` karena terdapat 3 huruf yang tidak bisa dibaca oleh orang cadel tersebut yaitu `R`, `S` dan `r`.

### Constraints

- Input `text` akan selalu berupa _string_ yang terdiri dari karakter huruf, dengan panjang karakter 1 <= `text` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
text = "Remaja muda yang berbakat"
```

**Expected Output / Behavior**:

```txt
3
```

**Explanation**:

Kita mendapatkan _output_ yang diharapkan adalah `3` karena terdapat 3 huruf yang tidak bisa dibaca oleh orang cadel tersebut yaitu `R`, `r` dan `t`.

#### Test Case 2

**Input**:

```txt
text = "Zebra Zig Zag"
```

**Expected Output / Behavior**:

```txt
4
```

**Explanation**:

Kita mendapatkan _output_ yang diharapkan adalah `4` karena terdapat 4 huruf yang tidak bisa dibaca oleh orang cadel tersebut yaitu 3 huruf `Z` dan 1 huruf `r`.
