// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	dtos "scoreBoard/models/dtos"

	mock "github.com/stretchr/testify/mock"

	models "scoreBoard/models"

	serviceerrors "scoreBoard/internal/errors"
)

// GameDAO is an autogenerated mock type for the GameDAO type
type GameDAO struct {
	mock.Mock
}

// GetLeaderBoardForAllRegion provides a mock function with given fields: gameId, offset, count
func (_m *GameDAO) GetLeaderBoardForAllRegion(gameId string, offset int, count int) ([]*models.ScoreBoard, *serviceerrors.Errors) {
	ret := _m.Called(gameId, offset, count)

	var r0 []*models.ScoreBoard
	if rf, ok := ret.Get(0).(func(string, int, int) []*models.ScoreBoard); ok {
		r0 = rf(gameId, offset, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ScoreBoard)
		}
	}

	var r1 *serviceerrors.Errors
	if rf, ok := ret.Get(1).(func(string, int, int) *serviceerrors.Errors); ok {
		r1 = rf(gameId, offset, count)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serviceerrors.Errors)
		}
	}

	return r0, r1
}

// GetLeaderBoardForRegion provides a mock function with given fields: gameId, region, offset, count
func (_m *GameDAO) GetLeaderBoardForRegion(gameId string, region string, offset int, count int) ([]*models.ScoreBoard, *serviceerrors.Errors) {
	ret := _m.Called(gameId, region, offset, count)

	var r0 []*models.ScoreBoard
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*models.ScoreBoard); ok {
		r0 = rf(gameId, region, offset, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ScoreBoard)
		}
	}

	var r1 *serviceerrors.Errors
	if rf, ok := ret.Get(1).(func(string, string, int, int) *serviceerrors.Errors); ok {
		r1 = rf(gameId, region, offset, count)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serviceerrors.Errors)
		}
	}

	return r0, r1
}

// GetTimedLeaderBoardForAllRegion provides a mock function with given fields: gameId, offset, count, fromTime, toTime
func (_m *GameDAO) GetTimedLeaderBoardForAllRegion(gameId string, offset int, count int, fromTime string, toTime string) ([]*models.ScoreBoard, *serviceerrors.Errors) {
	ret := _m.Called(gameId, offset, count, fromTime, toTime)

	var r0 []*models.ScoreBoard
	if rf, ok := ret.Get(0).(func(string, int, int, string, string) []*models.ScoreBoard); ok {
		r0 = rf(gameId, offset, count, fromTime, toTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ScoreBoard)
		}
	}

	var r1 *serviceerrors.Errors
	if rf, ok := ret.Get(1).(func(string, int, int, string, string) *serviceerrors.Errors); ok {
		r1 = rf(gameId, offset, count, fromTime, toTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serviceerrors.Errors)
		}
	}

	return r0, r1
}

// GetTimedLeaderBoardForRegion provides a mock function with given fields: gameId, region, offset, count, fromTime, toTime
func (_m *GameDAO) GetTimedLeaderBoardForRegion(gameId string, region string, offset int, count int, fromTime string, toTime string) ([]*models.ScoreBoard, *serviceerrors.Errors) {
	ret := _m.Called(gameId, region, offset, count, fromTime, toTime)

	var r0 []*models.ScoreBoard
	if rf, ok := ret.Get(0).(func(string, string, int, int, string, string) []*models.ScoreBoard); ok {
		r0 = rf(gameId, region, offset, count, fromTime, toTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ScoreBoard)
		}
	}

	var r1 *serviceerrors.Errors
	if rf, ok := ret.Get(1).(func(string, string, int, int, string, string) *serviceerrors.Errors); ok {
		r1 = rf(gameId, region, offset, count, fromTime, toTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serviceerrors.Errors)
		}
	}

	return r0, r1
}

// WithGameID provides a mock function with given fields: gameId
func (_m *GameDAO) WithGameID(gameId string) (*dtos.Game, *serviceerrors.Errors) {
	ret := _m.Called(gameId)

	var r0 *dtos.Game
	if rf, ok := ret.Get(0).(func(string) *dtos.Game); ok {
		r0 = rf(gameId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.Game)
		}
	}

	var r1 *serviceerrors.Errors
	if rf, ok := ret.Get(1).(func(string) *serviceerrors.Errors); ok {
		r1 = rf(gameId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serviceerrors.Errors)
		}
	}

	return r0, r1
}
