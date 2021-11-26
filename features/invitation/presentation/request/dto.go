package request

import "workuo/features/invitation"

type InvitationRequest struct {
	UserID uint `json: "userid"`
	JobID  uint `json: "jobid"`
}

func ToCore(data InvitationRequest) invitation.InvitationCore {
	return invitation.InvitationCore{
		UserID: data.UserID,
		JobID:  data.JobID,
	}
}
