package anything

import (
	"reflect"
	"testing"

	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var (
	registry = make(map[string]reflect.Type)
)

type MYTYPE map[string]interface{}

var myany = &MYTYPE{
	"a": 1,
	"b": "aa",
}

func init() {
	registry["myany"] = reflect.TypeOf(myany)
}

func UnmarshalAny(any *any.Any) (interface{}, error) {
	class := any.TypeUrl
	bytes := any.Value
	reflect.TypeOf(myany).Name()
	instance := reflect.New(registry[class]).Interface()
	err := proto.Unmarshal(bytes, instance.(proto.Message))
	if err != nil {
		return nil, err
	}
	fmt.Printf("instance: %v", instance)

	return instance, nil
}

func TestType(t *testing.T) {
	a := reflect.TypeOf(myany).Name()
	fmt.Println(a)
}

func TestAnything(t *testing.T) {

	t1 := &timestamp.Timestamp{
		Seconds: 5, // easy to verify
		Nanos:   6, // easy to verify
	}

	serialized, err := proto.Marshal(t1)
	if err != nil {
		t.Fatal("could not serialize timestamp")
	}

	// Blue was a great album by 3EB, before Cadgogan got kicked out
	// and Jenkins went full primadonna
	a := AnythingForYou{
		Anything: &any.Any{
			TypeUrl: "anteoy.site/" + proto.MessageName(t1),
			Value:   serialized,
		},
	}

	// marshal to simulate going on the wire:
	serializedA, err := proto.Marshal(&a)
	if err != nil {
		t.Fatal("could not serialize anything")
	}

	// unmarshal to simulate coming off the wire
	var a2 AnythingForYou
	if err := proto.Unmarshal(serializedA, &a2); err != nil {
		t.Fatal("could not deserialize anything")
	}

	// unmarshal the timestamp
	var t2 timestamp.Timestamp
	if err := ptypes.UnmarshalAny(a2.Anything, &t2); err != nil {
		t.Fatalf("Could not unmarshal timestamp from anything field: %s", err)
	}

	// Verify the values are as expected
	if !reflect.DeepEqual(t1, &t2) {
		t.Fatalf("Values don't match up:\n %+v \n %+v", t1, t2)
	}
}
