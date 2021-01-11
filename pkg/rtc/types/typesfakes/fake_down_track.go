// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

type FakeDownTrack struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	CreateSourceDescriptionChunksStub        func() []rtcp.SourceDescriptionChunk
	createSourceDescriptionChunksMutex       sync.RWMutex
	createSourceDescriptionChunksArgsForCall []struct {
	}
	createSourceDescriptionChunksReturns struct {
		result1 []rtcp.SourceDescriptionChunk
	}
	createSourceDescriptionChunksReturnsOnCall map[int]struct {
		result1 []rtcp.SourceDescriptionChunk
	}
	GetNACKSeqNoStub        func([]uint16) []uint16
	getNACKSeqNoMutex       sync.RWMutex
	getNACKSeqNoArgsForCall []struct {
		arg1 []uint16
	}
	getNACKSeqNoReturns struct {
		result1 []uint16
	}
	getNACKSeqNoReturnsOnCall map[int]struct {
		result1 []uint16
	}
	IsBoundStub        func() bool
	isBoundMutex       sync.RWMutex
	isBoundArgsForCall []struct {
	}
	isBoundReturns struct {
		result1 bool
	}
	isBoundReturnsOnCall map[int]struct {
		result1 bool
	}
	LastSSRCStub        func() uint32
	lastSSRCMutex       sync.RWMutex
	lastSSRCArgsForCall []struct {
	}
	lastSSRCReturns struct {
		result1 uint32
	}
	lastSSRCReturnsOnCall map[int]struct {
		result1 uint32
	}
	OnBindStub        func(func())
	onBindMutex       sync.RWMutex
	onBindArgsForCall []struct {
		arg1 func()
	}
	OnCloseHandlerStub        func(func())
	onCloseHandlerMutex       sync.RWMutex
	onCloseHandlerArgsForCall []struct {
		arg1 func()
	}
	SSRCStub        func() uint32
	sSRCMutex       sync.RWMutex
	sSRCArgsForCall []struct {
	}
	sSRCReturns struct {
		result1 uint32
	}
	sSRCReturnsOnCall map[int]struct {
		result1 uint32
	}
	SnOffsetStub        func() uint16
	snOffsetMutex       sync.RWMutex
	snOffsetArgsForCall []struct {
	}
	snOffsetReturns struct {
		result1 uint16
	}
	snOffsetReturnsOnCall map[int]struct {
		result1 uint16
	}
	TsOffsetStub        func() uint32
	tsOffsetMutex       sync.RWMutex
	tsOffsetArgsForCall []struct {
	}
	tsOffsetReturns struct {
		result1 uint32
	}
	tsOffsetReturnsOnCall map[int]struct {
		result1 uint32
	}
	WriteRTPStub        func(rtp.Packet) error
	writeRTPMutex       sync.RWMutex
	writeRTPArgsForCall []struct {
		arg1 rtp.Packet
	}
	writeRTPReturns struct {
		result1 error
	}
	writeRTPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownTrack) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeDownTrack) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeDownTrack) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeDownTrack) CreateSourceDescriptionChunks() []rtcp.SourceDescriptionChunk {
	fake.createSourceDescriptionChunksMutex.Lock()
	ret, specificReturn := fake.createSourceDescriptionChunksReturnsOnCall[len(fake.createSourceDescriptionChunksArgsForCall)]
	fake.createSourceDescriptionChunksArgsForCall = append(fake.createSourceDescriptionChunksArgsForCall, struct {
	}{})
	stub := fake.CreateSourceDescriptionChunksStub
	fakeReturns := fake.createSourceDescriptionChunksReturns
	fake.recordInvocation("CreateSourceDescriptionChunks", []interface{}{})
	fake.createSourceDescriptionChunksMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) CreateSourceDescriptionChunksCallCount() int {
	fake.createSourceDescriptionChunksMutex.RLock()
	defer fake.createSourceDescriptionChunksMutex.RUnlock()
	return len(fake.createSourceDescriptionChunksArgsForCall)
}

