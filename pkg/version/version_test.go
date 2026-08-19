/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the MIT license, see LICENSE in the project root for details.
*/
package version

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func __capture(f func()) string {
	originalStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = originalStdout

	var buf = make([]byte, 5096)
	n, _ := r.Read(buf)
	return string(buf[:n])
}

func releaseReturnsDevWhenVersionIsEmpty(t *testing.T) {
	Version = ""
	expected := "dev"
	assert.Equal(t, expected, Release())
}

func commitReturnsEmptyStringWhenGitCommitIsEmpty(t *testing.T) {
	GitCommit = ""
	expected := ""
	assert.Equal(t, expected, Commit())
}

func releaseReturnsVersionWhenVersionIsSet(t *testing.T) {
	Version = gofakeit.Numerify("v#.#.#")
	expected := Version
	assert.Equal(t, expected, Release())
}

func commitReturnsCommitHashWhenGitCommitIsSet(t *testing.T) {
	GitCommit = gofakeit.ProductISBN(nil)
	expected := GitCommit
	assert.Equal(t, expected, Commit())
}

func bannerReturnsLogo(t *testing.T) {
	expected := `
     _               _   _
  __| | _____  _____| |_| |
 / _  |/ _ \ \/ / __| __| |
| (_| |  __/>  < (__| |_| |
 \__,_|\___/_/\_\___|\__|_|
`

	assert.Equal(t, expected, Banner())
	assert.Len(t, Banner(), 140)
}

func printReturnsVersionAndCommit(t *testing.T) {
	Version = gofakeit.Numerify("v#.#.#")
	GitCommit = gofakeit.ProductISBN(nil)

	output := __capture(Print)
	assert.Contains(t, output, "Release: "+Release())
	assert.Contains(t, output, "Commit:  "+Commit())
	assert.Contains(t, output, "\033[34m")
	assert.Contains(t, output, "\033[0m")
}

func printReturnsVersionAndCommitWithNoColor(t *testing.T) {
	Version = gofakeit.Numerify("v#.#.#")
	GitCommit = gofakeit.ProductISBN(nil)

	_ = os.Setenv("NO_COLOR", "1")
	defer func() {
		_ = os.Unsetenv("NO_COLOR")
	}()
	output := __capture(Print)
	assert.Contains(t, output, "Release: "+Release())
	assert.Contains(t, output, "Commit:  "+Commit())
	assert.NotContains(t, output, "\033[34m")
	assert.NotContains(t, output, "\033[0m")
}

func TestVersion(t *testing.T) {
	t.Run("version.Release returns 'dev' when Version is empty", releaseReturnsDevWhenVersionIsEmpty)
	t.Run("version.Commit returns empty string when GitCommit is empty", commitReturnsEmptyStringWhenGitCommitIsEmpty)
	t.Run("version.Release returns Version when Version is set", releaseReturnsVersionWhenVersionIsSet)
	t.Run("version.Commit returns commit hash when GitCommit is set", commitReturnsCommitHashWhenGitCommitIsSet)
	t.Run("version.Banner returns the correct logo", bannerReturnsLogo)
	t.Run("version.Print prints logo, version and commit", printReturnsVersionAndCommit)
	t.Run("version.Print prints logo, version and commit with no color", printReturnsVersionAndCommitWithNoColor)

}
