package main

import (
  "fmt"
)

// Creating a struct/type

type gasEngine struct{
  mpg     uint8
  gallons uint8
  // owner // ownerInfo owner 
}

type eletrictEngine struct {
  mpkwh uint8
  kwh   uint8
}

// Another type inside a existent struct
// When you create your gasEngine you needed use the another struct to define that type

// type owner struct {
//   name string
// }

// Method to initialize the gasEngine
// Will be a function to my type gasEngine
func (e gasEngine) milesLeft() uint8 {
  return e.gallons * e.mpg
}

func (e eletrictEngine) milesLeft() uint8 {
  return e.kwh * e.mpkwh
}

type engine interface{
  milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
  if miles <= e.milesLeft() {
    fmt.Println("You can make it there!")
  }else{
    fmt.Println("Need to fuel up first!")
  }
}

func main(){
  // Define a struct with the values:
  // var myEngine gasEngine = gasEngine{25, 15, owner{"Marcos"}} // Or just put the values in order or can do like that:
  // myEngine.mpg = 20
  // myEngine.gallons = 20
  // // var myEngine gasEngine // To show the values of this type use .nameOfProperty like that:
  // fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

  var myGasEngine gasEngine = gasEngine{25, 20}
  var myEletrictEngine eletrictEngine = eletrictEngine{25, 15}
  // fmt.Println(myEngine.mpg, myEngine.gallons)
  // fmt.Println(myEngine.milesLeft())
  canMakeIt(myGasEngine, 20)
  canMakeIt(myEletrictEngine, 20)
}
