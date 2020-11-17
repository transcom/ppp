/*  react/jsx-props-no-spreading */
import React from 'react';
import { mount } from 'enzyme';
import { Provider } from 'react-redux';
import moment from 'moment';

import Home from '.';

import { MockProviders } from 'testUtils';
import { store } from 'shared/store';
import { formatCustomerDate } from 'utils/formatters';

const defaultProps = {
  serviceMember: {
    current_station: {},
    weight_allotment: {},
  },
  showLoggedInUser: jest.fn(),
  createServiceMember: jest.fn(),
  mtoShipments: [],
  mtoShipment: {},
  isLoggedIn: true,
  loggedInUserIsLoading: false,
  loggedInUserSuccess: true,
  loggedInUserError: false,
  isProfileComplete: true,
  moveSubmitSuccess: false,
  currentPpm: {},
  loadMTOShipments: jest.fn(),
  orders: {},
  history: {},
  location: {},
  move: {},
};

function mountHome(props = {}) {
  return mount(
    <Provider store={store}>
      <Home {...defaultProps} {...props} />
    </Provider>,
  );
}
describe('Home component', () => {
  describe('with default props', () => {
    const wrapper = mountHome();

    it('renders Home with the right amount of components', () => {
      expect(wrapper.find('Step').length).toBe(4);
      expect(wrapper.find('Helper').length).toBe(1);
      expect(wrapper.find('Contact').length).toBe(1);
    });

    it('Profile Step is editable', () => {
      const profileStep = wrapper.find('Step[step="1"]');
      expect(profileStep.prop('editBtnLabel')).toEqual('Edit');
    });
  });

  describe('contents of Step 3', () => {
    it('contains ppm and hhg cards if those shipments exist', () => {
      const props = {
        currentPpm: { id: '12345', createdAt: moment() },
        mtoShipments: [
          { id: '4321', createdAt: moment().add(1, 'days'), shipmentType: 'HHG' },
          { id: '4322', createdAt: moment().subtract(1, 'days'), shipmentType: 'HHG' },
        ],
      };

      const wrapper = mountHome(props);
      expect(wrapper.find('ShipmentListItem').length).toBe(3);
      expect(wrapper.find('ShipmentListItem').at(0).text()).toContain('HHG 1');
      expect(wrapper.find('ShipmentListItem').at(1).text()).toContain('PPM');
      expect(wrapper.find('ShipmentListItem').at(2).text()).toContain('HHG 2');
    });
  });

  describe('if the user does not have orders', () => {
    const wrapper = mountHome();

    it('renders the NeedsOrders helper', () => {
      expect(wrapper.find('HelperNeedsOrders').exists()).toBe(true);
    });

    it('Orders Step is not editable', () => {
      const ordersStep = wrapper.find('Step[step="2"]');
      expect(ordersStep.prop('editBtnLabel')).toEqual('');
    });
  });

  describe('if the user has orders but not shipments', () => {
    const wrapper = mountHome({
      orders: { testOrder: 'test', new_duty_station: { name: 'Test Duty Station' } },
      uploadedOrderDocuments: [{ filename: 'testOrder1.pdf' }],
    });

    it('renders the NeedsShipment helper', () => {
      expect(wrapper.find('HelperNeedsShipment').exists()).toBe(true);
    });

    it('Orders Step is editable', () => {
      const ordersStep = wrapper.find('Step[step="2"]');
      expect(ordersStep.prop('editBtnLabel')).toEqual('Edit');
    });
  });

  describe('if the user has orders and shipments but has not submitted their move', () => {
    const wrapper = mountHome({
      orders: { id: 'testOrder123', new_duty_station: { name: 'Test Duty Station' } },
      uploadedOrderDocuments: [{ filename: 'testOrder1.pdf' }],
      mtoShipments: [{ id: 'test123', shipmentType: 'HHG' }],
    });

    it('renders the NeedsSubmitMove helper', () => {
      expect(wrapper.find('HelperNeedsSubmitMove').exists()).toBe(true);
    });
  });

  describe('if the user has orders and a currentPpm but has not submitted their move', () => {
    const wrapper = mountHome({
      orders: { id: 'testOrder123', new_duty_station: { name: 'Test Duty Station' } },
      uploadedOrderDocuments: [{ filename: 'testOrder1.pdf' }],
      currentPpm: { id: 'testPpm123' },
    });

    it('renders the NeedsSubmitMove helper', () => {
      expect(wrapper.find('HelperNeedsSubmitMove').exists()).toBe(true);
    });
  });

  describe('if the user has submitted their move', () => {
    describe('for PPM moves', () => {
      const orders = {
        id: 'testOrder123',
        new_duty_station: {
          name: 'Test Duty Station',
        },
      };
      const uploadedOrderDocuments = [{ filename: 'testOrder1.pdf' }];
      const move = { status: 'SUBMITTED' };
      const currentPpm = { id: 'mockPpm ' };
      const wrapper = mount(
        <MockProviders initialEntries={['/']}>
          <Home
            {...defaultProps}
            orders={orders}
            uploadedOrderDocuments={uploadedOrderDocuments}
            move={move}
            currentPpm={currentPpm}
          />
        </MockProviders>,
      );

      it('renders the SubmittedMove helper', () => {
        expect(wrapper.find('HelperSubmittedMove').exists()).toBe(true);
      });

      it('Profile step is editable', () => {
        const profileStep = wrapper.find('Step[step="1"]');
        expect(profileStep.prop('editBtnLabel')).toEqual('Edit');
      });

      it('Orders Step is not editable', () => {
        const ordersStep = wrapper.find('Step[step="2"]');
        expect(ordersStep.prop('editBtnLabel')).toEqual('');
      });

      it('renders the SubmittedPPM helper', () => {
        expect(wrapper.find('HelperSubmittedPPM').exists()).toBe(true);
      });
    });

    describe('for HHG moves (no PPM)', () => {
      const wrapper = mountHome({
        orders: { id: 'testOrder123', new_duty_station: { name: 'Test Duty Station' } },
        uploadedOrderDocuments: [{ filename: 'testOrder1.pdf' }],
        mtoShipments: [{ id: 'test123', shipmentType: 'HHG' }],
        move: { status: 'SUBMITTED' },
      });

      it('renders the SubmittedMove helper', () => {
        expect(wrapper.find('HelperSubmittedMove').exists()).toBe(true);
      });

      it('Profile step is editable', () => {
        const profileStep = wrapper.find('Step[step="1"]');
        expect(profileStep.prop('editBtnLabel')).toEqual('Edit');
      });

      it('Orders Step is not editable', () => {
        const ordersStep = wrapper.find('Step[step="2"]');
        expect(ordersStep.prop('editBtnLabel')).toEqual('');
      });
    });

    describe('for NTS moves (no PPM)', () => {
      const wrapper = mountHome({
        orders: { id: 'testOrder123', new_duty_station: { name: 'Test Duty Station' } },
        uploadedOrderDocuments: [{ filename: 'testOrder1.pdf' }],
        mtoShipments: [{ id: 'test123', shipmentType: 'NTS' }],
        move: { status: 'SUBMITTED' },
      });

      it('renders the SubmittedMove helper', () => {
        expect(wrapper.find('HelperSubmittedMove').exists()).toBe(true);
      });

      it('Profile step is editable', () => {
        const profileStep = wrapper.find('Step[step="1"]');
        expect(profileStep.prop('editBtnLabel')).toEqual('Edit');
      });

      it('Orders Step is not editable', () => {
        const ordersStep = wrapper.find('Step[step="2"]');
        expect(ordersStep.prop('editBtnLabel')).toEqual('');
      });
    });

    describe('for HHG/PPM combo moves', () => {
      const submittedAt = new Date();
      const orders = {
        id: 'testOrder123',
        new_duty_station: {
          name: 'Test Duty Station',
        },
      };
      const uploadedOrderDocuments = [{ filename: 'testOrder1.pdf' }];
      const move = { status: 'SUBMITTED', submitted_at: submittedAt };
      const currentPpm = { id: 'mockCombo' };
      const wrapper = mount(
        <MockProviders initialEntries={['/']}>
          <Home
            {...defaultProps}
            orders={orders}
            uploadedOrderDocuments={uploadedOrderDocuments}
            move={move}
            currentPpm={currentPpm}
          />
        </MockProviders>,
      );

      it('renders submitted date at step 4', () => {
        expect(wrapper.find('[data-testid="move-submitted-description"]').text()).toBe(
          `Move submitted ${formatCustomerDate(submittedAt)}.Print the legal agreement`,
        );
      });

      it('renders secondary button when step 4 is completed', () => {
        expect(wrapper.find('[data-testid="review-and-submit-btn"]').at(1).hasClass('usa-button--secondary')).toBe(
          true,
        );
      });

      it('renders the SubmittedMove helper', () => {
        expect(wrapper.find('HelperSubmittedMove').exists()).toBe(true);
      });

      it('Profile step is editable', () => {
        const profileStep = wrapper.find('Step[step="1"]');
        expect(profileStep.prop('editBtnLabel')).toEqual('Edit');
      });

      it('Orders Step is not editable', () => {
        const ordersStep = wrapper.find('Step[step="2"]');
        expect(ordersStep.prop('editBtnLabel')).toEqual('');
      });

      it('renders the SubmittedPPM helper', () => {
        expect(wrapper.find('HelperSubmittedPPM').exists()).toBe(true);
      });
    });
  });
});
