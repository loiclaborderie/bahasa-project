package dialogue

import (
	"context"
)

type DialogueService struct {
	repo *DialogueRepository
}

func NewService(repo *DialogueRepository) *DialogueService {
	return &DialogueService{repo: repo}
}

func (d *DialogueService) GetAvailableSpeakers(c context.Context, dialogueId int) ([]string, error) {
	speakers, err := d.repo.GetAvailableSpeakers(c, dialogueId)
	if err != nil {
		return nil, err
	}

	return speakers, nil
}

func (d *DialogueService) GetDialogue(c context.Context, dialogueId int, speaker *string) (*Dialogue, error) {
	dialogue, err := d.repo.GetDialogue(c, dialogueId)
	if err != nil {
		return nil, err
	}

	if *speaker != "" {
		setUserRoleInDialogue(dialogue, speaker)
	}

	return dialogue, nil
}

func setUserRoleInDialogue(dialogue *Dialogue, speaker *string) {
	if *speaker == "" {
		return
	}

	for i := range dialogue.Lines {
		isUserTurn := dialogue.Lines[i].Speaker == *speaker
		dialogue.Lines[i].IsUserTurn = &isUserTurn
	}
}
