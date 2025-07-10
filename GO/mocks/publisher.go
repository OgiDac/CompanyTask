package mocks

import (
	"github.com/OgiDac/CompanyTask/domain"
)

type Publisher struct {
	Published []domain.UserEventEnvelope
}

func (p *Publisher) PublishEvent(envelope domain.UserEventEnvelope) error {
	p.Published = append(p.Published, envelope)
	return nil
}
