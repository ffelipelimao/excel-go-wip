package business

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
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

func (s GameServiceImpl) ReadExcel(file multipart.File) ([]domain.Game, error) {
	reader, err := excelize.OpenReader(file)
	if err != nil {
		return []domain.Game{}, err
	}

	rows, err := reader.GetRows(GameSheetName)
	if err != nil {
		return []domain.Game{}, err
	}

	games := []domain.Game{}
	for i, row := range rows {
		// skip header
		if i == 0 {
			continue
		}

		game, err := makeGameByLine(row)
		if err != nil {
			return []domain.Game{}, err
		}

		games = append(games, game)
	}

	return games, nil
}

func makeGameByLine(line []string) (domain.Game, error) {
	if len(line) < 4 {
		return domain.Game{}, errors.New("invalid columns")
	}

	game := domain.Game{
		ID:    line[0],
		Name:  line[1],
		Price: line[2],
		Thumb: line[3],
	}

	return game, nil
}
