package task

import (
	"testing"
)

func TestMemoryRepository(t *testing.T) {
	repositoryTest := RepositoryTest{
		repository: NewMemoryRepository(),
	}
	repositoryTest.TestAll(t)
}
