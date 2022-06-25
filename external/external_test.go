package external

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockExt := NewMockExternal(ctrl)
	var x int32 = 3
	var y int32 = 4
	var z int32 = 7
	mockExt.EXPECT().Calculate(x, y).Return(z)

	res := FinalCalculation(mockExt, x, y)
	assert.Equal(z, res, "result must be equal")
}
