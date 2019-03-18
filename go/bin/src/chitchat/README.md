# ChitChat

This is the simple forum application written with Go for the book "Go Web Programming" from Manning Publication Co. The source code for this application is the basis for the Chapter 2 - Go ChitChat.

However the code is a reference and is not a 1 - 1 match with the code in the book. There are portions of the code that is more detailed in here than in Chapter 2 (which is a simplified version of the source code here).

Some differences include:

* This version of ChitChat is configurable with a `config.json` file
* This version is able to log to a `chitchat.log` file
* There are test files in this code
* All structs are fully fleshed out (in the book chapter, they are only implied)
* Some of the functions are placed as methods for the struct types instead of being a part of the package

NOTE: I used this web app as a learning example and is not my original work

Made some changes from the original author:
(changes made by hashsequence)
* modified Db initialization in data.go for my local pc
```go
const (
  host     = "localhost"
  port     =  5432
  user     = "postgres"
  password = "password"
  dbname   = "chitchat"
)

var Db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
 fmt.Println("db init at port ", port)
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("db failed")
		log.Fatal(err)
	}
	fmt.Println("db success")
	return
}
```
