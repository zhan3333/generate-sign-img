package generate_sign_img_test

import (
	"fmt"
	"generate_sign_img"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	names := []string{
		"张三",
		"李四",
		"王五",
	}
	for _, name := range names {
		err := generate_sign_img.GenerateSignImg(name, name+".png")
		assert.Nil(t, err)
	}
}

func TestRandIntn(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
}
