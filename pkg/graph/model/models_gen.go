// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/google/uuid"
)

// Denotes type of a document that is attached to a 508/accessibility request,
// which can be one of a number of common types, or another user-specified type
type AccessibilityRequestDocumentType struct {
	CommonType           models.AccessibilityRequestDocumentCommonType `json:"commonType"`
	OtherTypeDescription *string                                       `json:"otherTypeDescription"`
}

type AccessibilityRequestEdge struct {
	Node *models.AccessibilityRequest `json:"node"`
}

type AccessibilityRequestsConnection struct {
	Edges []*AccessibilityRequestEdge `json:"edges"`
}

// Feedback intended for a business owner before they proceed to writing a
// business case for a system request
type AddGRTFeedbackInput struct {
	EmailBody string    `json:"emailBody"`
	Feedback  string    `json:"feedback"`
	IntakeID  uuid.UUID `json:"intakeID"`
}

// Payload for adding GRT feedback to a system request (contains the system
// request ID)
type AddGRTFeedbackPayload struct {
	ID *uuid.UUID `json:"id"`
}

// Input to add feedback to a system request
type BasicActionInput struct {
	Feedback string    `json:"feedback"`
	IntakeID uuid.UUID `json:"intakeId"`
}

// An IT governance requester's explanation of alternative solutions
// to their system, which involve leaving the system "as-is"
type BusinessCaseAsIsSolution struct {
	Cons        *string `json:"cons"`
	CostSavings *string `json:"costSavings"`
	Pros        *string `json:"pros"`
	Summary     *string `json:"summary"`
	Title       *string `json:"title"`
}

// A solution proposal within a business case
type BusinessCaseSolution struct {
	AcquisitionApproach     *string `json:"acquisitionApproach"`
	Cons                    *string `json:"cons"`
	CostSavings             *string `json:"costSavings"`
	HasUI                   *string `json:"hasUi"`
	HostingCloudServiceType *string `json:"hostingCloudServiceType"`
	HostingLocation         *string `json:"hostingLocation"`
	HostingType             *string `json:"hostingType"`
	Pros                    *string `json:"pros"`
	SecurityIsApproved      *bool   `json:"securityIsApproved"`
	SecurityIsBeingReviewed *string `json:"securityIsBeingReviewed"`
	Summary                 *string `json:"summary"`
	Title                   *string `json:"title"`
}

// Represents a date used for start and end dates on a contract
type ContractDate struct {
	Day   *string `json:"day"`
	Month *string `json:"month"`
	Year  *string `json:"year"`
}

// The input data used for adding a document to a 508/accessibility request
type CreateAccessibilityRequestDocumentInput struct {
	CommonDocumentType           models.AccessibilityRequestDocumentCommonType `json:"commonDocumentType"`
	MimeType                     string                                        `json:"mimeType"`
	Name                         string                                        `json:"name"`
	OtherDocumentTypeDescription *string                                       `json:"otherDocumentTypeDescription"`
	RequestID                    uuid.UUID                                     `json:"requestID"`
	Size                         int                                           `json:"size"`
	URL                          string                                        `json:"url"`
}

// The payload containing the input data used for adding a document to a
// 508/accessibility request
type CreateAccessibilityRequestDocumentPayload struct {
	AccessibilityRequestDocument *models.AccessibilityRequestDocument `json:"accessibilityRequestDocument"`
	UserErrors                   []*UserError                         `json:"userErrors"`
}

// The data needed to initialize a 508/accessibility request
type CreateAccessibilityRequestInput struct {
	IntakeID uuid.UUID `json:"intakeID"`
	Name     string    `json:"name"`
}

// The data used when adding a note to a 508/accessibility request
type CreateAccessibilityRequestNoteInput struct {
	RequestID       uuid.UUID `json:"requestID"`
	Note            string    `json:"note"`
	ShouldSendEmail bool      `json:"shouldSendEmail"`
}

// The payload for adding a note to a 508/accessibility request
type CreateAccessibilityRequestNotePayload struct {
	AccessibilityRequestNote *models.AccessibilityRequestNote `json:"accessibilityRequestNote"`
	UserErrors               []*UserError                     `json:"userErrors"`
}

