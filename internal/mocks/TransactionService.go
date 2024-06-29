// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	dto "github.com/saadozone/gin-gorm-rest/internal/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/saadozone/gin-gorm-rest/internal/model"
)

// TransactionService is an autogenerated mock type for the TransactionService type
type TransactionService struct {
	mock.Mock
}

// CountTransaction provides a mock function with given fields: userID
func (_m *TransactionService) CountTransaction(userID int) (int64, error) {
	ret := _m.Called(userID)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactions provides a mock function with given fields: userID, query
func (_m *TransactionService) GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	ret := _m.Called(userID, query)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(int, *dto.TransactionRequestQuery) ([]*model.Transaction, error)); ok {
		return rf(userID, query)
	}
	if rf, ok := ret.Get(0).(func(int, *dto.TransactionRequestQuery) []*model.Transaction); ok {
		r0 = rf(userID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(int, *dto.TransactionRequestQuery) error); ok {
		r1 = rf(userID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopUp provides a mock function with given fields: input
func (_m *TransactionService) TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error) {
	ret := _m.Called(input)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.TopUpRequestBody) (*model.Transaction, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*dto.TopUpRequestBody) *model.Transaction); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.TopUpRequestBody) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: input
func (_m *TransactionService) Transfer(input *dto.TransferRequestBody) (*model.Transaction, error) {
	ret := _m.Called(input)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.TransferRequestBody) (*model.Transaction, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*dto.TransferRequestBody) *model.Transaction); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.TransferRequestBody) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionService creates a new instance of TransactionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionService(t mockConstructorTestingTNewTransactionService) *TransactionService {
	mock := &TransactionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}