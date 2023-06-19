package validator

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
	"sync"
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

func (c *Validator) Validate(i interface{}) error {
	c.lazyInit()
	if err := c.validate.Struct(i); err != nil {
		var k []string
		for _, err := range err.(validator.ValidationErrors) {
			k = append(k, err.Translate(c.trans))
		}
		return errors.New(strings.Join(k, ","))
	}
	return nil
}

func (c *Validator) lazyInit() {
	c.once.Do(func() {
		zhCh := zh.New()
		uni := ut.New(zhCh)
		trans, _ := uni.GetTranslator("zh")
		validate := validator.New()
		//验证器注册翻译器
		_ = translations.RegisterDefaultTranslations(validate, trans)
		c.validate = validate
		c.trans = trans
	})
}
