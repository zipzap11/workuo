// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	news "workuo/features/news"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetData provides a mock function with given fields: keyword
func (_m *Repository) GetData(keyword string) ([]news.NewsCore, error) {
	ret := _m.Called(keyword)

	var r0 []news.NewsCore
	if rf, ok := ret.Get(0).(func(string) []news.NewsCore); ok {
		r0 = rf(keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]news.NewsCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
