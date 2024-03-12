package entities

import "errors"

type (
	Owner struct {
		ID   string `json:"id"`
		Name string `json:"name"`

		Albums []*Album `json:"albums,omitempty"`
		Musics []*Music `json:"musics,omitempty"`
	}

	NewOwnerParams struct {
		ID     string
		Name   string
		Albums []*Album
		Musics []*Music
	}
)

func NewOwner(props NewOwnerParams) (*Owner, error) {
	if props.ID == "" || props.Name == "" {
		return nil, errors.New("invalid owner")
	}

	owner := &Owner{
		ID:     props.ID,
		Name:   props.Name,
		Albums: props.Albums,
		Musics: props.Musics,
	}

	return owner, nil
}
