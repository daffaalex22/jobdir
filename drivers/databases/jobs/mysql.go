package jobs

import (
	"context"

	"gorm.io/gorm"
	"main.go/business/jobs"
)

type MysqlJobRepository struct {
	Conn *gorm.DB
}

func NewMysqlJobRepository(conn *gorm.DB) jobs.Repository {
	return &MysqlJobRepository{
		Conn: conn,
	}
}

func (rep *MysqlJobRepository) FillJobs(ctx context.Context, categories []jobs.CategoryDomain) ([]jobs.CategoryDomain, error) {

	for _, category := range categories {
		var job []Jobs
		result := rep.Conn.Where("categoryId = ?", category.Id).Find(&job)
		if result.Error != nil {
			return categories, result.Error
		}
		category.Jobs = append(category.Jobs, ListToDomain(job)...)
	}
	return categories, nil
}

func (rep *MysqlJobRepository) CreateJob(ctx context.Context, domain jobs.Domain) (jobs.Domain, error) {
	// job := Jobs{
	// 	Title:    domain.Title,
	// 	Category: domain.Category,
	// 	JobDesc:  domain.JobDesc,
	// }

	job := FromDomain(domain)

	result := rep.Conn.Create(&job)

	if result.Error != nil {
		return jobs.Domain{}, result.Error
	}

	return job.ToDomain(), nil
}

func (rep *MysqlJobRepository) GetAllJobs(ctx context.Context) ([]jobs.Domain, error) {
	var job []Jobs
	result := rep.Conn.Find(&job)

	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}

func (rep *MysqlJobRepository) GetJobById(ctx context.Context, id int) (jobs.Domain, error) {
	var job Jobs
	result := rep.Conn.First(&job, "id = ?", id)

	if result.Error != nil {
		return jobs.Domain{}, result.Error
	}

	return job.ToDomain(), nil
}

func (rep *MysqlJobRepository) DeleteAllJobs(ctx context.Context) error {
	var jobs []Jobs

	result := rep.Conn.Find(&jobs)
	if result.Error != nil {
		return result.Error
	}

	result = rep.Conn.Unscoped().Delete(&jobs)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlJobRepository) DeleteJobById(ctx context.Context, id int) (jobs.Domain, error) {
	var job Jobs
	result := rep.Conn.Where("id = ?", id).Delete(&job)

	if result.Error != nil {
		return jobs.Domain{}, result.Error
	}

	return job.ToDomain(), nil
}

func (rep *MysqlJobRepository) SearchJobs(ctx context.Context, title string) ([]jobs.Domain, error) {
	var job []Jobs
	result := rep.Conn.Where("title LIKE ?", title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	result = rep.Conn.Where("title LIKE ?", title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	result = rep.Conn.Where("title LIKE ?", "%"+title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}

func (rep *MysqlJobRepository) FilterJobByCategory(ctx context.Context, categoryId int) ([]jobs.Domain, error) {
	var job []Jobs
	result := rep.Conn.Where("category_id = ?", categoryId).Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}
