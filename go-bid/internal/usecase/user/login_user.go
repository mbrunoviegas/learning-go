package user

import (
	"context"
	"go-bid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "email must be a valid email")
	eval.CheckField(validator.NotBlank(req.Password), "password", "password cannot be blank")

	return eval
}
