package protocols

type EmailValidatorInterface interface {
	IsValid(email string) bool
}
