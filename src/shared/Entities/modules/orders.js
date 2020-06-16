import { orders } from '../schema';
import { ADD_ENTITIES } from '../actions';
import { denormalize } from 'normalizr';
import { swaggerRequest } from 'shared/Swagger/request';
import { formatDateForSwagger } from 'shared/dates';
import { getClient } from 'shared/Swagger/api';
import { get, filter, isEmpty, isNull } from 'lodash';
import { fetchActive } from 'shared/utils';

export const STATE_KEY = 'orders';
export const loadOrdersLabel = 'Orders.loadOrders';
const updateOrdersLabel = 'Orders.updateOrders';
const createOrdersLabel = 'Orders.createOrders';
export const getLatestOrdersLabel = 'Orders.showServiceMemberOrders';

export default function reducer(state = {}, action) {
  switch (action.type) {
    case ADD_ENTITIES:
      return {
        ...state,
        ...action.payload.orders,
      };

    default:
      return state;
  }
}

export function fetchLatestOrders(serviceMemberId, label = getLatestOrdersLabel) {
  const swaggerTag = 'service_members.showServiceMemberOrders';
  return swaggerRequest(getClient, swaggerTag, { serviceMemberId }, { label });
}

export function loadOrders(ordersId, label = loadOrdersLabel) {
  const swaggerTag = 'orders.showOrders';
  return swaggerRequest(getClient, swaggerTag, { ordersId }, { label });
}

export function updateOrders(ordersId, orders, label = updateOrdersLabel) {
  const swaggerTag = 'orders.updateOrders';
  return swaggerRequest(getClient, swaggerTag, { ordersId, updateOrders: orders }, { label });
}

export function createOrders(orders, label = createOrdersLabel) {
  const swaggerTag = 'orders.createOrders';
  orders.report_by_date = formatDateForSwagger(orders.report_by_date);
  orders.issue_date = formatDateForSwagger(orders.issue_date);
  return swaggerRequest(getClient, swaggerTag, { createOrders: orders }, { label });
}

// Selectors
export const selectUpload = (state, id) => {
  return denormalize([id], orders, state.entities)[0];
};

export function selectOrders(state, ordersId) {
  return get(state, `entities.orders.${ordersId}`) || {};
}

export function selectOrdersForMove(state, moveId) {
  const ordersId = get(state, `entities.moves.${moveId}.orders_id`);
  if (ordersId) {
    return selectOrders(state, ordersId);
  } else {
    return {};
  }
}

export function selectUploadsForOrders(state, ordersId) {
  const orders = selectOrders(state, ordersId);
  const uploadedOrders = get(state, `entities.documents.${orders.uploaded_orders}`);
  if (uploadedOrders) {
    return uploadedOrders.uploads.map((uploadId) => get(state, `entities.uploads.${uploadId}`));
  } else {
    return [];
  }
}

export function selectOrdersFromServiceMemberId(state, serviceMemberId) {
  const orders = Object.values(state.entities.orders);
  filter(orders, (order) => order.service_member_id === serviceMemberId);
  return orders[0] || {};
}

export function selectActiveOrders(state) {
  // temp until full redux refactor: gets active orders from entities if exist. If not, gets from orders.currentOrders.
  const serviceMember = get(state, 'user.userInfo.service_member', {});
  if (isNull(serviceMember)) {
    return null;
  }
  let activeOrders = selectOrdersFromServiceMemberId(state, serviceMember.id);
  if (isEmpty(activeOrders)) {
    activeOrders = fetchActive(get(state, 'user.userInfo.service_member.orders', {}));
  }
  return activeOrders;
}
