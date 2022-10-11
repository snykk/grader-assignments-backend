# HTTP Server

## TODO List App

### Description

Disini kamu diminta untuk membuat aplikasi bernama todo app yang berfungsi untuk membuat list beserta status pekerjaan dari aktifitas yang kita kerjakan. Fitur yang yang harus dibuat dari aplikasi ini adalah:

- Register
- Login
- Create Todo list
- Read Todo list
- Clear Todo List
- Logout

Buatlah chain middleware untuk menghandle Method dan Authentication dengan menggunakan metode session based token lalu simpan semua data user dan todo list di _in memory map_.

### Constraints

Berikut adalah hal-hal yang harus diperhatikan dalam mengerjakan aplikasi todo app ini:

- **Register** dengan method **POST** dan end point `/user/register`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan responnse message `{"error":"Method is not allowed!"}`
  - Harus memberikan request body dengan format JSON `{"username": "<username>", "password": <password>}`
    - Jika **tidak** memberikan request body, maka:
      - Berikan response code **400**
      - Berikan responnse message `{"error":"Internal Server Error"}`
    - Jika request body username dan password kosong(`""`), maka:
      - Berikan response code **400**
      - Berikan responnse message `{"error":"Username or Password empty"}`
    - Jika semua username sudah ada di database, maka:
      - Berikan response **409**
      - Berikan responnse message `{"error":"Username already exist"}`
    - Jika semua Oke, maka:
      - Berikan response **200**
      - Berikan response message `{"username":"<username>","message":"Register Success"}`
  - Simpan request body tersebut di memory map pada variable `Users` yaitu:

    ```go
    map[string]string{
        "<username>":"password"
    }
    ```

- **Login** dengan method **POST** end point `/user/login`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus memberikan request body dengan format JSON `{"username": "<username>", "password": <password>}`
    - Jika **tidak** memberikan request body, maka:
      - Berikan response code **400**
      - Berikan responnse message `{"error":"Internal Server Error"}`
    - Jika request body username dan password kosong(`""`), maka:
      - Berikan response code **400**
      - Berikan responnse message `{"error":"Username or Password empty"}`
  
    - jia semua Oke, maka check keberadaan username dan kesesuaian password pada request di memory map pada variable `Users` yaitu:

    ```go
    map[string]string{
        "<username>":"password"
    }
    ```

    - Jika _username_ dan _password_ **tidak ada**, maka:
      - Berikan response code **401**
      - Berikan response message `"Wrong User or Password!"`
    - Jika _username_ dan _password_ **ada**, maka:
      - Berikan response code **200**
      - Berikan response message `{"username":"<username>","message":"Login Success"}`
      - Buatlah cookie dengan ketentuan:
        - Name:    `"session_token"`
        - Value:   `<nilai _random session token_ menggunakan github.com/google/uuid>`
        - Expires: `<expired time dengan durasi 5 jam>`
      - Lakukan penyimpanan Cookie pada memory map pada variable `Sessions` yaitu:

        ```go
        map[string]model.Session{
            `<session_token>`: {
              Username `<username>`
              Expiry   `<expired_time>`
            }
        }
        ```

- **Logout** dengan method **GET**, **Authentication** dan end point `/user/logout`
  - Hanya boleh menggunakan method **GET** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus menggunakan cookie yang bernama `"session_token"` jika tidak, maka:
    - Berikan response code **401**
    - Berikan response message `{"error":"http: named cookie not present"}`
  - Jika menggunakan cookie maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<username>","message":"Logout success"}`
    - Hapus data sessions user yang melakukan logout pada memory dari variable `Sessions`
- **Create Todo List** dengan method **POST**, **Authentication**, end point `/todo/create` dan handle function `AddToDo`
  - Hanya boleh menggunakan method **POST** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus menggunakan cookie yang bernama `"session_token"` jika tidak, maka:
    - Berikan response code **401**
    - Berikan response message `{"error":"http: named cookie not present"}`
  - Harus memberikan request body dengan format JSON `{"task": "<task>", "done": <boolean>}` dan menggunakan cookie maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<username>","message":"Task <task-name> added!"}`
  - Simpan request body tersebut di memory map pada variable `Task` yaitu:

    ```go
    map[string][]model.Todo{
        `<username>`:[
            {
                "id"  : "uuid" `<auto_generate>`
                "task": "<task>"
                "done": <bool>
            }
        ]
    }
    ```

- **Read Todo List** dengan method **GET**, **Authentication**, end point `/todo/read` dan handle function `ListToDo`
  - Hanya boleh menggunakan method **GET** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus menggunakan cookie yang bernama `"session_token"` jika tidak, maka:
    - Berikan response code **401**
    - Berikan response message `{"error":"http: named cookie not present"}`
  - Jika menggunakan cookie dan namun todolist **tidak** ditemukan di memory map pada variable `Task`, maka:
    - Berikan response code **404**
    - Berikan response message `{"error": "Todolist not found!"}`
  - Jika menggunakan cookie dan todolist ditemukan di memory map pada variable `Task`, maka:
    - Berikan response code **200**
    - Berikan response message `[{"id":"<uuid auto generate>", "task":"<task>", "done": <bool>}]`
