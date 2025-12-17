package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"rentalgen/config"
	"rentalgen/logger"
	"rentalgen/types"
)

func Parse(flags *config.Flags) (*types.Tmpl, error) {
	file, err := os.Open(flags.FilePath)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao abrir arquivo de alugueres.")
	}
	defer file.Close()

	var tmplData types.Tmpl
	err = json.NewDecoder(file).Decode(&tmplData)
	if err != nil {
		return nil, fmt.Errorf("✗ erro ao processar arquivo.")
	}

	logger.Log(logger.ColorGreen, "✓ arquivo de alugueres selecionado: ")
	logger.Logln(logger.ColorDefault, flags.FilePath)
	logger.Logln(logger.ColorDefault, "✓ para realizar a impressão dos recibos, acesse http://localhost:8080\n")
	logger.Logln(logger.ColorDefault, "Para encerrar, pressione Ctrl + C")

	return &tmplData, nil
}
