package ejercicios

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Producto struct {
	id       int
	precio   int
	cantidad int
}

func GuardarArchivoCSV() {
	fmt.Println("Entrando a GuardarArchivoCSV()")
	data, _ := os.ReadFile("./myFileCSV.csv")

	/*
		if data == nil {
			fmt.Println("Archivo Vacío")
			fmt.Println(err.Error())
		} else {
	*/
	p1 := Producto{}
	p1.id = 4
	p1.precio = 4000
	p1.cantidad = 40

	csv := fmt.Sprintf("%v,%v,%v\n", p1.id, p1.precio, p1.cantidad)
	fmt.Println("Guardando...")
	fmt.Println(csv)
	data = append(data, csv...)
	os.WriteFile("./myFileCSV.csv", data, 0644)
	fmt.Println("*** Archivo Completo ***")
	fmt.Println(string(data))

	/*}*/
}

func LeerTexto() {
	data, err := os.ReadFile("./myFileCSV.csv")

	if data != nil {
		aux := string(data)
		lineaPorLinea := strings.Split(aux, "\n")
		fmt.Println("**********")
		for _, linea := range lineaPorLinea {
			if linea != "" {
				fmt.Println("Linea: ", linea)
				idAux := strings.Split(linea, ",")
				fmt.Println("Id:", idAux[0])
				fmt.Println("Precio:", idAux[1])
				fmt.Println("Cantidad:", idAux[2])
				precio, err := strconv.ParseInt(idAux[1], 10, 64)
				cantidad, err := strconv.ParseInt(idAux[2], 10, 64)
				if err != nil {
					fmt.Println("Precio Total: $", precio*cantidad)
				}
				listoParaGuardar := fmt.Sprintf("%v %v %v\n", idAux[0], precio, cantidad)

				data = append(data, listoParaGuardar...)
				os.WriteFile("./myFileCSV.csv", data, 0644)
				fmt.Println("*** Archivo Completo ***")
				fmt.Println(string(data))
			}
		}

	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Archivo vacío")
	}
}
