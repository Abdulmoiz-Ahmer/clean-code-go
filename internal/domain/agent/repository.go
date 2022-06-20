// Repository Interface for agent
package agent

import (
	"github.com/google/uuid"
)

type Repository interface {
  GetByID(id uuid.UUID) (*Agent, error)
  GetAll() ([]Agent, error)
  Add(agent Agent) error
  Update(agent Agent) error
  Delete(id uuid.UUID) error
}