// The payload containing the data needed to initialize an AccessibilityRequest
type CreateAccessibilityRequestPayload struct {
	AccessibilityRequest *models.AccessibilityRequest `json:"accessibilityRequest"`
	UserErrors           []*UserError                 `json:"userErrors"`
}

// The data needed to bookmark a cedar system
type CreateCedarSystemBookmarkInput struct {
	CedarSystemID uuid.UUID `json:"cedarSystemId"`
}

// The payload when bookmarking a cedar system
type CreateCedarSystemBookmarkPayload struct {
	CedarSystemBookmark *models.CedarSystemBookmark `json:"cedarSystemBookmark"`
	UserErrors          []*UserError                `json:"userErrors"`
}

// Input data for extending a system request's lifecycle ID
type CreateSystemIntakeActionExtendLifecycleIDInput struct {
	ID             uuid.UUID  `json:"id"`
	ExpirationDate *time.Time `json:"expirationDate"`
}

// Payload data for extending a system request's lifecycle ID
type CreateSystemIntakeActionExtendLifecycleIDPayload struct {
	SystemIntake *models.SystemIntake `json:"systemIntake"`
	UserErrors   []*UserError         `json:"userErrors"`
}

// The input data used to initialize an IT governance request for a system
type CreateSystemIntakeInput struct {
	RequestType models.SystemIntakeRequestType `json:"requestType"`
	Requester   *SystemIntakeRequesterInput    `json:"requester"`
}

// Input data for adding a note to a system request
type CreateSystemIntakeNoteInput struct {
	Content    string    `json:"content"`
	AuthorName string    `json:"authorName"`
	IntakeID   uuid.UUID `json:"intakeId"`
}

// The input required to add a test date/score to a 508/accessibility request
type CreateTestDateInput struct {
	Date      time.Time               `json:"date"`
	RequestID uuid.UUID               `json:"requestID"`
	Score     *int                    `json:"score"`
	TestType  models.TestDateTestType `json:"testType"`
}

// The payload for the input required to add a test date/score to a
// 508/accessibility request
type CreateTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

// The current user of the application
type CurrentUser struct {
	LaunchDarkly *LaunchDarklySettings `json:"launchDarkly"`
}

// The input used to delete a document from a 508/accessibility request
type DeleteAccessibilityRequestDocumentInput struct {
	ID uuid.UUID `json:"id"`
}

// The payload used to delete a document from a 508/accessibility request
type DeleteAccessibilityRequestDocumentPayload struct {
	ID *uuid.UUID `json:"id"`
}

// The input data needed to delete a 508/accessibility request
type DeleteAccessibilityRequestInput struct {
	ID     uuid.UUID                                 `json:"id"`
	Reason models.AccessibilityRequestDeletionReason `json:"reason"`
}

// The payload data sent when deleting a 508/accessibility request
type DeleteAccessibilityRequestPayload struct {
	ID         *uuid.UUID   `json:"id"`
	UserErrors []*UserError `json:"userErrors"`
}

// The payload when deleting a bookmark for a cedar system
type DeleteCedarSystemBookmarkPayload struct {
	UserErrors []*UserError `json:"userErrors"`
}

// The input required to delete a test date/score
type DeleteTestDateInput struct {
	ID uuid.UUID `json:"id"`
}

// The payload for the input required to delete a test date/score
type DeleteTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

// Input associated with a document to be uploaded to a 508/accessibility request
type GeneratePresignedUploadURLInput struct {
	FileName string `json:"fileName"`
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
}

// URL generated for a document to be uploaded to a 508/accessibility request
type GeneratePresignedUploadURLPayload struct {
	URL        *string      `json:"url"`
	UserErrors []*UserError `json:"userErrors"`
}

// The input data required to issue a lifecycle ID for a system's IT governance
// request
type IssueLifecycleIDInput struct {
	ExpiresAt    time.Time `json:"expiresAt"`
	Feedback     string    `json:"feedback"`
	IntakeID     uuid.UUID `json:"intakeId"`
	Lcid         *string   `json:"lcid"`
	NextSteps    *string   `json:"nextSteps"`
	Scope        string    `json:"scope"`
	CostBaseline *string   `json:"costBaseline"`
}

