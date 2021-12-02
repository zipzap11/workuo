package driver

import (
	AppModel "workuo/features/application/data"
	InvitationModel "workuo/features/invitation/data"
	JobModel "workuo/features/job/data"
	RecruiterModel "workuo/features/recruiter/data"
	UserModel "workuo/features/user/data"
)

type Entity struct {
	Model interface{}
}

func registerEntities() []Entity {
	return []Entity{
		{JobModel.Job{}},
		{JobModel.Requirement{}},
		{RecruiterModel.Recruiter{}},
		{UserModel.User{}},
		{UserModel.Experience{}},
		{UserModel.Skillset{}},
		{InvitationModel.Invitation{}},
		{AppModel.Application{}},
	}
}
