package repository

import (
	"context"
	"em/internal/model"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository struct {
	db *pgxpool.Pool
	sb squirrel.StatementBuilderType
}

func NewPersonRepo(db *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{
		db: db,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *PersonRepository) GetPeople(ctx context.Context, filter model.FilterPeopleRequest) ([]model.Person, error) {
	q := r.sb.Select(
		"id", "name", "surname", "patronymic", "gender", "age", "country_id", "created_at",
	).From("people")

	if filter.Name != "" {
		q = q.Where(squirrel.ILike{"name": "%" + filter.Name + "%"})
	}
	if filter.Surname != "" {
		q = q.Where(squirrel.ILike{"surname": "%" + filter.Surname + "%"})
	}
	if filter.Gender != "" {
		q = q.Where(squirrel.Eq{"gender": filter.Gender})
	}
	if filter.CountryID != "" {
		q = q.Where(squirrel.Eq{"country_id": filter.CountryID})
	}
	if filter.MinAge > 0 {
		q = q.Where(squirrel.GtOrEq{"age": filter.MinAge})
	}
	if filter.MaxAge > 0 {
		q = q.Where(squirrel.LtOrEq{"age": filter.MaxAge})
	}

	page := filter.Page
	if page == 0 {
		page = 1
	}
	limit := filter.Limit
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	q = q.Offset(uint64(offset)).Limit(uint64(limit))

	sqlStr, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var people []model.Person
	for rows.Next() {
		var p model.Person
		err := rows.Scan(
			&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Gender,
			&p.Age, &p.CountryID, &p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}


func (r *PersonRepository) CreatePerson(ctx context.Context, p *model.Person) error{
	q := r.sb.Insert("people").
	Columns("name", "surname", "patronymic", "gender", "age", "country_id").
	Values(p.Name, p.Surname, p.Patronymic,p.Gender,p.Age,p.CountryID).
	Suffix("RETURNING id, created_at") 

	sqlStr, args, err := q.ToSql()
	if err != nil{
		return err
	}

	return r.db.QueryRow(ctx, sqlStr, args...).Scan(&p.ID,&p.CreatedAt)
}