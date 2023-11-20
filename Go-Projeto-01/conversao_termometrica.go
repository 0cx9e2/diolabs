package main


import (
	"fmt"
)

const ebulicaoK = 373.0

func main(){
	var K = ebulicaoK
	var C = (K - 273)
	fmt.Println("Ponto de ebulição da água em °K é:", K)
	fmt.Println("Ponto de ebulição da água em °C é:", C)

}