package user

import (
	"context"

	"github.com/joaomarcosg/Projeto-Gobid/internal/validator"
)

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.Checkfield(validator.NotBlank(req.UserName), "user_name", "this field cannot be empty")
	eval.Checkfield(validator.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.Checkfield(validator.Matches(req.Email, validator.EmailRX), "email", "must be a valid email")
	eval.Checkfield(validator.NotBlank(req.Bio), "bio", "this field cannot be empty")
	eval.Checkfield(
		validator.MinChars(req.Bio, 10) &&
			validator.MaxChars(req.Bio, 255), "bio", "this field must have a length between 10 and 255")
	eval.Checkfield(validator.MinChars(req.Password, 8), "password", "password must be bigger than 8 chars")

	return eval
}
