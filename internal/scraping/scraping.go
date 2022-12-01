package scraping

import (
	"strings"

	"github.com/gocolly/colly"

	"new_employee_notice/internal/config"
)

const (
	divPeople         = "div.peple"
	ahref             = "a[href]"
	href              = "href"
	tablePeopleMenuTr = "table.peplemenu tr"
	tableProfData     = "table.profdate"
	secondTd          = "tr:nth-child(2) td"
	thirdTd           = "tr:nth-child(3) td"
)

type Employee struct {
	ID                string
	Name              string
	Furigana          string
	BirthDay          string
	TrainingStartDate string
}

type Scraping interface {
	FetchEmployeeData() ([]Employee, error)
	login() error
}

type ScrapingCli struct {
	c *colly.Collector
}

func New() ScrapingCli {
	return ScrapingCli{
		c: colly.NewCollector(colly.MaxDepth(1)),
	}
}

func (s ScrapingCli) FetchEmployeeData() ([]Employee, error) {
	emp := Employee{}
	emps := []Employee{}

	if err := s.login(); err != nil {
		return nil, err
	}

	s.onHTML(divPeople, func(e *colly.HTMLElement) {
		e.ForEach(ahref, func(_ int, e *colly.HTMLElement) {
			link := e.Attr(href)
			emp.ID = link[len(link)-3:]
			_ = s.c.Visit(e.Request.AbsoluteURL(link))
		})

		names := e.ChildText(tablePeopleMenuTr)
		arr := strings.Split(names, "\n")

		emp.Name = formatName(arr[0])
		emp.Furigana = formatName(arr[1])
		emps = append(emps, emp)
	})

	s.onHTML(tableProfData, func(e *colly.HTMLElement) {
		emp.BirthDay = e.ChildText(secondTd)
		emp.TrainingStartDate = e.ChildText(thirdTd)
	})

	if err := s.c.Visit(config.Conf.MyPageUrl); err != nil {
		return nil, err
	}

	return emps, nil
}

func (s *ScrapingCli) login() error {
	if err := s.c.Post(config.Conf.LoginUrl, map[string]string{
		"login_id":       config.Conf.LoginId,
		"login_password": config.Conf.LoginPassword,
	}); err != nil {
		return err
	}
	return nil
}

func (s ScrapingCli) onHTML(selector string, callback func(e *colly.HTMLElement)) {
	s.c.OnHTML(selector, callback)
}

func formatName(name string) string {
	return strings.Replace(strings.TrimSpace(name), "\u3000", " ", 1)
}
