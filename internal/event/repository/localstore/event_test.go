package localstore

import (
	"fmt"
	"testing"

	"github.com/pamost/calendar/internal/entities"
	"github.com/pamost/calendar/internal/event"
	"github.com/stretchr/testify/assert"
)

// Test Create Events in LocalStore
func TestCreate(t *testing.T) {
	localStore := NewEventLocalStore()

	for i := 1; i < 6; i++ {
		testEvent := &entities.Event{
			ID:    fmt.Sprintf("id%d", i),
			Title: fmt.Sprintf("Title%d", i),
		}

		err := localStore.Create(testEvent)
		assert.NoError(t, err)
	}

	returnedEvents, err := localStore.Read()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(returnedEvents))
}

// Test Update Events in LocalStore
func TestUpdate(t *testing.T) {
	localStore := NewEventLocalStore()

	eventTest1 := &entities.Event{ID: "id1", Title: "Title1"}
	err := localStore.Create(eventTest1)
	assert.NoError(t, err)

	eventTest2 := &entities.Event{ID: "id1", Title: "Title2"}
	err = localStore.Update(eventTest2)
	assert.NoError(t, err)

	assert.Equal(t, localStore.events["id1"].Title, "Title2")
}

// Test Delete Events from LocalStore
func TestDelete(t *testing.T) {
	localStore := NewEventLocalStore()

	eventTest1 := &entities.Event{ID: "id1"}

	err := localStore.Create(eventTest1)
	assert.NoError(t, err)

	err = localStore.Delete("id1")
	assert.NoError(t, err)

	// If Event Not Found
	err = localStore.Delete("")
	assert.Error(t, err)
	assert.Equal(t, err, event.ErrEventNotFound)
}
