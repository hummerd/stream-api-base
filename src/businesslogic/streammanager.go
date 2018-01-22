package businesslogic

import (
	"errors"
	"sync"

	"model"

	"github.com/pborman/uuid"
)

type StreamManager struct {
	streamsLock sync.RWMutex
	streams     map[string]*streamHandler
}

type streamHandler struct {
	stream     *model.Stream
	streamLock sync.RWMutex
}

func (sm *StreamManager) AddStream() (string, error) {
	sm.streamsLock.Lock()
	defer sm.streamsLock.Unlock()

	newStream := &model.Stream{
		ID:    uuid.New(),
		State: model.StreamStateCreated,
	}

	streamHandler := &streamHandler{
		stream:     newStream,
		streamLock: sync.RWMutex{},
	}

	sm.streams[newStream.ID] = streamHandler
	return newStream.ID, nil
}

func (sm *StreamManager) GetStream(id string) (*model.Stream, error) {
	sm.streamsLock.RLock()
	hStream, ok := sm.streams[id]
	sm.streamsLock.RUnlock()

	if !ok {
		return nil, errors.New("No stream with id: " + id)
	}

	hStream.streamLock.RLock()
	defer hStream.streamLock.RUnlock()

	return hStream.stream.Copy(), nil
}

func (sm *StreamManager) GetStreamList() {

}

func (sm *StreamManager) ActivateStream(id string) (string, error) {
	sm.streamsLock.RLock()
	defer sm.streamsLock.RUnlock()

	hStream, ok := sm.streams[id]
	if !ok {
		return "", errors.New("No stream with id: " + id)
	}

	result := model.StreamStateActive
	hStream.streamLock.Lock()
	defer hStream.streamLock.Unlock()

	if hStream.stream.CanCahngeStateTo(result) {
		hStream.stream.State = result
	} else {
		result = hStream.stream.State
	}

	return result, nil
}
