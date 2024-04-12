package main

import ("fmt")

type eletricEngine struct{
  kwh   float32
  mpkwh float32
}

type gasEngine struct{
  gallons float32
  mpg     float32
}

type car [T gasEngine | eletricEngine] struct{
  carMake string
  carModel string
  engine T
}

func main(){
  var gasCar = car[gasEngine]{
    carMake: "HONDA",
    carModel: "CIVIC",
    engine: gasEngine{
      gallons: 12.4,
      mpg: 40,
    },
  }

  var eletricCar = car[eletricEngine]{
    carMake: "Tesla",
    carModel: "Model 3",
    engine: eletricEngine{
      kwh: 57.5,
      mpkwh: 4.17,
    },
  }

  fmt.Println(gasCar)
  fmt.Println(eletricCar)
}
