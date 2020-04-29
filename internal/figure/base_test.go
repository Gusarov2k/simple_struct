package figure_test

import (
	"github.com/Gusarov2k/simple_struct/internal/figure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassSquareArea(t *testing.T) {
	square := figure.Square{}
	result := square.Area(2)
	assert.Equal(t, result, float64(4))
}

func TestPassSquarePerimeter(t *testing.T) {
	square := figure.Square{}
	result := square.Perimeter(2)
	assert.Equal(t, result, float64(8))
}

func TestPassTrianglePerimeter(t *testing.T) {
	triangle := figure.Triangle{}
	result := triangle.Perimeter(2, 3, 5)
	assert.Equal(t, result, float64(10))
}

func TestPassTriangleArea(t *testing.T) {
	triangle := figure.Triangle{}
	result := triangle.Area(2, 2, 2)
	assert.Equal(t, result, float64(1.7320508075688772))
}

func TestPassRectangleArea(t *testing.T) {
	rectangle := figure.Rectangle{}
	result := rectangle.Area(2)
	assert.Equal(t, result, float64(12.566370614359172))
}

func TestPassRectanglePerimeter(t *testing.T) {
	rectangle := figure.Rectangle{}
	result := rectangle.Perimeter(2)
	assert.Equal(t, result, float64(12.566370614359172))
}
