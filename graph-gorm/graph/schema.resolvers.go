package graph

import (
	"context"
	"graph-gorm/db"
	"graph-gorm/model"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err := db.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(c context.Context, id string) (bool, error) {
	userid, _ := strconv.Atoi(id)

	result := db.DB.Delete(&model.User{}, userid)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *queryResolver) GetUser(c context.Context) ([]*model.User, error) {
	// var users []models.User
	// db.DB.Find(*&users)

	var result []*model.User
	err := db.DB.Find(&result).Error
	return result, err

	// for _, u := range users {
	// 	result = append(result, &model.User{
	// 		ID:    strconv.Itoa(int(u.ID)),
	// 		Name:  u.Name,
	// 		Email: u.Email,
	// 	})
	// }

	// return result, nil
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
