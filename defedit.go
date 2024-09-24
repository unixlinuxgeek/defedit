// Running a default graphical text editor(not console) for text file
// Запуск графического текстового редактора по умолчанию (не консольного) для текстового файла

package defedit

import (
	"fmt"
	"io"
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
func Open() ([]byte, error) {
	tmp := "/tmp_file_" + strconv.Itoa(rand.Int())
	tmpFile, err := os.Create(os.TempDir() + tmp)
	if err != nil {
		os.Exit(1)
	}
	path, err := exec.LookPath("xdg-open")
	fmt.Printf("%s\n", path)

	cmd := exec.Command(path, tmpFile.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run(): %s\n", err)
		return nil, err
	}
	f, err := os.Open(os.TempDir() + tmp)
	if err != nil {
		fmt.Printf("os.Open: %s\n", err)
		return nil, err
	}
	g, _ := io.ReadAll(f)
	return g, nil
}
