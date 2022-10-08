# Exercise HTTP Server - Handler Middleware

## Task

Kita diminta untuk membuat aplikasi untuk mengakses halaman **student** dimana aplikasi tersebut hanya bisa di akses oleh method GET.

Maka, buatlah server yang memiliki endpoint `/student` yang memiliki middleware untuk mengatasi hanya method GET saja yang bisa mengakses halaman student tersebut yaitu `RequestMethodGet` dengan kondisi:

- Jika request method tidak sama dengan **GET**, maka return HTTP status code **405** dengan message "Method not allowed"
- Jika request dengan method GET, maka masuk ke halaman student dengan status code **200** dengan message "Welcome to Student page"

Berikut _template_ pengerjaannya, kerjakan pada fungsi `RequestMethodGet()`:

```go
func StudentHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Welcome to Student page"))
    }
}

func RequestMethodGet(next http.Handler) http.Handler {
    // kerjakan di sini
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
    // kerjakan di sini untuk routing middleware

    http.ListenAndServe("localhost:8080", nil)
}
```

## Input Format

- Method **GET**: `http://host:port/student`
- Method **POST**: `http://host:port/student`

## Output Format

- Message response for accessing with the right method is "`Welcome to Student page`" with response code **`200`**

- Message response for accessing with the wrong method is "`Method is not allowed`" with response code **`405`**

## Sample Input & Output

```bash
# Test 1:
Input  : curl -i -X GET http://localhost:8080/student
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 05:37:45 GMT
         Content-Length: 23
         Content-Type: text/plain; charset=utf-8

         Welcome to Student page

# Test 2:
Input  : curl -i -X POST http://localhost:8080/student
Output : HTTP/1.1 405 Method Not Allowed
         Date: Thu, 22 Sep 2022 05:38:21 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Method is not allowed

# Test 3:
Input  : curl -i -X DELETE http://localhost:8080/student
Output : HTTP/1.1 405 Method Not Allowed
         Date: Thu, 22 Sep 2022 05:38:21 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Method is not allowed

# Test 4:
Input  : curl -i -X PUT http://localhost:8080/student
Output : HTTP/1.1 405 Method Not Allowed
         Date: Thu, 22 Sep 2022 05:38:21 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Method is not allowed
```
