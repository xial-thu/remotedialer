package remotedialer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (sm *sessionManager) fakeGetDialer(clientKey string) (*Session, error) {
	sm.Lock()
	defer sm.Unlock()

	sessions := sm.clients[clientKey]
	if len(sessions) > 0 {
		return sessions[0], nil
	}
	return nil, fmt.Errorf("session not found")
}

func (sm *sessionManager) fakeGetRandomDialer(clientKey string) (*Session, error) {
	sm.Lock()
	defer sm.Unlock()

	sessions := sm.clients[clientKey]
	if len(sessions) == 1 {
		return sessions[0], nil
	}
	if len(sessions) > 1 {
		return sessions[rand.Intn(len(sessions))], nil
	}
	return nil, fmt.Errorf("session not found")
}

func newFakeSessionManager(n int) *sessionManager {
	sm := newSessionManager()
	sessions := make([]*Session, n)
	for i := 0; i < n; i++ {
		sessions[i] = &Session{}
	}
	sm.clients["client"] = sessions
	return sm
}

func BenchmarkGetDialer(b *testing.B) {
	sm := newFakeSessionManager(10)
	for i := 0; i < b.N; i++ {
		sm.fakeGetDialer("client")
	}
}

func BenchmarkGetRandomDialer(b *testing.B) {
	sm := newFakeSessionManager(10)
	for i := 0; i < b.N; i++ {
		sm.fakeGetRandomDialer("client")
	}
}
