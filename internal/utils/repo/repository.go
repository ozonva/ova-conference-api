package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils"
)

type repository struct {
	connection string
	db         *sqlx.DB
	chunkSize  int
}

func (repository repository) Close() error {
	return repository.db.Close()
}

func (repository *repository) AddEntities(ctx context.Context, entities []domain.Conference, batchCompletedCallback func([]domain.Conference)) ([]domain.Conference, error) {
	tr, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		rollbackLogError(err, tr)
		return nil, err
	}
	bulks, err := utils.SplitToBulksWithoutCopy(entities, repository.chunkSize)
	if err != nil {
		rollbackLogError(err, tr)
		return nil, err
	}
	for _, bulk := range bulks {
		err = insertBulk(tr, bulk)
		if err != nil {
			return nil, err
		}
		if batchCompletedCallback != nil {
			batchCompletedCallback(bulk)
		}
	}
	return entities, tr.Commit()
}

func insertBulk(tr *sqlx.Tx, entities []domain.Conference) error {
	rows, err := tr.NamedQuery(`INSERT INTO conferences (name, event_time, participant_count, speaker_count)
						VALUES (:name, :event_time , :participant_count, :speaker_count) RETURNING id`, entities)
	if err != nil {
		return err
	}
	var id int64
	for i := 0; rows.Next(); i++ {
		if err = rows.Scan(&id); err != nil {
			log.Err(err)
			continue
		}
		entities[i].Id = id
	}
	return nil
}

func (repository *repository) UpdateEntity(ctx context.Context, entity domain.Conference) error {
	tr, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		rollbackLogError(err, tr)
		return err
	}
	_, err = tr.NamedExecContext(ctx, `UPDATE conferences 
											 SET name = :name, event_time = :event_time , 
											      participant_count = :participant_count, speaker_count = :speaker_count
											 WHERE id = :id `, entity)
	if err != nil {
		rollbackLogError(err, tr)
		return err
	}
	return tr.Commit()
}

func (repository *repository) ListEntities(ctx context.Context, limit, offset int64) ([]domain.Conference, error) {
	var conferences []domain.Conference
	err := repository.db.SelectContext(ctx, &conferences, "SELECT * FROM conferences ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return conferences, nil
}

func (repository *repository) DescribeEntity(ctx context.Context, entityId int64) (*domain.Conference, error) {
	result := domain.Conference{}
	err := repository.db.GetContext(ctx, &result, "SELECT * FROM conferences WHERE id = $1", entityId)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return &result, nil
}

func (repository *repository) AddEntity(ctx context.Context, entity domain.Conference) (*domain.Conference, error) {
	tr, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		rollbackLogError(err, tr)
		return nil, err
	}

	var lastInserted int64
	err = tr.QueryRowx(`INSERT INTO conferences (name, event_time, participant_count, speaker_count)
						VALUES ($1, $2, $3, $4) RETURNING id`, entity.Name, entity.EventTime, entity.ParticipantCount, entity.SpeakerCount).Scan(&lastInserted)
	if err != nil {
		rollbackLogError(err, tr)
		return nil, err
	}

	entity.Id = lastInserted
	return &entity, tr.Commit()
}

func rollbackLogError(err error, tr *sqlx.Tx) {
	log.Err(err)
	err = tr.Rollback()
	if err != nil {
		log.Err(err)
	}
}

func (repository *repository) DeleteEntity(ctx context.Context, entityId int64) error {
	tr, err := repository.db.BeginTxx(ctx, nil)
	if err != nil {
		rollbackLogError(err, tr)
		return err
	}
	_, err = tr.ExecContext(ctx, `delete from conferences where id=$1`, entityId)
	if err != nil {
		rollbackLogError(err, tr)
		return err
	}
	return tr.Commit()
}

func (repository *repository) Open() error {
	db, err := sqlx.Open("postgres", repository.connection)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	repository.db = db
	return nil
}
