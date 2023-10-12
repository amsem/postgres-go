package main

import(
        "log"
        "github.com/amsem/examplePostgres/helper"
)


func main()  {
    _, err := helper.InitDB()

    if err != nil {
        log.Println(err)
    }
    log.Println("DB success")
}

