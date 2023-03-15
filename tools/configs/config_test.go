package configs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfigs(t *testing.T) {
	_, err := GetConfigs("config", "../../configs/")
	if err != nil {
		require.Nil(t, err)
	}
}
