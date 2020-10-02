package beater



import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// STATEFILE es el nombre del archivo que guarda el estado actual
const STATEFILE = "state.txt"

func storeState(id int) error {
	if fileExists(STATEFILE) {
		err := os.Remove(STATEFILE)
		if err != nil {
			log.Fatal("No se pudo guardar el estado", err)
			return err
		}
	}
	f, err := os.Create(STATEFILE)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write([]byte(strconv.Itoa(id)))
	return nil
}

func loadState() int {
	if fileExists(STATEFILE) {
		f, _ := os.Open(STATEFILE)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		num := 0
		for scanner.Scan() {
			lineStr := scanner.Text()
			num, _ = strconv.Atoi(lineStr)
		}
		return num
	}
	return 0
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
