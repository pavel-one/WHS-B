package models

import (
	"app/Services/Scrapper"
	"app/base"
	"app/requests"
	"gopkg.in/guregu/null.v4"
	"log"
	"os"
	"time"
)

type Dial struct {
	ID          uint
	UserID      uint
	Name        null.String
	Description null.String
	Screen      null.String
	Url         string
	Final       bool
	Type        int
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (dial *Dial) FillWithRequest(db *base.DB, request requests.DialRequestInterface) {
	dial.Url = request.GetUrl()
	dial.Description = request.GetDescription()
	dial.Name = request.GetName()
}

func (dial *Dial) CreateOrUpdateInfo(db *base.DB) {
	defer dial.SetProcessEnd(db)

	url, err := Scrapper.GetUrlInfo(dial.Url, dial.Url, dial.UserID)
	if err != nil {
		log.Println(err)
		return
	}

	if !dial.Name.Valid {
		dial.Name = null.StringFrom(url.Title)
	}

	if !dial.Screen.Valid {
		dial.Screen = null.StringFrom(url.Screen)
	}
}

func (dial *Dial) UpdatePhoto(db *base.DB) {
	defer dial.SetProcessEnd(db)

	url, err := Scrapper.GetUrlInfo(dial.Url, dial.Url, dial.UserID)
	if err != nil {
		log.Println(err)
		return
	}

	dial.Screen = null.StringFrom(url.Screen)
}

func (dial *Dial) SetProcess(db *base.DB) {
	if dial.Final == false {
		return
	}

	dial.Final = false
	db.Save(dial)
}

func (dial *Dial) SetProcessEnd(db *base.DB) {
	if dial.Final == true {
		return
	}

	dial.Final = true
	db.Save(dial)
}

func (dial *Dial) DropDialWithFiles(db *base.DB, fake bool) {
	if dial.Screen.Valid {
		_ = os.Remove(dial.Screen.String)
	}

	if fake {
		dial.Deleted = true
		db.Save(dial)
		return
	}

	db.Delete(dial)
}
