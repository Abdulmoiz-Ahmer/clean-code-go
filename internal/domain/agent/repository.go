// Repository Interface for agent
package domain

type Repository interface {
  GetByID(id uuid.UUID) (*Agent, error)
  GetAll() ([]Agent, error)
  Add(agent Agent) error
  Update(agent Agent) error
  Delete(id uuid.UUID) error
}