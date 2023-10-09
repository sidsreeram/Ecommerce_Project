package services

import (
	"math/rand"
	"strconv"
	"time"
)

type UserService interface {

}
func GenerateOTP() string {
	seed := time.Now().UnixNano()
	randomGenerator := rand.New(rand.NewSource(seed))
	return strconv.Itoa(randomGenerator.Intn(8999) + 1000)
}