- **Clear Todo List** dengan method **DELETE**, **Authentication**, end point `/todo/clear` dan handle function `ClearToDo`
  - Hanya boleh menggunakan method **DELETE** jika tidak, maka:
    - Berikan response code **405**
    - Berikan response message `{"error":"Method is not allowed!"}`
  - Harus menggunakan cookie yang bernama `"session_token"` jika tidak, maka:
    - Berikan response code **401**
    - Berikan response message `{"error":"http: named cookie not present"}`
  - Jika menggunakan cookie maka:
    - Berikan response code **200**
    - Berikan response message `{"username":"<username>","message":"Clear ToDo Success"}`
    - Hapus semua data task user yang melakukan clear pada memory dari variable `Task`

- **Detail Penting!**
  - Output dari setiap response harus dalam format JSON yang di encode dari struct:

    ```go
    type SuccessResponse struct {
      Username string `json:"username"`
      Message  string `json:"message"`
    }
    ```

    Jika **response success** dan struct:

    ```go
    type ErrorResponse struct {
      Error string `json:"error"`
    }
    ```

    Jika **response error**.

  - Gunakan contex untuk mendefinisikan state yang menunjukkan user sedang login dengan mengirimkan `username` sebagai success response.
  - Sessions token dikirim dari user setiap kali request ke endpoint:
    - `/todo/create`
    - `/todo/read`
    - `/todo/clear`

### Test Case Example

#### Test Case 1

**Input**:

```bash
> curl -i -X GET http://localhost:8080/user/register
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 405 Method Not Allowed
Date: Wed, 28 Sep 2022 06:33:20 GMT
Content-Length: 35
Content-Type: text/plain; charset=utf-8

{"error":"Method is not allowed!"}
```

#### Test Case 2

**Input**:

```bash
> curl -i -X POST http://localhost:8080/user/register -H "Content-Type:application/json" -d "{ \"username\":\"aditira\", \"password\":\"12345\" }" 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Sep 2022 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"aditira","message":"Register Success"}
```

#### Test Case 3

**Input**:

```bash
> curl -i -X POST http://localhost:8080/user/login -H "Content-Type:application/json" -d "{ \"username\":\"aditira\", \"password\":\"12345\" }" 
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Sep 2022 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"aditira","message":"Login Success"}
```

#### Test Case 4

**Input**:

```bash
> curl -X POST http://localhost:8080/login -H "Content-Type:application/json" -d "{ \"task\":\"Bersepeda\", \"done\": false }" --cookie "session_token=3e14a989-9a1b-4635-846d-d1ea2a7ebcb8"
```

**Expected Output / Behavior**:

```bash
HTTP/1.1 200 OK
Date: Wed, 28 Sep 2022 06:31:50 GMT
Content-Length: 52
Content-Type: text/plain; charset=utf-8

{"username":"aditira","message":"Task aditira added!"}
```

### Template

Project structure:

```txt
- 📁 todo-app-assignment
  - 📁 db
    - 📄 database.go
  - 📁 middleware
    - 📄 method.go
    - 📄 user.go
  - 📁 model
    - 📄 model.go
  - 📄 main.go
```

- db/database.go

  ```go
  var Users = map[string]string{}
  var Task = map[string][]model.Todo{}

  var Sessions = map[string]model.Session{}
  ```

- middleware/method.go
  
  ```go
  func Get(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // kerjakan di sini
    })
  }

  func Post(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // kerjakan di sini
    })
  }

  func Delete(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // kerjakan di sini
    })
  }
  ```

- middleware/user.go
  
  ```go
  func isExpired(s model.Session) bool {
    return s.Expiry.Before(time.Now())
  }

  func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // kerjakan di sini
    })
  }
  ```

- model/model.go

  ```go
  import "time"

  type Todo struct {
    Id   string `json:"id"`
    Task string `json:"task"`
    Done bool   `json:"done"`
  }

  type Session struct {
    Username string
    Expiry   time.Time
  }

  type Credentials struct {
    Password string `json:"password"`
    Username string `json:"username"`
  }

  type ErrorResponse struct {
    Error string `json:"error"`
  }

  type SuccessResponse struct {
    Username string `json:"username"`
    Message  string `json:"message"`
  }
  ```

- main.go

  ```go
  func Register(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func Login(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func AddToDo(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func ListToDo(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func ClearToDo(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func Logout(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
  }

  func ResetToDo(w http.ResponseWriter, r *http.Request) {
    db.Task = map[string][]model.Todo{}
    w.WriteHeader(http.StatusOK)
  }

  type API struct {
    mux *http.ServeMux
  }

  func NewAPI() API {
    mux := http.NewServeMux()
    api := API{
        mux,
    }

    mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
    mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
    mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

    //kerjakan di sini untuk routing /todo

    mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

    return api
  }

  func (api *API) Handler() *http.ServeMux {
    return api.mux
  }

  func (api *API) Start() {
    fmt.Println("starting web server at http://localhost:8080")
    http.ListenAndServe(":8080", api.Handler())
  }

  func main() {
    mainAPI := NewAPI()
    mainAPI.Start()
  }
  ```
