package coverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCOV(t *testing.T) {
	lcovParser := New(LCOV)
	report, err := lcovParser.Parse("./test/lcov.info")

	assert.NoError(t, err)
	assert.Equal(t, 67, report.CoveredBranches)
	assert.Equal(t, 1757, report.TotalBranches)
	assert.Equal(t, 412, report.CoveredFunctions)
	assert.Equal(t, 1695, report.TotalFunctions)
	assert.Equal(t, 2798, report.CoveredLines)
	assert.Equal(t, 6395, report.TotalLines)
	assert.Equal(t, 182, len(report.Files))

}

func TestCobertura(t *testing.T) {
	coberturaParser := New(COBERTURA)
	report, err := coberturaParser.Parse("./test/coverage.xml")

	assert.NoError(t, err)
	assert.Equal(t, 8, report.TotalLines)
	assert.Equal(t, 6, report.CoveredLines)
	assert.Equal(t, 4, report.TotalBranches)
	assert.Equal(t, 2, report.CoveredBranches)
}

func TestClover(t *testing.T) {
	cloverParser := New(CLOVER)
	report, err := cloverParser.Parse("./test/clover.xml")
	assert.NoError(t, err)

	assert.Equal(t, 20672, report.TotalLines)
	assert.Equal(t, 5337, report.CoveredLines)
	assert.Equal(t, 12764, report.TotalBranches)
	assert.Equal(t, 1123, report.CoveredBranches)
	assert.Equal(t, 2485, report.TotalFunctions)
	assert.Equal(t, 1282, report.CoveredFunctions)
}
