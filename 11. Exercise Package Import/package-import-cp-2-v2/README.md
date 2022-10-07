# Exercice package & import

## Mini Cashier app

### Description

Butlah sebuah aplikasi yang dapat digunakan untuk menjadi _cashier_ mini. Aplikasi ini dapat digunakan untuk mencatat list barang yang dibeli, menghitung total harga dan total kembalian dari pembayaran _customer_.

Aplikasi ini memiliki list barang yang harus disimpan berserta harga per barang, yaitu:

```txt
Kaos Polos      Rp 50.000
Kaos sablon     Rp 70.000
Baju Batik      Rp 200.000
Celana hitam    Rp 150.000
Celana pendek   Rp 100.000
Sabuk           RO 50.000
Topi            Rp 75.000
Galung Tangan   Rp 30.000
Sepatu          Rp 300.000
```

Berikut struktur directory dari project ini:

```txt
database
├── database.go
entity
├── entity.go
service
├── service.go
main.go
go.mod
```

### `entity`

Package `entity` akan berisi semua objek atau `struct` yang dibutuhkan untuk project ini. Di package ini terdapat 3 _struct_ yang dibutuhkan untuk project ini, yaitu:

- `CartItem` yang digunakan untuk tempat penyimpanan keranjang pembelian. _Struct_ ini berisi `ProductName`, `Price` dan `Quantity` yang dibeli oleh pelanggan

- `Product` yang digunakan untuk menyimpan data barang yang tersedia. _Struct_ ini berisi `Name` dan `Price`

- `PaymentInformation` yang digunakan untuk menyimpan data pembayaran. _Struct_ ini berisi _list_ produk yang dibeli, total harga, uang yang dibayarkan dan total kembalian.

### `database`

Package `database` akan menjadi tempat penyimpanan data-data yang dibutuhkan (berisi cart) oleh aplikasi ini. Data-data ini akan disimpan dalam bentuk _array of struct_ `CartItem`. Terdapat juga data produk yang tersedia dan sudah disimpan ke dalam _array of struct_ `Product`.

Terapat juga method yang tersedia di package ini untuk kebutuhan mendapatkan data atau menyimpan data ke database.

- `Load` untuk mendapatkan data dari database
- `Save` untuk menyimpan data baru ke database
- `GetProductData` untuk mendapatkan semua data produk yang tersedia
- `GetProductByName` untuk mendapatkan data produk berdasarkan nama produk

Kita akan menerapkan _interface_ agar package ini dapat digunakan untuk package lain sebagai _dependency injection_.

### `service`

Package ini akan kita gunakan untuk melakukan proses bisnis dari aplikasi ini. Package ini akan mengolah data-data yang ada di package `database` dan mengembalikan data-data yang sudah diolah ke package `main`. Buatlah semua proses bisnis yang dibutuhkan di package ini dengan menggunakan _interface_ dari package `database` dan melengkapi method yang sudah disediakan

- `AddCart` adalah fungsi yang digunakan untuk menambah keranjang produk ke database. Method ini akan menerima parameter `ProductName` dan `Quantity` yang akan ditambahkan ke database. Method ini akan mengembalikan `error` jika :
  - Produk yang ingin ditambahkan tidak ditemukan di database, kembalikan error `"product not found"`
  - `Quantity` bernilai 0 atau negatif, kembalikan error `"invalid quantity"`
  - Terjadi kesalahan saat melakukan _load_ data dari database
  - Terjadi kesalahan saat melakukan _save_ data ke database

- `RemoveCart` adalah method yang digunakan untuk menghapus produk dari keranjang. Method ini akan menerima parameter `productName` sebagai nama product yang akan dihapus dari database. Method ini akan mengembalikan `error` jika:
  - Produk yang ingin dihapus tidak ditemukan di database. kembalikan error `"product not found"`
  - Terjadi kesalahan saat melakukan _load_ data dari database
  - Terjadi kesalahan saat melakukan _save_ data ke database

- `ShowCart` adalah method yang digunakan untuk menampilkan data keranjang. Method ini akan mengembalikan _list_ `CartItem` yang ada di database. Method ini akan mengembalikan `error` jika:
  - Terjadi kesalahan saat melakukan _load_ data dari database

