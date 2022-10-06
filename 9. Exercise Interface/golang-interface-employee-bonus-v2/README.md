# Exercise interface 1

## Employee bonus

Terdapat perusahaan bernama "PT MencariRejekiSejati" yang sangat terkenal dan memiliki banyak karyawan. Kebetulan di bulan ini perusahaan tersebut sedang memberikan insentif berupa bonus kepada karyawan yang telah bekerja dengan baik selama setahun. Namun, karyawan yang belum bekerja selama setahun juga masih bisa mendapatkan bonus. Terdapat 3 kategori karyawan dengan perhitungan bonus yang berbeda-beda, yaitu Junior, Senior, dan Manager.

Buatlah struct yang dibutuhkan berdasarkan 3 kategori karyawan tersebut ( "Manager", "Senior", dan "Junior" ). Pada struct "Junior" memiliki atrubut sebagai berikut:

- `Name` bertipe `string`
- `BaseSalary` bertipe `int`
- `WorkingMonth` bertipe `int` yang berisi lama kerja dalam bulan

Struct "Senior" memiliki semua atribut yang dimiliki oleh struct "Junior" dan tambahan atribut `PerformanceRate` bertipe `float64`

Struct "Manager" memiliki semua atribut yang dimiliki oleh struct "Senior" dan tambahan atribut `BonusManagerRate` bertipe `float64`

Untuk menghitung bonus terdapat tiga aturan sebagai berikut:

- Bonus untuk Junior adalah 1 x BaseSalary x prorata lama kerja
- Bonus untuk Senior adalah 2 x BaseSalary x prorata lama kerja + (PerformanceRate x BaseSalary)
- Bonus untuk Manager adalah 2 x BaseSalary x prorata lama kerja + (PerformanceRate x BaseSalary) + (BonusManagerRate x BaseSalary)

Prorata lama kerja dihitung dengan mengambil lama kerja dalam bulan dan dibagi dengan 12, jika lama kerja lebih dari 12 bulan, maka prorata lama kerja menjadi `100%` / `1`.

Contoh, lama kerja adalah 6 bulan, maka prorata lama kerja adalah 6 / 12 = `0.5`

Terdapat _interface_ Employee yang memiliki method signature `GetBonus() float64` dan lakukan implementasi _interface_ ke setiap struct yang dibuat.

Buatlah fungsi bernama `EmployeeBonus` yang menerima parameter berupa _interface_ dengan mengembalikan nilai `float64` yang merupakan perhitungan bonus dari karyawan tersebut. Dan fungsi `TotalEmployeeBonus` yang menerima parameter berupa slice _interface_ dan mengembalikan nilai `float64` yang merupakan total bonus dari seluruh karyawan.

```go
type Employee interface{
    GetBonus() float64
}

// buatlah struct dan method yang dibutuhkan
type Junior struct {}
type Senior struct {}
type Manager struct {}

func EmployeeBonus(employee Employee) float64 {
    // kerjakan di sini
}

func TotalEmployeeBonus(employees []Employee) float64 {
    // kerjakan di sini
}
```

Contoh output yang diharapkan :

<!-- perlu direvisi karena ada perubahan rule rate dll -->

```txt
EmployeeBonus

Input: Junior { Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12 }
Output: 100000

Input: Senior { Name: "Senior 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5 }
Output: 250000

Input: Manager { Name: "Manager 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5 }
Output: 300000

--------------------------------------------
TotalEmployeeBonus

Input: [
    Junior { Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12 },
    Junior { Name: "Junior B", BaseSalary: 100000, WorkingMonth: 12 },
    Junior { Name: "Junior C", BaseSalary: 100000, WorkingMonth: 12 },
]
Output: 300000

input: [
    Senior { Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5 },
    Senior { Name: "Senior B", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5 },
]

Output: 250000

input: [
    Junior { Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12 },
    Senior { Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5 },
    Manager { Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5 },
]

Output: 650000
...
```
