package bar_test

import (
	bar "go118bug"
	"testing"
)

func TestContextMeta_crud(t *testing.T) {
	bar.GetData()
}
