package service_test

import (
	"errors"
	"kupon/features/users"
	"kupon/features/users/mocks"
	"kupon/features/users/service"
	eMock "kupon/helper/enkrip/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	enkrip := eMock.NewHashInterface(t)
	s := service.New(repo, enkrip)
	var inputData = users.User{Name: "lendra", Email: "lendra@gmail.com", Password: "lendra321"}
	var successReturndata = users.User{ID: uint(1), Name: "lendra", Email: "lendra@gmail.com"}
	var falseData = users.User{}
	t.Run("Succes Case", func(t *testing.T) {
		enkrip.On("HashPassword","lendra321").Return("some string", nil)
		inputData.Password = "some string"
		repo.On("Register", inputData).Return(successReturndata, nil).Once()
		res, err := s.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
		assert.Equal(t, "", res.Password)
		enkrip.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		s = service.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.Register(falseData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Name)
	})
}

func TestLogin(t *testing.T) {
	s := service.New(&MockRepository{}, &MockEnkrip{})
	t.Run("Success Case", func(t *testing.T) {
		res, err := s.Login("lendra@gmail.com", "lendra321")
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
	})

	t.Run("Failed Case", func(t *testing.T) {
		s = service.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.Login("wrong@gmail.com", "wrongpass")
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
	})
}

func TestUpdate(t *testing.T) {
	s := service.New(&MockRepository{}, &MockEnkrip{})
	var inputData = users.User{Name: "lendra", Email: "lendra@gmail.com", Password: "lendra321"}
	t.Run("Success Case", func(t *testing.T) {
		res, err := s.Update(nil, inputData)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
	})

	t.Run("Failed Case", func(t *testing.T) {
		s = service.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.Update(nil, inputData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
	})
}

func TestGetUser(t *testing.T) {
	s := service.New(&MockRepository{}, &MockEnkrip{})
	t.Run("Success Case", func(t *testing.T) {
		res, err := s.GetUser()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
	})

	t.Run("Failed Case", func(t *testing.T) {
		s = service.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.GetUser()
		assert.Error(t, err)
		assert.Equal(t, 0, len(res))
	})
}


type MockRepository struct {}

func (mr *MockRepository) Register(newUser users.User) (users.User, error) {
	return users.User{ID: uint(1), Name: "lendra", Email: "lendra@gmail.com"}, nil
}
func (mr *MockRepository)Login(email string) (users.User, error) {
	return users.User{}, nil
}
func (mr *MockRepository) Update(id uint, updateUser users.User) (users.User, error) {
    return users.User{}, nil
}
func (mr *MockRepository) GetUser() ([]users.User, error) {
    return []users.User{}, nil
}

type FalseMockRepository struct {}

func (fmr *FalseMockRepository) Register(newUser users.User) (users.User, error) {
	return users.User{}, errors.New("something hapenedd")
}
func (fmr *FalseMockRepository)Login(email string) (users.User, error) {
	return users.User{}, nil
}
func (fmr *FalseMockRepository) Update(id uint, updateUser users.User) (users.User, error) {
    return users.User{}, nil
}
func (fmr *FalseMockRepository) GetUser() ([]users.User, error) {
    return []users.User{}, nil
}

type MockEnkrip struct {}

func (me *MockEnkrip) Compare(hashed string, input string) error {
	return nil
}

func (me *MockEnkrip) HashPassword(input string) (string, error) {
	return "some string", nil
}

