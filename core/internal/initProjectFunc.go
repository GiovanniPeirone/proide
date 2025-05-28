package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitProjectPOI(route string) {
	if route == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error al obtener el directorio actual:", err)
			return
		}
		route = cwd
		fmt.Println("Ruta vacía. Usando directorio actual:", route)
	}

	err := createNewProject(route)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
	} else {
		fmt.Println("Proyecto inicializado correctamente en:", route)
	}
}

func createNewProject(route string) error {
	filePath := filepath.Join(route, "com.poi")

	// Verifica si ya existe
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("El archivo ya existe:", filePath)
		return nil
	}

	// Crear archivo vacío
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
