package service

import (
	"errors"
	"fmt"
	"workuo/features/application"
	"workuo/features/invitation"
	"workuo/features/job"
	"workuo/features/user"
)

type invitationService struct {
	invRepository invitation.Repository
	jobService    job.Service
	userService   user.Service
	appService    application.Service
}

func NewInvitationService(ir invitation.Repository, js job.Service, us user.Service, as application.Service) invitation.Service {
	return &invitationService{
		invRepository: ir,
		jobService:    js,
		userService:   us,
		appService:    as,
	}
}

func (is *invitationService) InviteUser(data invitation.InvitationCore) error {
	if data.Role != "recruiter" {
		return errors.New("only recruiter role allowed to invite user")
	}

	jobData, err := is.jobService.GetJobPostById(int(data.JobID))
	if err != nil {
		return err
	}
	if jobData.ID == 0 {
		msg := fmt.Sprintf("job with id %v not found", data.JobID)
		return errors.New(msg)
	}
	if jobData.RecruiterId != int(data.RecruiterID) {
		msg := fmt.Sprintf("recruiter with id %v didn't have job post with id %v", data.RecruiterID, data.JobID)
		return errors.New(msg)
	}

	userData, err := is.userService.GetUserById(int(data.UserID))
	if userData.Id == 0 {
		msg := fmt.Sprintf("user with id %v not found", data.UserID)
		return errors.New(msg)
	}
	if err != nil {
		return err
	}

	appData, err := is.appService.GetApplicationMultiParam(int(data.JobID), int(data.UserID))
	if appData.ID != 0 {
		msg := fmt.Sprintf("user with id %v has applied this job with status %v", data.UserID, appData.Status)
		return errors.New(msg)
	}
	if err != nil {
		return err
	}

	data.Status = "pending"

	err = is.invRepository.InviteUser(data)
	if err != nil {
		return err
	}

	return nil
}
