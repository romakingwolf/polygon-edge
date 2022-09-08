package samuel

import (
	"github.com/0xPolygon/polygon-edge/rootchain"
	"github.com/0xPolygon/polygon-edge/rootchain/proto"
)

type startDelegate func(uint64) error
type stopDelegate func() error
type subscribeDelegate func() <-chan rootchain.Event

type mockEventTracker struct {
	startFn     startDelegate
	stopFn      stopDelegate
	subscribeFn subscribeDelegate
}

func (m *mockEventTracker) Start(startBlock uint64) error {
	if m.startFn != nil {
		return m.startFn(startBlock)
	}

	return nil
}

func (m *mockEventTracker) Stop() error {
	if m.stopFn != nil {
		return m.stopFn()
	}

	return nil
}

func (m *mockEventTracker) Subscribe() <-chan rootchain.Event {
	if m.subscribeFn != nil {
		return m.subscribeFn()
	}

	return nil
}

type addMessageDelegate func(rootchain.SAM) error
type pruneDelegate func(uint64)
type peekDelegate func() rootchain.VerifiedSAM
type popDelegate func() rootchain.VerifiedSAM

type mockSAMP struct {
	addMessageFn addMessageDelegate
	pruneFn      pruneDelegate
	peekFn       peekDelegate
	popFn        popDelegate
}

func (m *mockSAMP) AddMessage(sam rootchain.SAM) error {
	if m.addMessageFn != nil {
		return m.addMessageFn(sam)
	}

	return nil
}

func (m *mockSAMP) Prune(index uint64) {
	if m.pruneFn != nil {
		m.pruneFn(index)
	}
}

func (m *mockSAMP) Peek() rootchain.VerifiedSAM {
	if m.peekFn != nil {
		return m.peekFn()
	}

	return nil
}

func (m *mockSAMP) Pop() rootchain.VerifiedSAM {
	if m.popFn != nil {
		return m.popFn()
	}

	return nil
}

type signDelegate func([]byte) ([]byte, uint64, error)
type verifySignatureDelegate func([]byte, []byte, uint64) error

type mockSigner struct {
	signFn            signDelegate
	verifySignatureFn verifySignatureDelegate
}

func (m *mockSigner) Sign(data []byte) ([]byte, uint64, error) {
	if m.signFn != nil {
		return m.signFn(data)
	}

	return nil, 0, nil
}

func (m *mockSigner) VerifySignature(
	rawData []byte,
	signature []byte,
	signedBlock uint64,
) error {
	if m.verifySignatureFn != nil {
		return m.verifySignatureFn(
			rawData,
			signature,
			signedBlock,
		)
	}

	return nil
}

type publishDelegate func(*proto.SAM) error
type subscribeTransportDelegate func(func(sam *proto.SAM)) error

type mockTransport struct {
	publishFn   publishDelegate
	subscribeFn subscribeTransportDelegate
}

func (m *mockTransport) Publish(sam *proto.SAM) error {
	if m.publishFn != nil {
		return m.publishFn(sam)
	}

	return nil
}

func (m *mockTransport) Subscribe(fn func(sam *proto.SAM)) error {
	if m.subscribeFn != nil {
		return m.subscribeFn(fn)
	}

	return nil
}
