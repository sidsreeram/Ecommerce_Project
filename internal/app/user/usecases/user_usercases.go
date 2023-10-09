package usercases

import (
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/ecommerce/internal/app/user/entities"
	"github.com/ecommerce/internal/app/user/repositories"
	"github.com/ecommerce/internal/app/user/services"
	"github.com/ecommerce/internal/model"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
	userService    services.UserService
}

func NewUserUseCase(userRepo repositories.UserRepository, userService services.UserService) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepo,
		userService:    userService,
	}
}

func (uc *UserUsecase) Register(email, password, confirmPassword string, user *entities.SignUpRequest) error {
	// Check if the email already exists
	

	exists, err  := uc.userRepository.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// Check if the password and confirm password match
	if password != confirmPassword {
		return errors.New("password doesn't match")
	}
	var otptoken entities.OTPToken
	otp := services.GenerateOTP()
	otptoken.OTPCode = otp 
	otptoken.ExpirationTime = time.Now().Add(time.Minute *5)
	
	SndErr := uc.SendOTPByEmail(email,otp)
	if SndErr != nil {
		return errors.New("Error in sending OTP")
	}
	
	// Hash the password
	var hashedPassword string 

	

	// Now you can use 'hashedPassword' for storing in the database or further processing

	// Create a new user entity with the provided details
	newUser := &model.User{
		Email:    email,
		Password: string(hashedPassword), // Store the hashed password
		// Other fields...
	}

	// Use the UserRepository to create the user in the database
	if err := uc.userRepository.Create(newUser); err != nil {
		return err
	}

	return nil
}
func (uc *UserUsecase) SendOTPByEmail(email, token string) error {
	auth := smtp.PlainAuth(
		"",
		"sidx141202@gmail.com", // Replace with your Gmail email
		"yhlm wzqg wobt deww",  // Replace with your Gmail password
		"smtp.gmail.com",
	)

	msg := "Subject:" + "OTP Verification" + "\n" + token

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"sidx141202@gmail.com", // Replace with your Gmail email
		[]string{email},
		[]byte(msg),
	)
	fmt.Println(err)
	return nil // Return an error if sending fails, or nil if it's successful
}


