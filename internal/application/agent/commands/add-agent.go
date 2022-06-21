package commands

import (
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/package/time"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/package/uuid"
)

//AddAgentRequest Model of CreateAgentRequestHandler
type AddAgentRequest struct {
  FirstName      string    
  LastName      string    
  Email      string    
  Password      string    
  Country   string   
}

//CreateAgentRequestHandler Struct that allows handling AddAgentRequest
type CreateAgentRequestHandler interface {
	Handle(command AddAgentRequest) error
}

type addAgentRequestHandler struct {
	uuidProvider        uuid.Provider
	timeProvider        time.Provider
	repository          agent.Repository
}

//Handle Handles the AddAgentRequest
func (handler addAgentRequestHandler) Handle(request AddAgentRequest) (error) {
	newAgent := agent.Agent{
		ID:        handler.uuidProvider.NewUUID(),
		FirstName:      request.FirstName,
		LastName:      request.LastName,
    	Email: request.Email,
		Password:      request.Password,
		Country:   request.Country,
		CreatedAt: handler.timeProvider.Now(),
	}
	error := handler.repository.Add(newAgent);

	if error != nil {
		return error
	}


	// return handler.notificationService.Notify(n)
}

//NewAddAgentRequestHandler Initializes an AddCommandHandler
func NewAddAgentRequestHandler(uuidProvider uuid.Provider, timeProvider time.Provider, repository agent.Repository) CreateAgentRequestHandler {
	return addAgentRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repository: repository}
}

