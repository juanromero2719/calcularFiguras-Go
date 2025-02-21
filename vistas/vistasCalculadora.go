package vistas

import (
	"calcularFiguras/figuras"
	"fmt"
)

type VistasCalculadora struct {
	interfazFigura figuras.Figura
}

func NuevaVistaCalculadora() *VistasCalculadora {
	return &VistasCalculadora{}
}

func (v *VistasCalculadora) MostrarMenu() {
	var opcion int16
	fmt.Println("Menu de Operaciones")
	fmt.Println("1. Calcular area y perimetro de figuras")
	fmt.Println("2. Calcular funcion cuadratica")
	fmt.Println("3. Salir")
	fmt.Scanln(&opcion)

	if opcion <= 0 || opcion > 3 {
		fmt.Println("Opcion no valida")
		v.MostrarMenu()
	}

	switch opcion {
	case 1:
		MenuOperacionesFiguras()
		break
	case 2:
		MenuOperacionCuadratica()
		break
	case 3:
		fmt.Println("Hasta luego")
	}

}

func MenuOperacionesFiguras() {
	var opcion int16
	fmt.Println("Seleccione una figura:")
	fmt.Println("1. Circulo")
	fmt.Println("2. Cuadrado")
	fmt.Println("3. Hexagono")
	fmt.Println("4. Paralelogramo")
	fmt.Println("5. Pentagono")
	fmt.Println("6. Rectangulo")
	fmt.Println("7. Rombo")
	fmt.Println("8. Trapecio")
	fmt.Println("9. Triangulo")
	fmt.Println("10. Salir")
	fmt.Scanln(&opcion)

	if opcion <= 0 || opcion > 10 {
		fmt.Println("Opcion no valida")
		MenuOperacionesFiguras()
	}

}

func MenuOperacionCuadratica() {

}

func MenuCirculo() {
	var radio float64

	fmt.Println("Ingrese el radio del circulo: ")
	fmt.Scanln(&radio)
	
	interfazFigura := figuras.NuevoCirculo(radio)
}