// The most recent note added by an admin to a system request
type LastAdminNote struct {
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
}

// The current user's Launch Darkly key
type LaunchDarklySettings struct {
	UserKey    string `json:"userKey"`
	SignedHash string `json:"signedHash"`
}

// Input data for rejection of a system's IT governance request
type RejectIntakeInput struct {
	Feedback  string    `json:"feedback"`
	IntakeID  uuid.UUID `json:"intakeId"`
	NextSteps *string   `json:"nextSteps"`
	Reason    string    `json:"reason"`
}

// Represents a request being made with the EASi system
type Request struct {
	ID              uuid.UUID   `json:"id"`
	Name            *string     `json:"name"`
	SubmittedAt     *time.Time  `json:"submittedAt"`
	Type            RequestType `json:"type"`
	Status          string      `json:"status"`
	StatusCreatedAt *time.Time  `json:"statusCreatedAt"`
	Lcid            *string     `json:"lcid"`
	NextMeetingDate *time.Time  `json:"nextMeetingDate"`
}

type RequestEdge struct {
	Node *Request `json:"node"`
}

type RequestsConnection struct {
	Edges []*RequestEdge `json:"edges"`
}

// Input to submit an intake for review
type SubmitIntakeInput struct {
	ID uuid.UUID `json:"id"`
}

type SystemConnection struct {
	Edges []*SystemEdge `json:"edges"`
}

type SystemEdge struct {
	Node *models.System `json:"node"`
}

// An action taken on a system intake, often resulting in a change in status.
type SystemIntakeAction struct {
	ID                   uuid.UUID                         `json:"id"`
	SystemIntake         *models.SystemIntake              `json:"systemIntake"`
	Type                 SystemIntakeActionType            `json:"type"`
	Actor                *SystemIntakeActionActor          `json:"actor"`
	Feedback             *string                           `json:"feedback"`
	LcidExpirationChange *SystemIntakeLCIDExpirationChange `json:"lcidExpirationChange"`
	CreatedAt            time.Time                         `json:"createdAt"`
}

// The contact who is associated with an action being done to a system request
type SystemIntakeActionActor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Represents the OIT business owner of a system
type SystemIntakeBusinessOwner struct {
	Component *string `json:"component"`
	Name      *string `json:"name"`
}

// The input data used to set the CMS business owner of a system
type SystemIntakeBusinessOwnerInput struct {
	Name      string `json:"name"`
	Component string `json:"component"`
}

// Represents a contact in OIT who is collaborating with the user
// creating a system IT governance request
type SystemIntakeCollaborator struct {
	Acronym      string `json:"acronym"`
	Collaborator string `json:"collaborator"`
	Key          string `json:"key"`
	Label        string `json:"label"`
	Name         string `json:"name"`
}

// The input data used to add an OIT collaborator for a system request
type SystemIntakeCollaboratorInput struct {
	Collaborator string `json:"collaborator"`
	Name         string `json:"name"`
	Key          string `json:"key"`
}

// Represents a contract for work on a system
type SystemIntakeContract struct {
	Contractor  *string       `json:"contractor"`
	EndDate     *ContractDate `json:"endDate"`
	HasContract *string       `json:"hasContract"`
	StartDate   *ContractDate `json:"startDate"`
	Vehicle     *string       `json:"vehicle"`
}

// Input data containing information about a contract related to a system request
type SystemIntakeContractInput struct {
	Contractor  *string    `json:"contractor"`
	EndDate     *time.Time `json:"endDate"`
	HasContract *string    `json:"hasContract"`
	StartDate   *time.Time `json:"startDate"`
	Vehicle     *string    `json:"vehicle"`
}

// Represents expectations about a system's additional costs
type SystemIntakeCosts struct {
	ExpectedIncreaseAmount *string `json:"expectedIncreaseAmount"`
	IsExpectingIncrease    *string `json:"isExpectingIncrease"`
}

// Input data for estimated system cost increases associated with a system request
type SystemIntakeCostsInput struct {
	ExpectedIncreaseAmount *string `json:"expectedIncreaseAmount"`
	IsExpectingIncrease    *string `json:"isExpectingIncrease"`
}

