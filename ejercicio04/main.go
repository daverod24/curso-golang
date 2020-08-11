package main

import (
  "fmt"
  "os"
  "bufio"
)




var numero1 int
var numero2 int
var resultado int
var leyenda string



func main(){
  PrimerCondicional()
  // fmt.Println("este es un separador de funciones")
  // SegundoCondicional()

}


func PrimerCondicional(){

  fmt.Println("Ingrese numero 1: ")
  fmt.Scanln(&numero1)


  fmt.Println("Ingrese numero 2: ")
  fmt.Scanln(&numero2)

  fmt.Println("Ingrese Descripcion: ")

  scanner := bufio.NewScanner(os.Stdin)
  if scanner.Scan(){
    leyenda = scanner.Text()

  }
  resultado=numero1+numero2
  fmt.Println(leyenda, resultado)

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
