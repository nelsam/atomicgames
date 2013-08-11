package datastore

import (
	"appengine"
	"appengine/datastore"
	"github.com/MCN-Healthcare/categorizer/models"
)

// Load a datastore value into a model and add the key
func GetModel(ctx appengine.Context, key *datastore.Key, model models.Model) error {
	err := datastore.Get(ctx, key, model)
	if err == nil {
		model.SetKey(key)
	}
	return err
}

// If the key has already been added to the model, we can just generate
// the key argument from the model.
func GetModelWithKey(ctx appengine.Context, model models.Model) error {
	key, err := model.Key()
	if err != nil {
		return err
	}
	return GetModel(ctx, key, model)
}

// Load multiple datastore values into a slice of models
func GetModels(ctx appengine.Context, keys []*datastore.Key, modelSlice []models.Model) error {
	err := datastore.GetMulti(ctx, keys, modelSlice)
	var (
		index int
		model models.Model
	)
	for index, model = range modelSlice {
		model.SetKey(keys[index])
	}
	return err
}

// If the keys have been added to the models already, we can just generate
// the keys argument from the models.
func GetModelsWithKeys(ctx appengine.Context, modelSlice []models.Model) error {
	keys := make([]*datastore.Key, len(modelSlice))
	for index, model := range modelSlice {
		key, err := model.Key()
		if err != nil {
			return err
		}
		keys[index] = key
	}
	return GetModels(ctx, keys, modelSlice)
}
