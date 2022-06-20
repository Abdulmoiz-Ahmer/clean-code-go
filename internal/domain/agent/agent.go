package agent

import (
	"time"

	"github.com/google/uuid"
)

//Agent Model that represents the Agent
type Agent struct {
  ID        uuid.UUID 
  FirstName      string    
  LastName      string    
  Email      string    
  Password      string    
  Country   string    
  CreatedAt time.Time 
  DeletedAt time.Time 
}