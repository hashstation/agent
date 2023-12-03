package utils

import (
	"github.com/gofrs/flock"
	"os"
	"path/filepath"
	"runtime"
)

type NamedFileLock struct {
	Name     string
	FileLock *flock.Flock
}

func NewNamedFileLock(lockName string) *NamedFileLock {
	var lockPath string
	if runtime.GOOS == "windows" {
		lockPath = filepath.Join(os.TempDir(), lockName)
	} else {
		lockPath = filepath.Join("/var/lib/boinc-client/", lockName)
	}
	
	return &NamedFileLock{
		FileLock: flock.New(lockPath),
	}
}

func (l *NamedFileLock) Lock() error {
	return l.FileLock.Lock()
}

func (l *NamedFileLock) Unlock() error {
	if !l.FileLock.Locked() {
		return nil
	}
	return l.FileLock.Unlock()
}
