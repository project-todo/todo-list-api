package task

import (
	"testing"
)

func TestMemoryRepository(t *testing.T) {
	repository := NewMemoryRepository()
	repositoryTest := RepositoryTest{
		repository: &repository,
	}
	repositoryTest.TestAll(t)
}
