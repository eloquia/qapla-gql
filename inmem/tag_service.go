package inmem

import (
	"context"
	"qaplagql/graph/model"
)

type TagServiceInmem struct {
	tags map[string]*model.MeetingNoteTag
}

func NewTagServiceInmem(tagMap map[string]*model.MeetingNoteTag) *TagServiceInmem {
	return &TagServiceInmem{
		tags: tagMap,
	}
}

func (t *TagServiceInmem) GetAll(ctx context.Context) ([]*model.MeetingNoteTag, error) {
	var tagArray []*model.MeetingNoteTag
	for _, tag := range t.tags {
		tagArray = append(tagArray, tag)
	}

	return tagArray, nil
}
