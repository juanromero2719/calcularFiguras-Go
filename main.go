package main

import (
	"calcularFiguras/figuras"
	"fmt"
)

func main() {

	// interfaz
	var interfazFigura figuras.Figura

	// triangulo
	triangulo := figuras.NuevoTriangulo(10, 5)
	interfazFigura = triangulo
	areaTriangulo := interfazFigura.CalculateArea()

	// cuadrado
	cuadrado := figuras.NuevoCuadrado(10)
	interfazFigura = cuadrado
	areaCuadrado := interfazFigura.CalculateArea()

	// rectangulo
	rectangulo := figuras.NuevoRectangulo(10, 2)
	interfazFigura = rectangulo
	areaRectangulo := interfazFigura.CalculateArea()

	// rombo
	rombo := figuras.NuevoRombo(6, 5)
	interfazFigura = rombo
	areaRombo := interfazFigura.CalculateArea()

	// pentagono
	pentagono := figuras.NuevoPentagono(4, 4)
	interfazFigura = pentagono
	areaPentagono := interfazFigura.CalculateArea()

	// hexagono
	hexagono := figuras.NuevoHexagono(6, 5)
	interfazFigura = hexagono
	areaHexagono := interfazFigura.CalculateArea()

	// circulo
	circulo := figuras.NuevoCirculo(3)
	interfazFigura = circulo
	areaCirculo := interfazFigura.CalculateArea()

	// trapecio
	trapecio := figuras.NuevoTrapecio(8, 6, 4)
	interfazFigura = trapecio
	areaTrapecio := interfazFigura.CalculateArea()

	fmt.Println("Área del triángulo:", areaTriangulo, "cm²")
	fmt.Println("Área del cuadrado:", areaCuadrado, "cm²")
	fmt.Println("Área del rectángulo:", areaRectangulo, "cm²")
	fmt.Println("Área del rombo:", areaRombo, "cm²")
	fmt.Println("Área del pentagono:", areaPentagono, "cm²")
	fmt.Println("Área del hexagono:", areaHexagono, "cm²")
	fmt.Println("Área del circulo:", areaCirculo, "cm²")
	fmt.Println("Área del trapecio:", areaTrapecio, "cm²")

}
