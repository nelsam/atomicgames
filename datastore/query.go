package datastore

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"fmt"
	"github.com/MCN-Healthcare/categorizer/models"
)

// An Iterator for models, which automatically adds an item's key
// if it implements Model
type ModelIterator struct {
	*datastore.Iterator
}

// Call the parent Next and set the model's key
func (iterator *ModelIterator) Next(destModel models.Model) (*datastore.Key, error) {
	modelKey, err := iterator.Iterator.Next(destModel)
	if err == nil {
		destModel.SetKey(modelKey)
	}
	return modelKey, err
}

// We prefer to have easy access to datastore keys in our models,
// so we're re-defining some query operations to handle that.
// Unfortunately, most Query methods return new Query instances,
// instead of executing in-place, so we'll have to have a lot of
// boring boilerplate code to make those calls return ModelQuery
// instances instead of datastore.Query instances.
type ModelQuery struct {
	*datastore.Query
	keysOnly bool
}

func NewModelQuery(kind string) *ModelQuery {
	return &ModelQuery{datastore.NewQuery(kind), false}
}

func (query *ModelQuery) Ancestor(ancestor *datastore.Key) *ModelQuery {
	return &ModelQuery{query.Query.Ancestor(ancestor), query.keysOnly}
}

func (query *ModelQuery) Distinct() *ModelQuery {
	return &ModelQuery{query.Query.Distinct(), query.keysOnly}
}

func (query *ModelQuery) End(cursor datastore.Cursor) *ModelQuery {
	return &ModelQuery{query.Query.End(cursor), query.keysOnly}
}

func (query *ModelQuery) Filter(filterStr string, value interface{}) *ModelQuery {
	return &ModelQuery{query.Query.Filter(filterStr, value), query.keysOnly}
}

// The only place where any real work is done.  Does the same thing as
// datastore.Query.GetAll, but loads each entry's datastore.Key into the
// BaseModel.SetKey() method if the entry implements Model.
//
// BIG FAT WARNING: Use query.Count() combined with make([]models.Model, count)
// to initialize the third parameter (except when running a keys only query).
// If the length of your slice of models.Model is different than the count
// of values the query is going to return, it will just return an error.
func (query *ModelQuery) GetAll(ctx appengine.Context, modelSlice []models.Model) ([]*datastore.Key, error) {
	if query.keysOnly {
		return query.Query.GetAll(ctx, nil)
	} else {
		// Without reflect.ValueOf (which uses the unsafe package and
		// can't be used on GAE), we can't really use
		// query.Query.GetAll() here.  So just use the iterator
		// instead.
		var (
			key   *datastore.Key
			err   error
			count int
		)
		count, err = query.Count(ctx)
		if err != nil {
			return nil, err
		} else if count != len(modelSlice) {
			// Without reflect, we can't find the base type, so we can't
			// create new elements, so the query and slice sizes have to
			// be equal.
			message := fmt.Sprintf("Query size is %d while length of []models.Model "+
				"argument is %d",
				count, len(modelSlice))
			return nil, errors.New(message)
		}
		keys := make([]*datastore.Key, len(modelSlice))
		iterator := query.Run(ctx)
		for index, model := range modelSlice {
			key, err = iterator.Next(model)
			if err != nil {
				break
			} else {
				keys[index] = key
				model.SetKey(key)
			}
		}
		if err == datastore.Done {
			err = nil
		}
		return keys, err
	}
}

func (query *ModelQuery) KeysOnly() *ModelQuery {
	return &ModelQuery{query.Query.KeysOnly(), true}
}

func (query *ModelQuery) Limit(limit int) *ModelQuery {
	return &ModelQuery{query.Query.Limit(limit), query.keysOnly}
}

func (query *ModelQuery) Offset(offset int) *ModelQuery {
	return &ModelQuery{query.Query.Offset(offset), query.keysOnly}
}

func (query *ModelQuery) Order(fieldName string) *ModelQuery {
	return &ModelQuery{query.Query.Order(fieldName), query.keysOnly}
}

func (query *ModelQuery) Project(fieldNames ...string) *ModelQuery {
	return &ModelQuery{query.Query.Project(fieldNames...), query.keysOnly}
}

// Return a ModelIterator instead of an Iterator.  ModelIterator does
// the same thing as ModelQuery, but for Iterator.
func (query *ModelQuery) Run(ctx appengine.Context) *ModelIterator {
	return &ModelIterator{query.Query.Run(ctx)}
}

func (query *ModelQuery) Start(cursor datastore.Cursor) *ModelQuery {
	return &ModelQuery{query.Query.Start(cursor), query.keysOnly}
}
