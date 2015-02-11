package zipcode

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_Lookup(t *testing.T) {
	assert.Equal(t, "", Lookup("00000"))
	assert.Equal(t, "39.882703,-74.972036", Lookup("08003"))
}
