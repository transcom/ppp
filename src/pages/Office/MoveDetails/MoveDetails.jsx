import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import classnames from 'classnames';
import { GridContainer, Grid, Tag } from '@trussworks/react-uswds';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { queryCache, useMutation } from 'react-query';
import { func } from 'prop-types';

import styles from '../TXOMoveInfo/TXOTab.module.scss';

import 'styles/office.scss';
import { updateMoveStatus, updateMTOShipmentStatus } from 'services/ghcApi';
import LeftNav from 'components/LeftNav';
import CustomerInfoTable from 'components/Office/CustomerInfoTable';
import RequestedShipments from 'components/Office/RequestedShipments/RequestedShipments';
import AllowancesTable from 'components/Office/AllowancesTable/AllowancesTable';
import OrdersTable from 'components/Office/OrdersTable/OrdersTable';
import { useMoveDetailsQueries } from 'hooks/queries';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import { MOVES, MTO_SHIPMENTS } from 'constants/queryKeys';

const sectionLabels = {
  'requested-shipments': 'Requested shipments',
  'approved-shipments': 'Approved shipments',
  orders: 'Orders',
  allowances: 'Allowances',
  'customer-info': 'Customer info',
};

const MoveDetails = ({ setUnapprovedShipmentCount }) => {
  const { moveCode } = useParams();

  const [activeSection, setActiveSection] = useState('');

  const { move, order, mtoShipments, mtoServiceItems, isLoading, isError } = useMoveDetailsQueries(moveCode);

  let sections = ['orders', 'allowances', 'customer-info'];

  const handleScroll = () => {
    const distanceFromTop = window.scrollY;
    let newActiveSection;

    sections.forEach((section) => {
      const sectionEl = document.querySelector(`#${section}`);
      if (sectionEl?.offsetTop <= distanceFromTop && sectionEl?.offsetTop + sectionEl?.offsetHeight > distanceFromTop) {
        newActiveSection = section;
      }
    });

    if (activeSection !== newActiveSection) {
      setActiveSection(newActiveSection);
    }
  };

  useEffect(() => {
    // attach scroll listener
    window.addEventListener('scroll', handleScroll);

    // remove scroll listener
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });

  // use mutation calls
  const [mutateMoveStatus] = useMutation(updateMoveStatus, {
    onSuccess: (data) => {
      queryCache.setQueryData([MOVES, data.locator], data);
    },
  });

  const [mutateMTOShipmentStatus] = useMutation(updateMTOShipmentStatus, {
    onSuccess: (updatedMTOShipment) => {
      mtoShipments[mtoShipments.findIndex((shipment) => shipment.id === updatedMTOShipment.id)] = updatedMTOShipment;
      queryCache.setQueryData([MTO_SHIPMENTS, updatedMTOShipment.moveTaskOrderID, false], mtoShipments);
    },
  });

  const submittedShipments = mtoShipments.filter((shipment) => shipment.status === 'SUBMITTED');

  useEffect(() => {
    const shipmentCount = submittedShipments.length;
    setUnapprovedShipmentCount(shipmentCount);
  }, [mtoShipments, submittedShipments, setUnapprovedShipmentCount]);

  if (isLoading) return <LoadingPlaceholder />;
  if (isError) return <SomethingWentWrong />;

  const { customer, entitlement: allowances } = order;

  const approvedShipments = mtoShipments.filter((shipment) => shipment.status === 'APPROVED');

  if (submittedShipments.length > 0 && approvedShipments.length > 0) {
    sections = ['requested-shipments', 'approved-shipments', ...sections];
  } else if (approvedShipments.length > 0) {
    sections = ['approved-shipments', ...sections];
  } else if (submittedShipments.length > 0) {
    sections = ['requested-shipments', ...sections];
  }

  const ordersInfo = {
    newDutyStation: order.destinationDutyStation,
    currentDutyStation: order.originDutyStation,
    issuedDate: order.date_issued,
    reportByDate: order.report_by_date,
    departmentIndicator: order.department_indicator,
    ordersNumber: order.order_number,
    ordersType: order.order_type,
    ordersTypeDetail: order.order_type_detail,
    tacMDC: order.tac,
    sacSDN: order.sac,
  };
  const allowancesInfo = {
    branch: customer.agency,
    rank: order.grade,
    weightAllowance: allowances.totalWeight,
    authorizedWeight: allowances.authorizedWeight,
    progear: allowances.proGearWeight,
    spouseProgear: allowances.proGearWeightSpouse,
    storageInTransit: allowances.storageInTransit,
    dependents: allowances.dependentsAuthorized,
  };
  const customerInfo = {
    name: `${customer.last_name}, ${customer.first_name}`,
    dodId: customer.dodID,
    phone: `+1 ${customer.phone}`,
    email: customer.email,
    currentAddress: customer.current_address,
    backupContact: customer.backup_contact,
  };

  const requiredOrdersInfo = {
    ordersNumber: order.order_number,
    ordersType: order.order_type,
    ordersTypeDetail: order.order_type_detail,
    tacMDC: order.tac,
  };

  const hasMissingOrdersInfo = () => {
    return Object.values(requiredOrdersInfo).some((value) => {
      return !value || value === '';
    });
  };

  const defineSectionLink = (section) => {
    let showErrorTag = false;

    // TODO This will likely become a switch statement or be refactored as more values are considered required
    if (section === 'orders' && hasMissingOrdersInfo()) {
      showErrorTag = true;
    }

    return (
      <a key={`sidenav_${section}`} href={`#${section}`} className={classnames({ active: section === activeSection })}>
        {sectionLabels[`${section}`]}
        {showErrorTag && (
          <Tag className="usa-tag usa-tag--alert">
            <FontAwesomeIcon icon="exclamation" />
          </Tag>
        )}
      </a>
    );
  };

  return (
    <div className={styles.tabContent}>
      <div className={styles.container}>
        <LeftNav className={styles.sidebar}>
          {sections.map((s) => {
            return defineSectionLink(s);
          })}
        </LeftNav>

        <GridContainer className={styles.gridContainer} data-testid="too-move-details">
          <h1>Move details</h1>
          {/* TODO - RequestedShipments could be simplified, if extra time we could tackle this or just write a story to track */}
          {submittedShipments.length > 0 && (
            <div className={styles.section} id="requested-shipments">
              <RequestedShipments
                mtoShipments={submittedShipments}
                ordersInfo={ordersInfo}
                allowancesInfo={allowancesInfo}
                customerInfo={customerInfo}
                mtoServiceItems={mtoServiceItems}
                shipmentsStatus="SUBMITTED"
                approveMTO={mutateMoveStatus}
                approveMTOShipment={mutateMTOShipmentStatus}
                moveTaskOrder={move}
              />
            </div>
          )}
          {approvedShipments.length > 0 && (
            <div className={styles.section} id="approved-shipments">
              <RequestedShipments
                mtoShipments={approvedShipments}
                ordersInfo={ordersInfo}
                allowancesInfo={allowancesInfo}
                customerInfo={customerInfo}
                mtoServiceItems={mtoServiceItems}
                shipmentsStatus="APPROVED"
                moveTaskOrder={move}
              />
            </div>
          )}
          <div className={styles.section} id="orders">
            <GridContainer>
              <Grid row gap>
                <Grid col>
                  <OrdersTable ordersInfo={ordersInfo} />
                </Grid>
              </Grid>
            </GridContainer>
          </div>
          <div className={styles.section} id="allowances">
            <GridContainer>
              <Grid row gap>
                <Grid col>
                  <AllowancesTable info={allowancesInfo} showEditBtn />
                </Grid>
              </Grid>
            </GridContainer>
          </div>
          <div className={styles.section} id="customer-info">
            <GridContainer>
              <Grid row gap>
                <Grid col>
                  <CustomerInfoTable customerInfo={customerInfo} />
                </Grid>
              </Grid>
            </GridContainer>
          </div>
        </GridContainer>
      </div>
    </div>
  );
};

MoveDetails.propTypes = {
  setUnapprovedShipmentCount: func.isRequired,
};

export default MoveDetails;
