// +build !appengine

package datastore

import (
	"github.com/stretchr/testify/assert"
)

func (suite *datastoreTestSuite) TestPutModel() {
	test := new(testModel)
	test.TestStr = testStr
	test.TestInt = testInt

	PutModel(suite.GAEContext, test)

	key, err := test.Key()
	assert.Nil(suite.T, err, "Error should be nil.")
	assert.NotNil(suite.T, key, "PutModel should set test's key value.")
	assert.Equal(suite.T, test.TestStr, testStr, "PutModel should not alter values other than key")
	assert.Equal(suite.T, test.TestInt, testInt, "PutModel should not alter values other than key")
}

func (suite *datastoreTestSuite) TestPutModels() {
	models := testData()
	PutModels(suite.GAEContext, models)

	query := NewModelQuery(kindForTest())
	count, err := query.Count(suite.GAEContext)
	assert.Nil(suite.T, err, "The query should be countable.")
	assert.Equal(suite.T, count, len(models), "The number of models available from the DB "+
		"should be equal to the number models saved")
}
