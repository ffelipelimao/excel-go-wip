package business

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ffelipelimao/excel-go/internal/domain"
	"github.com/ffelipelimao/excel-go/internal/integrations"
	"github.com/xuri/excelize/v2"
)

const GameSheetName = "Games"

type GameServiceImpl struct {
	integration integrations.GameIntegration
}

func NewGameServiceImpl() GameService {
	return GameServiceImpl{
		integrations.NewIntegration(),
	}
}

func (s GameServiceImpl) List() ([]domain.Game, error) {
	games, err := s.integration.Get()
	if err != nil {
		return []domain.Game{}, err
	}
	return games, nil
}

func (s GameServiceImpl) ListExcel() (*bytes.Buffer, error) {
	games, err := s.integration.Get()
	if err != nil {
		return &bytes.Buffer{}, err
	}

	// Initialize a file
	file := excelize.NewFile()
	index := file.NewSheet(GameSheetName)

	makeHeaderFile(file)

	for i, game := range games {
		file.SetCellValue(GameSheetName, fmt.Sprintf("A%s", strconv.Itoa(i+2)), game.ID)
		file.SetCellValue(GameSheetName, fmt.Sprintf("B%s", strconv.Itoa(i+2)), game.Name)
		file.SetCellValue(GameSheetName, fmt.Sprintf("C%s", strconv.Itoa(i+2)), game.Price)
		file.SetCellValue(GameSheetName, fmt.Sprintf("D%s", strconv.Itoa(i+2)), game.Thumb)
	}

	// Flush the file
	file.SetActiveSheet(index)
	fileBuffer, err := file.WriteToBuffer()
	if err != nil {
		return &bytes.Buffer{}, err
	}

	return fileBuffer, nil
}

func makeHeaderFile(file *excelize.File) {

	file.SetCellValue(GameSheetName, "A1", "ID")
	file.SetCellValue(GameSheetName, "B1", "Name")
	file.SetCellValue(GameSheetName, "C1", "Price")
	file.SetCellValue(GameSheetName, "D1", "Thumb")
}
