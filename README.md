
# Golang CMS

Build fast and flexible websites in Golang.

# How to use

Change your mysql credentials and database name in "config/db.go".

# MYSQL Structure

Create database:

```sql
CREATE DATABASE golangcms;
```

Create users table:

```sql
CREATE TABLE users (
    Id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(50) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Admin BOOLEAN
);
```

Create articles table:

```sql
CREATE TABLE articles (
    Id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Teaser TEXT,
    Body LONGTEXT,
    Image TEXT,
    Tags TEXT,
    Author VARCHAR(255),
    Url TEXT,
    Changed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Uncomment the line "demo.User()" inside of main.go.

"demo.User()" will gerenate a demo user:

> **Name:** admin

> **Email:** admin@admin.com

> **Password:** admin

Execute:

```sh
$ go run main.go
```

After running for the first time, comment "demo.User()" otherwise that will generate a new user again.

**Enjoy!**

# TODO

* CRUD User validation
* CRUD Article validation
* User login validation

# Packages

[github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

[github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

[github.com/gorilla/securecookie](https://github.com/gorilla/securecookie)

[golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

[gopkg.in/gomail.v2](https://gopkg.in/gomail.v2)

[github.com/go-redis/redis](https://github.com/go-redis/redis)