func (fake *FakeDownTrack) CreateSourceDescriptionChunksCalls(stub func() []rtcp.SourceDescriptionChunk) {
	fake.createSourceDescriptionChunksMutex.Lock()
	defer fake.createSourceDescriptionChunksMutex.Unlock()
	fake.CreateSourceDescriptionChunksStub = stub
}

func (fake *FakeDownTrack) CreateSourceDescriptionChunksReturns(result1 []rtcp.SourceDescriptionChunk) {
	fake.createSourceDescriptionChunksMutex.Lock()
	defer fake.createSourceDescriptionChunksMutex.Unlock()
	fake.CreateSourceDescriptionChunksStub = nil
	fake.createSourceDescriptionChunksReturns = struct {
		result1 []rtcp.SourceDescriptionChunk
	}{result1}
}

func (fake *FakeDownTrack) CreateSourceDescriptionChunksReturnsOnCall(i int, result1 []rtcp.SourceDescriptionChunk) {
	fake.createSourceDescriptionChunksMutex.Lock()
	defer fake.createSourceDescriptionChunksMutex.Unlock()
	fake.CreateSourceDescriptionChunksStub = nil
	if fake.createSourceDescriptionChunksReturnsOnCall == nil {
		fake.createSourceDescriptionChunksReturnsOnCall = make(map[int]struct {
			result1 []rtcp.SourceDescriptionChunk
		})
	}
	fake.createSourceDescriptionChunksReturnsOnCall[i] = struct {
		result1 []rtcp.SourceDescriptionChunk
	}{result1}
}

