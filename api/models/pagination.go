package models

type PageInfo struct {
	Page       int
	Size       int
	TotalCount int64
}

type List struct {
	Data       interface{}
	Pagination PageInfo
}
