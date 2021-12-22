package main

import (
	"fmt"
)

func main() {

	// *****    TURNO MAÑANA    ******
	// ################################### //
	//Ejercicio 1
	//sueldoTotal := ejercicios.Ejercicio1()
	//fmt.Println("SUELDO: ", sueldoTotal)

	// ################################### //
	//Ejercicio 2
	//fmt.Println(ejercicios.Ejercicio2(9, 7, 6))

	// ################################### //
	//Ejercicio 3
	//fmt.Println(ejercicios.Ejercicio3(9600, "A"))

	// ################################### //
	//Ejercicio 4
	//fmt.Println(ejercicios.Ejercicio4())

	// ################################### //
	//Ejercicio 5
	//fmt.Println(ejercicios.Ejercicio5())

	// *****    TURNO MAÑANA - LECTURA ARCHIVOS   ******
	//ejercicios.EscribirTexto()
	//fmt.Println("Hola")
	//ejercicios.GuardarArchivoCSV()
	//ejercicios.LeerTexto()
	//fmt.Println("Texto escrito")

	// *****    TURNO TARDE - CONCURRENCIA   ******
	user := usuario{"Jose", "Pasini", 26, "fdas", "fdsa@gmail.com"}
	fmt.Printf("%vUSUARIO:\n", user)
	user = CambiarNombre(&user)
	fmt.Printf("%vUSUARIO:", user)

}

type usuario struct {
	Nombre      string `json:"nombre"`
	Apellido    string `json:"Apellido"`
	Edad        int    `json:"edad"`
	Correo      string `json:"correo"`
	Contrasenia string `json:"contrasenia"`
}

func CambiarNombre(u *usuario) usuario {
	var nombre string
	var apellido string
	fmt.Print("Ingrese el nuevo nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Ingrese el nuevo apellido: ")
	fmt.Scanln(&apellido)
	u.Nombre = nombre
	u.Apellido = apellido
	return *u
}

func CambiarEdad(u *usuario) usuario {
	edad := 0
	fmt.Print("Ingrese la nueva edad: ")
	fmt.Scanln(&edad)
	u.Edad = edad
	return *u
}

func CambiarCorreo(u *usuario) usuario {
	var correo string
	fmt.Print("Ingrese el nuevo correo: ")
	fmt.Scanln(&correo)
	u.Correo = correo
	return *u
}

func CambiarContrasenia(u *usuario) usuario {
	var contrasenia string
	fmt.Print("Ingrese la nueva Contrasenia: ")
	fmt.Scanln(&contrasenia)
	u.Contrasenia = contrasenia
	return *u
}
