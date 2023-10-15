package domain

import (
	"go-gin-clean-arch/domain/vobj"
	"go-gin-clean-arch/packages/context"
	"go-gin-clean-arch/resource/request"
)

type User struct {
	SoftDeleteModel
	Email    string        `json:"email" validate:"required" gorm:"index;unique"`
	Password vobj.Password `json:"-"`

	RecoveryToken *vobj.RecoveryToken `json:"-" gorm:"index"`
}

func NewUser(ctx context.Context, dto *request.UserCreate) (*User, error) {
	user := User{
		Email:         dto.Email,
		RecoveryToken: vobj.NewRecoveryToken(""),
	}

	ctx.Validate(user)

	password, err := vobj.NewPassword(ctx, dto.Password, dto.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	user.Password = *password

	return &user, nil
}

func (u *User) ResetPassword(ctx context.Context, dto *request.UserResetPassword) error {
	if !u.RecoveryToken.IsValid() {
		ctx.FieldError("RecoveryToken", "リカバリートークンが無効です")
		return nil
	}

	password, err := vobj.NewPassword(ctx, dto.Password, dto.PasswordConfirm)
	if err != nil {
		return err
	}

	u.Password = *password

	u.RecoveryToken.Clear()
	return nil
}
