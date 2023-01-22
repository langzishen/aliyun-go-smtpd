package service

import "go-smtpd/models"

type Receives struct {
	Base
}

func (this *Receives) Add(from string, to []string, content, code string) {
	m := models.Receives{}
	m.From = from
	m.Content = content
	m.Code = code
	for _, vo_to := range to {
		m.To = vo_to
		InitDB().Create(&m)
	}
}