// Represents the source of funding for a system
type SystemIntakeFundingSource struct {
	FundingNumber *string `json:"fundingNumber"`
	IsFunded      *bool   `json:"isFunded"`
	Source        *string `json:"source"`
}

// Input data detailing how a system is funded
type SystemIntakeFundingSourceInput struct {
	FundingNumber *string `json:"fundingNumber"`
	IsFunded      *bool   `json:"isFunded"`
	Source        *string `json:"source"`
}

// Contains multiple system request collaborators, if any
type SystemIntakeGovernanceTeam struct {
	IsPresent *bool                       `json:"isPresent"`
	Teams     []*SystemIntakeCollaborator `json:"teams"`
}

// The input data used to set the list of OIT collaborators for a system request
type SystemIntakeGovernanceTeamInput struct {
	IsPresent *bool                            `json:"isPresent"`
	Teams     []*SystemIntakeCollaboratorInput `json:"teams"`
}

// The Information System Security Officer (ISSO) that is
// assicuated with a system request, if any
type SystemIntakeIsso struct {
	IsPresent *bool   `json:"isPresent"`
	Name      *string `json:"name"`
}

// The input data used to set the ISSO associated with a system request, if any
type SystemIntakeISSOInput struct {
	IsPresent *bool   `json:"isPresent"`
	Name      *string `json:"name"`
}

// Contains the data needed to change the expiration date of a system request's
// lifecycle ID
type SystemIntakeLCIDExpirationChange struct {
	PreviousDate time.Time `json:"previousDate"`
	NewDate      time.Time `json:"newDate"`
}

// A note added to a system request
type SystemIntakeNote struct {
	Author    *SystemIntakeNoteAuthor `json:"author"`
	Content   string                  `json:"content"`
	CreatedAt time.Time               `json:"createdAt"`
	ID        uuid.UUID               `json:"id"`
}

// The author of a note added to a system request
type SystemIntakeNoteAuthor struct {
	Eua  string `json:"eua"`
	Name string `json:"name"`
}

// The product manager associated with a system
type SystemIntakeProductManager struct {
	Component *string `json:"component"`
	Name      *string `json:"name"`
}

// The input data used to set the CMS product manager/lead of a system
type SystemIntakeProductManagerInput struct {
	Name      string `json:"name"`
	Component string `json:"component"`
}

// The contact who made an IT governance request for a system
type SystemIntakeRequester struct {
	Component *string `json:"component"`
	Email     *string `json:"email"`
	Name      string  `json:"name"`
}

// The input data used to set the requester of a system request
type SystemIntakeRequesterInput struct {
	Name string `json:"name"`
}

// The input data used to set the requester for a system request along with the
// requester's business component
type SystemIntakeRequesterWithComponentInput struct {
	Name      string `json:"name"`
	Component string `json:"component"`
}

// Parameters for updating a 508/accessibility request's status
type UpdateAccessibilityRequestStatus struct {
	RequestID uuid.UUID                         `json:"requestID"`
	Status    models.AccessibilityRequestStatus `json:"status"`
}

// Result of updating a 508/accessibility request's status
type UpdateAccessibilityRequestStatusPayload struct {
	ID         uuid.UUID                         `json:"id"`
	RequestID  uuid.UUID                         `json:"requestID"`
	Status     models.AccessibilityRequestStatus `json:"status"`
	EuaUserID  string                            `json:"euaUserId"`
	UserErrors []*UserError                      `json:"userErrors"`
}

// Input data used to update the admin lead assigned to a system IT governance
// request
type UpdateSystemIntakeAdminLeadInput struct {
	AdminLead string    `json:"adminLead"`
	ID        uuid.UUID `json:"id"`
}

// The input data used to update the contact details of the people associated with
// a system request
type UpdateSystemIntakeContactDetailsInput struct {
	ID              uuid.UUID                                `json:"id"`
	Requester       *SystemIntakeRequesterWithComponentInput `json:"requester"`
	BusinessOwner   *SystemIntakeBusinessOwnerInput          `json:"businessOwner"`
	ProductManager  *SystemIntakeProductManagerInput         `json:"productManager"`
	Isso            *SystemIntakeISSOInput                   `json:"isso"`
	GovernanceTeams *SystemIntakeGovernanceTeamInput         `json:"governanceTeams"`
}

