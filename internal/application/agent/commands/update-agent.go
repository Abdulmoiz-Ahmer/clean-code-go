package commands

import (
	"fmt"

	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
	"github.com/google/uuid"
)

//UpdateAgentRequest Update Model
type UpdateAgentRequest struct {
	ID      uuid.UUID
    FirstName      string    
    LastName      string   
	Country string
}

//UpdateAgentRequestHandler Contains the dependencies of the handler
type UpdateAgentRequestHandler interface {
	Handle(command UpdateAgentRequest) error
}

type updateAgentRequestHandler struct {
	repository agent.Repository
}

//Handle Handles the update request
func (handler updateAgentRequestHandler) Handle(command UpdateAgentRequest) error {
	agent, error := handler.repository.GetByID(command.ID)
	if agent == nil {
		return fmt.Errorf("the provided agent id does not exist")
	}
	if error != nil {
		return error
	}

	agent.FirstName = command.FirstName
	agent.LastName = command.LastName
	agent.Country = command.Country

	return handler.repository.Update(*agent)

}

//NewUpdateAgentRequestHandler Constructor
func NewUpdateAgentRequestHandler(repository agent.Repository) UpdateAgentRequestHandler {
	return updateAgentRequestHandler{repository: repository}
}