- `ResetCart` adalah method yang digunakan untuk menghapus semua data keranjang. Method ini akan mengembalikan `error` jika:
  - Terjadi kesalahan saat melakukan _load_ data dari database
  - Terjadi kesalahan saat melakukan _save_ data ke database

- `GetAllProduct` adalah method yang digunakan untuk mendapatkan semua data produk yang tersedia. Method ini akan mengembalikan _list_ `Product` yang ada di database. Method ini akan mengembalikan `error` jika:
  - Terjadi kesalahan saat melakukan _load_ data dari database

- `Paid` adalah method yang digunakan untuk melakukan pembayaran. Method ini akan menerima parameter `money` bertipe `int` yang berupa uang yang dibayarkan oleh pelanggan. Method ini akan mengembalikan `PaymentInformation` yang berisi data pembayaran. Selain itu, pemanggilan method `Paid` yang berhasil (tidak error) akan membersihkan / me-reset `cart`. Method ini akan mengembalikan `error` jika :
  - Uang yang dibayarkan kurang dari total harga, kembalikan error `"money is not enough"`
  - Terjadi kesalahan saat melakukan _load_ data dari database
  - Terjadi kesalahan saat melakukan _save_ data ke database

### Test Case Examples

#### Test Case AddCart

- input

```txt
productName = "Kaos Polos"
quantity = 2
```

Ketika kita menambahkan produk `Kaos Polos` sebanyak 2 ke keranjang, maka data di database harusnya menjadi seperti berikut:

```txt
database = [
    { productName: "Kaos Polos", price: 50000, quantity: 2 }
]
```

#### Test Case ShowCart

- output

```txt
[
    { productName: "Kaos Polos", price: 50000, quantity: 2 }
]
```

Fungsi ini akan mengembalikan data keranjang yang ada di database. Karena sebelumnya kita sudah menambahkan produk `Kaos Polos` sebanyak 2 ke keranjang, maka data yang dikembalikan harusnya sama dengan data di database.

#### Test Case RemoveCart

- input

```txt
productName = "Kaos Polos"
```

Ketika kita menghapus produk `Kaos Polos` dari keranjang, maka data di database harusnya menjadi seperti berikut:

```txt
[]
```

Data di database harusnya kosong karena di database cuma ada produk `Kaos Polos` yng sudah dihapus dari keranjang.

#### Test Case ResetCart

- output

```txt
[]
```

Fungsi ini akan menghapus semua data keranjang yang ada di database. Seberapa banyakpun datanya, setelah fungsi ini dijalankan, data di database harusnya kosong.

#### Test Case GetAllProduct

- output

```txt
[{
    productName: "Kaos Polos",
    price: 50000

},{
    productName: "Kemeja Polos",
    price: 100000

},{
    productName: "Celana Panjang",
    price: 150000

},{
    productName: "Celana Pendek",
    price: 100000
    ...
]
```

Fungsi ini akan mengembalikan semua data produk yang ada di database. Data yang dikembalikan harusnya sama dengan data di database.

#### Test Case Paid

Anggap kita menambahkan 2 produk `Kaos Polos` dan 1 produk `Kaos Sablon` ke keranjang.

- `AddCart` input

```txt
productName = "Kaos Polos"
quantity = 2
```

- `AddCart` input

```txt
productName = "Kaos Sablon"
quantity = 1
```

Kemudian kita jalankan `Paid` dengan input:

```txt
money = 500000
```

- output

```txt
{
    ListProduct: [
        { productName: "Kaos Polos", price: 50000, quantity: 2 },
        { productName: "Kaos Sablon", price: 70000, quantity: 1 }
    ],
    TotalPrice: 170000,
    MoneyPaid: 500000,
    Change: 330000
}
```

Karena kita menambahkan 2 produk `Kaos Polos` dan 1 produk `Kaos Sablon` ke keranjang, maka total harga yang harus dibayar adalah `170000` (2 x 50000 + 1 x 70000). Karena kita membayar dengan uang `500000`, maka uang kembalian yang harus diberikan adalah `330000` (500000 - 170000).
