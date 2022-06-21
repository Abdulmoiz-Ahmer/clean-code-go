package application

import (
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/application/agent/commands"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/application/agent/queries"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/domain/agent"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/package/time"
	"github.com/Abdulmoiz-Ahmer/live-agents-services/internal/package/uuid"
)

//Queries Contains all available query handlers of this app
type Queries struct {
	GetAllAgentsHandler queries.GetAllAgentsRequestHandler
	GetAgentHandler     queries.GetAgentRequestHandler
}

//Commands Contains all available command handlers of this app
type Commands struct {
	// CreateAgentHandler commands.CreateAgentRequestHandler
	UpdateAgentHandler commands.UpdateAgentRequestHandler
	DeleteAgentHandler commands.DeleteAgentRequestHandler
}

//AgentServices Contains the grouped queries and commands of the app layer
type AgentServices struct {
	Queries  Queries
	Commands Commands
}

//Services contains all exposed services of the application layer
type Services struct {
	AgentServices AgentServices
}

// NewServices Bootstraps Application Layer dependencies
func NewServices(agentRepo agent.Repository, up uuid.Provider, tp time.Provider) Services {
	return Services{
		AgentServices: AgentServices{
			Queries: Queries{
				GetAllAgentsHandler: queries.NewGetAllAgentsRequestHandler(agentRepo),
				GetAgentHandler:     queries.NewGetAgentRequestHandler(agentRepo),
			},
			Commands: Commands{
				// CreateAgentHandler: commands.NewAddAgentRequestHandler(up, tp, agentRepo, ns),
				UpdateAgentHandler: commands.NewUpdateAgentRequestHandler(agentRepo),
				DeleteAgentHandler: commands.NewDeleteAgentRequestHandler(agentRepo),
			},
		}}
}