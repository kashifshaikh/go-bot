package profile

import (
	"bot"
	"time"

	"github.com/asdine/storm/v3"
	"go.uber.org/zap"
)

type Service struct {
	log *zap.SugaredLogger
	db  *storm.DB
}

func NewService(log *zap.SugaredLogger, db *storm.DB) *Service {
	return &Service{log: log, db: db}
}
func (s *Service) Add(profile *Profile) error {

	profile.CreatedAt = time.Now().Unix()
	profile.UpdatedAt = time.Now().Unix()
	profile.DeletedAt = bot.TIMESTAMP_NOT_SET
	return s.db.Save(profile)
}

func (s *Service) Get(id bot.PK) (profile *Profile, err error) {
	err = s.db.One("ID", id, profile)
	if err != nil {
		return nil, err
	}
	profile.CreditCard.Number = ""
	profile.CreditCard.CVV = ""
	return profile, nil
}

func (s *Service) GetAll(includeDeleted bool) ([]Profile, error) {

	var profiles []Profile
	var err error
	if includeDeleted {
		err = s.db.All(&profiles)
	} else {
		err = s.db.Find("DeletedAt", bot.TIMESTAMP_NOT_SET, &profiles)
		s.log.Infof("Profiles: %v", profiles)
	}
	if bot.NotFoundError(err) {
		return []Profile{}, nil
	} else if err != nil {
		return nil, err
	}
	for _, p := range profiles {
		p.CreditCard.Number = ""
		p.CreditCard.CVV = ""
	}
	return profiles, nil
}

func (s *Service) Update(profile *Profile) error {
	var org Profile
	err := s.db.One("ID", profile.ID, &org)
	if err != nil {
		return err
	}
	profile.CreatedAt = org.CreatedAt
	profile.UpdatedAt = time.Now().Unix()
	profile.DeletedAt = org.DeletedAt

	// CreditCard is selective-update, so if not specified use existing
	if profile.CreditCard == nil {
		profile.CreditCard = org.CreditCard
	}
	return s.db.Save(profile)
}

func (s *Service) Delete(id bot.PK) error {
	return s.db.UpdateField(&Profile{ID: id}, "DeletedAt", time.Now().Unix())
}

// DeleteAll marks all profiles as deleted
func (s *Service) DeleteAll() error {

	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var profiles []Profile
	err = tx.Find("DeletedAt", bot.TIMESTAMP_NOT_SET, &profiles)
	if bot.NotFoundError(err) {
		// Nothing to delete
		return nil
	} else if err != nil {
		return err
	}

	now := time.Now().Unix()
	for _, p := range profiles {
		s.log.Infof("Deleting profile: %d", p.ID)
		p.DeletedAt = now
		tx.Save(&p)
	}

	return tx.Commit()
}

// PurgeAll permanently deletes all deleted profiles
func (s *Service) PurgeAll() error {
	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var profiles []Profile
	//err = s.db.Select(q.Gt("DeletedAt", time.Time{}.)).Find(&profiles)
	err = tx.Range("DeletedAt", 0, time.Now().Unix(), &profiles)
	if bot.NotFoundError(err) {
		// Nothing to delete
		return nil
	} else if err != nil {
		return err
	}

	for _, p := range profiles {
		s.log.Infof("Purging deleted profile: %d", p.ID)
		tx.DeleteStruct(&p)
	}

	return tx.Commit()
	// err := s.db.Drop(&Profile{})
	// if err != nil {
	// 	return err
	// }
	// // return s.db.Init(&Profile{})
	// return nil
}
