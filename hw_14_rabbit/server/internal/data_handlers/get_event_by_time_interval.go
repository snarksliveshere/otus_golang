package data_handlers

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func CheckGetEventByTimeIntervalFromProtoTimestamp(from, till *timestamp.Timestamp) (time.Time, time.Time, error) {
	fromT, err := ptypes.Timestamp(from)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	tillT, err := ptypes.Timestamp(till)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return fromT, tillT, nil
}
