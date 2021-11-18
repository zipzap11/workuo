package driver

import (
	JobModel "workuo/features/job/data"
	UserModel "workuo/features/user/data"
)

type Entity struct {
	Model interface{}
}

func registerEntities() []Entity {
	return []Entity{
		{Model: JobModel.Job{}},
		{Model: JobModel.Requirement{}},
		{Model: UserModel.User{}},
		{Model: UserModel.Skillset{}},
		{Model: UserModel.Experience{}},
	}
}
