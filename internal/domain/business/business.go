package business

import (
	"time"

	"github.com/google/uuid"
)

//Business Model that represents the Business
type Business struct {
  ID        uuid.UUID 
  FirstName      string    
  LastName      string    
  Email      string    
  Password      string    
  Country   string    
  CreatedAt time.Time 
  DeletedAt time.Time 
}