// Input data for updating contract details related to a system request
type UpdateSystemIntakeContractDetailsInput struct {
	ID            uuid.UUID                       `json:"id"`
	CurrentStage  *string                         `json:"currentStage"`
	FundingSource *SystemIntakeFundingSourceInput `json:"fundingSource"`
	Costs         *SystemIntakeCostsInput         `json:"costs"`
	Contract      *SystemIntakeContractInput      `json:"contract"`
}

// The payload for updating a system's IT governance request
type UpdateSystemIntakePayload struct {
	SystemIntake *models.SystemIntake `json:"systemIntake"`
	UserErrors   []*UserError         `json:"userErrors"`
}

// Input to update some fields on a system request
type UpdateSystemIntakeRequestDetailsInput struct {
	ID               uuid.UUID `json:"id"`
	RequestName      *string   `json:"requestName"`
	BusinessNeed     *string   `json:"businessNeed"`
	BusinessSolution *string   `json:"businessSolution"`
	NeedsEaSupport   *bool     `json:"needsEaSupport"`
}

// Input data used to update GRT and GRB dates for a system request
type UpdateSystemIntakeReviewDatesInput struct {
	GrbDate *time.Time `json:"grbDate"`
	GrtDate *time.Time `json:"grtDate"`
	ID      uuid.UUID  `json:"id"`
}

// The input required to update a test date/score
type UpdateTestDateInput struct {
	Date     time.Time               `json:"date"`
	ID       uuid.UUID               `json:"id"`
	Score    *int                    `json:"score"`
	TestType models.TestDateTestType `json:"testType"`
}

// The payload for the input required to update a test date/score
type UpdateTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

// UserError represents application-level errors that are the result of
// either user or application developer error.
type UserError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

// Indicates the type of a request being made with the EASi system
type RequestType string

const (
	RequestTypeAccessibilityRequest RequestType = "ACCESSIBILITY_REQUEST"
	RequestTypeGovernanceRequest    RequestType = "GOVERNANCE_REQUEST"
)

var AllRequestType = []RequestType{
	RequestTypeAccessibilityRequest,
	RequestTypeGovernanceRequest,
}

func (e RequestType) IsValid() bool {
	switch e {
	case RequestTypeAccessibilityRequest, RequestTypeGovernanceRequest:
		return true
	}
	return false
}

func (e RequestType) String() string {
	return string(e)
}

func (e *RequestType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestType", str)
	}
	return nil
}

