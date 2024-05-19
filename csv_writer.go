package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Fahmi", "Arzalega", "Seazz"})
	_ = writer.Write([]string{"Budi", "Nugraha", "Wijaya"})
	_ = writer.Write([]string{"Odetta", "Firstly", "Nourma"})

	writer.Flush()
}
