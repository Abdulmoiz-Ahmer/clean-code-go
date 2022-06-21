package queries

import (
	"time"

	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
	"github.com/google/uuid"
)

// GetAllAgentsResult is the result of the GetAllAgentsRequest Query
type GetAllAgentsResult struct {
  ID        uuid.UUID 
  FirstName      string    
  LastName      string    
  Email      string    
  Password      string    
  Country   string    
  CreatedAt time.Time 
  DeletedAt time.Time 
}

//GetAllAgentsRequestHandler Contains the dependencies of the Handler
type GetAllAgentsRequestHandler interface {
	Handle() ([]GetAllAgentsResult, error)
}

type getAllAgentsRequestHandler struct {
	repository agent.Repository
}

//Handle Handles the query
func (handler getAllAgentsRequestHandler) Handle() ([]GetAllAgentsResult, error) {

	response, error := handler.repository.GetAll()
	if error != nil {
		return nil, error
	}
	var result []GetAllAgentsResult
	for _, agent := range response {
		result = append(result, GetAllAgentsResult{ID: agent.ID, FirstName: agent.FirstName, LastName: agent.LastName, Country: agent.Country, CreatedAt: agent.CreatedAt, DeletedAt: agent.DeletedAt})
	}
	return result, nil
}


//NewGetAllAgentsRequestHandler Handler constructor
func NewGetAllAgentsRequestHandler(repository agent.Repository) GetAllAgentsRequestHandler {
	return getAllAgentsRequestHandler{repository: repository}
}