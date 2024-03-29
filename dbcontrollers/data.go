package dbcontrollers

import (
	"time"

	"github.com/artofimagination/timescaledb-project-log-go-interface/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrFailedToAddData = errors.New("Failed to save data")

func (c *TimescaleController) AddData(data []models.Data) error {
	if err := c.DBFunctions.AddData(data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDataByViewer(viewerID int) error {
	if err := c.DBFunctions.DeleteByViewerID(viewerID); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDataByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteByTime(intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetDataByViewerAndTime(viewerID *uuid.UUID, time time.Time) ([]models.Data, error) {
	data, err := c.DBFunctions.GetDataByViewerAndTime(viewerID, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}
