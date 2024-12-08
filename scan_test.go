package depexplorer_test

import (
	"github.com/artarts36/depexplorer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScanProjectThisRepository(t *testing.T) {
	file, err := depexplorer.ScanProjectDir("./")
	require.NoError(t, err)
	assert.Equal(t, "go.mod", file.Name)
	assert.Equal(t, depexplorer.DependencyManagerGoMod, file.DependencyManager)
}
