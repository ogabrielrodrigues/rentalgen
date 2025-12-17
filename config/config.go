package config

import (
	"flag"
)

type Flags struct {
	FilePath string
}

func GetFlags() *Flags {
	filePath := flag.String("file", "alugueres.json", "caminho do arquivo de alugueres.")
	flag.Parse()

	flags := Flags{
		FilePath: *filePath,
	}

	return &flags
}
