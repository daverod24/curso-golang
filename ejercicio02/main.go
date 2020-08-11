package main

import (
	"strconv"
	"fmt"
)


var numero int
var texto string
var status bool = true

func main(){
  numero2, numero3, numero4, texto, status := 2, 5, 67, "este es mi texto de golang", false

  var moneda int64 = 0
  numero2 = int(moneda)
  //texto = fmt.Sprintf("%d", moneda)
  texto = strconv.Itoa(int(moneda))


  fmt.Println(numero2)
  fmt.Println(numero3)
  fmt.Println(numero4)
  fmt.Println(texto)
  fmt.Println(status)


  MostrarStatus()

}

func MostrarStatus(){
  fmt.Println(status)
}
