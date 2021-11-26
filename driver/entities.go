package driver

import (
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
		{UserModel.User{}},
		{UserModel.Experience{}},
		{UserModel.Skillset{}},
		{RecruiterModel.Recruiter{}},
		{InvitationModel.Invitation{}},
	}
}
