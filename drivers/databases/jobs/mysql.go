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

func (rep *MysqlJobRepository) CreateJob(ctx context.Context, domain jobs.Domain) (jobs.Domain, error) {
	job := FromDomain(domain)

	result := rep.Conn.Preload("Category").Create(&job)

	if result.Error != nil {
		return jobs.Domain{}, result.Error
	}

	return job.ToDomain(), nil
}

func (rep *MysqlJobRepository) GetAllJobs(ctx context.Context) ([]jobs.Domain, error) {
	var job []Jobs
	result := rep.Conn.Preload("Applications").Preload("Category").Find(&job)

	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}

func (rep *MysqlJobRepository) GetJobById(ctx context.Context, id int) (jobs.Domain, error) {
	var job Jobs
	result := rep.Conn.Preload("Applications").Preload("Category").First(&job, "id = ?", id)

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
	result := rep.Conn.Preload("Applications").Preload("Category").Where("title LIKE ?", title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	result = rep.Conn.Preload("Applications").Preload("Category").Where("title LIKE ?", title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	result = rep.Conn.Preload("Applications").Preload("Category").Where("title LIKE ?", "%"+title+"%").Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}

func (rep *MysqlJobRepository) FilterJobByCategory(ctx context.Context, categoryId int) ([]jobs.Domain, error) {
	var job []Jobs
	result := rep.Conn.Preload("Applications").Preload("Category").Where("category_id = ?", categoryId).Find(&job)
	if result.Error != nil {
		return []jobs.Domain{}, result.Error
	}

	return ListToDomain(job), nil
}
