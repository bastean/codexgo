package captcha

import (
	"github.com/mojocn/base64Captcha"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type Captcha struct {
	ID, Image string
}

type Data struct {
	CaptchaID, CaptchaAnswer string
}

func Generate() (*Captcha, error) {
	id, image, _, err := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore).Generate()

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Generate",
			What:  "Failure to generate captcha",
			Who:   err,
		})
	}

	return &Captcha{
		ID:    id,
		Image: image,
	}, nil
}

func Verify(id, answer string) error {
	if !base64Captcha.DefaultMemStore.Verify(id, answer, false) && !env.IsServerGinModeTest() {
		return errors.New[errors.Failure](&errors.Bubble{
			Where: "Verify",
			What:  "Wrong captcha answer",
			Why: errors.Meta{
				"ID": id,
			},
		})
	}

	return nil
}

func Clear(id string) {
	_ = base64Captcha.DefaultMemStore.Get(id, true)
}
