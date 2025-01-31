/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { GRTFeedbackType, SystemIntakeRequestType, SystemIntakeStatus, SystemIntakeDocumentCommonType, SystemIntakeDocumentStatus } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL query operation: GetSystemIntake
// ====================================================

export interface GetSystemIntake_systemIntake_businessOwner {
  __typename: "SystemIntakeBusinessOwner";
  component: string | null;
  name: string | null;
}

export interface GetSystemIntake_systemIntake_contract_endDate {
  __typename: "ContractDate";
  day: string | null;
  month: string | null;
  year: string | null;
}

export interface GetSystemIntake_systemIntake_contract_startDate {
  __typename: "ContractDate";
  day: string | null;
  month: string | null;
  year: string | null;
}

export interface GetSystemIntake_systemIntake_contract {
  __typename: "SystemIntakeContract";
  contractor: string | null;
  endDate: GetSystemIntake_systemIntake_contract_endDate;
  hasContract: string | null;
  startDate: GetSystemIntake_systemIntake_contract_startDate;
  vehicle: string | null;
  number: string | null;
}

export interface GetSystemIntake_systemIntake_costs {
  __typename: "SystemIntakeCosts";
  isExpectingIncrease: string | null;
  expectedIncreaseAmount: string | null;
}

export interface GetSystemIntake_systemIntake_annualSpending {
  __typename: "SystemIntakeAnnualSpending";
  currentAnnualSpending: string | null;
  plannedYearOneSpending: string | null;
}

export interface GetSystemIntake_systemIntake_grtFeedbacks {
  __typename: "GRTFeedback";
  feedback: string | null;
  feedbackType: GRTFeedbackType | null;
  createdAt: Time;
}

export interface GetSystemIntake_systemIntake_governanceTeams_teams {
  __typename: "SystemIntakeCollaborator";
  acronym: string;
  collaborator: string;
  key: string;
  label: string;
  name: string;
}

export interface GetSystemIntake_systemIntake_governanceTeams {
  __typename: "SystemIntakeGovernanceTeam";
  isPresent: boolean | null;
  teams: GetSystemIntake_systemIntake_governanceTeams_teams[] | null;
}

export interface GetSystemIntake_systemIntake_isso {
  __typename: "SystemIntakeISSO";
  isPresent: boolean | null;
  name: string | null;
}

export interface GetSystemIntake_systemIntake_fundingSources {
  __typename: "SystemIntakeFundingSource";
  source: string | null;
  fundingNumber: string | null;
}

export interface GetSystemIntake_systemIntake_productManager {
  __typename: "SystemIntakeProductManager";
  component: string | null;
  name: string | null;
}

export interface GetSystemIntake_systemIntake_requester {
  __typename: "SystemIntakeRequester";
  component: string | null;
  email: string | null;
  name: string;
}

export interface GetSystemIntake_systemIntake_lastAdminNote {
  __typename: "LastAdminNote";
  content: string | null;
  createdAt: Time | null;
}

export interface GetSystemIntake_systemIntake_documents_documentType {
  __typename: "SystemIntakeDocumentType";
  commonType: SystemIntakeDocumentCommonType;
  otherTypeDescription: string | null;
}

export interface GetSystemIntake_systemIntake_documents {
  __typename: "SystemIntakeDocument";
  documentType: GetSystemIntake_systemIntake_documents_documentType;
  id: UUID;
  fileName: string;
  status: SystemIntakeDocumentStatus;
  uploadedAt: Time;
  url: string;
}

export interface GetSystemIntake_systemIntake {
  __typename: "SystemIntake";
  id: UUID;
  adminLead: string | null;
  businessNeed: string | null;
  businessSolution: string | null;
  businessOwner: GetSystemIntake_systemIntake_businessOwner;
  contract: GetSystemIntake_systemIntake_contract;
  costs: GetSystemIntake_systemIntake_costs | null;
  annualSpending: GetSystemIntake_systemIntake_annualSpending | null;
  currentStage: string | null;
  decisionNextSteps: string | null;
  grbDate: Time | null;
  grtDate: Time | null;
  grtFeedbacks: GetSystemIntake_systemIntake_grtFeedbacks[];
  governanceTeams: GetSystemIntake_systemIntake_governanceTeams;
  isso: GetSystemIntake_systemIntake_isso;
  existingFunding: boolean | null;
  fundingSources: GetSystemIntake_systemIntake_fundingSources[];
  lcid: string | null;
  lcidExpiresAt: Time | null;
  lcidScope: string | null;
  lcidCostBaseline: string | null;
  needsEaSupport: boolean | null;
  productManager: GetSystemIntake_systemIntake_productManager;
  rejectionReason: string | null;
  requester: GetSystemIntake_systemIntake_requester;
  requestName: string | null;
  requestType: SystemIntakeRequestType;
  status: SystemIntakeStatus;
  grtReviewEmailBody: string | null;
  decidedAt: Time | null;
  businessCaseId: UUID | null;
  submittedAt: Time | null;
  updatedAt: Time | null;
  createdAt: Time;
  archivedAt: Time | null;
  euaUserId: string;
  lastAdminNote: GetSystemIntake_systemIntake_lastAdminNote;
  hasUiChanges: boolean | null;
  documents: GetSystemIntake_systemIntake_documents[];
}

export interface GetSystemIntake {
  systemIntake: GetSystemIntake_systemIntake | null;
}

export interface GetSystemIntakeVariables {
  id: UUID;
}
