package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateArchetypeSetupFile(project string) error {
	// Read the content of the source setup.go file
	sourceCiFilePath := "app/shared/archetype/setup.go"
	ciFileContentBytes, err := ioutil.ReadFile(sourceCiFilePath)
	if err != nil {
		err := fmt.Errorf("error reading source setup.go file at %s: %s", sourceCiFilePath, err)
		fmt.Println(err)
		return err
	}

	ciFileContent := string(ciFileContentBytes)

	// Write the content to the destination setup.go file
	ciFilePath := filepath.Join(project, "app/shared/archetype/setup.go")
	err = ioutil.WriteFile(ciFilePath, []byte(ciFileContent), 0644)
	if err != nil {
		err := fmt.Errorf("error writing setup.go file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("setup.go file generated successfully at %s.\n", ciFilePath)
	return nil
}
