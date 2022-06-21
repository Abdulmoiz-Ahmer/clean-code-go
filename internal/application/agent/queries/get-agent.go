package queries

import (
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
	"github.com/google/uuid"

	"time"
)

//GetAgentRequest Model of the Handler
type GetAgentRequest struct {
	AgentID uuid.UUID
}

// GetAgentResult is the return model of Agent Query Handlers
type GetAgentResult struct {
  ID        uuid.UUID 
  FirstName      string    
  LastName      string    
  Email      string    
  Password      string    
  Country   string    
  CreatedAt time.Time 
  DeletedAt time.Time 
}

//GetAgentRequestHandler provides an interfaces to handle a GetAgentRequest and return a *GetAgentResult
type GetAgentRequestHandler interface {
	Handle(query GetAgentRequest) (*GetAgentResult, error)
}

type getAgentRequestHandler struct {
	repository agent.Repository
}

//Handle Handlers the GetAgentRequest query
func (handler getAgentRequestHandler) Handle(query GetAgentRequest) (*GetAgentResult, error) {
	agent, error := handler.repository.GetByID(query.AgentID)
	var result *GetAgentResult
	if agent != nil && error == nil {
		result = &GetAgentResult{ID: agent.ID, FirstName: agent.FirstName, LastName: agent.LastName, Country: agent.Country, CreatedAt: agent.CreatedAt, DeletedAt: agent.DeletedAt}
	}
	return result, error
}

//NewGetAgentRequestHandler Handler Constructor
func NewGetAgentRequestHandler(repository agent.Repository) GetAgentRequestHandler {
	return getAgentRequestHandler{repository: repository}
}