# Exercise HTTP Server - Response Code

## Task

Di sekolah, kita diminta membuat aplikasi untuk mengecek apakah nama yang di input sudah terdaftar di _list of slice_. Jadi kita membuat API agar memberikan response code sesuai dengan status dari ketersediaan data tersebut.

Maka, buatlah server yang memiliki endpoint `/students` dengan paremeter `name` yang mengembalikan response code dari beberapa **kondisi** berikut ini:

- Jika pada url parameter `name` **tidak ditemukan nama** pada data slice string `students`, maka return response code **404** dengan message "Data not found"
- Jika pada url parameter `name` **ditemukan nama** pada data slice string `students`, maka return response code **200** dengan message "Name is exists"
- Jika request method selain **GET**, maka return response code **405** dengan message "Method is not allowed"

Berikut _template_ pengerjaannya, kerjakan pada fungsi `CheckStudentName()`:

```go
var students = []string{
  "Aditira",
  "Dito",
  "Afis",
  "Eddy",
}

func IsNameExists(name string) bool {
  for _, n := range students {
    if n == name {
      return true
    }
  }

  return false
}

func CheckStudentName() http.HandlerFunc {
  // kerjakan di sini
  return nil // TODO: replace this
}

func GetMuxResCode() *http.ServeMux {
  mux := http.NewServeMux()
  // kerjakan di sini untuk routing
  return mux
}

func main() {
  http.ListenAndServe("localhost:8080", GetMux())
}
```

## Input Format

- Method **GET**: `http://host:port/students`
- Method **GET**: `http://host:port/students?name=<name value>`
- Method **POST**: `http://host:port/students?name=<name value>`

## Output Format

- Message not found is "`Data not found`" with response code **`404`**
- Message data exists is "`Name is exists`" with response code **`200`**
- Message wrong method is "`Method is not allowed`" with response code **`405`**

## Sample Input & Output

```bash
# Test 1:
Input  : curl -i -X GET http://localhost:8080/students
Output : HTTP/1.1 404 Not Found
         Date: Wed, 21 Sep 2022 15:34:10 GMT
         Content-Length: 14
         Content-Type: text/plain; charset=utf-8

         Data not found

# Test 2:
Input  : curl -i -X GET http://localhost:8080/students?name=Aditira
Output : HTTP/1.1 200 OK
         Date: Wed, 21 Sep 2022 15:37:40 GMT
         Content-Length: 14
         Content-Type: text/plain; charset=utf-8

         Name is exists

# Test 3:
Input  : curl -i -X GET http://localhost:8080/students?name=Dito
Output : HTTP/1.1 200 OK
         Date: Wed, 21 Sep 2022 15:37:40 GMT
         Content-Length: 14
         Content-Type: text/plain; charset=utf-8

         Name is exists

# Test 4:
Input  : curl -i -X GET http://localhost:8080/students?name=Afis
Output : HTTP/1.1 200 OK
         Date: Wed, 21 Sep 2022 15:37:40 GMT
         Content-Length: 14
         Content-Type: text/plain; charset=utf-8

         Name is exists

# Test 5:
Input  : curl -i -X GET http://localhost:8080/students?name=Eddy
Output : HTTP/1.1 200 OK
         Date: Wed, 21 Sep 2022 15:37:40 GMT
         Content-Length: 14
         Content-Type: text/plain; charset=utf-8

         Name is exists

# Test 6:
Input  : curl -i -X POST http://localhost:8080/students?name=Aditira
Output : HTTP/1.1 405 Method Not Allowed
         Date: Wed, 21 Sep 2022 15:38:29 GMT
         Content-Length: 21
         Content-Type: text/plain; charset=utf-8

         Method is not allowed
```
