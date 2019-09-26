package test_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/cf-deployment/units/helpers"
	"gopkg.in/yaml.v2"
)

const testDirectory = "operations/test"

func TestTest(t *testing.T) {
	content, err := ioutil.ReadFile("operations.yml")
	if err != nil {
		fmt.Printf("Error reading operations file: %s", err)
	}
	testTests := make(map[string]helpers.OpsFileTestParams)
	yaml.Unmarshal(content, &testTests)

	cfDeploymentHome, err := helpers.SetPath()
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	suite := helpers.NewSuiteTest(cfDeploymentHome, testDirectory, testTests)
	suite.EnsureTestCoverage(t)
	suite.InterpolateTest(t)
}
