package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	coll *mongo.Collection
}

func NewCollection(name string, db *mongo.Database) *Collection {
	return &Collection{
		coll: db.Collection(name),
	}
}

// CreateOnly
//
// Mongo update-one with setOnInsert + upsert operation
func (c *Collection) CreateOnly(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)

	theUpdate := bson.M{
		"$setOnInsert": update,
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateOne(ctx, filter, theUpdate, opts)
	if err != nil {
		return nil, err
	}

	return r, err
}

// UpdateOneOnly
//
// Mongo update-one with set + no-upsert operation
func (c *Collection) UpdateOneOnly(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)

	theUpdate := bson.M{
		"$set": update,
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateOne(ctx, filter, theUpdate, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateOneOnlyNoSet
//
// Mongo update-one with no-upsert operation
func (c *Collection) UpdateOneOnlyNoSet(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)

	theUpdate := update

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateOne(ctx, filter, theUpdate, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// FindOneAndUpdate
//
// Mongo update-one with set + no-upsert operation
func (c *Collection) FindOneAndUpdate(filter interface{}, update interface{}, isNew bool) (r *mongo.SingleResult, err error) {
	opts := &options.FindOneAndUpdateOptions{}
	opts.SetUpsert(false)
	if isNew {
		opts.SetReturnDocument(options.After)
	}

	theUpdate := bson.M{
		"$set": update,
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r = c.coll.FindOneAndUpdate(ctx, filter, theUpdate, opts)
	err = r.Err()
	if err != nil {
		return nil, err
	}

	return r, nil
}

// FindOneAndUpdate
//
// Mongo update-one with set + no-upsert operation
func (c *Collection) FindOneAndUpdateNoSet(filter interface{}, update interface{}, isNew bool) (r *mongo.SingleResult, err error) {
	opts := &options.FindOneAndUpdateOptions{}
	opts.SetUpsert(false)
	if isNew {
		opts.SetReturnDocument(options.After)
	}

	theUpdate := update

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r = c.coll.FindOneAndUpdate(ctx, filter, theUpdate, opts)
	err = r.Err()
	if err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateManyOnly
//
// Mongo update-many with set + no-upsert operation
func (c *Collection) UpdateManyOnly(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)

	theUpdate := bson.M{
		"$set": update,
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, theUpdate, opts)
	if err != nil {
		return nil, err
	}

	time.Sleep(1 * time.Millisecond)

	return r, nil
}

// UpdateManyOnly
//
// Mongo update-many with set + no-upsert operation
func (c *Collection) UpdateManyOnlyNoSet(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	time.Sleep(1 * time.Millisecond)

	return r, nil
}

// Update
//
// Mongo update with set + upsert operation
func (c *Collection) Update(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)

	theUpdate := bson.M{
		"$set": update,
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, theUpdate, opts)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BulkCreateOnly
//
// Mongo update with setOnInsert + upsert operation
func (c *Collection) BulkCreateOnly(theList []*UpdatePair) (r *mongo.BulkWriteResult, err error) {
	theList_b := make([]mongo.WriteModel, len(theList))
	for idx, each := range theList {
		theUpdate := bson.M{
			"$setOnInsert": each.Update,
		}
		theList_b[idx] = mongo.NewUpdateOneModel().SetFilter(each.Filter).SetUpdate(theUpdate).SetUpsert(true)
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := options.BulkWrite().SetOrdered(false)

	r, err = c.coll.BulkWrite(ctx, theList_b, opts)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BulkUpdateOneOnly
//
// Mongo update with set + no-upsert operation
func (c *Collection) BulkUpdateOneOnly(theList []*UpdatePair) (r *mongo.BulkWriteResult, err error) {
	theList_b := make([]mongo.WriteModel, len(theList))
	for idx, each := range theList {
		theUpdate := bson.M{
			"$set": each.Update,
		}
		theList_b[idx] = mongo.NewUpdateOneModel().SetFilter(each.Filter).SetUpdate(theUpdate).SetUpsert(false)
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := options.BulkWrite().SetOrdered(false)

	r, err = c.coll.BulkWrite(ctx, theList_b, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// BulkUpdateOneOnlyNoSet
//
// Mongo update without set and no-upsert operation
// WARNING!!! Must ensure that the update part is with $set, $unset, $setOnInsert.
func (c *Collection) BulkUpdateOneOnlyNoSet(theList []*UpdatePair) (r *mongo.BulkWriteResult, err error) {
	theList_b := make([]mongo.WriteModel, len(theList))
	for idx, each := range theList {
		theList_b[idx] = mongo.NewUpdateOneModel().SetFilter(each.Filter).SetUpdate(each.Update).SetUpsert(false)
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := options.BulkWrite().SetOrdered(false)

	r, err = c.coll.BulkWrite(ctx, theList_b, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// BulkUpdate
//
// Mongo update with set + upsert operation
func (c *Collection) BulkUpdate(theList []*UpdatePair) (r *mongo.BulkWriteResult, err error) {
	theList_b := make([]mongo.WriteModel, len(theList))
	for idx, each := range theList {
		theUpdate := bson.M{
			"$set": each.Update,
		}
		theList_b[idx] = mongo.NewUpdateOneModel().SetFilter(each.Filter).SetUpdate(theUpdate).SetUpsert(true)
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := options.BulkWrite().SetOrdered(false)

	r, err = c.coll.BulkWrite(ctx, theList_b, opts)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Find
//
// Never return error with normal operations. need to check len for not-found.
//
// Params:
//
//	ret: return values, requiring passing with pointer (malloced in cur.All)
//	project: the empty struct of the return type.
//
// ex:
//
//	query := make(map[string]interface{})
//	query["test"] = 1
//
//	var ret []*Temp //!!! declare but initiate
//
//	Find(query, 4, &ret, &Temp{})
func (c *Collection) Find(filter interface{}, n int64, ret interface{}, project interface{}, sort interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := &options.FindOptions{}
	if n > 0 {
		opts.SetLimit(n)
	}

	if project != nil {
		opts.SetProjection(project)
	}

	if sort != nil {
		opts.SetSort(sort)
	}

	cur, err := c.coll.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, ret); err != nil {
		logrus.Warnf("Find: unable to find data: e: %v", err)
		return err
	}

	return nil
}

// Count
func (c *Collection) Count(filter interface{}, n int64) (count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := &options.CountOptions{}
	if n > 0 {
		opts.SetLimit(n)
	}

	return c.coll.CountDocuments(ctx, filter, opts)
}

// FindOne
//
// Params:
//
//	ret: return value.
//	project: the empty struct of the return type.
//
// ex:
//
//	query := make(map[string]interface{})
//	query["test"] = 1
//
//	ret := &Temp{}
//
//	FindOne(query, ret, nil)
//
// Err:
//
//	mongo.ErrNoDocuments if not found.
func (c *Collection) FindOne(filter interface{}, ret interface{}, project interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	opts := &options.FindOneOptions{}

	if project != nil {
		opts.SetProjection(project)
	}

	return c.coll.FindOne(ctx, filter, opts).Decode(ret)
}

// Remove
//
// Params:
//
//	filter: filter. Must pass with non-empty map.
func (c *Collection) Remove(filter bson.M) (err error) {
	if len(filter) == 0 {
		return ErrEmptyInRemove
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	_, err = c.coll.DeleteMany(ctx, filter)

	return err
}

func (c *Collection) Drop() (err error) {
	if !isTest {
		return ErrInvalidOp
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	_, err = c.coll.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Collection) CreateIndex(keys *bson.D, opts *options.IndexOptions) (err error) {
	iv := c.coll.Indexes()

	if opts == nil {
		opts = options.Index()
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	model := mongo.IndexModel{Keys: keys, Options: opts}

	_, err = iv.CreateOne(ctx, model)

	return err
}

func (c *Collection) CreateUniqueIndex(keys *bson.D) (err error) {
	opts := options.Index().SetUnique(true)
	return c.CreateIndex(keys, opts)
}

func (c *Collection) Aggregate(filter interface{}, group interface{}) (ret []bson.M, err error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$group", Value: group}},
	}

	opts := options.Aggregate().SetMaxTime(TIMEOUT_MILLI_TS * time.Millisecond)

	cursor, err := c.coll.Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
