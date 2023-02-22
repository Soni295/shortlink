package utils_test

import (
	"fmt"
	"testing"

	"github.com/Soni295/shortlink/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortNameHandler(t *testing.T) {

	reference := make([]string, 1000)

	setLength := 12
	t.Run(fmt.Sprintf("it length is should be :%v", setLength), func(t *testing.T) {

		for i := range reference {
			reference[i] = utils.GenerateShorNameHandler(setLength)

			if len(reference[i]) != setLength {
				assert.Equal(t, len(reference[i]), setLength)
			}
		}
	})
}
