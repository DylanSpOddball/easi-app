package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

// SystemIntakeStatus represents the status of a system intake
type SystemIntakeStatus string

// SystemIntakeRequestType represents the type of a system intake
type SystemIntakeRequestType string

// SystemIntakeStatusFilter represents a filter in GETting system intakes
type SystemIntakeStatusFilter string

const (
	// SystemIntakeStatusINTAKEDRAFT captures enum value "INTAKE_DRAFT"
	SystemIntakeStatusINTAKEDRAFT SystemIntakeStatus = "INTAKE_DRAFT"
	// SystemIntakeStatusINTAKESUBMITTED captures enum value "INTAKE_SUBMITTED"
	SystemIntakeStatusINTAKESUBMITTED SystemIntakeStatus = "INTAKE_SUBMITTED"
	// SystemIntakeStatusNEEDBIZCASE captures enum value "NEED_BIZ_CASE"
	SystemIntakeStatusNEEDBIZCASE SystemIntakeStatus = "NEED_BIZ_CASE"
	// SystemIntakeStatusCLOSED captures enum value "CLOSED"
	SystemIntakeStatusCLOSED SystemIntakeStatus = "CLOSED"
	// SystemIntakeStatusAPPROVED captures enum value "APPROVED"
	SystemIntakeStatusAPPROVED SystemIntakeStatus = "APPROVED"
	// SystemIntakeStatusREADYFORGRT captures enum value "READY_FOR_GRT"
	SystemIntakeStatusREADYFORGRT SystemIntakeStatus = "READY_FOR_GRT"
	// SystemIntakeStatusREADYFORGRB captures enum value "READY_FOR_GRB"
	SystemIntakeStatusREADYFORGRB SystemIntakeStatus = "READY_FOR_GRB"
	// SystemIntakeStatusWITHDRAWN captures enum value "WITHDRAWN"
	SystemIntakeStatusWITHDRAWN SystemIntakeStatus = "WITHDRAWN"
	// SystemIntakeStatusNOTITREQUEST captures enum value "NOT_IT_REQUEST"
	SystemIntakeStatusNOTITREQUEST SystemIntakeStatus = "NOT_IT_REQUEST"
	// SystemIntakeStatusLCIDISSUED captures enum value "LCID_ISSUED"
	SystemIntakeStatusLCIDISSUED SystemIntakeStatus = "LCID_ISSUED"
	// SystemIntakeStatusBIZCASEDRAFT captures enum value "BIZ_CASE_DRAFT"
	SystemIntakeStatusBIZCASEDRAFT SystemIntakeStatus = "BIZ_CASE_DRAFT"
	// SystemIntakeStatusBIZCASEDRAFTSUBMITTED captures enum value "BIZ_CASE_DRAFT_SUBMITTED"
	SystemIntakeStatusBIZCASEDRAFTSUBMITTED SystemIntakeStatus = "BIZ_CASE_DRAFT_SUBMITTED"
	// SystemIntakeStatusBIZCASEFINALSUBMITTED captures enum value "BIZ_CASE_FINAL_SUBMITTED"
	SystemIntakeStatusBIZCASEFINALSUBMITTED SystemIntakeStatus = "BIZ_CASE_FINAL_SUBMITTED"
	// SystemIntakeStatusBIZCASECHANGESNEEDED captures enum value "BIZ_CASE_CHANGES_NEEDED"
	SystemIntakeStatusBIZCASECHANGESNEEDED SystemIntakeStatus = "BIZ_CASE_CHANGES_NEEDED"
	// SystemIntakeStatusBIZCASEFINALNEEDED captures enum value "BIZ_CASE_FINAL_NEEDED"
	SystemIntakeStatusBIZCASEFINALNEEDED SystemIntakeStatus = "BIZ_CASE_FINAL_NEEDED"
	// SystemIntakeStatusNOTAPPROVED captures enum value "NOT_APPROVED"
	SystemIntakeStatusNOTAPPROVED SystemIntakeStatus = "NOT_APPROVED"
	// SystemIntakeStatusNOGOVERNANCE captures enum value "NO_GOVERNANCE"
	SystemIntakeStatusNOGOVERNANCE SystemIntakeStatus = "NO_GOVERNANCE"
	// SystemIntakeStatusSHUTDOWNINPROGRESS captures enum value "SHUTDOWN_IN_PROGRESS"
	SystemIntakeStatusSHUTDOWNINPROGRESS SystemIntakeStatus = "SHUTDOWN_IN_PROGRESS"
	// SystemIntakeStatusSHUTDOWNCOMPLETE captures enum value "SHUTDOWN_COMPLETE"
	SystemIntakeStatusSHUTDOWNCOMPLETE SystemIntakeStatus = "SHUTDOWN_COMPLETE"

	// SystemIntakeRequestTypeNEW captures enum value of "NEW"
	SystemIntakeRequestTypeNEW SystemIntakeRequestType = "NEW"
	// SystemIntakeRequestTypeMAJORCHANGES captures enum value of "MAJOR_CHANGES"
	SystemIntakeRequestTypeMAJORCHANGES SystemIntakeRequestType = "MAJOR_CHANGES"
	// SystemIntakeRequestTypeRECOMPETE captures enum value of "RECOMPETE"
	SystemIntakeRequestTypeRECOMPETE SystemIntakeRequestType = "RECOMPETE"
	// SystemIntakeRequestTypeSHUTDOWN captures enum value of "SHUTDOWN"
	SystemIntakeRequestTypeSHUTDOWN SystemIntakeRequestType = "SHUTDOWN"

	// SystemIntakeStatusFilterOPEN captures enum value "OPEN"
	SystemIntakeStatusFilterOPEN SystemIntakeStatusFilter = "OPEN"
	// SystemIntakeStatusFilterCLOSED captures enum value "CLOSED"
	SystemIntakeStatusFilterCLOSED SystemIntakeStatusFilter = "CLOSED"
)

