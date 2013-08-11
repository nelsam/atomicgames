package datastore

import (
	"appengine"
	"appengine/datastore"
	"github.com/MCN-Healthcare/categorizer/models"
)

func DeleteModel(ctx appengine.Context, model models.Model) error {
	key, err := model.Key()
	if err != nil {
		return err
	}
	return datastore.Delete(ctx, key)
}

func DeleteMultiModel(ctx appengine.Context, modelSlice []models.Model) error {
	keys := make([]*datastore.Key, len(modelSlice))
	var (
		key *datastore.Key
		err error
	)
	for index, model := range modelSlice {
		key, err = model.Key()
		if err != nil {
			return err
		}
		keys[index] = key
	}
	return datastore.DeleteMulti(ctx, keys)
}
