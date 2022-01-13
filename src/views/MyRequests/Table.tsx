import React, { useMemo } from 'react';
import { useTranslation } from 'react-i18next';
import { useSortBy, useTable } from 'react-table';
import { useQuery } from '@apollo/client';
import { Table as UswdsTable } from '@trussworks/react-uswds';
import classnames from 'classnames';
import { DateTime } from 'luxon';

import UswdsReactLink from 'components/LinkWrapper';
import Spinner from 'components/Spinner';
import GetRequestsQuery from 'queries/GetRequestsQuery';
import { GetRequests, GetRequestsVariables } from 'queries/types/GetRequests';
import { RequestType } from 'types/graphql-global-types';
import { accessibilityRequestStatusMap } from 'utils/accessibilityRequest';
import { formatDate } from 'utils/date';

import './index.scss';

const Table = () => {
  const { t } = useTranslation(['home', 'intake', 'accessibility']);
  const { loading, error, data: tableData } = useQuery<
    GetRequests,
    GetRequestsVariables
  >(GetRequestsQuery, {
    variables: { first: 20 },
    fetchPolicy: 'cache-and-network'
  });

  const columns: any = useMemo(() => {
    return [
      {
        Header: t('requestsTable.headers.name'),
        accessor: 'name',
        Cell: ({ row, value }: any) => {
          let link: string;
          switch (row.original.type) {
            case RequestType.ACCESSIBILITY_REQUEST:
              link = `/508/requests/${row.original.id}`;
              break;
            case RequestType.GOVERNANCE_REQUEST:
              link = `/governance-task-list/${row.original.id}`;
              break;
            default:
              link = '/';
          }
          return (
            <UswdsReactLink link={link}>
              {value || t('requestsTable.defaultName')}
            </UswdsReactLink>
          );
        },
        maxWidth: 350
      },
      {
        Header: t('requestsTable.headers.type'),
        accessor: 'type',
        Cell: ({ value }: any) => {
          return t(`requestsTable.types.${value}`);
        }
      },
      {
        Header: t('requestsTable.headers.submittedAt'),
        accessor: 'submittedAt',
        Cell: ({ value }: any) => {
          if (value) {
            return formatDate(value);
          }
          return t('requestsTable.defaultSubmittedAt');
        }
      },
      {
        Header: t('requestsTable.headers.status'),
        accessor: 'status',
        width: '200px',
        Cell: ({ row, value }: any) => {
          let statusString;
          switch (row.original.type) {
            case RequestType.ACCESSIBILITY_REQUEST:
              // Status hasn't changed if the status record created at is the same
              // as the 508 request's submitted at
              if (
                row.original.submittedAt.toISO() ===
                row.original.statusCreatedAt.toISO()
              ) {
                return <span>{accessibilityRequestStatusMap[value]}</span>;
              }

              return (
                <span>
                  {accessibilityRequestStatusMap[value]}{' '}
                  <span className="text-base-dark font-body-3xs">{`changed on ${formatDate(
                    row.original.statusCreatedAt
                  )}`}</span>
                </span>
              );
            case RequestType.GOVERNANCE_REQUEST:
              statusString = t(`intake:statusMap.${value}`);
              if (row.original.lcid) {
                return `${statusString}: ${row.original.lcid}`;
              }
              return statusString;
            default:
              return '';
          }
        }
      },
      {
        Header: t('requestsTable.headers.nextMeetingDate'),
        accessor: 'nextMeetingDate',
        className: 'next-meeting-date',
        width: '180px',
        Cell: ({ value }: any) => {
          if (value) {
            return formatDate(value);
          }
          return 'None';
        }
      }
    ];
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const data = useMemo(() => {
    const requests =
      tableData &&
      tableData.requests &&
      tableData.requests.edges.map(edge => {
        return edge.node;
      });

    const mappedData = requests?.map(request => {
      const submittedAt = request.submittedAt
        ? DateTime.fromISO(request.submittedAt)
        : null;

      const statusCreatedAt = request.statusCreatedAt
        ? DateTime.fromISO(request.statusCreatedAt)
        : null;
      return { ...request, submittedAt, statusCreatedAt };
    });

    return mappedData || [];
  }, [tableData]);

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow
  } = useTable(
    {
      columns,
      data,
      sortTypes: {
        alphanumeric: (rowOne, rowTwo, columnName) => {
          const rowOneElem = rowOne.values[columnName];
          const rowTwoElem = rowTwo.values[columnName];

          // If item is a string, enforce capitalization (temporarily) and then compare
          if (typeof rowOneElem === 'string') {
            return rowOneElem.toUpperCase() > rowTwoElem.toUpperCase() ? 1 : -1;
          }

          // If item is a DateTime, convert to Number and compare
          if (rowOneElem instanceof DateTime) {
            return Number(rowOneElem) > Number(rowTwoElem) ? 1 : -1;
          }

          // If neither string nor DateTime, return bare comparison
          return rowOneElem > rowTwoElem ? 1 : -1;
        }
      },
      initialState: {
        sortBy: [{ id: 'submittedAt', desc: true }]
      }
    },
    useSortBy
  );

  const getHeaderSortIcon = (isDesc: boolean | undefined) => {
    return classnames('margin-left-1', {
      'fa fa-caret-down fa-lg caret': isDesc,
      'fa fa-caret-up fa-lg caret': !isDesc
    });
  };

  const getColumnSortStatus = (
    column: any
  ): 'descending' | 'ascending' | 'none' => {
    if (column.isSorted) {
      if (column.isSortedDesc) {
        return 'descending';
      }
      return 'ascending';
    }

    return 'none';
  };

  if (loading) {
    return (
      <div className="text-center" data-testid="table-loading">
        <Spinner size="xl" />;
      </div>
    );
  }

  if (error) {
    return <div>{JSON.stringify(error)}</div>;
  }

  if (data.length === 0) {
    return <p>{t('requestsTable.empty')}</p>;
  }

  return (
    <div className="accessibility-requests-table">
      <UswdsTable bordered={false} {...getTableProps()} fullWidth scrollable>
        <caption className="usa-sr-only">{t('requestsTable.caption')}</caption>
        <thead>
          {headerGroups.map(headerGroup => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map(column => (
                <th
                  {...column.getHeaderProps()}
                  aria-sort={getColumnSortStatus(column)}
                  style={{}}
                  scope="col"
                >
                  <button
                    className="usa-button usa-button--unstyled"
                    type="button"
                    {...column.getSortByToggleProps()}
                  >
                    {column.render('Header')}
                    {column.isSorted && (
                      <span
                        className={getHeaderSortIcon(column.isSortedDesc)}
                      />
                    )}
                    {!column.isSorted && (
                      <span className="margin-left-1 fa fa-sort caret" />
                    )}
                  </button>
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
                      <th {...cell.getCellProps()} scope="row">
                        {cell.render('Cell')}
                      </th>
                    );
                  }
                  return (
                    <td
                      {...cell.getCellProps()}
                      style={{ width: cell.column.width }}
                    >
                      {cell.render('Cell')}
                    </td>
                  );
                })}
              </tr>
            );
          })}
        </tbody>
      </UswdsTable>
      <div
        className="usa-sr-only usa-table__announcement-region"
        aria-live="polite"
      >
        {currentTableSortDescription(headerGroups[0])}
      </div>
    </div>
  );
};

const currentTableSortDescription = headerGroup => {
  const sortedHeader = headerGroup.headers.find(header => header.isSorted);

  if (sortedHeader) {
    const direction = sortedHeader.isSortedDesc ? 'descending' : 'ascending';
    return `Requests table sorted by ${sortedHeader.Header} ${direction}`;
  }
  return 'Requests table reset to default sort order';
};

export default Table;
