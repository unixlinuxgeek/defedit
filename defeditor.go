// Running a default graphical text editor(not console) for text file
// Запуск графического текстового редактора по умолчанию (не консольного) для текстового файла

package defeditor

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"strconv"
)

// Open creates a temporary file and opens your preferred graphical file editor.
// (in ubuntu it is "Image viewer".)
//
// Open создаёт временный файл и открывает предпочтительный файловый редактор
// (в ubuntu это "Image viewer".)
func Open() (string, error) {
	tmpFile, err := os.Create(os.TempDir() + "/tmp_file_" + strconv.Itoa(rand.Int()))
	if err != nil {
		os.Exit(1)
	}
	path, err := exec.LookPath("xdg-open")
	fmt.Printf("%s\n", path)

	wd, _ := os.Getwd()
	cmd := exec.Command(path, tmpFile.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err)
		return "", err
	}
	return wd + tmpFile.Name(), nil
}
