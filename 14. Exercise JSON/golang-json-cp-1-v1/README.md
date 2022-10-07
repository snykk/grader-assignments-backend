# Exercise JSON

## Grade point average (IP / Indeks Prestasi)

### Description

Perguruan tinggi ingin membuat sistem untuk menghitung nilai mahasiswa menggunakan perhitungan Indeks prestasi (IP). Sistem penilaian di pergurun tinggi menggunakan huruf, seperti A, B, C, dan seterusnya. Nilai huruf tersebut diubah menjadi skala penilaian perguruan tinggi dengan ketentuan sebagai berikut:

| Huruf | Skala Penilaian |
| ----- | ----- |
| A     | 4     |
| AB    | 3.5   |
| B     | 3     |
| BC    | 2.5   |
| C     | 2     |
| CD    | 1.5   |
| D     | 1     |
| DE    | 0.5   |
| E     | 0     |

IP dapat dihitung dengan rumus sebagai berikut:

```txt
IP = ( Jumlah Penilaian x study credit )  / total study credit
```

Terdapat data dari setiap mahasiswa dalam bentuk `json` yang memiliki _key_ `id` berisi kode mahasiwa, `name` berisi nama mahasiswa, `date` berisi tanggal rilis data tersebut, `semester` berisi data semester mahasiswa. Dan terdapat _key_ `studies` yang berisi data mata kuliah dalam bentuk _object_ yang memiliki _key_ sebagai berikut :

- `study_name` berisi nama mata kuliah
- `study_credit` berisi SKS (sistem kredit semester) dari mata kuliah tersebut
- `grade` berisi nilai yang didapat mahasiswa dalam mata kuliah tersebut

Contoh data JSON seperti berikut :

```json
{
    "id" : "M-20220606",
    "name" : "Imam Jaya Permana",
    "date" : "12/08/2022",
    "semester": 4,
    "studies": [
        {
            "study_name": "Mata kuliah 1",
            "study_credit": 2,
            "grade" : "A"
        },
        {
            "study_name": "Mata kuliah 2",
            "study_credit": 3,
            "grade" : "B"
        },
        {
            "study_name": "Mata kuliah 3",
            "study_credit": 3,
            "grade" : "AB"
        },
        {
            "name": "Mata kuliah 4",
            "study_credit": 3,
            "grade" : "A"
        }
    ]
}
```

Maka perhitungan nilai IP untuk mahasiswa tersebut adalah sebagai berikut :

- Mata kuliah 1 : dapat grader "A", maka nilai skala penilaian adalah 4, dan study credit adalah 2. Maka jumlah nilai mata kuliah 1 adalah 4 x 2 = 8
- Mata kuliah 2 : dapat grader "B", maka nilai skala penilaian adalah 3, dan study credit adalah 3. Maka jumlah nilai mata kuliah 1 adalah 3 x 3 = 9
- Mata kuliah 3 : dapat grader "AB", maka nilai skala penilaian adalah 3.5, dan study credit adalah 3. Maka jumlah nilai mata kuliah 1 adalah 3.5 x 3 = 10.5
- Mata kuliah 4 : dapat grader "A", maka nilai skala penilaian adalah 4, dan study credit adalah 3. Maka jumlah nilai mata kuliah 1 adalah 4 x 3 = 12

Kita mendapat total nilai dari semua mata kuliah adalah 8 + 9 + 10.5 + 12 = 39.5 kemudian tinggal dibagi dengan total study credit yang diambil yaitu 2 + 3 + 3 + 3 = 11. Maka nilai IP dari mahasiswa tersebut adalah 39.5 / 11 = 3.59

Buatlah program untuk menghitung nilai IP dari mahasiswa tersebut. Diberikan fungsi `ReadJSON` untuk membaca data dari file `json`. Fungsi ini memiliki parameter `filename` yang berisi alamat file `json` yang akan dibaca. Terakhir  kembalikan 2 data yaitu `struct` dan `error`:

- nilai `struct` akan berisi data `json` yang di tampung ke dalam `struct` yang sudah disediakan, yaitu `Report`.
- nilai `error` akan berisi nilai `nil` jika tidak terjadi error, dan akan berisi nilai `error` jika terjadi error.

### Test Case Examples

#### Test Case 1

**Input**:

- dari file json

```json
{
    "id" : "M-20220606",
    "name" : "Imam Jaya Permana",
    "date" : "12/08/2022",
    "semester": 5,
    "studies": [
        {
            "study_name": "Mata kuliah 1",
            "study_credit": 2,
            "grade" : "A"
        },
        {
            "study_name": "Mata kuliah 2",
            "study_credit": 2,
            "grade" : "A"
        },
        {
            "study_name": "Mata kuliah 3",
            "study_credit": 3,
            "grade" : "C"
        },
        {
            "study_name": "Mata kuliah 3",
            "study_credit": 3,
            "grade" : "C"
        },
    ]
}
```

**Expected Output / Behavior**:

```txt
2.80
```

**Explanation**:

total nilai yang didapat adalah 28 (hasil dari konversi _grade_ ke skala penilaian x _study credit_). Kemudian dibagi dengan total _study credit_ yaitu 10. Maka nilai IP dari mahasiswa tersebut adalah 28 / 10 = 2.8

#### Test Case 2

**Input**:

- dari file json

```json
{
    "id" : "M-20220606",
    "name" : "Imam Jaya Permana",
    "date" : "12/08/2022",
    "semester": 5,
    "studies": []
}
```

**Expected Output / Behavior**:

```txt
0.0
```

**Explanation**:

total nilai IP adalah 0 karena tidak ada mata kuliah yang diambil di semester itu.

### Template

Berikut _template_ pengerjaannya :

```go

// buat property dan method yang dibutuhkan
type Report struct {}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
    // kerjakan di sini
}

func GradePoint(report Report) float64 {
    // kerjakan di sini
}
```
