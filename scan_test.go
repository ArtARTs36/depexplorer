package depexplorer_test

import (
	"github.com/artarts36/depexplorer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScanProjectThisRepository(t *testing.T) {
	files, err := depexplorer.ScanProjectDir("./")
	require.NoError(t, err)
	require.NotNil(t, files[depexplorer.DependencyManagerGoMod])
}
