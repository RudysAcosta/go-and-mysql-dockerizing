package db

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func TestArchivo() {
	// Obtener el nombre del directorio
	dirName := "/usr/src/app/migrations"

	// Abrir el directorio
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println("Error al abrir el directorio:", err)
		return
	}
	defer dir.Close()

	// Leer el contenido del directorio
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return
	}

	// Imprimir el nombre de cada archivo y directorio
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// RunMigrations executes SQL migration files found in the specified directory
func RunMigrations(database *Database, migrationsDir string) error {
	// Open the directory
	dir, err := os.Open(migrationsDir)

	if err != nil {
		return fmt.Errorf("failed to open migration directory 1 : %v", err)
	}
	defer dir.Close()

	// Read the contents of the directory
	files, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("failed to read migration directory 2: %v", err)
	}

	// Iterate over migration files
	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}

		// Open the migration file
		filePath := filepath.Join(migrationsDir, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open migration file %s: %v", file.Name(), err)
		}
		defer f.Close()

		// Read the content of the migration file
		scanner := bufio.NewScanner(f)
		var content string
		for scanner.Scan() {
			content += scanner.Text() + "\n"
		}

		// Execute the migration query
		_, err = database.Exec(content)
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file.Name(), err)
		}

		// Print a message indicating the migration was successful
		fmt.Printf("Migration %s applied successfully\n", file.Name())
	}

	return nil
}
