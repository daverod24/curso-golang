package main

import (
	"fmt"
)



var estado bool

func main(){
  PrimerCondicional()
  fmt.Println("este es un separador de funciones")
  SegundoCondicional()

}


func PrimerCondicional(){
  estado=true
  if estado=false; estado==true{
    fmt.Println(estado)
  }else{
    fmt.Println("es falso")
  }


}
func SegundoCondicional(){
  switch numero := 10; numero{
  case 1:
    fmt.Println(1)
  case 2:
    fmt.Println(2)
  case 3:
    fmt.Println(3)
  case 4:
    fmt.Println(4)
  case 5:
    fmt.Println(5)
  default:
    fmt.Println("Mayor a 5")
  }


}
