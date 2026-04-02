package storage

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathtransformFunc(t *testing.T) {
	key := "momsbestpicture"
	path := CASPathTransformFunc(key)
	fmt.Println(path)
	assert.Equal(t, path.PathName, "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff")
	assert.Equal(t, path.Original, "6804429f74181a63c50c3d81d733a12f14a353ff")

}
func TestStorage(t *testing.T) {
	storageopts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(storageopts)
	err := store.WriteStream("C:/Users/Satyajit Samal/OneDrive/Desktop/go/Top_Projects/", strings.NewReader("Hello from This Side"))
	if err != nil {
		t.Error(err)

	}
	assert.Nil(t, err)

}
