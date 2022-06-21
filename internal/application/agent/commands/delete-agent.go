package commands

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
)

//DeleteAgentRequest Command Model
type DeleteAgentRequest struct {
	AgentID uuid.UUID
}

//DeleteAgentRequestHandler Handler Struct with Dependencies
type DeleteAgentRequestHandler interface {
	Handle(command DeleteAgentRequest) error
}

type deleteAgentRequestHandler struct {
	repository agent.Repository
}

//Handle Handlers the DeleteAgentRequest request
func (handler deleteAgentRequestHandler) Handle(command DeleteAgentRequest) error {
	agent, err := handler.repository.GetByID(command.AgentID)
	if agent == nil {
		return fmt.Errorf("the provided agent id does not exist")
	}
	if err != nil {
		return err
	}
	return handler.repository.Delete(command.AgentID)

}

//NewDeleteAgentRequestHandler Handler constructor
func NewDeleteAgentRequestHandler(repository agent.Repository) DeleteAgentRequestHandler {
	return deleteAgentRequestHandler{repository: repository}
}