// SystemIntakeState represents whether the intake is open or closed
type SystemIntakeState string

const (
	// SystemIntakeStateOPEN captures enum value "OPEN"
	SystemIntakeStateOPEN SystemIntakeState = "OPEN"
	// SystemIntakeStateCLOSED captures enum value "CLOSED"
	SystemIntakeStateCLOSED SystemIntakeState = "CLOSED"
)

// SystemIntakeStep represents the current step in the intake process
type SystemIntakeStep string

const (
	// SystemIntakeStepINITIALFORM captures enum value "INITIAL_REQUEST_FORM"
	SystemIntakeStepINITIALFORM SystemIntakeStep = "INITIAL_REQUEST_FORM"
	// SystemIntakeStepDRAFTBIZCASE captures enum value "DRAFT_BUSINESS_CASE"
	SystemIntakeStepDRAFTBIZCASE SystemIntakeStep = "DRAFT_BUSINESS_CASE"
	// SystemIntakeStepGRTMEETING captures enum value "GRT_MEETING"
	SystemIntakeStepGRTMEETING SystemIntakeStep = "GRT_MEETING"
	// SystemIntakeStepFINALBIZCASE captures enum value "FINAL_BUSINESS_CASE"
	SystemIntakeStepFINALBIZCASE SystemIntakeStep = "FINAL_BUSINESS_CASE"
	// SystemIntakeStepGRBMEETING captures enum value "GRB_MEETING"
	SystemIntakeStepGRBMEETING SystemIntakeStep = "GRB_MEETING"
	// SystemIntakeStepDECISION captures enum value "DECISION_AND_NEXT_STEPS"
	SystemIntakeStepDECISION SystemIntakeStep = "DECISION_AND_NEXT_STEPS"
)

