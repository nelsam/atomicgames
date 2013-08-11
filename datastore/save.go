package datastore

import (
	"appengine"
	"appengine/datastore"
	"github.com/MCN-Healthcare/categorizer/models"
)

// Store a Model, and call SetKey after generating its key (if it doesn't
// have one already).
func PutModel(ctx appengine.Context, model models.Model) error {
	modelKey, err := model.Key()
	if err == models.NoKeyLoaded {
		modelKey = datastore.NewIncompleteKey(ctx, model.Kind(), nil)
	} else if err != nil {
		panic(err)
	}
	key, err := datastore.Put(ctx, modelKey, model)
	if err == nil {
		model.SetKey(key)
	}
	return err
}

// Store multiple models, calling SetKey after generating the key (if it
// doesn't have one already)
func PutModels(ctx appengine.Context, modelSlice []models.Model) error {
	keys := make([]*datastore.Key, len(modelSlice))
	var err error
	for index, model := range modelSlice {
		keys[index], err = model.Key()
		if err == models.NoKeyLoaded {
			keys[index] = datastore.NewIncompleteKey(ctx, model.Kind(), nil)
		}
	}
	newKeys, err := datastore.PutMulti(ctx, keys, modelSlice)
	for index, model := range modelSlice {
		model.SetKey(newKeys[index])
	}
	return err
}
