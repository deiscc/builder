package storage

import (
	"testing"

	"github.com/arschles/assert"
	"github.com/deis/builder/pkg/sys"
)

func TestGetAuthEmptyAuth(t *testing.T) {
	fs := sys.NewFakeFS()
	creds, err := getAuth(fs)
	assert.NoErr(t, err)
	assert.Equal(t, *creds, emptyCreds, "returned credentials")
}

func TestGetAuthMissingSecret(t *testing.T) {
	fs := sys.NewFakeFS()
	fs.Files[accessSecretKeyFile] = []byte("hello world")
	creds, err := getAuth(fs)
	assert.Err(t, err, errMissingKey)
	assert.True(t, creds == nil, "returned credentials were not nil")
}

func TestGetAuthMissingKey(t *testing.T) {
	fs := sys.NewFakeFS()
	fs.Files[accessKeyIDFile] = []byte("hello world")
	creds, err := getAuth(fs)
	assert.Err(t, err, errMissingSecret)
	assert.True(t, creds == nil, "returned credentials were not nil")
}

func TestGetAuthSuccess(t *testing.T) {
	fs := sys.NewFakeFS()
	fs.Files[accessKeyIDFile] = []byte("stuff")
	fs.Files[accessSecretKeyFile] = []byte("other stuff")
	creds, err := getAuth(fs)
	assert.NoErr(t, err)
	assert.True(t, creds != nil, "creds were nil when they shouldn't have been")
}

func TestCredsOKFail(t *testing.T) {
	fs := sys.NewFakeFS()
	assert.False(t, CredsOK(fs), "true returned when there were no credentials")
}

func TestCredsOKSuccess(t *testing.T) {
	fs := sys.NewFakeFS()
	fs.Files[accessKeyIDFile] = []byte("stuff")
	fs.Files[accessSecretKeyFile] = []byte("other stuff")
	assert.True(t, CredsOK(fs), "false returned when there were valid credentials")
}