// SystemIntake is the model for the system intake form
type SystemIntake struct {
	ID                          uuid.UUID                    `json:"id"`
	EUAUserID                   null.String                  `json:"euaUserId" db:"eua_user_id"`
	Status                      SystemIntakeStatus           `json:"status"`
	State                       SystemIntakeState            `json:"state" db:"state"`
	Step                        SystemIntakeStep             `json:"step" db:"step"`
	RequestType                 SystemIntakeRequestType      `json:"requestType" db:"request_type"`
	Requester                   string                       `json:"requester"`
	Component                   null.String                  `json:"component"`
	BusinessOwner               null.String                  `json:"businessOwner" db:"business_owner"`
	BusinessOwnerComponent      null.String                  `json:"businessOwnerComponent" db:"business_owner_component"`
	ProductManager              null.String                  `json:"productManager" db:"product_manager"`
	ProductManagerComponent     null.String                  `json:"productManagerComponent" db:"product_manager_component"`
	ISSO                        null.String                  `json:"isso"`
	ISSOName                    null.String                  `json:"issoName" db:"isso_name"`
	TRBCollaborator             null.String                  `json:"trbCollaborator" db:"trb_collaborator"`
	TRBCollaboratorName         null.String                  `json:"trbCollaboratorName" db:"trb_collaborator_name"`
	OITSecurityCollaborator     null.String                  `json:"oitSecurityCollaborator" db:"oit_security_collaborator"`
	OITSecurityCollaboratorName null.String                  `json:"oitSecurityCollaboratorName" db:"oit_security_collaborator_name"`
	EACollaborator              null.String                  `json:"eaCollaborator" db:"ea_collaborator"`
	EACollaboratorName          null.String                  `json:"eaCollaboratorName" db:"ea_collaborator_name"`
	ProjectName                 null.String                  `json:"projectName" db:"project_name"`
	ProjectAcronym              null.String                  `json:"projectAcronym" db:"project_acronym"`
	BusinessNeed                null.String                  `json:"businessNeed" db:"business_need"`
	Solution                    null.String                  `json:"solution"`
	ProcessStatus               null.String                  `json:"processStatus" db:"process_status"`
	EASupportRequest            null.Bool                    `json:"eaSupportRequest" db:"ea_support_request"`
	ExistingContract            null.String                  `json:"existingContract" db:"existing_contract"`
	CostIncrease                null.String                  `json:"costIncrease" db:"cost_increase"`
	CostIncreaseAmount          null.String                  `json:"costIncreaseAmount" db:"cost_increase_amount"`
	CurrentAnnualSpending       null.String                  `json:"currentAnnualSpending" db:"current_annual_spending"`
	PlannedYearOneSpending      null.String                  `json:"plannedYearOneSpending" db:"planned_year_one_spending"`
	Contractor                  null.String                  `json:"contractor" db:"contractor"`
	ContractVehicle             null.String                  `json:"contractVehicle" db:"contract_vehicle"`
	ContractNumber              null.String                  `json:"contractNumber" db:"contract_number"` // replaces contract vehicle - see EASI-1977
	ContractStartDate           *time.Time                   `json:"contractStartDate" db:"contract_start_date"`
	ContractStartMonth          null.String                  `json:"contractStartMonth" db:"contract_start_month"`
	ContractStartYear           null.String                  `json:"contractStartYear" db:"contract_start_year"`
	ContractEndDate             *time.Time                   `json:"contractEndDate" db:"contract_end_date"`
	ContractEndMonth            null.String                  `json:"contractEndMonth" db:"contract_end_month"`
	ContractEndYear             null.String                  `json:"contractEndYear" db:"contract_end_year"`
	CreatedAt                   *time.Time                   `json:"createdAt" db:"created_at"`
	UpdatedAt                   *time.Time                   `json:"updatedAt" db:"updated_at"`
	SubmittedAt                 *time.Time                   `json:"submittedAt" db:"submitted_at"`
	DecidedAt                   *time.Time                   `json:"decidedAt" db:"decided_at"`
	ArchivedAt                  *time.Time                   `json:"archivedAt" db:"archived_at"`
	GRTDate                     *time.Time                   `json:"grtDate" db:"grt_date"`
	GRBDate                     *time.Time                   `json:"grbDate" db:"grb_date"`
	AlfabetID                   null.String                  `json:"alfabetID" db:"alfabet_id"`
	GrtReviewEmailBody          null.String                  `json:"grtReviewEmailBody" db:"grt_review_email_body"`
	RequesterEmailAddress       null.String                  `json:"requesterEmailAddress" db:"requester_email_address"`
	BusinessCaseID              *uuid.UUID                   `json:"businessCase" db:"business_case_id"`
	LifecycleID                 null.String                  `json:"lcid" db:"lcid"`
	LifecycleExpiresAt          *time.Time                   `json:"lcidExpiresAt" db:"lcid_expires_at" gqlgen:"lcidExpiresAt"`
	LifecycleScope              null.String                  `json:"lcidScope" db:"lcid_scope"`
	LifecycleCostBaseline       null.String                  `json:"lcidCostBaseline" db:"lcid_cost_baseline"`
	LifecycleExpirationAlertTS  *time.Time                   `json:"lcidExpirationAlertTS" db:"lcid_expiration_alert_ts"`
	DecisionNextSteps           null.String                  `json:"decisionNextSteps" db:"decision_next_steps"`
	RejectionReason             null.String                  `json:"rejectionReason" db:"rejection_reason"`
	AdminLead                   null.String                  `json:"adminLead" db:"admin_lead"`
	LastAdminNoteContent        null.String                  `json:"lastAdminNoteContent" db:"last_admin_note_content"`      // TODO break this out into it's own resolver, as this isn't actually a stored column in the DB
	LastAdminNoteCreatedAt      *time.Time                   `json:"lastAdminNoteCreatedAt" db:"last_admin_note_created_at"` // TODO break this out into it's own resolver, as this isn't actually a stored column in the DB
	CedarSystemID               null.String                  `json:"cedarSystemId" db:"cedar_system_id"`
	ExistingFunding             null.Bool                    `json:"existingFunding" db:"existing_funding"`
	FundingSource               null.String                  `json:"fundingSource" db:"funding_source"`
	FundingNumber               null.String                  `json:"fundingNumber" db:"funding_number"`
	FundingSources              []*SystemIntakeFundingSource `json:"fundingSources"`
	HasUIChanges                null.Bool                    `json:"hasUiChanges" db:"has_ui_changes"`
	RequestFormState            SystemIntakeFormState        `json:"requestFormState" db:"request_form_state"`
	DraftBusinessCaseState      SystemIntakeFormState        `json:"draftBusinessCaseState" db:"draft_business_case_state"`
	FinalBusinessCaseState      SystemIntakeFormState        `json:"finalBusinessCaseState" db:"final_business_case_state"`
	DecisionState               SystemIntakeDecisionState    `json:"decisionState" db:"decision_state"`
}

