package snowflake

import (
	"sync/atomic"
	"time"
)

var (
	Epoch = time.UnixMilli(1700000000000)
)

const (
	MachineID       = 1
	SequenceMask    = 0xFFF
	MachineMask     = 0x3FF
	MachineShift    = 12
	Time41Shift     = 10 + MachineShift
	JSNumberMaxMask = 1<<53 - 1
)

var (
	Snowflake = NewSnowflakeID()
)

func ID() int64 {
	return Snowflake.ID()
}

type SnowflakeID struct {
	seq atomic.Uint32
}

func NewSnowflakeID() *SnowflakeID {
	return &SnowflakeID{}
}

func (s *SnowflakeID) ID() int64 {
	seq12 := int64(s.seq.Add(1) & SequenceMask)
	machineid10 := int64(MachineID & MachineMask)
	time41 := time.Since(Epoch).Milliseconds()
	return (time41<<Time41Shift | machineid10<<MachineShift | seq12) & JSNumberMaxMask
}