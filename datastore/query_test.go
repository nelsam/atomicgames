// +build !appengine

package datastore

import (
	"appengine/datastore"
	"github.com/MCN-Healthcare/categorizer/models"
	"github.com/stretchr/testify/assert"
)

func (suite *datastoreTestSuite) TestModelQuery_GetAll() {
	loadTestData(suite.GAEContext)

	query := NewModelQuery(kindForTest())
	count, err := query.Count(suite.GAEContext)
	assert.Nil(suite.T, err, "Query should be countable.")
	assert.NotEqual(suite.T, count, 0, "Data should exist in the database by now.")

	var modelSlice []models.Model = make([]models.Model, count)
	for index := range modelSlice {
		modelSlice[index] = new(testModel)
	}
	query.GetAll(suite.GAEContext, modelSlice)
	var key *datastore.Key
	for _, entry := range modelSlice {
		key, err = entry.Key()
		assert.NotNil(suite.T, key, "ModelQuery.GetAll should populate entry keys.")
	}
}

func (suite *datastoreTestSuite) TestModelQuery_GetAll_KeysOnly() {
	loadTestData(suite.GAEContext)

	query := NewModelQuery(kindForTest()).
		KeysOnly()
	var modelSlice []models.Model
	keys, err := query.GetAll(suite.GAEContext, modelSlice)

	assert.Nil(suite.T, err, "GetAll should not return error.")
	assert.NotEqual(suite.T, len(keys), 0, "Keys should have been returned.")
}

func (suite *datastoreTestSuite) TestModelIterator() {
	loadTestData(suite.GAEContext)

	query := NewModelQuery(kindForTest())
	count, err := query.Count(suite.GAEContext)
	assert.Nil(suite.T, err, "Query should be countable.")
	assert.NotEqual(suite.T, count, 0, "Data should exist in the database by now.")

	iterator := query.Run(suite.GAEContext)

	var (
		entry models.Model = new(testModel)
		key   *datastore.Key
	)
	for _, err := iterator.Next(entry); err != datastore.Done; _, err = iterator.Next(entry) {
		if err != nil {
			panic(err)
		}
		key, err = entry.Key()
		assert.NotNil(suite.T, key, "Query iterators should populate entry keys.")
	}
}
