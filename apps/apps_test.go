package apps

import (
	"github.com/dustinkirkland/golang-petname"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Create(t *testing.T) {
	name := "test-" + petname.Name()
	tests := []struct {
		name         string
		ShouldCreate bool
	}{
		{name, true},
		{name, false},
	}
	for _, test := range tests {
		err := App{test.name}.Create()
		if test.ShouldCreate == true {
			assert.Nil(t, err, "Namespace should be created")
		} else {
			assert.NotNil(t, err, "Namespace shouldn't be created")
		}
	}
}
