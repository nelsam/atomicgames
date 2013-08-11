// +build !appengine

package datastore

import (
	"appengine/datastore"
	"github.com/MCN-Healthcare/categorizer/models"
	"github.com/stretchr/testify/assert"
)

func (suite *datastoreTestSuite) TestGetModel() {
	models := loadTestData(suite.GAEContext)

	src := models[0]
	key, err := src.Key()
	assert.NotNil(suite.T, key, "The src model should have its key set")

	dest := new(testModel)
	err = GetModel(suite.GAEContext, key, dest)
	assert.Nil(suite.T, err, "GetModel shouldn't return an error")
	assert.ObjectsAreEqual(src, dest)
}

func (suite *datastoreTestSuite) TestGetModelWithKey() {
	models := loadTestData(suite.GAEContext)

	src := models[0]
	key, err := src.Key()
	assert.NotNil(suite.T, key, "The src model should have its key set")

	dest := new(testModel)
	dest.SetKey(key)
	err = GetModelWithKey(suite.GAEContext, dest)
	assert.Nil(suite.T, err, "GetModelWithKey shouldn't return an error")
	assert.ObjectsAreEqual(src, dest)
}

func (suite *datastoreTestSuite) TestGetModels() {
	srcModels := loadTestData(suite.GAEContext)

	destModels := make([]models.Model, 5)
	destTest := make([]*testModel, 5)
	keys := make([]*datastore.Key, 5)
	var (
		key *datastore.Key
		err error
	)
	for index, _ := range destTest {
		key, err = srcModels[index].Key()
		keys[index] = key
		destTest[index] = new(testModel)
		destModels[index] = destTest[index]
	}

	err = GetModels(suite.GAEContext, keys, destModels)
	assert.Nil(suite.T, err, "GetModels shouldn't return an error")
	for index, model := range destTest {
		assert.ObjectsAreEqual(model, srcModels[index])
		key, err = model.Key()
		assert.Nil(suite.T, err, "Models should have keys after running Get")
		assert.NotNil(suite.T, key, "Models should have keys after running Get")
	}
}

func (suite *datastoreTestSuite) TestGetModelsWithKeys() {
	srcModels := loadTestData(suite.GAEContext)

	destModels := make([]models.Model, 5)
	destTest := make([]*testModel, 5)
	var (
		key *datastore.Key
		err error
	)
	for index, _ := range destTest {
		key, err = srcModels[index].Key()
		destTest[index] = new(testModel)
		destTest[index].SetKey(key)
		destModels[index] = destTest[index]
	}

	err = GetModelsWithKeys(suite.GAEContext, destModels)
	assert.Nil(suite.T, err, "GetModelsWithKeys shouldn't error")
	for index, model := range destTest {
		assert.ObjectsAreEqual(model, srcModels[index])
		key, err = model.Key()
		assert.Nil(suite.T, err, "Models should have keys after running Get")
		assert.NotNil(suite.T, key, "Models should have keys after running Get")
	}
}
