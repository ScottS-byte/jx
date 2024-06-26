package upgrade_test

import (
	"testing"

	"github.com/jenkins-x/jx/pkg/cmd/upgrade"
	"github.com/stretchr/testify/require"
)

func TestUpgrade(t *testing.T) {
	t.Parallel()

	cmd, _ := upgrade.NewCmdUpgradePlugins()

	err := cmd.Execute()
	require.NoError(t, err, "failed to run upgrade command")
}
