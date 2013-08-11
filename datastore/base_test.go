// +build !appengine

package datastore

import (
	"appengine"
	"github.com/MCN-Healthcare/categorizer/models"
	"github.com/MCN-Healthcare/categorizer/tests"
	"github.com/najeira/testbed"
	"github.com/remogatto/prettytest"
	"testing"
)

const (
	testStr string = "foo bar"
	testInt int    = 17
)

var (
	testStrSlice []string = []string{
		"foo",
		"bar",
		"foo bar",
		"baz",
		"fizz buzz",
		"pancake",
		"cherry",
	}
)

type testModel struct {
	models.BaseModel
	TestStr string
	TestInt int
}

func kindForTest() string {
	return "test"
}

func (model *testModel) Kind() string {
	return kindForTest()
}

func testData() []models.Model {
	testData := make([]models.Model, len(testStrSlice))
	var next *testModel
	for index, strVal := range testStrSlice {
		next = new(testModel)
		next.TestStr = strVal
		next.TestInt = index
		testData[index] = next
	}
	return testData
}

func loadTestData(ctx appengine.Context) []models.Model {
	testData := testData()
	PutModels(ctx, testData)
	return testData
}

type datastoreTestSuite struct {
	prettytest.Suite
	GAEContext appengine.Context
	TestBed    *testbed.Testbed
}

func TestDatastoreSuite(t *testing.T) {
	prettytest.Run(t, new(datastoreTestSuite))
}

func (suite *datastoreTestSuite) Before() {
	suite.TestBed, suite.GAEContext = tests.MakeBedAndGAEContext()
}

func (suite *datastoreTestSuite) After() {
	suite.TestBed.Close()
}
