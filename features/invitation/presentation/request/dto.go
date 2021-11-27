package request

import "workuo/features/invitation"

type InvitationRequest struct {
	UserID      uint `json: "userid"`
	JobID       uint `json: "jobid"`
	RecruiterID uint
	Role        string
}

func ToCore(data InvitationRequest) invitation.InvitationCore {
	return invitation.InvitationCore{
		UserID:      data.UserID,
		JobID:       data.JobID,
		RecruiterID: data.RecruiterID,
		Role:        data.Role,
	}
}
