# Exercise Condition

## Ticket Discount

### Description

Terdapat sebuah bioskop yang menjual tiket berdasarkan kelas dengan harga sebagai berikut:

- _Ticket_ VIP : $30
- _Ticket_ Regular : $20
- _Ticket_ Student : $10

Bioskop ini memberikan diskon khusus jika terdapat pembelian _ticket_ dalam jumlah banyak, dengan total harga _ticket_ minimal $100 keatas berdasarkan kondisi berikut :

- Jika tanggal pembelian _ticket_ tersebut adalah tanggal ganjil
  - Jika _ticket_ yang dibeli kurang dari 5 _ticket_, maka mendapat diskon 15% dari total pembelian
  - Jika _ticket_ yang dibeli minimal 5 _ticket_ keatas, maka mendapat diskon 25% dari total pembelian
- Jika tanggal pembelian _ticket_ tersebut adalah tanggal genap
  - Jika _ticket_ yang dibeli kurang dari 5 _ticket_, maka mendapat diskon 10% dari total pembelian
  - Jika _ticket_ yang dibeli minimal 5 _ticket_ keatas, maka mendapat diskon 20% dari total pembelian

Terdapat fungsi bernama `GetTicketPrice` yang menerima 4 parameter, yaitu:

- `VIP` bertipe `int` yang berisi total pembelian _ticket_ VIP
- `regular` bertipe `int` yang berisi total pembelian _ticket_ Regular
- `student` bertipe `int` yang berisi total pembelian _ticket_ Student
- `day` bertipe `int` yang berisi tanggal pembelian _ticket_ dalam bentuk angka (1 - 31)

Buatlah sebuah _code_ yang dapat menghitung berapa uang yang harus dibayarkan berdasarkan _ticket_ yang dibeli dari tanggal hari ini. Kembalikan _output_ berupa total harga _ticket_ yang harus dibayar dalam tipe data `float32`.

### Constraints

- Input `VIP`, `regular`, dan `student` akan selalu bernilai positif, dengan nilai maksimal 1000
- Input `day` akan selalu bernilai positif, dengan nilai maksimal 31 (1 - 31)

### Test Case Examples

#### Test Case 1

**Input**:

```txt
VIP = 1, regular = 1, student = 1, day = 20
```

**Expected Output / Behavior**:

```txt
60
```

**Explanation**:

Total harga _ticket_ adalah (1 x $30) + (1 x $20) + (1 x $10) = $60, karena total harga tersebut tidak lebih dari $100, maka tidak ada diskon yang diberikan. Sehingga total harga yang harus dibayar adalah $60.

#### Test Case 2

**Input**:

```txt
VIP = 3, regular = 3, student = 3, day = 20
```

**Expected Output / Behavior**:

```txt
144
```

**Explanation**:

Total harga _ticket_ adalah (3 x $30) + (3 x $20) + (3 x $10) = $180. Total harga tersebut lebih dari $100, maka akan diberikan diskon dari pembelian tersebut.

Karena saat pembelian _ticket_ tersebut tanggal `20` (genap), dan total _ticket_ yang dibeli lebih dari `5` maka mendapat diskon sebesar 20%. Sehingga total harga yang harus dibayar adalah $180 - (20% x $180) = $144.

#### Test Case 3

**Input**:

```txt
VIP = 4, regular = 0, student = 0, day = 21
```

**Expected Output / Behavior**:

```txt
102
```

**Explanation**:

Total harga _ticket_ adalah (4 x $30) + (0 x $20) + (0 x $10) = $120. Total harga tersebut lebih dari $100, maka akan diberikan diskon dari pembelian tersebut.

Karena saat pembelian _ticket_ tersebut tanggal `21` (ganjil), dan total _ticket_ yang dibeli kurang dari `5` maka mendapat diskon sebesar 15%. Sehingga total harga yang harus dibayar adalah $120 - (15% x $120) = $102
