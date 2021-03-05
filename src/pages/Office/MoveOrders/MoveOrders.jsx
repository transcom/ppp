/* eslint-disable camelcase */
import React, { useState, useEffect } from 'react';
import { Link, useHistory, useParams } from 'react-router-dom';
import { Button } from '@trussworks/react-uswds';
import * as Yup from 'yup';
import { Formik } from 'formik';
import { queryCache, useMutation } from 'react-query';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import styles from './MoveOrders.module.scss';

import { getTacValid, updateMoveOrder } from 'services/ghcApi';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import OrdersDetailForm from 'components/Office/OrdersDetailForm/OrdersDetailForm';
import { dropdownInputOptions, formatSwaggerDate } from 'shared/formatters';
import { DEPARTMENT_INDICATOR_OPTIONS } from 'constants/departmentIndicators';
import { ORDERS_TYPE_DETAILS_OPTIONS, ORDERS_TYPE_OPTIONS } from 'constants/orders';
import { MOVE_ORDERS } from 'constants/queryKeys';
import { useOrdersDocumentQueries } from 'hooks/queries';

const deptIndicatorDropdownOptions = dropdownInputOptions(DEPARTMENT_INDICATOR_OPTIONS);
const ordersTypeDropdownOptions = dropdownInputOptions(ORDERS_TYPE_OPTIONS);
const ordersTypeDetailsDropdownOptions = dropdownInputOptions(ORDERS_TYPE_DETAILS_OPTIONS);

const validationSchema = Yup.object({
  originDutyStation: Yup.object().defined('Required'),
  newDutyStation: Yup.object().required('Required'),
  issueDate: Yup.date()
    .typeError('Enter a complete date in DD MMM YYYY format (day, month, year).')
    .required('Required'),
  reportByDate: Yup.date()
    .typeError('Enter a complete date in DD MMM YYYY format (day, month, year).')
    .required('Required'),
  departmentIndicator: Yup.string().required('Required'),
  ordersNumber: Yup.string().required('Required'),
  ordersType: Yup.string().required('Required'),
  ordersTypeDetail: Yup.string().required('Required'),
  tac: Yup.string().min(4, 'Enter a 4-character TAC').required('Required'),
  sac: Yup.string().required('Required'),
});

const MoveOrders = () => {
  const history = useHistory();
  const { moveCode } = useParams();
  const [isValidTac, setIsValidTac] = useState(true);
  const { move, moveOrders, isLoading, isError } = useOrdersDocumentQueries(moveCode);
  const orderId = move?.ordersId;

  const handleClose = () => {
    history.push(`/moves/${moveCode}/details`);
  };

  const [mutateOrders] = useMutation(updateMoveOrder, {
    onSuccess: (data, variables) => {
      const updatedOrder = data.moveOrders[variables.orderID];
      queryCache.setQueryData([MOVE_ORDERS, variables.orderID], {
        moveOrders: {
          [`${variables.orderID}`]: updatedOrder,
        },
      });
      queryCache.invalidateQueries(MOVE_ORDERS);
      handleClose();
    },
    onError: (error) => {
      const errorMsg = error?.response?.body;
      // TODO: Handle error some how
      // RA Summary: eslint: no-console - System Information Leak: External
      // RA: The linter flags any use of console.
      // RA: This console displays an error message from unsuccessful mutation.
      // RA: TODO: As indicated, this error needs to be handled and needs further investigation and work.
      // RA: POAM story here: https://dp3.atlassian.net/browse/MB-5597
      // RA Developer Status: Known Issue
      // RA Validator Status: Known Issue
      // RA Modified Severity: CAT II
      // eslint-disable-next-line no-console
      console.log(errorMsg);
    },
  });

  const handleTacValidation = (value) => {
    if (value && value.length === 4) {
      getTacValid({ tac: value }).then((response) => setIsValidTac(response.isValid));
    }
  };

  const moveOrder = Object.values(moveOrders)?.[0];

  useEffect(() => {
    // if the initial value === value, and it's 4 digits, run validator and show warning if invalid
    if (moveOrder?.tac) handleTacValidation(moveOrder.tac);
  }, [moveOrder]);

  if (isLoading) return <LoadingPlaceholder />;
  if (isError) return <SomethingWentWrong />;

  const onSubmit = (values) => {
    const { originDutyStation, newDutyStation, ...fields } = values;
    const body = {
      ...fields,
      originDutyStationId: values.originDutyStation.id,
      newDutyStationId: values.newDutyStation.id,
      issueDate: formatSwaggerDate(values.issueDate),
      reportByDate: formatSwaggerDate(values.reportByDate),
    };
    mutateOrders({ orderID: orderId, ifMatchETag: moveOrder.eTag, body });
  };

  const tacWarningMsg =
    'This TAC does not appear in TGET, so it might not be valid. Make sure it matches what‘s on the orders before you continue.';

  const initialValues = {
    agency: moveOrder?.agency,
    originDutyStation: moveOrder?.originDutyStation,
    newDutyStation: moveOrder?.destinationDutyStation,
    issueDate: moveOrder?.date_issued,
    reportByDate: moveOrder?.report_by_date,
    departmentIndicator: moveOrder?.department_indicator,
    ordersNumber: moveOrder?.order_number,
    ordersType: moveOrder?.order_type,
    ordersTypeDetail: moveOrder?.order_type_detail,
    tac: moveOrder?.tac,
    sac: moveOrder?.sac,
  };

  return (
    <div className={styles.sidebar}>
      <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={onSubmit}>
        {(formik) => {
          // onBlur, if the value has 4 digits, run validator and show warning if invalid
          const tacWarning = isValidTac ? '' : tacWarningMsg;
          return (
            <form onSubmit={formik.handleSubmit}>
              <div className={styles.orderDetails}>
                <div className={styles.top}>
                  <Button
                    className={styles.closeButton}
                    data-testid="closeSidebar"
                    type="button"
                    onClick={handleClose}
                    unstyled
                  >
                    <FontAwesomeIcon icon="times" title="Close sidebar" aria-label="Close sidebar" />
                  </Button>
                  <h2 className={styles.header}>View Orders</h2>
                  <div>
                    <Link className={styles.viewAllowances} data-testid="view-allowances" to="allowances">
                      View Allowances
                    </Link>
                  </div>
                </div>
                <div className={styles.body}>
                  <OrdersDetailForm
                    deptIndicatorOptions={deptIndicatorDropdownOptions}
                    ordersTypeOptions={ordersTypeDropdownOptions}
                    ordersTypeDetailOptions={ordersTypeDetailsDropdownOptions}
                    tacWarning={tacWarning}
                    validateTac={handleTacValidation}
                  />
                </div>
                <div className={styles.bottom}>
                  <div className={styles.buttonGroup}>
                    <Button type="submit" disabled={formik.isSubmitting}>
                      Save
                    </Button>
                    <Button type="button" secondary onClick={handleClose}>
                      Cancel
                    </Button>
                  </div>
                </div>
              </div>
            </form>
          );
        }}
      </Formik>
    </div>
  );
};

export default MoveOrders;
