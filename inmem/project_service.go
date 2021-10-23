package inmem

import (
	"context"
	"errors"
	"qaplagql/common"
	"qaplagql/graph/model"
	"strings"
)

type ProjectServiceInmem struct {
	projects    map[string]*model.Project
	projectUser map[string][]string
	users       map[string]*model.User
}

func NewProjectServiceInmem(userMap map[string]*model.User, projectMap map[string]*model.Project, projectUserMap map[string][]string) *ProjectServiceInmem {
	return &ProjectServiceInmem{
		users:       userMap,
		projects:    projectMap,
		projectUser: projectUserMap,
	}
}

func (projectService *ProjectServiceInmem) CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error) {
	var sameNameProject *model.Project
	for id, project := range projectService.projects {
		if project.Name == input.Name {
			sameNameProject = projectService.projects[id]
		}
	}

	if sameNameProject != nil {
		return &model.Project{}, errors.New("Project with name already exists")
	}

	projectID := common.RandomId()
	newProject := &model.Project{
		ID:          projectID,
		Name:        input.Name,
		Description: input.Description,
		Slug:        kebabify(input.Name),
	}

	projectService.projects[projectID] = newProject

	return newProject, nil
}

func (projectService *ProjectServiceInmem) GetAll(ctx context.Context) ([]*model.ProjectDetails, error) {
	var projects []*model.ProjectDetails
	for _, project := range projectService.projects {
		projectDetail := &model.ProjectDetails{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Slug:        project.Slug,
		}

		// get personnel for each project
		var projectPersonnel []*model.User
		projectPersonnelIDs := projectService.projectUser[projectDetail.ID]
		for _, uID := range projectPersonnelIDs {
			user := projectService.users[uID]
			projectPersonnel = append(projectPersonnel, user)
		}

		projectDetail.Personnel = projectPersonnel

		projects = append(projects, projectDetail)
	}
	return projects, nil
}

func (projectService *ProjectServiceInmem) GetById(ctx context.Context, id string) (*model.Project, error) {
	foundProject := projectService.projects[id]
	if foundProject == nil {
		return &model.Project{}, errors.New("Project with ID not found")
	}
	return foundProject, nil
}

func (projectService *ProjectServiceInmem) GetProjectDetails(ctx context.Context, slug string) (*model.ProjectDetails, error) {
	if slug == "" {
		return &model.ProjectDetails{}, errors.New("Slug needs to be defined")
	}

	var foundProject *model.Project
	for _, project := range projectService.projects {
		if project.Slug == slug {
			foundProject = project
		}
	}

	if foundProject == nil {
		return &model.ProjectDetails{}, errors.New("Project with slug not found")
	}

	projectDetails := &model.ProjectDetails{
		ID:          foundProject.ID,
		Name:        foundProject.Name,
		Description: foundProject.Description,
		Slug:        foundProject.Slug,
	}

	return projectDetails, nil
}

func (p *ProjectServiceInmem) AddUserToProject(ctx context.Context, userId string, projectId string) (*model.User, error) {
	// check to see that user and project exist
	user := p.users[userId]
	if user == nil {
		return &model.User{}, errors.New("User with ID does not exist.")
	}

	project := p.projects[projectId]
	if project == nil {
		return &model.User{}, errors.New("Project with ID does not exist.")
	}

	// if they exist, then check project-user map
	projectUsers := p.projectUser[projectId]
	if projectUsers == nil || len(projectUsers) == 0 {
		projectUsers = make([]string, 1)
	}

	isAssigned := false
	for i := 0; i < len(projectUsers); i++ {
		if projectUsers[i] == userId {
			isAssigned = true
		}
	}

	// if the association doesn't exist, add it
	if !isAssigned {
		projectUsers = append(projectUsers, userId)
	}

	p.projectUser[projectId] = projectUsers

	return user, nil
}

func (p *ProjectServiceInmem) GetProjectPersonnel(ctx context.Context, projectID string) ([]*model.User, error) {
	// Check projects to see if it exists
	project := p.projects[projectID]
	if project == nil {
		return []*model.User{}, errors.New("Project with ID does not exist")
	}

	// get list of userIDs assigned to project
	projectUserIDs := p.projectUser[projectID]

	var users []*model.User
	// get user info from User ID
	for i := 0; i < len(projectUserIDs); i++ {
		userId := projectUserIDs[i]
		user := p.users[userId]

		users = append(users, user)
	}

	return users, nil
}

func (p *ProjectServiceInmem) GetAssignedProjects(ctx context.Context, userID string) ([]*model.Project, error) {
	var assignedProjects []*model.Project
	if userID == "" {
		return assignedProjects, errors.New("User ID must be defined")
	}

	// see if user exists in database
	foundUser := p.users[userID]

	if foundUser == nil {
		return assignedProjects, errors.New("User with ID not found in the database")
	}

	// get project IDs associated with user
	foundProjectIDs := p.projectUser[userID]

	// get projects associated with each ID
	for i := 0; i < len(foundProjectIDs); i++ {
		foundProjectID := foundProjectIDs[i]
		foundProject := p.projects[foundProjectID]
		assignedProjects = append(assignedProjects, foundProject)
	}

	return assignedProjects, nil
}

func (p *ProjectServiceInmem) GetProjectListItems(ctx context.Context) ([]*model.ProjectListItem, error) {
	var projects []*model.ProjectListItem
	for _, project := range p.projects {
		projectListItem := &model.ProjectListItem{
			ID:   project.ID,
			Name: project.Name,
			Slug: project.Slug,
		}

		// get personnel for each project
		var projectPersonnel []*model.User
		projectPersonnelIDs := p.projectUser[projectListItem.ID]
		for _, uID := range projectPersonnelIDs {
			user := p.users[uID]
			projectPersonnel = append(projectPersonnel, user)
		}

		projectListItem.Personnel = projectPersonnel

		projects = append(projects, projectListItem)
	}
	return projects, nil
}

func kebabify(input string) string {
	hyphens := strings.ReplaceAll(input, " ", "-")
	lowercase := strings.ToLower(hyphens)
	return lowercase
}

func (p *ProjectServiceInmem) GetAllProjectDetails(ctx context.Context) ([]*model.ProjectDetails, error) {
	var projectDetailsList []*model.ProjectDetails
	for _, project := range p.projects {
		projectDetails := &model.ProjectDetails{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Slug:        project.Slug,
		}

		var personnel []*model.User
		personnelIDs := p.projectUser[project.ID]
		for _, pID := range personnelIDs {
			p := p.users[pID]
			personnel = append(personnel, p)
		}

		projectDetails.Personnel = personnel
		projectDetailsList = append(projectDetailsList, projectDetails)
	}
	return projectDetailsList, nil
}
