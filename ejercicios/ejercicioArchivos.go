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
				idAux := strings.Split(linea, ",")
				fmt.Println("Id:", idAux[0])
				fmt.Println("Precio:", idAux[1])
				fmt.Println("Cantidad:", idAux[2])
				precio, _ := strconv.ParseFloat(idAux[1], 64)
				cantidad, _ := strconv.ParseFloat(idAux[2], 64)
				precioTotal := precio * cantidad
				fmt.Println("Precio Total: $", precio*cantidad)
				listoParaGuardar := fmt.Sprintf("%v,%v,%v\n", idAux[0], precio, cantidad)
				archivoConPreciosFinales := fmt.Sprintf("%v,%v,%v\n", idAux[0], precioTotal, cantidad)

				// myFileCSV.csv uardar queda el archivo idéntico.
				// myFileCSVprecioActualizado.csv queda el archivo con los precios finales.
				fmt.Println("LISTO PARA GUARDAR", string(listoParaGuardar))
				os.WriteFile("./myFileCSV.csv", []byte(listoParaGuardar), 0644)
				os.WriteFile("./myFileCSVprecioActualizado.csv", []byte(archivoConPreciosFinales), 0644)
			}
		}

	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Archivo vacío")
	}
}
