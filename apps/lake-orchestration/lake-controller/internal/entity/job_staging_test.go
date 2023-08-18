package entity

import (
     "testing"
     // "github.com/stretchr/testify/assert"
     "github.com/stretchr/testify/suite"
)

type JobStagingSuite struct {
     suite.Suite
}

func TestJobStagingSuite(t *testing.T) {
     suite.Run(t, new(JobStagingSuite))
}
