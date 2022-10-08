# Exercise HTTP Server - Error Handling

## Task

Buatlah **error handling** dan lakukan pengecekan terhadap handler berikut:

- **Method**:
  - **Method Handler** dengan kondisi jika request method != **GET** maka return error code **405** dengan message "`Method not allowed`".
  - **Method Handler** dengan kondisi jika request == **GET** maka return status code **200** dengan message "`Method handler passed`".
- **Data**
  - **Data Handler** dengan kondisi jika request tidak memiliki parameter, maka return error code **404** dengan message "`Data not found`".
  - **Data Handler** dengan kondisi jika request memiliki parameter `data` == `<text bebas>`, maka return status code **200** dengan message "`Data handler passed`".
- **Open File**
  - **Open File Handler** dengan kondisi jika request tidak memiliki parameter, maka handler akan return error code **500** dengan message "`File not found`".
  - **Open File Handler** dengan kondisi jika request memiliki parameter `filename` == `wrong.txt` , maka handler akan return status code **500** dengan message "`File not found`".
  - **Open File Handler** dengan kondisi jika request memiliki parameter `filename` == `hello.txt` , maka handler akan return status code **200** dengan message "`Error handler passed`".

Berikut _template_ pengerjaannya:

```go
func MethodGet(r *http.Request) error {
    if r.Method != http.MethodGet {
        return fmt.Errorf("Method not allowed")
    }
    return nil
}

func CheckDataRequest(r *http.Request) error {
    data := r.URL.Query().Get("data")
    if len(data) == 0 {
        return fmt.Errorf("Data not found")
    }
    return nil
}

func CheckOpenFile(r *http.Request) error {
    filename := r.URL.Query().Get("filename")
    _, err := os.Open(filename)
    if err != nil {
        return err
    }
    return nil
}

func MethodHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := MethodGet(r)
        fmt.Println(err) // TODO: replace this
    }
}

func DataHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := CheckDataRequest(r)
        fmt.Println(err) // TODO: replace this
    }
}

func OpenFileHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := CheckOpenFile(r)
        fmt.Println(err) // TODO: replace this
    }
}
```

## Sample Input & Output

```bash
# Test 1 Method Handler
Input  : Request Method != GET 
Output : Status Code: 405
         Message: Method not allowed

# Test 2 Method Handler
Input  : Request Method == GET
Output : Status Code: 200
         Message: Method handler passed

# Test 3 Data Handler
Input  : Request without parameter
Output : Status Code: 404
         Message: Data not found

# Test 4 Data Handler
Input  : Request with parameter data=Aditira
Output : Status Code: 200
         Message: Data handler passed

# Test 5 Open File Handler
Input  : Request without parameter
Output : Status Code: 500
         Message: File not found

# Test 6 Open File Handler
Input  : Request with parameter filename=wrong.txt
Output : Status Code: 500
         Message: File not found

# Test 7 Open File Handler
Input  : Request with parameter filename=hello.txt
Output : Status Code: 200
         Message: Error handler passed
```
