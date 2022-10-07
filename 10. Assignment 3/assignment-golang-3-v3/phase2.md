# Phase 2

Kita akan mengimplementasikan _interface_ untuk _struct_ inovice departemen warehouse. Buatlah method yang dibutuhkan di file `warehouse.go` agar dapat melakukan konversi data invoice warehouse ke format yang sebelumnya sudah dijelaskan di **phase 1**.

Untuk `TotalInvoice` dari invoice warehouse, diakumulasikan pada setiap product dengan rumus sebagai berikut:

```txt
unit x price - discout price
```

Contoh, jika terdapat 2 product dengan total 2 dan 3, harga 10000 dan 20000, dan diskon `0.1` dan `0.2`, maka total invoice adalah:

- product pertama, `2 x 10000 = 18000` dikurangi diskon `0.1` atau 10% (`1800`) maka totalnya adalah `16200`
- product kedua, `3 x 20000 = 60000` dikurangi diskon `0.2` atau 20% (`12000`) maka totalnya adalah `48000`

Maka total invoice adalah `16200 + 48000 = 64200`

Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong `"DD-month-YYYY"` maka return `error` dengan pesan `"invoice date is empty"`
- Jika `InvoiceType` berisi string kosong atau selain "purchase" dan "sales", maka return `error` dengan pesan `"invoice type is not valid"`
- Jika isi dari `Products` kosong maka return `error` dengan pesan `"invoice products is empty"`
- Jika terdapat total pembelian product (`Unit`) yang bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"unit product is not valid"`
- Jika terdapat harga product (`Price`) yang bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"price product is not valid"`

## Test Case

**Input**:

```txt
WarehouseInvoice{
    Date: "01-January-2022",
    InvoiceType: "purchase",
    Approved: true,
    Products: []Product{
        {
            Unit: 2,
            Price: 10000,
            Discount: 0.1,
        },
        {
            Unit: 3,
            Price: 20000,
            Discount: 0.2,
        },
    },
}
```

**Output**:

```txt
InvoiceData{
    Date: "01-January-2022",
    TotalInvoice: 150000,
    Department: "warehouse",
}
```

Format date akan tetap menjadi `DD-month-YYYY` seperti contoh di atas, dan hasil `TotalInvoice` akan diakumulasikan dari `(unit x price - discout price )` yang ada di `Products` yaitu `(2 x 10000 - 1000) + (3 x 20000 - 4000) = 150000`.  Dan terakhir `Department` akan bernilai `"warehouse"`.
