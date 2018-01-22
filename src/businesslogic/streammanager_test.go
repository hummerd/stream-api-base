package businesslogic

import (
	"sync"
	"testing"
)

func TestStreamManager_Race(t *testing.T) {
	manager := &StreamManager{
		streamsLock: sync.RWMutex{},
		streams:     map[string]*streamHandler{},
	}

	idMap := sync.Map{}
	wait := sync.WaitGroup{}
	wait.Add(3)

	go func() {
		for index := 0; index < 1000; index++ {
			id, _ := manager.AddStream()
			idMap.Store(id, nil)
		}
		wait.Done()
	}()

	go func() {
		for index := 0; index < 1000; index++ {
			i := 0
			idMap.Range(func(k, v interface{}) bool {
				id := k.(string)
				manager.GetStream(id)
				i++
				return i <= 100
			})

		}
		wait.Done()
	}()

	go func() {
		for index := 0; index < 1000; index++ {
			i := 0
			idMap.Range(func(k, v interface{}) bool {
				id := k.(string)
				manager.ActivateStream(id)
				i++
				return i <= 100
			})
		}
		wait.Done()
	}()

	wait.Wait()
}

// func TestStreamManager_AddStream(t *testing.T) {
// 	type fields struct {
// 		streamsLock sync.RWMutex
// 		streams     map[string]*streamHandler
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    string
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sm := &StreamManager{
// 				streamsLock: tt.fields.streamsLock,
// 				streams:     tt.fields.streams,
// 			}
// 			got, err := sm.AddStream()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("StreamManager.AddStream() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("StreamManager.AddStream() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestStreamManager_GetStream(t *testing.T) {
// 	type fields struct {
// 		streamsLock sync.RWMutex
// 		streams     map[string]*streamHandler
// 	}
// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *model.Stream
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sm := &StreamManager{
// 				streamsLock: tt.fields.streamsLock,
// 				streams:     tt.fields.streams,
// 			}
// 			got, err := sm.GetStream(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("StreamManager.GetStream() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("StreamManager.GetStream() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestStreamManager_GetStreamList(t *testing.T) {
// 	type fields struct {
// 		streamsLock sync.RWMutex
// 		streams     map[string]*streamHandler
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sm := &StreamManager{
// 				streamsLock: tt.fields.streamsLock,
// 				streams:     tt.fields.streams,
// 			}
// 			sm.GetStreamList()
// 		})
// 	}
// }

// func TestStreamManager_ActivateStream(t *testing.T) {
// 	type fields struct {
// 		streamsLock sync.RWMutex
// 		streams     map[string]*streamHandler
// 	}
// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sm := &StreamManager{
// 				streamsLock: tt.fields.streamsLock,
// 				streams:     tt.fields.streams,
// 			}
// 			got, err := sm.ActivateStream(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("StreamManager.ActivateStream() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("StreamManager.ActivateStream() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
