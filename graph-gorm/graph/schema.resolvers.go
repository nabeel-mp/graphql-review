package graph

import (
	"context"
	"graph-gorm/db"
	"graph-gorm/graph/model"
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

func (r *mutationResolver) DeleteUser(ctx context.Context, id uint) (bool, error) {
	// userid, _ := strconv.ParseUint(id, 10, 32)

	result := db.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
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

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
