import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { get } from 'lodash';
import { getEntitlements } from 'shared/Entities/modules/moveTaskOrders';

class TOO extends React.Component {
  componentDidMount() {
    this.props.getEntitlements('fake_move_task_order_id');
  }

  render() {
    const { entitlements } = this.props;
    console.log(entitlements);
    const NTS = entitlements && entitlements.nonTemporaryStorage ? 'Y' : 'N';
    const POV = entitlements && entitlements.privatelyOwnedVehicle ? 'Y' : 'N';
    return (
      <>
        <h1>TOO Placeholder Page</h1>
        {entitlements && (
          <>
            <h2>Customer Entitlements</h2>
            <dl>
              <dt>Weight Entitlement</dt>
              <dd>{entitlements.totalWeightSelf}</dd>
              <dt>SIT Entitlement</dt>
              <dd>{entitlements.storageInTransit}</dd>
              <dt>NTS Entitlement</dt>
              <dd>{NTS}</dd>
              <dt>POV Entitlement</dt>
              <dd>{POV}</dd>
            </dl>
          </>
        )}
      </>
    );
  }
}
const mapStateToProps = state => {
  const entitlements = get(state, 'entities.entitlements');
  return {
    entitlements: entitlements && Object.values(entitlements).length > 0 ? Object.values(entitlements)[0] : null,
  };
};
const mapDispatchToProps = dispatch => bindActionCreators({ getEntitlements }, dispatch);

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(TOO);
