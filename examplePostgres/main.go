package main

import(
        "log"
        "github.com/amsem/examplePostgres/helper"
)


func main()  {
    _, err := helper.InitDB()

    if err != nil {
        log.Println(err)
    }else {
        log.Println("DB success")        
    }
}

