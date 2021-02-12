// This is for the CMS 508 project flow

const accessibility = {
  documentTable: {
    caption: 'Documents uploaded for',
    header: {
      actions: 'Actions',
      documentName: 'Document',
      uploadedAt: 'Uploaded Date'
    },
    remove: 'Remove',
    view: 'View'
  },
  requestTable: {
    caption: 'Active 508 Requests',
    header: {
      requestName: 'Request Name',
      submissionDate: 'Submission Date',
      businessOwner: 'Business Owner',
      testDate: 'Test Date',
      pointOfContact: 'Point of Contact',
      status: 'Status'
    },
    lastUpdated: 'last updated on'
  },
  tabs: {
    accessibilityRequests: '508 Requests'
  },
  requestDetails: {
    documents: {
      label: 'Documents',
      none: 'No documents added to this request yet.'
    },
    documentUpload: 'Upload a document',
    other: 'Other request details',
    remove: 'Remove this request from EASi'
  },
  createTestDate: {
    addTestDateHeader: 'Add a test date for {{requestName}}',
    testTypeHeader: 'What type of test?',
    dateHeader: 'TestDate',
    dateHelpText: 'For example: 4 28 2020',
    scoreHeader: 'Does this test have a score?',
    scoreValueHeader: 'Test Score',
    scoreValueSRHelpText: 'Enter the test score without the percentage symbol',
    submitButton: 'Add date',
    cancel: "Don't add and go back to request page",
    confirmation: 'Test date {{date}} added to {{requestName}} page'
  }
};

export default accessibility;