func (e RequestType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A user role associated with a job code
type Role string

const (
	// A 508 Tester
	RoleEasi508Tester Role = "EASI_508_TESTER"
	// A 508 request program team member
	RoleEasi508User Role = "EASI_508_USER"
	// A 508 request program team member or tester
	RoleEasi508TesterOrUser Role = "EASI_508_TESTER_OR_USER"
	// A member of the GRT
	RoleEasiGovteam Role = "EASI_GOVTEAM"
	// A generic EASi user
	RoleEasiUser Role = "EASI_USER"
)

var AllRole = []Role{
	RoleEasi508Tester,
	RoleEasi508User,
	RoleEasi508TesterOrUser,
	RoleEasiGovteam,
	RoleEasiUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleEasi508Tester, RoleEasi508User, RoleEasi508TesterOrUser, RoleEasiGovteam, RoleEasiUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Represents the type of an action that is being done to a system request
type SystemIntakeActionType string

const (
	SystemIntakeActionTypeBizCaseNeedsChanges            SystemIntakeActionType = "BIZ_CASE_NEEDS_CHANGES"
	SystemIntakeActionTypeCreateBizCase                  SystemIntakeActionType = "CREATE_BIZ_CASE"
	SystemIntakeActionTypeGUIDEReceivedClose             SystemIntakeActionType = "GUIDE_RECEIVED_CLOSE"
	SystemIntakeActionTypeExtendLcid                     SystemIntakeActionType = "EXTEND_LCID"
	SystemIntakeActionTypeIssueLcid                      SystemIntakeActionType = "ISSUE_LCID"
	SystemIntakeActionTypeNeedBizCase                    SystemIntakeActionType = "NEED_BIZ_CASE"
	SystemIntakeActionTypeNoGovernanceNeeded             SystemIntakeActionType = "NO_GOVERNANCE_NEEDED"
	SystemIntakeActionTypeNotItRequest                   SystemIntakeActionType = "NOT_IT_REQUEST"
	SystemIntakeActionTypeNotRespondingClose             SystemIntakeActionType = "NOT_RESPONDING_CLOSE"
	SystemIntakeActionTypeProvideFeedbackNeedBizCase     SystemIntakeActionType = "PROVIDE_FEEDBACK_NEED_BIZ_CASE"
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft SystemIntakeActionType = "PROVIDE_GRT_FEEDBACK_BIZ_CASE_DRAFT"
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal SystemIntakeActionType = "PROVIDE_GRT_FEEDBACK_BIZ_CASE_FINAL"
	SystemIntakeActionTypeReadyForGrb                    SystemIntakeActionType = "READY_FOR_GRB"
	SystemIntakeActionTypeReadyForGrt                    SystemIntakeActionType = "READY_FOR_GRT"
	SystemIntakeActionTypeReject                         SystemIntakeActionType = "REJECT"
	SystemIntakeActionTypeSendEmail                      SystemIntakeActionType = "SEND_EMAIL"
	SystemIntakeActionTypeSubmitBizCase                  SystemIntakeActionType = "SUBMIT_BIZ_CASE"
	SystemIntakeActionTypeSubmitFinalBizCase             SystemIntakeActionType = "SUBMIT_FINAL_BIZ_CASE"
	SystemIntakeActionTypeSubmitIntake                   SystemIntakeActionType = "SUBMIT_INTAKE"
)

var AllSystemIntakeActionType = []SystemIntakeActionType{
	SystemIntakeActionTypeBizCaseNeedsChanges,
	SystemIntakeActionTypeCreateBizCase,
	SystemIntakeActionTypeGUIDEReceivedClose,
	SystemIntakeActionTypeExtendLcid,
	SystemIntakeActionTypeIssueLcid,
	SystemIntakeActionTypeNeedBizCase,
	SystemIntakeActionTypeNoGovernanceNeeded,
	SystemIntakeActionTypeNotItRequest,
	SystemIntakeActionTypeNotRespondingClose,
	SystemIntakeActionTypeProvideFeedbackNeedBizCase,
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft,
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal,
	SystemIntakeActionTypeReadyForGrb,
	SystemIntakeActionTypeReadyForGrt,
	SystemIntakeActionTypeReject,
	SystemIntakeActionTypeSendEmail,
	SystemIntakeActionTypeSubmitBizCase,
	SystemIntakeActionTypeSubmitFinalBizCase,
	SystemIntakeActionTypeSubmitIntake,
}

func (e SystemIntakeActionType) IsValid() bool {
	switch e {
	case SystemIntakeActionTypeBizCaseNeedsChanges, SystemIntakeActionTypeCreateBizCase, SystemIntakeActionTypeGUIDEReceivedClose, SystemIntakeActionTypeExtendLcid, SystemIntakeActionTypeIssueLcid, SystemIntakeActionTypeNeedBizCase, SystemIntakeActionTypeNoGovernanceNeeded, SystemIntakeActionTypeNotItRequest, SystemIntakeActionTypeNotRespondingClose, SystemIntakeActionTypeProvideFeedbackNeedBizCase, SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft, SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal, SystemIntakeActionTypeReadyForGrb, SystemIntakeActionTypeReadyForGrt, SystemIntakeActionTypeReject, SystemIntakeActionTypeSendEmail, SystemIntakeActionTypeSubmitBizCase, SystemIntakeActionTypeSubmitFinalBizCase, SystemIntakeActionTypeSubmitIntake:
		return true
	}
	return false
}

func (e SystemIntakeActionType) String() string {
	return string(e)
}

func (e *SystemIntakeActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SystemIntakeActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SystemIntakeActionType", str)
	}
	return nil
}

func (e SystemIntakeActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
