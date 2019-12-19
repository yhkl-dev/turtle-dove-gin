package utils

import (
	"fmt"
	"math"
	"strings"
)

// Pager struct
type Pager struct {
	Page        int
	PageSize    int
	TotalPage   int
	PrePageURL  string
	NextPageURL string
}

// NewPager function
func NewPager(page, pagesize, totalNumber int, currentURL string) Pager {
	pager := new(Pager)

	pager.Page = page
	pager.PageSize = pagesize
	pager.TotalPage = int(math.Ceil(float64(totalNumber) / float64(pagesize)))

	pagerURL := strings.Split(currentURL, "?")
	if pager.TotalPage <= 1 {
		pager.NextPageURL = ""
		pager.PrePageURL = ""
	} else if pager.Page == 1 {
		pager.PrePageURL = ""
		pager.NextPageURL = fmt.Sprintf("%s?page=%d&pagesize=%d", pagerURL[0], pager.Page, pager.PageSize)
	} else {
		pager.NextPageURL = fmt.Sprintf("%s?page=%d&pagesize=%d", pagerURL[0], pager.Page, pager.PageSize)
		pager.PrePageURL = fmt.Sprintf("%s?page=%d&pagesize=%d", pagerURL[0], pager.Page-1, pager.PageSize)
	}
	return *pager
}
