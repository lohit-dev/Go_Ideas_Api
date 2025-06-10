package storage

import (
	"test_project/test/internal/model"
	utils "test_project/test/pkg"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ps *PostgresStore) AddVote(userID uuid.UUID, ideaID uuid.UUID) utils.Result[string] {
	vote := model.Vote{
		ID:        uuid.New().String(),
		UserID:    userID,
		IdeaID:    ideaID,
		CreatedAt: time.Now(),
	}
	if err := ps.db.Create(&vote).Error; err != nil {
		return utils.NewResult("", err)
	}

	if err := ps.db.Model(&model.Idea{}).Where("id = ?", ideaID).
		UpdateColumn("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
		return utils.NewResult("", err)
	}

	return utils.NewResult("Successfully Added vote", nil)
}

func (ps *PostgresStore) RemoveVote(userID uuid.UUID, ideaID uuid.UUID) utils.Result[string] {
	if err := ps.db.Where("user_id = ? AND idea_id = ?", userID, ideaID).
		Delete(&model.Vote{}).Error; err != nil {
		return utils.NewResult("", err)
	}

	if err := ps.db.Model(&model.Idea{}).Where("id = ?", ideaID).
		UpdateColumn("votes", gorm.Expr("votes - ?", 1)).Error; err != nil {
		return utils.NewResult("", err)
	}

	return utils.NewResult("Successfully Removed Vote", nil)
}

func (ps *PostgresStore) HasUserVoted(userID uuid.UUID, ideaID uuid.UUID) utils.Result[bool] {
	var count int64
	err := ps.db.Model(&model.Vote{}).
		Where("user_id = ? AND idea_id = ?", userID, ideaID).
		Count(&count).Error

	return utils.NewResult(count > 0, err)
}

func (ps *PostgresStore) GetVoteCount(ideaID uuid.UUID) utils.Result[int] {
	var count int64
	err := ps.db.Model(&model.Vote{}).
		Where("idea_id = ?", ideaID).
		Count(&count).Error

	return utils.NewResult(int(count), err)
}
