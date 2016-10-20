package forms

import (
	"github.com/revel/revel"
	"matrix/modules/auth/models"
	"strings"
)

type UserDetailForm struct {
	User models.User

	PasswordAgain    string
	NewPassword      string
	NewPasswordAgain string
}

func (f *UserDetailForm) IsCreate() bool {
	return f.User.Id == 0
}

func (f *UserDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.User.LoginName).Message("登录名不能为空！")
	validation.MinSize(f.User.LoginName, 3).Message("登录名长度不能小于3！")
	validation.MaxSize(f.User.LoginName, 10).Message("登录名长度不能大于10！")

	validation.Required(f.User.NickName).Message("用户名不能为空！")
	validation.MinSize(f.User.NickName, 3).Message("用户名长度不能小于3！")
	validation.MaxSize(f.User.NickName, 10).Message("用户名长度不能大于10！")

	if f.IsCreate() {
		validation.Required(f.User.Password).Message("密码不能为空！")
		validation.MinSize(f.User.Password, 6).Message("密码长度不能小于6！")
		validation.MaxSize(f.User.Password, 12).Message("密码长度不能大于12！")

		if f.User.Password != f.PasswordAgain {
			validation.Errors = append(validation.Errors, &revel.ValidationError{
				Key:     "password_again_not_match",
				Message: "两次输入的密码不一致！",
			})
		}
	} else {

		if strings.Trim(f.NewPassword, " ") != "" || strings.Trim(f.NewPasswordAgain, " ") != "" {
			validation.MinSize(f.NewPassword, 6).Message("密码长度不能小于6！")
			validation.MaxSize(f.NewPassword, 12).Message("密码长度不能大于12！")

			if f.NewPassword != f.NewPasswordAgain {
				validation.Errors = append(validation.Errors, &revel.ValidationError{
					Key:     "password_again_not_match",
					Message: "两次输入的密码不一致！",
				})
			}
		}
	}

	return validation.HasErrors() == false
}

//-----------------------------------------------------------------------------------------------------------

type LoginForm struct {
	LoginName string
	Password  string
}

func (f *LoginForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.LoginName).Message("登录名不能为空！")
	validation.MinSize(f.LoginName, 3).Message("登录名长度不能小于3！")
	validation.MaxSize(f.LoginName, 10).Message("登录名长度不能大于10！")

	validation.Required(f.Password).Message("密码不能为空！")
	validation.MinSize(f.Password, 6).Message("密码长度不能小于6！")
	validation.MaxSize(f.Password, 12).Message("密码长度不能大于12！")

	return validation.HasErrors() == false
}

//-----------------------------------------------------------------------------------------------------------

type ProfileForm struct {
	Password      string
	PasswordAgain string
}

func (f *ProfileForm) Valid(validation *revel.Validation) bool {
	if strings.Trim(f.Password, " ") != "" || strings.Trim(f.PasswordAgain, " ") != "" {
		validation.Required(f.Password).Message("密码不能为空！")
		validation.MinSize(f.Password, 6).Message("密码长度不能小于6！")
		validation.MaxSize(f.Password, 12).Message("密码长度不能大于12！")

		if f.Password != f.PasswordAgain {
			validation.Errors = append(validation.Errors, &revel.ValidationError{
				Key:     "password_again_not_match",
				Message: "两次输入的密码不一致！",
			})
		}
	}

	return validation.HasErrors() == false
}
