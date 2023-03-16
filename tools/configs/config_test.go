package configs

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestGetConfigs(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		_, err := GetConfigs("config", "../../configs/")
		require.Nil(t, err)
	})
	t.Run("error finding file", func(t *testing.T) {
		_, err := GetConfigs("configs", "../../configs/")
		require.Error(t, viper.ConfigFileNotFoundError{}, err)
	})
}
