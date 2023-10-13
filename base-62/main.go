package main

import (
    "log"
    "github.com/amsem/base-62/utils"
)


func main()  {
    x := 4366452375678231
    str := utils.ToEncode(x)
    log.Println(str)
    num := utils.ToDeCode(str)
    log.Println(num)
}
