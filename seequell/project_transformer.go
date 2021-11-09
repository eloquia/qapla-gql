package seequell

import "qaplagql/graph/model"

func ProjectToListItemDomain(project CoreQaplaProject) *model.ProjectListItem {
	return &model.ProjectListItem{
		ID:   int(project.ProjectID),
		Name: project.ProjectName,
		Slug: project.Slug,
	}
}

func ProjectToDomain(projectModel CoreQaplaProject) *model.Project {
	return &model.Project{
		ID:          int(projectModel.ProjectID),
		Name:        projectModel.ProjectName,
		Description: &projectModel.ProjectDesc.String,
		Slug:        projectModel.Slug,
	}
}

func ProjectToDetailsDomain(projectModel CoreQaplaProject) *model.ProjectDetails {
	return &model.ProjectDetails{
		ID:          int(projectModel.ProjectID),
		Name:        projectModel.ProjectName,
		Description: projectModel.ProjectDesc.String,
		Slug:        projectModel.Slug,
	}
}
