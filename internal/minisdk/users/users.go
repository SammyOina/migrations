package users

import (
	"context"
	"encoding/json"

	"github.com/mainflux/mainflux/auth"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/mainflux/users"
	"github.com/mainflux/mainflux/users/postgres"
)

var (
	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("entity not found")
	// ErrViewEntity indicates error in viewing entity or entities.
	ErrViewEntity = errors.New("view entity failed")
)

type UserRepository interface {
	GetUsers(ctx context.Context, size, offset int) ([]users.User, error)
}

var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	db postgres.Database
}

// NewUserRepo instantiates a PostgreSQL implementation of user
// repository.
func NewUserRepo(db postgres.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur userRepository) GetUsers(ctx context.Context, limit, offset int) ([]users.User, error) {
	q := `SELECT id, email, metadata FROM users ORDER BY id LIMIT :limit OFFSET :offset;`
	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	rows, err := ur.db.NamedQueryContext(ctx, q, params)
	if err != nil {
		return []users.User{}, errors.Wrap(ErrViewEntity, err)
	}
	defer rows.Close()

	var items []users.User
	for rows.Next() {
		dbusr := dbUser{}
		if err := rows.StructScan(&dbusr); err != nil {
			return []users.User{}, errors.Wrap(ErrViewEntity, err)
		}

		user, err := toUser(dbusr)
		if err != nil {
			return []users.User{}, err
		}

		items = append(items, user)
	}

	return items, nil
}

type dbUser struct {
	ID       string       `db:"id"`
	Email    string       `db:"email"`
	Password string       `db:"password"`
	Metadata []byte       `db:"metadata"`
	Groups   []auth.Group `db:"groups"`
	Status   string       `db:"status"`
}

func toUser(dbu dbUser) (users.User, error) {
	var metadata map[string]interface{}
	if dbu.Metadata != nil {
		if err := json.Unmarshal([]byte(dbu.Metadata), &metadata); err != nil {
			return users.User{}, errors.Wrap(errors.ErrMalformedEntity, err)
		}
	}

	return users.User{
		ID:       dbu.ID,
		Email:    dbu.Email,
		Password: dbu.Password,
		Metadata: metadata,
	}, nil
}
