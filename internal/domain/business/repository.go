// Repository Interface for business
package domain

type Repository interface {
  GetByID(id uuid.UUID) (*Business, error)
  GetAll() ([]Business, error)
  Add(business Business) error
  Update(business Business) error
  Delete(id uuid.UUID) error
}