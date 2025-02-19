package main

import (
	"calcularFiguras/figuras"
	"calcularFiguras/funciones"
	"fmt"
)

func main() {
	// Crear un objeto de tipo FuncionCuadratica
	fc := funciones.NuevaFuncionCuadratica(1, -3, 2)

	// Calcular las raíces de la ecuación cuadrática
	x1, x2, err := fc.CalcularFuncionCuadratica()
	if err != nil {
		fmt.Println("Error al calcular la función cuadrática:", err)
		return
	}

	// Imprimir los resultados
	fmt.Println("Raíz 1:", x1)
	fmt.Println("Raíz 2:", x2)

	// Interfaz para las figuras
	var interfazFigura figuras.Figura

	// Triángulo
	triangulo := figuras.NuevoTriangulo(10, 5)
	interfazFigura = triangulo
	areaTriangulo := interfazFigura.CalculateArea()
	perimetroTriangulo := interfazFigura.CalcularPerimetro()

	// Cuadrado
	cuadrado := figuras.NuevoCuadrado(10)
	interfazFigura = cuadrado
	areaCuadrado := interfazFigura.CalculateArea()
	perimetroCuadrado := interfazFigura.CalcularPerimetro()

	// Rectángulo
	rectangulo := figuras.NuevoRectangulo(10, 2)
	interfazFigura = rectangulo
	areaRectangulo := interfazFigura.CalculateArea()
	perimetroRectangulo := interfazFigura.CalcularPerimetro()

	// Rombo
	rombo := figuras.NuevoRombo(6, 5)
	interfazFigura = rombo
	areaRombo := interfazFigura.CalculateArea()
	perimetroRombo := interfazFigura.CalcularPerimetro()

	// Pentágono
	pentagono := figuras.NuevoPentagono(4, 4)
	interfazFigura = pentagono
	areaPentagono := interfazFigura.CalculateArea()
	perimetroPentagono := interfazFigura.CalcularPerimetro()

	// Hexágono
	hexagono := figuras.NuevoHexagono(6, 5)
	interfazFigura = hexagono
	areaHexagono := interfazFigura.CalculateArea()
	perimetroHexagono := interfazFigura.CalcularPerimetro()

	// Círculo
	circulo := figuras.NuevoCirculo(3)
	interfazFigura = circulo
	areaCirculo := interfazFigura.CalculateArea()
	perimetroCirculo := interfazFigura.CalcularPerimetro()

	// Trapecio
	trapecio := figuras.NuevoTrapecio(8, 6, 4)
	interfazFigura = trapecio
	areaTrapecio := interfazFigura.CalculateArea()
	perimetroTrapecio := interfazFigura.CalcularPerimetro()

	// Impresión de resultados
	fmt.Println("Área del triángulo:", areaTriangulo, "cm²")
	fmt.Println("Perímetro del triángulo:", perimetroTriangulo, "cm")
	fmt.Println("Área del cuadrado:", areaCuadrado, "cm²")
	fmt.Println("Perímetro del cuadrado:", perimetroCuadrado, "cm")
	fmt.Println("Área del rectángulo:", areaRectangulo, "cm²")
	fmt.Println("Perímetro del rectángulo:", perimetroRectangulo, "cm")
	fmt.Println("Área del rombo:", areaRombo, "cm²")
	fmt.Println("Perímetro del rombo:", perimetroRombo, "cm")
	fmt.Println("Área del pentágono:", areaPentagono, "cm²")
	fmt.Println("Perímetro del pentágono:", perimetroPentagono, "cm")
	fmt.Println("Área del hexágono:", areaHexagono, "cm²")
	fmt.Println("Perímetro del hexágono:", perimetroHexagono, "cm")
	fmt.Println("Área del círculo:", areaCirculo, "cm²")
	fmt.Println("Perímetro (circunferencia) del círculo:", perimetroCirculo, "cm")
	fmt.Println("Área del trapecio:", areaTrapecio, "cm²")
	fmt.Println("Perímetro del trapecio:", perimetroTrapecio, "cm")

	//Prueba de la funcion Cuadratica
}
