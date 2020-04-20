package link

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
}

type Linker interface {
	Create(userID string) error
	FetchLinksForUser(userID string, sortByDateCreated bool) error
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(userID string, createLinkRequest *CreateLinkRequest) (*LinkData, error) {
	linkID := uuid.New().String()
	linkData := LinkData{
		LinkID:      linkID,
		UserID:      userID,
		DateCreated: time.Now(),
		Type:        createLinkRequest.Type,
		Data:        createLinkRequest.Data,
	}
	// TODO insert into dynamodb here

	return &linkData, nil
}

func (s *Service) FetchLinksForUser(userID string, sortByDateCreated bool) (*[]LinkData, error) {
	// TODO fetch from dynamodb here, if sortByDateCreated is true then use the sortkey LinkDateCreated in the query

	// This just returns dummy data
	linksData := []LinkData{
		{
			LinkID:      uuid.New().String(),
			UserID:      userID,
			DateCreated: time.Now(),
			Type:        "classic",
			Data:        "",
		},
		{
			LinkID:      uuid.New().String(),
			UserID:      userID,
			DateCreated: time.Now(),
			Type:        "shows_list",
			Data:        "",
		},
		{
			LinkID:      uuid.New().String(),
			UserID:      userID,
			DateCreated: time.Now(),
			Type:        "music_player",
			Data:        "",
		},
	}
	return &linksData, nil
}
