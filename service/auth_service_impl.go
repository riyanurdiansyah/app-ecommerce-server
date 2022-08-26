package service

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/data/entity"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/repository"
	"app-ecommerce-server/validation"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}

}

// SignUp implements AuthService
func (service *AuthServiceImpl) SignUp(ctx *gin.Context, request *dto.UserCreateDTO) *dto.UserResponseDTO {
	errorValidation := service.Validate.Struct(request)
	println("ERROR IKI ", errorValidation)
	if errorValidation != nil {
		msgError := validation.SignUpValidation(errorValidation.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.CategoryValidation(tx.Error.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		user := entity.User{
			Username:    request.Username,
			Email:       request.Email,
			Password:    request.Password,
			PhoneNumber: request.PhoneNumber,
			SignupWith:  request.SignupWith,
			CreatedAt:   time.Now().Local().String(),
			UpdatedAt:   time.Now().Local().String(),
		}

		userResponse := service.AuthRepository.SignUp(ctx, tx, &user)

		return helper.ToAuthResponseDTO(userResponse)
	}
}
