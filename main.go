//Name: Vincent G.
//Date Completed: 05/08/2020
//Description: Planets (2 points) 4.2.1 Structs

package main

import "fmt"

//Create a struct for planet (name, gravity, number of Moons).

type person struct{

  name string
  weight float64

}
  
type planet struct{
  name string
  gravity float64
  numberOfmoons int
}

func main() {

//Use the struct to Calculate the weight of a person on the planet and earth .

var earth = planet{name: "earth", gravity: 1.0, numberOfmoons: 1}

var mars = planet{name: "mars", gravity: .38, numberOfmoons: 2}

var Tim = person{name: "Tim", weight: 179.0}
//Tims gravity on mars
a := Tim.weight * mars.gravity 
//Tims gravity on earth
b := Tim.weight * earth.gravity

fmt.Println(a, b)

}