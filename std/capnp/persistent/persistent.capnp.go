// Code generated by capnpc-go. DO NOT EDIT.

package persistent

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

const Persistent_ = uint64(0xf622595091cafb67)

type Persistent capnp.Client

// Persistent_TypeID is the unique identifier for the type Persistent.
const Persistent_TypeID = 0xc8cb212fcd9f5691

func (c Persistent) Save(ctx context.Context, params func(Persistent_SaveParams) error) (Persistent_SaveResults_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "persistent.capnp:Persistent",
			MethodName:    "save",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Persistent_SaveParams(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Persistent_SaveResults_Future{Future: ans.Future()}, release

}

func (c Persistent) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Persistent) String() string {
	return "Persistent(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Persistent) AddRef() Persistent {
	return Persistent(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Persistent) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Persistent) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Persistent) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Persistent) DecodeFromPtr(p capnp.Ptr) Persistent {
	return Persistent(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Persistent) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Persistent) IsSame(other Persistent) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Persistent) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Persistent) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A Persistent_Server is a Persistent with a local implementation.
type Persistent_Server interface {
	Save(context.Context, Persistent_save) error
}

// Persistent_NewServer creates a new Server from an implementation of Persistent_Server.
func Persistent_NewServer(s Persistent_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Persistent_Methods(nil, s), s, c)
}

// Persistent_ServerToClient creates a new Client from an implementation of Persistent_Server.
// The caller is responsible for calling Release on the returned Client.
func Persistent_ServerToClient(s Persistent_Server) Persistent {
	return Persistent(capnp.NewClient(Persistent_NewServer(s)))
}

// Persistent_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Persistent_Methods(methods []server.Method, s Persistent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Save(ctx, Persistent_save{call})
		},
	})

	return methods
}

// Persistent_save holds the state for a server call to Persistent.save.
// See server.Call for documentation.
type Persistent_save struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Persistent_save) Args() Persistent_SaveParams {
	return Persistent_SaveParams(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Persistent_save) AllocResults() (Persistent_SaveResults, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Persistent_SaveResults(r), err
}

// Persistent_List is a list of Persistent.
type Persistent_List = capnp.CapList[Persistent]

// NewPersistent creates a new list of Persistent.
func NewPersistent_List(s *capnp.Segment, sz int32) (Persistent_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Persistent](l), err
}

type Persistent_SaveParams capnp.Struct

// Persistent_SaveParams_TypeID is the unique identifier for the type Persistent_SaveParams.
const Persistent_SaveParams_TypeID = 0xf76fba59183073a5

func NewPersistent_SaveParams(s *capnp.Segment) (Persistent_SaveParams, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Persistent_SaveParams(st), err
}

func NewRootPersistent_SaveParams(s *capnp.Segment) (Persistent_SaveParams, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Persistent_SaveParams(st), err
}

func ReadRootPersistent_SaveParams(msg *capnp.Message) (Persistent_SaveParams, error) {
	root, err := msg.Root()
	return Persistent_SaveParams(root.Struct()), err
}

func (s Persistent_SaveParams) String() string {
	str, _ := text.Marshal(0xf76fba59183073a5, capnp.Struct(s))
	return str
}

func (s Persistent_SaveParams) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Persistent_SaveParams) DecodeFromPtr(p capnp.Ptr) Persistent_SaveParams {
	return Persistent_SaveParams(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Persistent_SaveParams) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Persistent_SaveParams) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Persistent_SaveParams) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Persistent_SaveParams) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Persistent_SaveParams) SealFor() (capnp.Ptr, error) {
	return capnp.Struct(s).Ptr(0)
}

func (s Persistent_SaveParams) HasSealFor() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Persistent_SaveParams) SetSealFor(v capnp.Ptr) error {
	return capnp.Struct(s).SetPtr(0, v)
}

// Persistent_SaveParams_List is a list of Persistent_SaveParams.
type Persistent_SaveParams_List = capnp.StructList[Persistent_SaveParams]

// NewPersistent_SaveParams creates a new list of Persistent_SaveParams.
func NewPersistent_SaveParams_List(s *capnp.Segment, sz int32) (Persistent_SaveParams_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Persistent_SaveParams](l), err
}

// Persistent_SaveParams_Future is a wrapper for a Persistent_SaveParams promised by a client call.
type Persistent_SaveParams_Future struct{ *capnp.Future }

func (f Persistent_SaveParams_Future) Struct() (Persistent_SaveParams, error) {
	p, err := f.Future.Ptr()
	return Persistent_SaveParams(p.Struct()), err
}
func (p Persistent_SaveParams_Future) SealFor() *capnp.Future {
	return p.Future.Field(0, nil)
}

type Persistent_SaveResults capnp.Struct

// Persistent_SaveResults_TypeID is the unique identifier for the type Persistent_SaveResults.
const Persistent_SaveResults_TypeID = 0xb76848c18c40efbf

func NewPersistent_SaveResults(s *capnp.Segment) (Persistent_SaveResults, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Persistent_SaveResults(st), err
}

func NewRootPersistent_SaveResults(s *capnp.Segment) (Persistent_SaveResults, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Persistent_SaveResults(st), err
}

