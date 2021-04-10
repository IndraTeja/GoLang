# GoLang MySQL 

* High performance Databases needs connection-pooling. 
* Go provides built-in standard package `database/sql` and features automatic connection-pooling. 
* Connections are setup at the start of the application from pool of connections and these connections are reused.
* This means when a query is performed, we are not creating and destroying new connections but reusing from same pool.
* Install MySQL drivers using command prompt:
  
    ``` go get -u github.com/go-sql-driver/mysql ```

* Import the packages. ```database/sql``` is in-built in Go.

    ```go
    package main

    // imported MySQL drivers
    import (
        "fmt"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
    )
    ```
