# Exercise Function

## Count vowel and consonant

### Description

Kita ingin menghitung jumlah konsonan dan vokal dalam sebuah kalimat. Huruf vokal atau huruf hidup terdiri dari a, i, u, e, o. Selain dari itu dianggap sebagai huruf konsonan (baik itu huruf besar atau kecil). Perlu diperhatikan bahwa spasi tidak dihitung sebagai huruf.

Buatlah code yang dibutuhkan di fungsi `CountVowelConsonant` agar dapat menghitung jumlah konsonan dan vokal dalam sebuah kalimat. Terdapat parameter `str` bertipe `string` yang akan menerima input berupa kalimat. Terdapat 3 return yang dapat diisi sesuai kriteria sebagai berikut:

1. _Return_ pertama adalah jumlah huruf vokal (vowel)
2. _Return_ kedua adalah jumlah huruf konsonan (consonant)
3. _Return_ ketiga adalah pengecekan. Jika dalam kalimat tersebut tidak ada vokal atau / dan tidak ada konsonan maka _return_ ketiga akan berisi `true`, jika ada maka `false`.

### Constraints

- Input `str` akan berisi karakter tanpa simbol dan angka, dengan panjang karakter 1 <= `str` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
"kopi"
```

**Expected Output / Behavior**:

```txt
2, 2, false
```

**Explanation**:

Output yang dihasilkan adalah `2, 2, false`, dengan penjelasan sebagai berikut:

- _return_ pertama adalah `2` karena jumlah huruf vokal adalah `2` (o, i)
- _return_ kedua adalah `2` karena jumlah huruf konsonan adalah `2` (k, p)
- _return_ ketiga adalah `false` karena ada huruf vokal dan huruf konsonan

#### Test Case 2

**Input**:

```txt
"bbbbb ccccc"
```

**Expected Output / Behavior**:

```txt
0, 10, true
```

**Explanation**:

Output yang dihasilkan adalah `0, 10, true`, dengan penjelasan sebagai berikut:

- _return_ pertama adalah `0` karena tidak mengandung huruf vokal
- _return_ kedua adalah `10` karena jumlah huruf konsonan adalah `10` (b, b, b, b, b, c, c, c, c, c)
- _return_ ketiga adalah `true` karena tidak ada huruf vokal sama sekali

#### Test Case 3

**Input**:

```txt
"Hidup Itu Indah"
```

**Expected Output / Behavior**:

```txt
6, 7, false
```

**Explanation**:

Output yang dihasilkan adalah `6, 7, false`, dengan penjelasan sebagai berikut:

- _return_ pertama adalah `6` karena jumlah huruf vokal adalah `6`
- _return_ kedua adalah `7` karena jumlah huruf konsonan adalah `7`
- _return_ ketiga adalah `false`
