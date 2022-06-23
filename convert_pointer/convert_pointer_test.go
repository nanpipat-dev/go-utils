package convert_pointer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringPointerToString(t *testing.T) {
	expect := ""

	result := StringPointerToString(nil)

	assert.Equal(t, result, expect)
}

func TestTimePointerToTime(t *testing.T) {
	expect := time.Time{}

	result := TimePointerToTime(nil)

	assert.Equal(t, result, expect)
}

func TestBooleanPointerToBoolean(t *testing.T) {
	expect := false

	result := BooleanPointerToBoolean(nil)

	assert.Equal(t, result, expect)
}

func TestIntPointerToInt(t *testing.T) {
	expect := 0

	result := IntPointerToInt(nil)

	assert.Equal(t, result, expect)
}

func TestInt32PointerToInt32(t *testing.T) {
	expect := int32(0)

	result := Int32PointerToInt32(nil)

	assert.Equal(t, result, expect)
}
func TestInt64PointerToInt64(t *testing.T) {
	expect := int64(0)

	result := Int64PointerToInt64(nil)

	assert.Equal(t, result, expect)
}
