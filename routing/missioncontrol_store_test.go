package routing

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/decred/dcrlnd/lnwire"

	"github.com/decred/dcrlnd/routing/route"
	bolt "go.etcd.io/bbolt"
)

const testMaxRecords = 2

func TestMissionControlStore(t *testing.T) {
	// Set time zone explictly to keep test deterministic.
	time.Local = time.UTC

	file, err := ioutil.TempFile("", "*.db")
	if err != nil {
		t.Fatal(err)
	}

	dbPath := file.Name()

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(dbPath)

	store, err := newMissionControlStore(db, testMaxRecords)
	if err != nil {
		t.Fatal(err)
	}

	results, err := store.fetchAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 0 {
		t.Fatal("expected no results")
	}

	testRoute := route.Route{
		SourcePubKey: route.Vertex{1},
		Hops: []*route.Hop{
			{
				PubKeyBytes:   route.Vertex{2},
				LegacyPayload: true,
			},
		},
	}

	failureSourceIdx := 1

	result1 := paymentResult{
		route:            &testRoute,
		failure:          lnwire.NewFailIncorrectDetails(100, 1000),
		failureSourceIdx: &failureSourceIdx,
		id:               99,
		timeReply:        testTime,
		timeFwd:          testTime.Add(-time.Minute),
	}

	result2 := result1
	result2.timeReply = result1.timeReply.Add(time.Hour)
	result2.timeFwd = result1.timeReply.Add(time.Hour)
	result2.id = 2

	// Store result.
	err = store.AddResult(&result2)
	if err != nil {
		t.Fatal(err)
	}

	// Store again to test idempotency.
	err = store.AddResult(&result2)
	if err != nil {
		t.Fatal(err)
	}

	// Store second result which has an earlier timestamp.
	err = store.AddResult(&result1)
	if err != nil {
		t.Fatal(err)
	}

	results, err = store.fetchAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatal("expected two results")
	}

	// Check that results are stored in chronological order.
	if !reflect.DeepEqual(&result1, results[0]) {
		t.Fatalf("the results differ: %v vs %v", spew.Sdump(&result1),
			spew.Sdump(results[0]))
	}
	if !reflect.DeepEqual(&result2, results[1]) {
		t.Fatalf("the results differ: %v vs %v", spew.Sdump(&result2),
			spew.Sdump(results[1]))
	}

	// Recreate store to test pruning.
	store, err = newMissionControlStore(db, testMaxRecords)
	if err != nil {
		t.Fatal(err)
	}

	// Add a newer result.
	result3 := result1
	result3.timeReply = result1.timeReply.Add(2 * time.Hour)
	result3.timeFwd = result1.timeReply.Add(2 * time.Hour)
	result3.id = 3

	err = store.AddResult(&result3)
	if err != nil {
		t.Fatal(err)
	}

	// Check that results are pruned.
	results, err = store.fetchAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatal("expected two results")
	}

	if !reflect.DeepEqual(&result2, results[0]) {
		t.Fatalf("the results differ: %v vs %v", spew.Sdump(&result2),
			spew.Sdump(results[0]))
	}
	if !reflect.DeepEqual(&result3, results[1]) {
		t.Fatalf("the results differ: %v vs %v", spew.Sdump(&result3),
			spew.Sdump(results[1]))
	}
}
