package ejercicios

import (
	"fmt"
)

func Ejercicio1() float64 {
	/*Una empresa de chocolates necesita calcular el impuesto de sus empleados
	al momento de depositar el sueldo, para cumplir el objetivo es necesario
	crear una función que devuelva el impuesto de un salario.
	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará
	un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%
	*/

	sueldo := 0.0
	fmt.Print("Ingrese su sueldo: $")
	fmt.Scanln(&sueldo)

	if sueldo > 50000 && sueldo < 150000 {
		return sueldo - (sueldo * 0.17)
	} else if sueldo > 150000 {
		return sueldo - (sueldo * 0.27)
	} else {
		return sueldo
	}

}

func Ejercicio2(notas ...float64) (float64, error) {

	count := 0.0

	for _, a := range notas {
		if a < 0 {
			return 0, fmt.Errorf("Error, ingresaron un nùmero negativo en las notas")
		}
		count += a
	}

	fmt.Println("CONTADOR: ", count)
	return count / float64(len(notas)), nil
}

func Ejercicio3(minutos float64, categoria string) (float64, error) {
	/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la
	  cantidad de horas trabajadas por mes y la categoría.
	  Si es categoría C, su salario es de $1.000 por hora
	  Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	  Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
	  Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
	*/

	horas := minutos / 60
	//fmt.Println("Ingrese cantidad de horas trabajadas en el mes: ")
	//fmt.Scanln(&horas)
	//fmt.Println("categoria: A - B - C ")
	//fmt.Scanln(&categoria)

	switch categoria {
	case "C":
		fmt.Println("Su sueldo es de: $", horas*1000)
		return horas * 1000, nil
	case "B":
		fmt.Println("Su sueldo es de: $", (horas*1500)*1.2)
		return (horas * 1500) * 1.2, nil
	case "A":
		fmt.Println("Su sueldo es de: $", (horas*3000)*1.5)
		return (horas * 3000) * 1.5, nil
	default:
		fmt.Println("Debe ingresar una opción correcta")
	}
	return 0.0, fmt.Errorf("ERROR")
}

func Ejercicio4() float64 {

	op := 0
	fmt.Println("Indique que opción desea realizar: \n1. Minimo\n2. Máximo\n3. Promedio")
	fmt.Scanln(&op)

	oper := ElegirOperacion(op)
	res := oper(3, 4, 5, 10, 15, 12, -5, 8)
	fmt.Println("R:", res)
	return res
}

func ElegirOperacion(op int) func(...float64) float64 {

	switch op {
	case 1:
		return OperacionMin
	case 2:
		return OperacionMax
	case 3:
		return OperacionProm
	default:
		return nil
	}
}

func OperacionMin(notas ...float64) float64 {
	notaMenor := notas[0]
	for _, a := range notas {
		if a < notaMenor {
			notaMenor = a
		}
	}
	return notaMenor
}
func OperacionMax(notas ...float64) float64 {
	notaMayor := notas[0]
	for _, a := range notas {
		if a > notaMayor {
			notaMayor = a
		}
	}
	return notaMayor
}

func OperacionProm(notas ...float64) float64 {
	count := 0.0
	for _, a := range notas {
		count += a
	}
	fmt.Println("CONTADOR: ", count)
	return count / float64(len(notas)) //, nil
}

func Ejercicio5() float64 {
	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	a := Animal(gato)
	ress, err := a(5)
	if err != nil {
		fmt.Println("Animal no encontrado:")
		return 0.0
	} else {
		fmt.Println("RESULTADO:", ress)
		return ress
	}
}

func Animal(name string) func(int) (float64, error) {
	if name == "perro" {
		return AnimalPerro
	} else if name == "gato" {
		return AnimalGato
	} else if name == "hamster" {
		return AnimalHamster
	} else if name == "tarantula" {
		return AnimalTarantula
	}
	return AnimalNoEncontrado
}

func AnimalPerro(cantidad int) (float64, error) {
	return float64(cantidad * 10), nil
}
func AnimalGato(cantidad int) (float64, error) {
	return float64(cantidad * 100), nil
}
func AnimalHamster(cantidad int) (float64, error) {
	return float64(cantidad * 1000), nil
}
func AnimalTarantula(cantidad int) (float64, error) {
	return float64(cantidad * 10000), nil
}

func AnimalNoEncontrado(cantidad int) (float64, error) {
	return 0.0, fmt.Errorf("Animal no encontrado")
}