func (fake *FakeDownTrack) GetNACKSeqNo(arg1 []uint16) []uint16 {
	var arg1Copy []uint16
	if arg1 != nil {
		arg1Copy = make([]uint16, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getNACKSeqNoMutex.Lock()
	ret, specificReturn := fake.getNACKSeqNoReturnsOnCall[len(fake.getNACKSeqNoArgsForCall)]
	fake.getNACKSeqNoArgsForCall = append(fake.getNACKSeqNoArgsForCall, struct {
		arg1 []uint16
	}{arg1Copy})
	stub := fake.GetNACKSeqNoStub
	fakeReturns := fake.getNACKSeqNoReturns
	fake.recordInvocation("GetNACKSeqNo", []interface{}{arg1Copy})
	fake.getNACKSeqNoMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) GetNACKSeqNoCallCount() int {
	fake.getNACKSeqNoMutex.RLock()
	defer fake.getNACKSeqNoMutex.RUnlock()
	return len(fake.getNACKSeqNoArgsForCall)
}

func (fake *FakeDownTrack) GetNACKSeqNoCalls(stub func([]uint16) []uint16) {
	fake.getNACKSeqNoMutex.Lock()
	defer fake.getNACKSeqNoMutex.Unlock()
	fake.GetNACKSeqNoStub = stub
}

func (fake *FakeDownTrack) GetNACKSeqNoArgsForCall(i int) []uint16 {
	fake.getNACKSeqNoMutex.RLock()
	defer fake.getNACKSeqNoMutex.RUnlock()
	argsForCall := fake.getNACKSeqNoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownTrack) GetNACKSeqNoReturns(result1 []uint16) {
	fake.getNACKSeqNoMutex.Lock()
	defer fake.getNACKSeqNoMutex.Unlock()
	fake.GetNACKSeqNoStub = nil
	fake.getNACKSeqNoReturns = struct {
		result1 []uint16
	}{result1}
}

func (fake *FakeDownTrack) GetNACKSeqNoReturnsOnCall(i int, result1 []uint16) {
	fake.getNACKSeqNoMutex.Lock()
	defer fake.getNACKSeqNoMutex.Unlock()
	fake.GetNACKSeqNoStub = nil
	if fake.getNACKSeqNoReturnsOnCall == nil {
		fake.getNACKSeqNoReturnsOnCall = make(map[int]struct {
			result1 []uint16
		})
	}
	fake.getNACKSeqNoReturnsOnCall[i] = struct {
		result1 []uint16
	}{result1}
}

func (fake *FakeDownTrack) IsBound() bool {
	fake.isBoundMutex.Lock()
	ret, specificReturn := fake.isBoundReturnsOnCall[len(fake.isBoundArgsForCall)]
	fake.isBoundArgsForCall = append(fake.isBoundArgsForCall, struct {
	}{})
	stub := fake.IsBoundStub
	fakeReturns := fake.isBoundReturns
	fake.recordInvocation("IsBound", []interface{}{})
	fake.isBoundMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) IsBoundCallCount() int {
	fake.isBoundMutex.RLock()
	defer fake.isBoundMutex.RUnlock()
	return len(fake.isBoundArgsForCall)
}

func (fake *FakeDownTrack) IsBoundCalls(stub func() bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = stub
}

func (fake *FakeDownTrack) IsBoundReturns(result1 bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = nil
	fake.isBoundReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDownTrack) IsBoundReturnsOnCall(i int, result1 bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = nil
	if fake.isBoundReturnsOnCall == nil {
		fake.isBoundReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isBoundReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDownTrack) LastSSRC() uint32 {
	fake.lastSSRCMutex.Lock()
	ret, specificReturn := fake.lastSSRCReturnsOnCall[len(fake.lastSSRCArgsForCall)]
	fake.lastSSRCArgsForCall = append(fake.lastSSRCArgsForCall, struct {
	}{})
	stub := fake.LastSSRCStub
	fakeReturns := fake.lastSSRCReturns
	fake.recordInvocation("LastSSRC", []interface{}{})
	fake.lastSSRCMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) LastSSRCCallCount() int {
	fake.lastSSRCMutex.RLock()
	defer fake.lastSSRCMutex.RUnlock()
	return len(fake.lastSSRCArgsForCall)
}

func (fake *FakeDownTrack) LastSSRCCalls(stub func() uint32) {
	fake.lastSSRCMutex.Lock()
	defer fake.lastSSRCMutex.Unlock()
	fake.LastSSRCStub = stub
}

func (fake *FakeDownTrack) LastSSRCReturns(result1 uint32) {
	fake.lastSSRCMutex.Lock()
	defer fake.lastSSRCMutex.Unlock()
	fake.LastSSRCStub = nil
	fake.lastSSRCReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) LastSSRCReturnsOnCall(i int, result1 uint32) {
	fake.lastSSRCMutex.Lock()
	defer fake.lastSSRCMutex.Unlock()
	fake.LastSSRCStub = nil
	if fake.lastSSRCReturnsOnCall == nil {
		fake.lastSSRCReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.lastSSRCReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) OnBind(arg1 func()) {
	fake.onBindMutex.Lock()
	fake.onBindArgsForCall = append(fake.onBindArgsForCall, struct {
		arg1 func()
	}{arg1})
	stub := fake.OnBindStub
	fake.recordInvocation("OnBind", []interface{}{arg1})
	fake.onBindMutex.Unlock()
	if stub != nil {
		fake.OnBindStub(arg1)
	}
}

func (fake *FakeDownTrack) OnBindCallCount() int {
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	return len(fake.onBindArgsForCall)
}

func (fake *FakeDownTrack) OnBindCalls(stub func(func())) {
	fake.onBindMutex.Lock()
	defer fake.onBindMutex.Unlock()
	fake.OnBindStub = stub
}

func (fake *FakeDownTrack) OnBindArgsForCall(i int) func() {
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	argsForCall := fake.onBindArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownTrack) OnCloseHandler(arg1 func()) {
	fake.onCloseHandlerMutex.Lock()
	fake.onCloseHandlerArgsForCall = append(fake.onCloseHandlerArgsForCall, struct {
		arg1 func()
	}{arg1})
	stub := fake.OnCloseHandlerStub
	fake.recordInvocation("OnCloseHandler", []interface{}{arg1})
	fake.onCloseHandlerMutex.Unlock()
	if stub != nil {
		fake.OnCloseHandlerStub(arg1)
	}
}

func (fake *FakeDownTrack) OnCloseHandlerCallCount() int {
	fake.onCloseHandlerMutex.RLock()
	defer fake.onCloseHandlerMutex.RUnlock()
	return len(fake.onCloseHandlerArgsForCall)
}

func (fake *FakeDownTrack) OnCloseHandlerCalls(stub func(func())) {
	fake.onCloseHandlerMutex.Lock()
	defer fake.onCloseHandlerMutex.Unlock()
	fake.OnCloseHandlerStub = stub
}

func (fake *FakeDownTrack) OnCloseHandlerArgsForCall(i int) func() {
	fake.onCloseHandlerMutex.RLock()
	defer fake.onCloseHandlerMutex.RUnlock()
	argsForCall := fake.onCloseHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownTrack) SSRC() uint32 {
	fake.sSRCMutex.Lock()
	ret, specificReturn := fake.sSRCReturnsOnCall[len(fake.sSRCArgsForCall)]
	fake.sSRCArgsForCall = append(fake.sSRCArgsForCall, struct {
	}{})
	stub := fake.SSRCStub
	fakeReturns := fake.sSRCReturns
	fake.recordInvocation("SSRC", []interface{}{})
	fake.sSRCMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) SSRCCallCount() int {
	fake.sSRCMutex.RLock()
	defer fake.sSRCMutex.RUnlock()
	return len(fake.sSRCArgsForCall)
}

func (fake *FakeDownTrack) SSRCCalls(stub func() uint32) {
	fake.sSRCMutex.Lock()
	defer fake.sSRCMutex.Unlock()
	fake.SSRCStub = stub
}

func (fake *FakeDownTrack) SSRCReturns(result1 uint32) {
	fake.sSRCMutex.Lock()
	defer fake.sSRCMutex.Unlock()
	fake.SSRCStub = nil
	fake.sSRCReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) SSRCReturnsOnCall(i int, result1 uint32) {
	fake.sSRCMutex.Lock()
	defer fake.sSRCMutex.Unlock()
	fake.SSRCStub = nil
	if fake.sSRCReturnsOnCall == nil {
		fake.sSRCReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.sSRCReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) SnOffset() uint16 {
	fake.snOffsetMutex.Lock()
	ret, specificReturn := fake.snOffsetReturnsOnCall[len(fake.snOffsetArgsForCall)]
	fake.snOffsetArgsForCall = append(fake.snOffsetArgsForCall, struct {
	}{})
	stub := fake.SnOffsetStub
	fakeReturns := fake.snOffsetReturns
	fake.recordInvocation("SnOffset", []interface{}{})
	fake.snOffsetMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) SnOffsetCallCount() int {
	fake.snOffsetMutex.RLock()
	defer fake.snOffsetMutex.RUnlock()
	return len(fake.snOffsetArgsForCall)
}

func (fake *FakeDownTrack) SnOffsetCalls(stub func() uint16) {
	fake.snOffsetMutex.Lock()
	defer fake.snOffsetMutex.Unlock()
	fake.SnOffsetStub = stub
}

func (fake *FakeDownTrack) SnOffsetReturns(result1 uint16) {
	fake.snOffsetMutex.Lock()
	defer fake.snOffsetMutex.Unlock()
	fake.SnOffsetStub = nil
	fake.snOffsetReturns = struct {
		result1 uint16
	}{result1}
}

func (fake *FakeDownTrack) SnOffsetReturnsOnCall(i int, result1 uint16) {
	fake.snOffsetMutex.Lock()
	defer fake.snOffsetMutex.Unlock()
	fake.SnOffsetStub = nil
	if fake.snOffsetReturnsOnCall == nil {
		fake.snOffsetReturnsOnCall = make(map[int]struct {
			result1 uint16
		})
	}
	fake.snOffsetReturnsOnCall[i] = struct {
		result1 uint16
	}{result1}
}

func (fake *FakeDownTrack) TsOffset() uint32 {
	fake.tsOffsetMutex.Lock()
	ret, specificReturn := fake.tsOffsetReturnsOnCall[len(fake.tsOffsetArgsForCall)]
	fake.tsOffsetArgsForCall = append(fake.tsOffsetArgsForCall, struct {
	}{})
	stub := fake.TsOffsetStub
	fakeReturns := fake.tsOffsetReturns
	fake.recordInvocation("TsOffset", []interface{}{})
	fake.tsOffsetMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) TsOffsetCallCount() int {
	fake.tsOffsetMutex.RLock()
	defer fake.tsOffsetMutex.RUnlock()
	return len(fake.tsOffsetArgsForCall)
}

func (fake *FakeDownTrack) TsOffsetCalls(stub func() uint32) {
	fake.tsOffsetMutex.Lock()
	defer fake.tsOffsetMutex.Unlock()
	fake.TsOffsetStub = stub
}

func (fake *FakeDownTrack) TsOffsetReturns(result1 uint32) {
	fake.tsOffsetMutex.Lock()
	defer fake.tsOffsetMutex.Unlock()
	fake.TsOffsetStub = nil
	fake.tsOffsetReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) TsOffsetReturnsOnCall(i int, result1 uint32) {
	fake.tsOffsetMutex.Lock()
	defer fake.tsOffsetMutex.Unlock()
	fake.TsOffsetStub = nil
	if fake.tsOffsetReturnsOnCall == nil {
		fake.tsOffsetReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.tsOffsetReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeDownTrack) WriteRTP(arg1 rtp.Packet) error {
	fake.writeRTPMutex.Lock()
	ret, specificReturn := fake.writeRTPReturnsOnCall[len(fake.writeRTPArgsForCall)]
	fake.writeRTPArgsForCall = append(fake.writeRTPArgsForCall, struct {
		arg1 rtp.Packet
	}{arg1})
	stub := fake.WriteRTPStub
	fakeReturns := fake.writeRTPReturns
	fake.recordInvocation("WriteRTP", []interface{}{arg1})
	fake.writeRTPMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDownTrack) WriteRTPCallCount() int {
	fake.writeRTPMutex.RLock()
	defer fake.writeRTPMutex.RUnlock()
	return len(fake.writeRTPArgsForCall)
}

func (fake *FakeDownTrack) WriteRTPCalls(stub func(rtp.Packet) error) {
	fake.writeRTPMutex.Lock()
	defer fake.writeRTPMutex.Unlock()
	fake.WriteRTPStub = stub
}

func (fake *FakeDownTrack) WriteRTPArgsForCall(i int) rtp.Packet {
	fake.writeRTPMutex.RLock()
	defer fake.writeRTPMutex.RUnlock()
	argsForCall := fake.writeRTPArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownTrack) WriteRTPReturns(result1 error) {
	fake.writeRTPMutex.Lock()
	defer fake.writeRTPMutex.Unlock()
	fake.WriteRTPStub = nil
	fake.writeRTPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownTrack) WriteRTPReturnsOnCall(i int, result1 error) {
	fake.writeRTPMutex.Lock()
	defer fake.writeRTPMutex.Unlock()
	fake.WriteRTPStub = nil
	if fake.writeRTPReturnsOnCall == nil {
		fake.writeRTPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeRTPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createSourceDescriptionChunksMutex.RLock()
	defer fake.createSourceDescriptionChunksMutex.RUnlock()
	fake.getNACKSeqNoMutex.RLock()
	defer fake.getNACKSeqNoMutex.RUnlock()
	fake.isBoundMutex.RLock()
	defer fake.isBoundMutex.RUnlock()
	fake.lastSSRCMutex.RLock()
	defer fake.lastSSRCMutex.RUnlock()
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	fake.onCloseHandlerMutex.RLock()
	defer fake.onCloseHandlerMutex.RUnlock()
	fake.sSRCMutex.RLock()
	defer fake.sSRCMutex.RUnlock()
	fake.snOffsetMutex.RLock()
	defer fake.snOffsetMutex.RUnlock()
	fake.tsOffsetMutex.RLock()
	defer fake.tsOffsetMutex.RUnlock()
	fake.writeRTPMutex.RLock()
	defer fake.writeRTPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownTrack) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.DownTrack = new(FakeDownTrack)
