# Exercise Conditon

## Graduate student

### Description

Buatlah sebuah kondisi yang menentukan apakah murid tersebut lulus atau tidak berdasarkan kondisi sebagai berikut :

- Jika nilai 70 keatas dan absennya di bawah 5 maka tampilkan `"lulus"`
- Jika nilai di bawah 70 atau absennya minimal 5 keatas maka tampilkan `"tidak lulus"`

Terdapat fungsi bernama `GraduateStudent` yang berisi 2 parameter yang bisa digunakan, pertama `score` bertipe `int` berisi nilai akhir, dan kedua `absent` bertipe `int` berisi jumlah ketidakhadiran murid.

### Constraints

- Input `score` akan selalu berupa bilangan positif dari 0 sampai 100
- Input `absent` akan selalu berupa bilangan positif dari 0 sampai 10

### Test Case Examples

#### Test Case 1

**Input**:

```txt
score = 100, absent = 4
```

**Expected Output / Behavior**:

```txt
"lulus"
```

**Explanation**:

Output yang diharapkan adalah `"lulus"` karena nilai 100 lebih besar atau sama dengan 70 dan jumlah ketidakhadiran 4 lebih kecil dari 5.

#### Test Case 2

**Input**:

```txt
score = 80, absent = 5
```

**Expected Output / Behavior**:

```txt
"tidak lulus"
```

**Explanation**:

Output yang diharapkan adalah `"tidak lulus"` karena absent 5 lebih besar atau sama dengan 5.
