import React, { useMemo, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Cell, Column, useTable } from 'react-table';
import { useMutation } from '@apollo/client';
import { Button, Link, Table } from '@trussworks/react-uswds';
import { DeleteAccessibilityRequestDocumentQuery } from 'queries/AccessibilityRequestDocumentQueries';
import {
  DeleteAccessibilityRequestDocument,
  DeleteAccessibilityRequestDocumentVariables
} from 'queries/types/DeleteAccessibilityRequestDocument';

import Modal from 'components/Modal';
import PageHeading from 'components/PageHeading';
import {
  AccessibilityRequestDocumentCommonType,
  AccessibilityRequestDocumentStatus
} from 'types/graphql-global-types';
import { translateDocumentType } from 'utils/accessibilityRequest';
import { formatDate } from 'utils/date';

type Document = {
  id: string;
  status: AccessibilityRequestDocumentStatus;
  url: string;
  uploadedAt: string;
  documentType: {
    commonType: AccessibilityRequestDocumentCommonType;
    otherTypeDescription: string | null;
  };
};

type DocumentsListProps = {
  documents: Document[];
  requestName: string;
  refetchRequest: () => any;
  setConfirmationText: (text: string) => void;
};

const AccessibilityDocumentsList = ({
  documents,
  requestName,
  refetchRequest,
  setConfirmationText
}: DocumentsListProps) => {
  const { t } = useTranslation('accessibility');
  const [document, setDocument] = useState<Document | null>(null);

  const getDocType = (documentType: {
    commonType: AccessibilityRequestDocumentCommonType;
    otherTypeDescription: string | null;
  }) => {
    if (documentType.commonType !== 'OTHER') {
      return translateDocumentType(
        documentType.commonType as AccessibilityRequestDocumentCommonType
      );
    }
    return documentType.otherTypeDescription || '';
  };

  const [mutate] = useMutation<
    DeleteAccessibilityRequestDocument,
    DeleteAccessibilityRequestDocumentVariables
  >(DeleteAccessibilityRequestDocumentQuery);

  const submitDelete = (id: string) => {
    mutate({
      variables: {
        input: {
          id
        }
      }
    }).then(() => {
      refetchRequest();
      if (document) {
        setConfirmationText(
          `${getDocType(document.documentType)} removed from ${requestName}`
        );
      }
      setDocument(null);
    });
  };

  const columns = useMemo<Column<Document>[]>(() => {
    return [
      {
        Header: t<string>('documentTable.header.documentName'),
        accessor: 'documentType',
        Cell: ({ value }: any) => getDocType(value)
      },
      {
        Header: t<string>('documentTable.header.uploadedAt'),
        accessor: 'uploadedAt',
        Cell: ({ value }: any) => {
          if (value) {
            return formatDate(value);
          }
          return '';
        },
        width: '25%'
      },
      {
        Header: t<string>('documentTable.header.actions'),
        Cell: ({ row }: Cell<Document>) => (
          <>
            {row.original.status === 'PENDING' && (
              <em>Virus scan in progress...</em>
            )}
            {row.original.status === 'AVAILABLE' && (
              <>
                <Link
                  className="margin-right-3"
                  target="_blank"
                  rel="noreferrer"
                  href={row.original.url}
                  aria-label={`View ${getDocType(
                    row.original.documentType
                  )} in a new tab or window`}
                >
                  {t('documentTable.view')}
                </Link>
                <Button
                  aria-label={`Remove ${getDocType(row.original.documentType)}`}
                  type="button"
                  unstyled
                  onClick={() => setDocument(row.original)}
                >
                  {t('documentTable.remove')}
                </Button>
                <Modal
                  isOpen={document === row.original}
                  closeModal={() => setDocument(null)}
                >
                  <PageHeading
                    headingLevel="h2"
                    className="margin-top-0 line-height-heading-2 margin-bottom-2"
                  >
                    {t('documentTable.modal.header', {
                      name: getDocType(row.original.documentType)
                    })}
                  </PageHeading>
                  <span>{t('documentTable.modal.warning')}</span>
                  <div className="display-flex margin-top-2">
                    <Button
                      type="button"
                      className="margin-right-5"
                      onClick={() => submitDelete(row.original.id)}
                    >
                      {t('documentTable.modal.proceedButton')}
                    </Button>
                    <Button
                      type="button"
                      unstyled
                      onClick={() => setDocument(null)}
                    >
                      {t('documentTable.modal.declineButton')}
                    </Button>
                  </div>
                </Modal>
              </>
            )}
            {row.original.status === 'UNAVAILABLE' && (
              <>
                <i className="fa fa-exclamation-circle text-secondary" />{' '}
                Document failed virus scan
              </>
            )}
          </>
        )
      }
    ];
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow
  } = useTable({
    columns,
    data: documents,
    documents,
    initialState: {
      sortBy: [{ id: 'uploadedAt', desc: true }]
    }
  });

  if (documents.length === 0) {
    return <span>{t('requestDetails.documents.none')}</span>;
  }

  return (
    <>
      <Table bordered={false} {...getTableProps()} fullWidth>
        <caption className="usa-sr-only">
          {`${t('documentTable.caption')} ${requestName}`}
        </caption>
        <thead>
          {headerGroups.map(headerGroup => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map(column => (
                <th
                  {...column.getHeaderProps()}
                  style={{ whiteSpace: 'nowrap', width: column.width }}
                  scope="col"
                >
                  {column.render('Header')}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map(row => {
            prepareRow(row);
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map((cell, i) => {
                  if (i === 0) {
                    return (
                      <th
                        {...cell.getCellProps()}
                        scope="row"
                        style={{ maxWidth: '16rem' }}
                      >
                        {cell.render('Cell')}
                      </th>
                    );
                  }
                  return (
                    <td {...cell.getCellProps()} style={{ maxWidth: '16rem' }}>
                      {cell.render('Cell')}
                    </td>
                  );
                })}
              </tr>
            );
          })}
        </tbody>
      </Table>
      <Modal isOpen={!!document} closeModal={() => setDocument(null)}>
        {document && (
          <>
            <PageHeading
              headingLevel="h2"
              className="margin-top-0 line-height-heading-2 margin-bottom-2"
            >
              {t('documentTable.modal.header', {
                name: getDocType(document.documentType)
              })}
            </PageHeading>
            <span>{t('documentTable.modal.warning')}</span>
            <div className="display-flex margin-top-2">
              <Button
                type="button"
                className="margin-right-5"
                onClick={() => submitDelete(document.id)}
              >
                {t('documentTable.modal.proceedButton')}
              </Button>
              <Button type="button" unstyled onClick={() => setDocument(null)}>
                {t('documentTable.modal.declineButton')}
              </Button>
            </div>
          </>
        )}
      </Modal>
    </>
  );
};

export default AccessibilityDocumentsList;
