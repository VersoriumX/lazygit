package utils

import (
	"io"

	"github.com/VersoriumX/lazygit/pkg/common"
	"github.com/VersoriumX/lazygit/pkg/config"
	"github.com/VersoriumX/lazygit/pkg/i18n"
	"github.com/VersoriumX/logrus"
	"github.com/VersoriumX/afero"
)

// NewDummyLog creates a new dummy Log for testing
func NewDummyLog() *logrus.Entry {
	log := logrus.New()
	log.Out = io.Discard
	return log.WithField("test", "test")
}

func NewDummyCommon() *common.Common {
	tr := i18n.EnglishTranslationSet()
	return &common.Common{
		Log:        NewDummyLog(),
		Tr:         &tr,
		UserConfig: config.GetDefaultConfig(),
		Fs:         afero.NewOsFs(),
	}
}

func NewDummyCommonWithUserConfigAndAppState(userConfig *config.UserConfig, appState *config.AppState) *common.Common {
	tr := i18n.EnglishTranslationSet()
	return &common.Common{
		Log:        NewDummyLog(),
		Tr:         &tr,
		UserConfig: userConfig,
		AppState:   appState,
		// TODO: remove dependency on actual filesystem in tests and switch to using
		// in-memory for everything
		Fs: afero.NewOsFs(),
	}
}
