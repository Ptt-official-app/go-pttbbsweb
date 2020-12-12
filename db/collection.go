package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
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

//CreateOnly
//
//Mongo update-one with setOnInsert + upsert operation
func (c *Collection) CreateOnly(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {

	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)

	theUpdate := make(map[string]interface{})
	theUpdate["$setOnInsert"] = update

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, theUpdate, opts)
	if r.UpsertedCount == 0 && r.ModifiedCount == 0 {
		return r, ErrNoUpdate
	}

	return r, err
}

//UpdateOneOnly
//
//Mongo update-one with set + no-upsert operation
func (c *Collection) UpdateOneOnly(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {

	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)

	theUpdate := make(map[string]interface{})
	theUpdate["$set"] = update

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, theUpdate, opts)
	if r.UpsertedCount == 0 && r.ModifiedCount == 0 {
		return r, ErrNoUpdate
	}

	return r, nil
}

//Update
//
//Mongo update with set + upsert operation
func (c *Collection) Update(filter interface{}, update interface{}) (r *mongo.UpdateResult, err error) {

	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)

	theUpdate := make(map[string]interface{})
	theUpdate["$set"] = update

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	r, err = c.coll.UpdateMany(ctx, filter, theUpdate, opts)
	if r.UpsertedCount == 0 && r.ModifiedCount == 0 {
		return r, ErrNoUpdate
	}
	return r, nil
}

//Find
//
//Never return error with normal operations. need to check len for not-found.
//
//Params:
//    ret: return values, requiring passing with pointer (malloced in cur.All)
//    project: the empty struct of the return type.
//
//ex:
//    query := make(map[string]interface{})
//    query["test"] = 1
//
//    var ret []*Temp //!!! declare but initiate
//
//    Find(query, 4, &ret, &Temp{})
func (c *Collection) Find(filter interface{}, n int64, ret interface{}, project map[string]bool) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	opts := &options.FindOptions{}
	opts.SetLimit(n)

	if project != nil {
		opts.SetProjection(project)
	}

	cur, err := c.coll.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, ret); err != nil {
		log.Warnf("Find: unable to find data: e: %v", err)
		return err
	}

	return nil
}

//Count
//
func (c *Collection) Count(filter interface{}, n int64) (count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	opts := &options.CountOptions{}
	if n > 0 {
		opts.SetLimit(n)
	}

	return c.coll.CountDocuments(ctx, filter, opts)
}

//FindOne
//
//Params:
//    ret: return value.
//    project: the empty struct of the return type.
//
//ex:
//    query := make(map[string]interface{})
//    query["test"] = 1
//
//    ret := &Temp{}
//
//    FindOne(query, ret, nil)
func (c *Collection) FindOne(filter interface{}, ret interface{}, project map[string]bool) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	opts := &options.FindOneOptions{}

	if project != nil {
		opts.SetProjection(project)
	}

	return c.coll.FindOne(ctx, filter, opts).Decode(ret)
}

//Remove
//
//Params:
//    filter: filter. Must pass with non-empty map.
func (c *Collection) Remove(filter map[string]interface{}) (err error) {
	if filter == nil {
		return ErrEmptyInRemove
	}

	if len(filter) == 0 {
		return ErrEmptyInRemove
	}

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if ctxErr != nil {
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
		if ctxErr != nil {
			err = ctxErr
		}
	}()

	err = c.coll.Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}