// SystemIntakes is a list of System Intakes
type SystemIntakes []SystemIntake

// SystemIntakeMetrics is a model for storing metrics related to system intake
type SystemIntakeMetrics struct {
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	Started            int       `json:"started"`
	CompletedOfStarted int       `json:"completedOfStarted"`
	Completed          int       `json:"completed"`
	Funded             int       `json:"funded"`
}

// GetStatusesByFilter returns a list of status corresponding to a /system_intakes/ filter
func GetStatusesByFilter(filter SystemIntakeStatusFilter) ([]SystemIntakeStatus, error) {
	switch filter {
	case SystemIntakeStatusFilterOPEN:
		return []SystemIntakeStatus{
			SystemIntakeStatusINTAKESUBMITTED,
			SystemIntakeStatusNEEDBIZCASE,
			SystemIntakeStatusBIZCASEDRAFT,
			SystemIntakeStatusBIZCASEDRAFTSUBMITTED,
			SystemIntakeStatusBIZCASECHANGESNEEDED,
			SystemIntakeStatusBIZCASEFINALNEEDED,
			SystemIntakeStatusBIZCASEFINALSUBMITTED,
			SystemIntakeStatusREADYFORGRT,
			SystemIntakeStatusREADYFORGRB,
			SystemIntakeStatusSHUTDOWNINPROGRESS,
		}, nil
	case SystemIntakeStatusFilterCLOSED:
		return []SystemIntakeStatus{
			SystemIntakeStatusLCIDISSUED,
			SystemIntakeStatusWITHDRAWN,
			SystemIntakeStatusNOTITREQUEST,
			SystemIntakeStatusNOTAPPROVED,
			SystemIntakeStatusNOGOVERNANCE,
			SystemIntakeStatusSHUTDOWNCOMPLETE,
		}, nil
	default:
		return []SystemIntakeStatus{}, errors.New("unexpected system intake status filter name")
	}
}

// SystemIntakeFormState represents the possible states of of any System Intake form types.
type SystemIntakeFormState string

// These are the options for SystemIntakeRequestFormState
const (
	SIRFSNotStarted     SystemIntakeFormState = "NOT_STARTED"
	SIRFSInProgress     SystemIntakeFormState = "IN_PROGRESS"
	SIRFSEditsRequested SystemIntakeFormState = "EDITS_REQUESTED"
	SIRFSSubmitted      SystemIntakeFormState = "SUBMITTED"
)

// SystemIntakeDecisionState represents the types of SystemIntakeDecisionState types.
type SystemIntakeDecisionState string

// These are the options for SystemIntakeDecisionState
const (
	SIDSNoDecision   SystemIntakeDecisionState = "NO_DECISION"
	SIDSLcidIssued   SystemIntakeDecisionState = "LCID_ISSUED"
	SIDSNotApproved  SystemIntakeDecisionState = "NOT_APPROVED"
	SIDSNoGovernance SystemIntakeDecisionState = "NO_GOVERNANCE"
)

// SystemIntakeMeetingState is the state for any meeting for a system intake
type SystemIntakeMeetingState string

// These are the options for SystemIntakeMeetingState
const (
	SIMSScheduled    SystemIntakeMeetingState = "SCHEDULED"
	SIMSNotScheduled SystemIntakeMeetingState = "NOT_SCHEDULED"
)

// GRTMeetingState returns if a GRTMeeting has been scheduled or not
func (si *SystemIntake) GRTMeetingState() SystemIntakeMeetingState {
	return isMeetingScheduled(si.GRTDate)
}

// GRBMeetingState returns if a GRBMeeting has been scheduled or not
func (si *SystemIntake) GRBMeetingState() SystemIntakeMeetingState {
	return isMeetingScheduled(si.GRBDate)
}

func isMeetingScheduled(date *time.Time) SystemIntakeMeetingState {
	if date == nil {
		return SIMSNotScheduled
	}
	return SIMSScheduled
}
