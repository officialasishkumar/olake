package protocol

import (
	"github.com/piyushsingariya/shift/drivers/base"
	"github.com/piyushsingariya/shift/types"
)

type Connector interface {
	Setup(config any, base *base.Driver) error
	Spec() (any, error)
	Check() error

	Type() string
}

type Driver interface {
	Connector
	Discover() ([]*types.Stream, error)
	Read(stream Stream, channel chan<- types.Record) error
}

type Adapter interface {
	Connector
	Write(channel <-chan types.Record) error
	Create(streamName string) error
}

type Stream interface {
	Self() *types.ConfiguredStream
	Name() string
	Namespace() string
	JSONSchema() *types.Schema
	GetStream() *types.Stream
	GetSyncMode() types.SyncMode
	SupportedSyncModes() []types.SyncMode
	Cursor() string
	InitialState() any
	GetState() any
	SetState(value any)
	BatchSize() int64
	SetBatchSize(size int64)

	// WithSyncModes(modes []types.SyncMode) Stream
	// WithPrimaryKeys(keys []string) Stream
	// WithCursorFields(columns []string) Stream
	// WithJSONSchema(schema types.Schema) Stream
}
