/*  no-only-tests/no-only-tests */
/*  react/jsx-props-no-spreading */
import React from 'react';
import { shallow, mount } from 'enzyme';

import ConnectedOffice, { OfficeApp } from './index';

import { MockProviders } from 'testUtils';
import { roleTypes } from 'constants/userRoles';

describe('Office App', () => {
  const mockOfficeProps = {
    loadUser: jest.fn(),
    loadInternalSchema: jest.fn(),
    loadPublicSchema: jest.fn(),
  };

  describe('component', () => {
    let wrapper;

    beforeEach(() => {
      wrapper = shallow(<OfficeApp {...mockOfficeProps} />);
    });

    it('renders without crashing or erroring', () => {
      const officeWrapper = wrapper.find('div');
      expect(officeWrapper).toBeDefined();
      expect(wrapper.find('SomethingWentWrong')).toHaveLength(0);
    });

    it('renders the basic header by default', () => {
      expect(wrapper.find('QueueHeader')).toHaveLength(1);
    });

    it('fetches initial data', () => {
      expect(mockOfficeProps.loadUser).toHaveBeenCalled();
      expect(mockOfficeProps.loadInternalSchema).toHaveBeenCalled();
      expect(mockOfficeProps.loadPublicSchema).toHaveBeenCalled();
    });

    describe('if an error occurs', () => {
      it('renders the fail whale', () => {
        wrapper.setState({ hasError: true });
        expect(wrapper.find('SomethingWentWrong')).toHaveLength(1);
      });
    });
  });

  describe('if the user is logged in with multiple roles', () => {
    const multiRoleState = {
      auth: {
        activeRole: roleTypes.TOO,
      },
      user: {
        isLoading: false,
        userInfo: {
          isLoggedIn: true,
          roles: [
            {
              roleType: roleTypes.CONTRACTING_OFFICER,
            },
            {
              roleType: roleTypes.TOO,
            },
            {
              roleType: roleTypes.TIO,
            },
          ],
        },
      },
    };

    describe('on a page that isn’t the Select Application page', () => {
      it('renders the Select Application link', () => {
        const app = mount(
          <MockProviders initialState={multiRoleState} initialEntries={['/']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        expect(app.containsMatchingElement(<a href="/select-application">Change user role</a>)).toEqual(true);
      });
    });

    describe('on the Select Application page', () => {
      it('does not render the Select Application link', () => {
        const app = mount(
          <MockProviders initialState={multiRoleState} initialEntries={['/select-application']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        expect(app.containsMatchingElement(<a href="/select-application">Change user role</a>)).toEqual(false);
      });
    });
  });

  describe('routing', () => {
    // TODO - expects should look for actual component content instead of the route path
    // Might have to add testing-library for this because something about enzyme + Suspense + routes are not rendering content
    // I FIGURED OUT HOW - need to mock the loadUser (this sets loading back to true and prevents content from rendering)

    const loggedInState = {
      auth: {
        activeRole: roleTypes.PPM,
      },
      user: {
        isLoading: false,
        userInfo: {
          isLoggedIn: true,
          roles: [
            {
              roleType: roleTypes.PPM,
            },
          ],
        },
      },
    };

    const loggedOutState = {
      auth: {
        activeRole: null,
      },
      user: {
        isLoading: false,
        userInfo: {
          isLoggedIn: false,
        },
      },
    };

    it('handles the SignIn URL', () => {
      const app = mount(
        <MockProviders initialState={loggedOutState} initialEntries={['/sign-in']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('Route');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/sign-in');
    });

    it('handles the root URL', () => {
      const app = mount(
        <MockProviders initialState={loggedInState} initialEntries={['/']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('PrivateRoute');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/');
    });

    it('handles the Select Application URL', () => {
      const app = mount(
        <MockProviders initialState={loggedInState} initialEntries={['/select-application']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('PrivateRoute');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/select-application');
    });

    describe('PPM routes', () => {
      const loggedInPPMState = {
        auth: {
          activeRole: roleTypes.PPM,
        },
        user: {
          isLoading: false,
          userInfo: {
            isLoggedIn: true,
            roles: [
              {
                roleType: roleTypes.PPM,
              },
            ],
          },
        },
      };

      it('handles a MoveInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInPPMState} initialEntries={['/queues/new/moves/123']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/queues/:queueType/moves/:moveId');
      });

      it('handles a Queues URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInPPMState} initialEntries={['/queues/new']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/queues/:queueType');
      });

      it('handles a OrdersInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInPPMState} initialEntries={['/moves/123/orders']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveId/orders');

        // OrdersInfo does NOT render the header
        expect(app.find('QueueHeader')).toHaveLength(0);
      });

      it('handles a DocumentViewer URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInPPMState} initialEntries={['/moves/123/documents/abc']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveId/documents/:moveDocumentId?');

        // DocumentViewer does NOT render the header
        expect(app.find('QueueHeader')).toHaveLength(0);
      });
    });

    describe('TOO routes', () => {
      const loggedInTOOState = {
        auth: {
          activeRole: roleTypes.TOO,
        },
        user: {
          isLoading: false,
          userInfo: {
            isLoggedIn: true,
            roles: [
              {
                roleType: roleTypes.TOO,
              },
            ],
          },
        },
      };

      it('handles the moves queue URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTOOState} initialEntries={['/moves/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/queue');
      });

      it('handles the TXOMoveInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTOOState} initialEntries={['/moves/123']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveOrderId');
      });

      it('handles the CustomerDetails URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTOOState} initialEntries={['/too/123/customer/abc']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/too/:moveOrderId/customer/:customerId');
      });
    });

    describe('TIO routes', () => {
      const loggedInTIOState = {
        auth: {
          activeRole: roleTypes.TIO,
        },
        user: {
          isLoading: false,
          userInfo: {
            isLoggedIn: true,
            roles: [
              {
                roleType: roleTypes.TIO,
              },
            ],
          },
        },
      };

      it('handles the invoicing queue URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTIOState} initialEntries={['/invoicing/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/invoicing/queue');
      });

      it('handles the TXOMoveInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTIOState} initialEntries={['/moves/123']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveOrderId');
      });
    });
  });
});
