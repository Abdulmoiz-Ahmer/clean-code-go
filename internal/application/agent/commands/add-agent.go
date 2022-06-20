package application

import "time"

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