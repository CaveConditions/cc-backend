package db

import (
	"CaveConditions/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TabelCave - database table 'caves'
	TabelCave = "caves"
)

// FindCaveByUK queries the cave by unique key 'title'
func (s *Handler) FindCaveByUK(title string) (*model.Cave, error) {
	c := model.Cave{}

	// query the cave rows by title
	res := s.db.Table(TabelCave).Where("title = ?", title).First(&c)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &c, nil
}

// CreateCave insert the cave into database
func (s *Handler) CreateCave(u *model.Cave) error {
	return s.db.Table(TabelCave).Create(u).Error
}
