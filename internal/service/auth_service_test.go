package service

// import (
// 	"testing"
// 	"time"

// 	"github.com/saadozone/gin-gorm-rest/internal/dto"
// 	"github.com/saadozone/gin-gorm-rest/internal/mocks"
// 	"github.com/saadozone/gin-gorm-rest/internal/model"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestNewAuthService(t *testing.T) {
// 	var authService = NewAuthService(&ASConfig{UserRepository: mocks.NewUserRepository(t), PasswordResetRepository: mocks.NewPassowrdResetRepository(t)})

// 	assert.NotNil(t, authService)
// }

// func Test_authService_Attempt(t *testing.T) {
// 	userRepository := mocks.NewUserRepository(t)
// 	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: mocks.NewPassowrdResetRepository(t)})

// 	t.Run("test success attempt user", func(t *testing.T) {
// 		userRepository.Mock.On("FindByEmail", "saad@user.com").Return(&model.User{ID: 1, Name: "saad", Email: "saad@user.com", Password: "$2a$04$93AZUXoqhOu6TNb481MYke3iDbM8UAzizOHmKSEf36bQtzV3kffwm"}, nil).Once()

// 		input := &dto.LoginRequestBody{}
// 		input.Email = "saad@user.com"
// 		input.Password = "12345"
// 		user, err := authService.Attempt(input)

// 		assert.Nil(t, err)
// 		assert.Equal(t, uint(1), user.ID)
// 		assert.Equal(t, "saad", user.Name)
// 		assert.Equal(t, "saad@user.com", user.Email)
// 	})
// }

// func Test_authService_ForgotPass(t *testing.T) {
// 	userRepository := mocks.NewUserRepository(t)
// 	passwordResetRepository := mocks.NewPassowrdResetRepository(t)
// 	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})

// 	t.Run("test success forgot password", func(t *testing.T) {
// 		userRepository.Mock.On("FindByEmail", "saad@user.com").Return(&model.User{ID: 1, Name: "saad", Email: "saad@user.com", Password: "$2a$04$93AZUXoqhOu6TNb481MYke3iDbM8UAzizOHmKSEf36bQtzV3kffwm"}, nil).Once()
// 		passwordResetRepository.Mock.On("FindByUserId", 1).Return(&model.PasswordReset{}, nil).Once()
// 		passwordResetRepository.Mock.On("Save", mock.Anything).Return(&model.PasswordReset{ID: 1, UserID: 1, Token: "brondol", ExpiredAt: time.Now().Add(time.Minute * 15)}, nil).Once()

// 		input := &dto.ForgotPasswordRequestBody{}
// 		input.Email = "saad@user.com"
// 		passwordReset, err := authService.ForgotPass(input)

// 		assert.Nil(t, err)
// 		assert.Equal(t, uint(1), passwordReset.ID)
// 		assert.Equal(t, "brondol", passwordReset.Token)
// 	})
// }

// func Test_authService_ResetPass(t *testing.T) {
// 	userRepository := mocks.NewUserRepository(t)
// 	passwordResetRepository := mocks.NewPassowrdResetRepository(t)
// 	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})

// 	t.Run("test success reset password", func(t *testing.T) {
// 		user := model.User{ID: 1, Name: "saad", Email: "saad@user.com", Password: "12345"}
// 		passwordReset := &model.PasswordReset{ID: 1, UserID: 1, User: user, Token: "brondol", ExpiredAt: time.Now().Add(time.Minute * 15)}
// 		passwordResetRepository.Mock.On("FindByToken", "brondol").Return(passwordReset, nil).Once()
// 		userRepository.Mock.On("Update", mock.Anything).Return(&user, nil).Once()
// 		passwordResetRepository.Mock.On("Delete", passwordReset).Return(passwordReset, nil).Once()

// 		input := &dto.ResetPasswordRequestBody{}
// 		input.Token = "brondol"
// 		input.Password = "12345"
// 		input.ConfirmPassword = "12345"
// 		passwordReset, err := authService.ResetPass(input)

// 		assert.Nil(t, err)
// 		assert.Equal(t, uint(1), passwordReset.ID)
// 		assert.Equal(t, "brondol", passwordReset.Token)
// 	})
// }
