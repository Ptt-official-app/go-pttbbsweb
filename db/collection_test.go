package db

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCollection_CreateOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "2"

	type testUpdate struct {
		Test1 string `json:"test1"`
		Test3 bool   `json:"test3"`
	}
	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.UpsertedCount = 1

	expectedData1 := make(map[string]interface{})
	expectedData1["test"] = 1
	expectedData1["test1"] = "2"
	expectedData1["test3"] = true

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected2,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.CreateOnly(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.CreateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedID = gotR.UpsertedID
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.CreateOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_UpdateOneOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `json:"test1"`
		Test3 bool   `json:"test3"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 1
	filter0["test1"] = "4"

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 0

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := &testUpdate{Test1: "2", Test3: true}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 1
	expected2.UpsertedCount = 0

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
			wantErr:   true,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.UpdateOneOnly(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.UpdateOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_Update(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `json:"test1"`
		Test3 bool   `json:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "3", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 1

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := &testUpdate{Test1: "2", Test3: true}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 1
	expected2.UpsertedCount = 0

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.Update(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedID = gotR.UpsertedID
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.Update() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_Find(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	find1 := make(map[string]interface{})
	find1["test"] = 2

	project := make(map[string]bool)
	project["test"] = true
	project["test5"] = true
	project["test3"] = true

	var ret1 []*testFind1
	expected1 := []*testFind1{{Find1: 2, Find2: nil, Find4: true}}

	find2 := make(map[string]interface{})
	var ret2 []*testFind1
	expected2 := []*testFind1{{Find1: 0, Find2: nil, Find4: true}}

	type queryFind1 struct {
		Find1 int `bson:"test"`
	}

	find3 := &queryFind1{Find1: 2}
	var ret3 []*testFind1

	find4 := &queryFind1{Find1: 5}
	var ret4 []*testFind1

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter  interface{}
		ret     []*testFind1
		n       int64
		project map[string]bool
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedRet []*testFind1
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name:        "find specific with map[string]interface{}",
			fields:      fields{coll: coll},
			args:        args{filter: find1, ret: ret1, n: 4, project: project},
			expectedRet: expected1,
		},
		{
			name:        "find all, but limit 1",
			fields:      fields{coll: coll},
			args:        args{filter: find2, ret: ret2, n: 1, project: project},
			expectedRet: expected2,
		},
		{
			name:        "use struct for query",
			fields:      fields{coll: coll},
			args:        args{filter: find3, ret: ret3, n: 4, project: project},
			expectedRet: expected1,
		},
		{
			name:        "not found",
			fields:      fields{coll: coll},
			args:        args{filter: find4, ret: ret4, n: 4, project: project},
			expectedRet: nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			err := c.Find(tt.args.filter, tt.args.n, &tt.args.ret, tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(tt.args.ret) != len(tt.expectedRet) {
				t.Errorf("collection.Find: ret1: %v ret: %v expected: %v", len(ret1), len(tt.args.ret), len(tt.expectedRet))
			}
			log.Infof("collection.Find: ret1: %v ret: %v expected: %v", len(ret1), len(tt.args.ret), len(tt.expectedRet))

			for idx, each := range tt.args.ret {
				if idx >= len(tt.expectedRet) {
					t.Errorf("Collection.Find: (%v/%v): %v", idx, len(tt.args.ret), each)
				}
				types.TDeepEqual(t, tt.args.ret[idx], tt.expectedRet[idx])
			}
		})
	}
}

func TestCollection_Remove(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})

	filter2 := make(map[string]interface{})
	filter2["test"] = 2

	filter3 := make(map[string]interface{})
	filter3["test2"] = "12"

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		expect  int
	}{
		// TODO: Add test cases.
		{
			name:    "nil, should not remove",
			fields:  fields{coll},
			args:    args{},
			wantErr: true,
		},
		{
			name:    "empty, should not remove",
			fields:  fields{coll},
			args:    args{filter1},
			wantErr: true,
		},
		{
			name:   "remove matching 1",
			fields: fields{coll},
			args:   args{filter2},
			expect: 2,
		},
		{
			name:   "remove matching all",
			fields: fields{coll},
			args:   args{filter3},
			expect: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			err = c.Remove(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			var ret []*testFind1
			find := struct{}{}
			err := c.Find(find, 4, &ret, nil)
			if err != nil {
				t.Errorf("Collection.Remove(): unable find: e: %v", err)
				return
			}
			if len(ret) != tt.expect {
				t.Errorf("Collection.Remove(): len(ret): %v want: %v", len(ret), tt.expect)
			}
		})
	}
}

func TestCollection_FindOne(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	ret1 := &testFind1{}

	filter2 := make(map[string]interface{})
	filter2["test2"] = "12"
	ret2 := &testFind1{}

	filter3 := make(map[string]interface{})
	filter3["test"] = 5
	ret3 := &testFind1{}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter  interface{}
		ret     interface{}
		project map[string]bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "find test",
			fields: fields{coll},
			args:   args{filter: filter1, ret: ret1},
		},
		{
			name:   "find multiple",
			fields: fields{coll},
			args:   args{filter: filter2, ret: ret2},
		},
		{
			name:    "find none",
			fields:  fields{coll},
			args:    args{filter: filter3, ret: ret3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			if err := c.FindOne(tt.args.filter, tt.args.ret, tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("Collection.FindOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_Count(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1

	filter2 := make(map[string]interface{})
	filter2["test2"] = "12"

	filter3 := make(map[string]interface{})
	filter3["test2"] = "123"

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		n      int64
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedCount int64
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter1, n: 4},
			expectedCount: 1,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter2},
			expectedCount: 3,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter2, n: 1},
			expectedCount: 1,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter3, n: 4},
			expectedCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotCount, err := c.Count(tt.args.filter, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.expectedCount {
				t.Errorf("Collection.Count() = %v, want %v", gotCount, tt.expectedCount)
			}
		})
	}
}
