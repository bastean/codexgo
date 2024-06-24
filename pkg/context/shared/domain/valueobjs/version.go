package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Version struct {
	Version string `validate:"number"`
}

func (version *Version) Value() string {
	return version.Version
}

func (version *Version) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(version)
}

func NewVersion(version string) (models.ValueObject[string], error) {
	version = strings.TrimSpace(version)

	versionVO := &Version{
		Version: version,
	}

	if versionVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewVersion",
			What:  "version must be numeric only",
			Why: errors.Meta{
				"Version": version,
			},
		})
	}

	return versionVO, nil
}