func ReadRootPersistent_SaveResults(msg *capnp.Message) (Persistent_SaveResults, error) {
	root, err := msg.Root()
	return Persistent_SaveResults(root.Struct()), err
}

func (s Persistent_SaveResults) String() string {
	str, _ := text.Marshal(0xb76848c18c40efbf, capnp.Struct(s))
	return str
}

func (s Persistent_SaveResults) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Persistent_SaveResults) DecodeFromPtr(p capnp.Ptr) Persistent_SaveResults {
	return Persistent_SaveResults(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Persistent_SaveResults) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Persistent_SaveResults) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Persistent_SaveResults) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Persistent_SaveResults) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Persistent_SaveResults) SturdyRef() (capnp.Ptr, error) {
	return capnp.Struct(s).Ptr(0)
}

func (s Persistent_SaveResults) HasSturdyRef() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Persistent_SaveResults) SetSturdyRef(v capnp.Ptr) error {
	return capnp.Struct(s).SetPtr(0, v)
}

// Persistent_SaveResults_List is a list of Persistent_SaveResults.
type Persistent_SaveResults_List = capnp.StructList[Persistent_SaveResults]

// NewPersistent_SaveResults creates a new list of Persistent_SaveResults.
func NewPersistent_SaveResults_List(s *capnp.Segment, sz int32) (Persistent_SaveResults_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Persistent_SaveResults](l), err
}

// Persistent_SaveResults_Future is a wrapper for a Persistent_SaveResults promised by a client call.
type Persistent_SaveResults_Future struct{ *capnp.Future }

func (f Persistent_SaveResults_Future) Struct() (Persistent_SaveResults, error) {
	p, err := f.Future.Ptr()
	return Persistent_SaveResults(p.Struct()), err
}
func (p Persistent_SaveResults_Future) SturdyRef() *capnp.Future {
	return p.Future.Field(0, nil)
}

const schema_b8630836983feed7 = "x\xdat\x90\xbfk\x14A\x1c\xc5\xdf\x9b\xd9u\xef " +
	"\xc1\x9b\x1b\xc1\x14\x81\xa8\x8d\"\x18\x7f\x81E\x9a\xbb\x8d" +
	"\xa0^!\xee\xec\xa1\x10\xbb5\x8e?\xe0\xb2Yv6" +
	"\x09V\xfe\x036\xe9\xec\xc4\xc2\xd2\xc2J\xb1\x11;Q" +
	",\x82\x85\x95M\xfe\x00\xed\x14\xb1\x18Y0\x97\xe5\xbc" +
	"\xb43\xdf\xc7\xfb|^\xe7K?8?\xbb%!\xcc" +
	"\xa9\xf0\x90\x7f\xf7\xa3\xff\xe4\xfd\xb5\x07\xaf\xa1\xe6\xe9\xb7" +
	"o=\xfb|\xf6\xf8\xa7\x0f\x08\x19ux\xf19\x97\xa9" +
	"_1\x02\xf4K\xf6\xd0\xf8WJ\xfa\xaf\xdf{O/" +
	"\xb5V\xdf\x00\xe8P\xefpW\x7f\xe3I@\xff\xe6U" +
	"=\x10\x91\x1e\x88\xa3\xfe\xfe\x9f\x8f\xdb\xc9\xca\x89\x9f\xd8" +
	"Q\xe116\"\xd4\x03\xb1\xabo\x8a\x08\x18&B\x12" +
	"\xf4/\xdc\xb9\xb9\x95\xb7\xeb\xbf\xa6\x91\xc4b\x89\xda\xd4" +
	"\xd7\xfa\xba\xe8\xe1\xb2/l\xe9\x1e\xba\xca\x06y\xb5\xb8" +
	"\x9a\x15y\xb1\x94\xfc{\xc9\xab\xc5a\xb6iS\xeb6" +
	"F\x95CB\x9a@\x06@@@\xcd\xa6\x80\x99\x914" +
	"s\x82\xdeU\x1b\xe5\xddG\xa9\x05\xef\xb1\xcbF'\xc0" +
	".8\xee\x10\x93\x1d\x91\xcd+\xd3b\x13\xb9}\xbb\xb1" +
	"d\xfb\x8e\xaf\x11\x92\xac\xcc \xd7\x9c\xdf\xe3A4\xaa" +
	"\x9c\x09d\x08\x8c\xa3\xdc\x8b)u\x1a\x88g\x18\xcfS" +
	"\x9d\x89\x0e\xbbl\xd3*.\x98@4\xc0X;L{" +
	"\xec3!\xe3\x16U\x98\xaa\xf6\x05?\xdc7[\xb8\xb1" +
	"\x95\xdbr\x8aK\xb1\xef\x92\x90\x90\xe3\x13y\xd0\xa4\xbd" +
	"Zh\xcdM,\xba\x0c\x98\x96\xa49\"\xf8\xd8\xd9l" +
	"te\xbdd7\xfc\x7f\xce\xbf\x01\x00\x00\xff\xff\x92\xdc" +
	"\xc1$"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_b8630836983feed7,
		Nodes: []uint64{
			0xb76848c18c40efbf,
			0xc8cb212fcd9f5691,
			0xf622595091cafb67,
			0xf76fba59183073a5,
		},
		Compressed: true,
	})
}
