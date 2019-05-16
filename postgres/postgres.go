package alpostgres

import (
    "database/sql"
    "fmt"
    _"github.com/lib/pq"
) 

var db *sql.DB

func checkErr(err error){
    if err != nil{
       panic(err)
   }
   fmt.Println("success")
}

func SqlOpen() *sql.DB{
    var err error
    db, err = sql.Open("postgres", "port=5432 user=postgres password=mysecretpassword dbname=aimonitor sslmode=disable")
    checkErr(err)
    return db
}

func SqlInsert(db *sql.DB, name string, value string, change string){
    stmt, err := db.Prepare("INSERT INTO aisj(name,value,change) VALUES($1,$2,$3) RETURNING uid")
    checkErr(err)
 
    res, err := stmt.Exec(name,value,change)
    //这里的三个参数就是对应上面的$1,$2,$3了
 
    checkErr(err)
 
    affect, err := res.RowsAffected()
    checkErr(err)
 
    fmt.Println("rows affect:", affect)

}

func SqlSelect(db *sql.DB) [3][]string{
    var data [3][]string
    var values []string
    var values1 []string
    var values2 []string

    rows, err := db.Query("select * from aisj where change between '2019-05-05 19:59:43' and '2019-05-05 21:59:43'")
    checkErr(err)
    for rows.Next(){
        var uid int
        var name, value, change string
        err = rows.Scan(&uid, &name, &value, &change)
        checkErr(err)
        values = append(values, value)

    } 
    rows, err = db.Query("select * from aisj where change between '2019-05-04 19:59:43' and '2019-05-04 23:59:43'")
    checkErr(err)
    for rows.Next(){
        var uid int
        var name, value, change string
        err = rows.Scan(&uid, &name, &value, &change)
        checkErr(err)
        values1 = append(values1, value)

    }
    rows, err = db.Query("select * from aisj where change between '2019-05-01 19:59:43' and '2019-05-01 23:59:43'")
    checkErr(err)
    for rows.Next(){
        var uid int
        var name, value, change string
        err = rows.Scan(&uid, &name, &value, &change)
        checkErr(err)
        values2 = append(values2, value)

    }
    data[0] = values
    data[1] = values1
    data[2] = values2
    return data


}

func SqlClose(db *sql.DB){
    db.Close()

}


