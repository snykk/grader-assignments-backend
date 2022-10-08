# Exercise HTTP Server - Chaining Middleware

## Task

Kita diminta untuk membuat aplikasi yang memiliki halaman **admin** sekolah dimana aplikasi tersebut hanya bisa di akses oleh role ADMIN dan hanya method GET yang bisa mengakses aplikasi tersebut.

Maka, buatlah server yang memiliki endpoint `/admin` dengan menggunakan metode chaining middleware dengan 2 handler middleware yaitu:

- `RequestMethodGetMiddleware` yang akan melakukan pengecekan terhadap request method dimana:
  - Jika request method selain **GET**, maka return error dengan status code **405** dengan message "Method is not allowed"
  - Jika method **GET**, lanjut ke handler berikutnya
- `AdminMiddleware` yang akan melakukan pengecekan terhadap header "role" dimana:
  - Jika nilai header "role" tidak sama dengan **ADMIN**, maka return error dengan status code **401** dengan message "Role not authorized"
  - Jika nilai header "role" sama dengan **ADMIN**, maka lanjut ke handler berikutnya
- Jika kedua middleware di atas lolos, maka akan masuk ke handler function `AdminHandler()` dengan status code **200** dengan message "Welcome to Admin page"

Berikut _template_ pengerjaannya, kerjakan pada fungsi `RequestMethodGetMiddleware()` dan `AdminMiddleware()`:

```go
func AdminHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to Admin page"))
  }
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
    // kerjakan di sini
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
    // kerjakan di sini
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
    // kerjakan di sini untuk routing chanining middleware

    http.ListenAndServe("localhost:8080", nil)
}
```

## Input Format

- Method **POST**: `http://host:port/admin`
- Method **GET**: `http://host:port/admin` with Header role USER
- Method **GET**: `http://host:port/admin` with Header role ADMIN

## Output Format

- Message wrong method is "`Method is not allowed`" with response code **`405`**
- Message wrong role is "`Role not authorized`" with response code **`401`**
- Message data exists is "`Welcome to Admin page`" with response code **`200`**

## Sample Input & Output

```bash
# Test 1:
Input  : curl -i -X POST http://localhost:8080/admin
Output : HTTP/1.1 405 Method Not Allowed
         Date: Wed, 21 Sep 2022 17:21:15 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Method is not allowed

# Test 2:
Input  : curl -i -X GET http://localhost:8080/admin -H role:USER
Output : HTTP/1.1 401 Unauthorized
         Date: Wed, 21 Sep 2022 17:22:49 GMT
         Content-Length: 19
         Content-Type: text/plain; charset=utf-8

         Role not authorized

# Test 3:
Input  : curl -i -X GET http://localhost:8080/admin -H role:ADMIN
Output : HTTP/1.1 200 OK
         Date: Wed, 21 Sep 2022 17:23:47 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Welcome to Admin page
```
