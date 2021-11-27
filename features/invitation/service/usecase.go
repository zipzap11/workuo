package service

import (
	"errors"
	"fmt"
	"workuo/features/invitation"
	"workuo/features/job"
)

type invitationService struct {
	invRepository invitation.Repository
	jobService    job.Service
}

func NewInvitationService(ir invitation.Repository, js job.Service) invitation.Service {
	return &invitationService{
		invRepository: ir,
		jobService:    js,
	}
}

func (is *invitationService) InviteUser(data invitation.InvitationCore) error {
	jobData, err := is.jobService.GetJobPostById(int(data.JobID))
	if err != nil {
		return err
	}

	if jobData.RecruiterId != int(data.RecruiterID) {
		msg := fmt.Sprintf("recruiter with id %v didn't have job post with id %v", data.RecruiterID, data.JobID)
		return errors.New(msg)
	}

	err = is.invRepository.InviteUser(data)
	if err != nil {
		return err
	}

	return nil
}
