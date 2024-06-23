package gorm

import (
	"github.com/awgst/datings/internal/entity/model"
	"github.com/awgst/datings/internal/usecase/repo"
	"github.com/awgst/datings/pkg/pagination"
	"gorm.io/gorm"
)

type userRepo struct {
	*base
}

type userFinder struct {
	*userRepo
}

func NewGormUserFinder(db *gorm.DB) repo.UserFinder {
	return &userFinder{
		userRepo: &userRepo{
			base: &base{
				db: db,
			},
		},
	}
}

func (u *userFinder) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).Preload("Premium").First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userFinder) FindByID(id int) (model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", id).Preload("Profile").First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userFinder) FindAllProfile(user model.User, paging *pagination.Paginator) ([]model.Profile, error) {
	// TODO: filter profile by user preferences
	var (
		profiles        []model.Profile
		concurrentCount = 2
		errChan         = make(chan error, concurrentCount)
	)

	go func() {
		var count int64
		err := u.db.Raw(`
				SELECT
					count(p.id)
				FROM
					profiles p
				LEFT JOIN
					swipes s
				ON
					s.profile_id = p.id AND DATE(s.created_at) = CURDATE()
				LEFT JOIN
					premiums pr
				ON
					pr.user_id = p.user_id AND pr.feature = 'verified_label'
				WHERE
					p.user_id != ?
				AND
					s.id IS NULL
			`, user.ID).Scan(&count).
			Error

		if err != nil {
			errChan <- err
			return
		}

		paging.SetTotal(count)
		errChan <- nil
	}()

	go func() {
		err := u.db.Raw(
			`
				SELECT
					p.id,
					p.name,
					case
						when pr.feature = 'verified_label' then true
						else false
					end as is_verified
				FROM
					profiles p
				LEFT JOIN
					swipes s
				ON
					s.profile_id = p.id AND DATE(s.created_at) = CURDATE()
				LEFT JOIN
					premiums pr
				ON
					pr.user_id = p.user_id AND pr.feature = 'verified_label'
				WHERE
					p.user_id != ?
				AND
					s.id IS NULL
				LIMIT ?
				OFFSET ?
			`,
			user.ID,
			paging.Limit,
			paging.Offset,
		).Scan(&profiles).Error
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	var errors []error
	for i := 0; i < concurrentCount; i++ {
		err := <-errChan
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return []model.Profile{}, errors[0]
	}

	return profiles, nil
}
