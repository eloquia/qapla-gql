package seequell

import "qaplagql/graph/model"

func ProjectToListItemDomain(project CoreQaplaProject) *model.ProjectListItem {
	return &model.ProjectListItem{
		ID:   int(project.ProjectID),
		Name: project.ProjectName,
		Slug: project.Slug,
	}
}
