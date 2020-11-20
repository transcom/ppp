import React from 'react';
import { withRouter } from 'react-router-dom';

import { usePaymentRequestQueueQueries, useUserQueries } from 'hooks/queries';
import { createHeader } from 'components/Table/utils';
import { HistoryShape } from 'types/router';
import {
  formatDateFromIso,
  formatAgeToDays,
  paymentRequestStatusReadable,
  serviceMemberAgencyLabel,
} from 'shared/formatters';
import MultiSelectCheckBoxFilter from 'components/Table/Filters/MultiSelectCheckBoxFilter';
import SelectFilter from 'components/Table/Filters/SelectFilter';
import DateSelectFilter from 'components/Table/Filters/DateSelectFilter';
import { BRANCH_OPTIONS, BRANCH_OPTIONS_NO_MARINES, GBLOC, PAYMENT_REQUEST_STATUS_OPTIONS } from 'constants/queues';
import TableQueue from 'components/Table/TableQueue';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';

const columns = (includeBranchOptionMarines = true) => [
  createHeader('ID', 'id'),
  createHeader(
    'Customer name',
    (row) => {
      return `${row.customer.last_name}, ${row.customer.first_name}`;
    },
    {
      id: 'lastName',
      isFilterable: true,
    },
  ),
  createHeader('DoD ID', 'customer.dodID', {
    id: 'dodID',
    isFilterable: true,
  }),
  createHeader(
    'Status',
    (row) => {
      return paymentRequestStatusReadable(row.status);
    },
    {
      id: 'status',
      isFilterable: true,
      // eslint-disable-next-line react/jsx-props-no-spreading
      Filter: (props) => <MultiSelectCheckBoxFilter options={PAYMENT_REQUEST_STATUS_OPTIONS} {...props} />,
    },
  ),
  createHeader(
    'Age',
    (row) => {
      return formatAgeToDays(row.age);
    },
    'age',
  ),
  createHeader(
    'Submitted',
    (row) => {
      return formatDateFromIso(row.submittedAt, 'DD MMM YYYY');
    },
    {
      id: 'submittedAt',
      isFilterable: true,
      Filter: DateSelectFilter,
    },
  ),
  createHeader('Move Code', 'locator', {
    id: 'moveID',
    isFilterable: true,
  }),
  createHeader(
    'Branch',
    (row) => {
      return serviceMemberAgencyLabel(row.customer.agency);
    },
    {
      id: 'branch',
      isFilterable: true,
      Filter: (props) => (
        // eslint-disable-next-line react/jsx-props-no-spreading
        <SelectFilter options={includeBranchOptionMarines ? BRANCH_OPTIONS : BRANCH_OPTIONS_NO_MARINES} {...props} />
      ),
    },
  ),
  createHeader('Origin GBLOC', 'originGBLOC'),
];

const PaymentRequestQueue = ({ history }) => {
  const {
    // eslint-disable-next-line camelcase
    data: { office_user },
    isLoading,
    isError,
  } = useUserQueries();

  const includeBranchOptionMarines = office_user?.transportation_office?.gbloc === GBLOC.USMC;

  const handleClick = (values) => {
    history.push(`/moves/MOVE_CODE/payment-requests/${values.id}`);
  };

  if (isLoading) return <LoadingPlaceholder />;
  if (isError) return <SomethingWentWrong />;

  return (
    <TableQueue
      showFilters
      showPagination
      columns={columns(includeBranchOptionMarines)}
      title="Payment requests"
      handleClick={handleClick}
      useQueries={usePaymentRequestQueueQueries}
    />
  );
};

PaymentRequestQueue.propTypes = {
  history: HistoryShape.isRequired,
};

export default withRouter(PaymentRequestQueue);
