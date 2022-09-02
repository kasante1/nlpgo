package checkfile

import (
	"errors"
	"os"
)

func checkfile(filepath string)bool{
	_, error := os.Stat(filepath)
	return !errors.Is(error, os.ErrNotExist)
}