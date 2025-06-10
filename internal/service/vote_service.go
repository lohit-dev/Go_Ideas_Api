package service

import (
	"fmt"
	"test_project/test/internal/storage"
	utils "test_project/test/pkg"

	"github.com/google/uuid"
)

type VoteService struct {
	store storage.VoteStorage
}

func NewVoteService(store storage.VoteStorage) *VoteService {
	return &VoteService{store}
}

func (s *VoteService) AddVote(userId uuid.UUID, ideaId uuid.UUID) utils.Result[string] {
	// if user has already voted or not
	hasVoted := s.store.HasUserVoted(userId, ideaId)
	if hasVoted.Err != nil {
		return utils.Result[string]{
			Err: fmt.Errorf(hasVoted.Err.Error()),
		}
	}

	if hasVoted.Data {
		return utils.Result[string]{
			Err: fmt.Errorf("User has already voted"),
		}
	}

	return s.store.AddVote(userId, ideaId)
}

func (s *VoteService) RemoveVote(userId uuid.UUID, ideaId uuid.UUID) utils.Result[string] {
	hasVoted := s.store.HasUserVoted(userId, ideaId)
	if hasVoted.Err != nil {
		return utils.Result[string]{
			Err: fmt.Errorf(hasVoted.Err.Error()),
		}
	}
	if !hasVoted.Data {
		return utils.Result[string]{
			Err: fmt.Errorf("User has not voted for this idea"),
		}
	}

	return s.store.RemoveVote(userId, ideaId)
}

func (s *VoteService) HasUserVoted(userId uuid.UUID, ideaId uuid.UUID) utils.Result[bool] {
	return s.store.HasUserVoted(userId, ideaId)
}

func (s *VoteService) GetVoteCount(ideaId uuid.UUID) utils.Result[int] {
	return s.store.GetVoteCount(ideaId)
}
