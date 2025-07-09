package domain

type UserEventEnvelope struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type UserCreatedEvent struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserUpdatedEvent struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserDeletedEvent struct {
	ID uint `json:"id"`
}

type EventPublisher interface {
	PublishEvent(envelope UserEventEnvelope) error
}
