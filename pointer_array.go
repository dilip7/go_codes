package main

import "fmt"

const MAX int = 3

func main(){
    a := []int {1,2,3}

    var i int
    var ptr[MAX] *int

    for i=0; i< MAX;i++{
        ptr[i] = &a[i]
    }

    for i=0; i< MAX;i++{
        fmt.Printf("Value of a[%d] is %d \n",i,*ptr[i])
    }

}


