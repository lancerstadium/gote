# Gote
- A vote system in go.

## 1 env
1. lang: gin, gorm, go-1.22
2. db: MySQL
3. dc: Redis

> Source:
>   1. gin  : `go get -u github.com/gin-gonic/gin`
>   2. gorm : `go get -u gorm.io/gorm`
>   3. mysql: `go get -u github.com/go-sql-driver/mysql gorm.io/driver/mysql`

## 2 use

```shell
make            # build & run program
make clean      # delete program
```

## 3 MVC

```
app
 ├─ logic               # [C]: Controller
 │ ├─ index.go          
 │ └─ login.go
 ├─ model               # [M]: Data Model
 │ ├─ db.go
 │ └─ user.go
 ├─ router              # Web Router
 │ └─ router.go
 ├─ tools               # Tool Packages
 ├─ view                # [V]: Web Pages
 │ ├─ index.tmpl
 │ └─ login.tmpl
 └─ app.go              # << App Entry

```


## 4 func

### 4.1 login


### 4.2 cookie


### 4.3 vote

- Relation of data tables: 
    - `<vote, vote_opt>` = 1:n
    - `<user, vote>` = n:n (must 3 tbls)


