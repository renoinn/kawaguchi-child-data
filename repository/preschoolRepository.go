package repository

import (
	"log/slog"
	"strconv"

	"github.com/renoinn/kawaguchi-child-data/datasource"
	"github.com/renoinn/kawaguchi-child-data/entity"
)

type PreschoolRepository interface {
	GetData() ([]entity.Preschool, error)
}

type preschoolRepository struct {
	source datasource.PreschoolCsv
}

// GetData implements PreschoolRepository.
func (p *preschoolRepository) GetData() ([]entity.Preschool, error) {
	csv, err := p.source.LoadFromCsv()
	if err != nil {
		slog.Error("faild load csv")
		return nil, err
	}
	slog.Info("load csv success")

	data := []entity.Preschool{}
	for _, record := range csv {
		locationGovernmentCode, _ := strconv.Atoi(record[7])
		lat, _ := strconv.ParseFloat(record[15], 64)
		lon, _ := strconv.ParseFloat(record[16], 64)
		e := entity.Preschool{
			Code:                   record[0],
			Id:                     record[1],
			Name:                   record[3],
			KanaName:               record[4],
			EnglishName:            record[5],
			Kind:                   record[6],
			LocationGovernmentCode: locationGovernmentCode,
			TownId:                 record[8],
			LocationFull:           record[9],
			Prefectures:            record[10],
			Municipalities:         record[11],
			Town:                   record[12],
			StreetAddress:          record[13],
			BuildingName:           record[14],
			Latitude:               lat,
			Longitude:              lon,
		}
		data = append(data, e)
	}
	return data, nil
}

func NewPreschoolRepository(source datasource.PreschoolCsv) PreschoolRepository {
	return &preschoolRepository{source}
}
