package model

type Moment struct {
	Id              string `bson:"_id,omitempty"`
	Title           string
	Link            string
	MediaUrl        string
	UserId          string
	IsVideo         bool
	DatetimeCreated string
}

type GeoLocation struct {
	Lat  float64
	Long float64
}

type GetMomentsOutput struct {
	Moments []*Moment
}

type GetMomentsReportedOutput struct {
	MomentIds []string
}

type GetMomentOutput struct {
	Moment *Moment
}

type GetRetailersNearbyInput struct {
	UserId      string
	GeoLocation *GeoLocation
	PageNumber  int32
	PageSize    int32
}

type GetRetailersBlockedOutput struct {
	UserIds []string
}

type CreateMomentInput struct {
	Moment *Moment `bson:"inline"`
}

type UpdateMomentInput struct {
	Moment *Moment
}

type SelectionOption struct {
	Id      string
	Name    string
	Options []*SelectionOption `bson:",omitempty"`
}

type MomentReported struct {
	ReportedUserId  string
	UserId          string
	MomentId        string
	Reason          string
	DatetimeCreated string
}

type CreateMomentReportedInput struct {
	MomentReported *MomentReported `bson:"inline"`
}

type RetailerBlocked struct {
	BlockedUserId   string
	UserId          string
	Period          int32
	DatetimeCreated string
}

type CreateRetailerBlockedInput struct {
	RetailerBlocked *RetailerBlocked `bson:"inline"`
}

type DeleteRetailerBlockedInput struct {
	BlockedUserId string
	UserId        string
}

type ShareContent struct {
	Text       string
	ImageUrl   string
	WebpageUrl string
}

type Task struct {
	Id              string `bson:"_id,omitempty"`
	Name            string
	PictureUrl      string
	DetailsUrl      string
	RedemptionUrl   string `bson:"redemptionurl,omitempty"`
	StartDate       string
	EndDate         string
	Type            int32
	Status          int32
	ShareContent    *ShareContent      `bson:",omitempty"`
	Participation   *TaskParticipation `bson:"participation,omitempty"`
	DatetimeCreated string
}

type TaskParticipation struct {
	Id                           string `bson:"_id"`
	UserId                       string
	TaskId                       string
	Status                       int32
	RedemptionCode               string
	DatetimeCreated              string
	DatetimeAchieved             string
	DatetimeRedeemed             string
	RedemptionConfirmationUserId string
}

type ContestEnrollment struct {
	Status int32
	Msg    string
}

type Contest struct {
	Id        string `bson:"_id,omitempty"`
	Name      string
	StartDate string
	EndDate   string
	PosterUrl string
	Active    bool
}

type CreateContestInput struct {
	Contest *Contest `bson:"inline"`
}

type UpdateContestInput struct {
	Contest *Contest
}

type GetContestsOutput struct {
	Contests []*Contest
}

type GetContestOutput struct {
	Contest *Contest
}

type VotingResult struct {
	Voted         bool  `bson:"voted"`
	NumberOfVotes int32 `bson:"numberofvotes"`
}

type CandidateFilter struct {
	EnrollmentStatus int32 // -1 all, 0 pending, 1 approved
	MatchingText     string
}

type ContestCandidate struct {
	Id                 string `bson:"_id,omitempty"`
	UserId             string
	ContestId          string
	PetId              string
	PetName            string
	PetOwnerName       string
	PetOwnerPhone      string
	PetPhotoUrl        string
	PetSnsUrl          string
	Enrollment         *ContestEnrollment
	VotingResult       *VotingResult `bson:"votingresult,omitempty"`
	Sequence           int32
	ExtraNumberOfVotes int32
	Position           int32
}

type GetContestCandidatesInput struct {
	ContestId  string
	UserId     string
	Filter     *CandidateFilter
	PageNumber int32
	PageSize   int32
}

type GetContestCandidatesOutput struct {
	Candidates []*ContestCandidate
}

type GetContestCandidatesByVoterUserIdOutput struct {
	Candidates []*ContestCandidate
}

type GetContestCandidateInput struct {
	ContestId  string
	UserId     string
	PageSize   int32
	PageNumber int32
}

type GetContestCandidateOutput struct {
	Candidate *ContestCandidate
}

type CreateContestCandidateInput struct {
	Candidate *ContestCandidate `bson:"inline"`
}

type UpdateContestCandidateInput struct {
	Candidate *ContestCandidate
}

type UpdateContestCandidateEnrollmentInput struct {
	Id         string
	Enrollment *ContestEnrollment
}

type UpdateContestCandidateExtraNumberOfVotesInput struct {
	Id                 string
	ExtraNumberOfVotes int32
}

type Vote struct {
	Id          string `bson:"_id,omitempty"`
	UserId      string
	CandidateId string
	PetId       string
	ContestId   string
}

type CreateVoteInput struct {
	Vote *Vote `bson:"inline"`
}

type GetVotesOutput struct {
	Votes []*Vote
}

type DeleteVoteInput struct {
	Vote *Vote
}

type Counter struct {
	ReferenceId string
	Sequence    int32
}
