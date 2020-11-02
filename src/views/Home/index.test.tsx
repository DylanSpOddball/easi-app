import React from 'react';
import { act } from 'react-dom/test-utils';
import { Provider } from 'react-redux';
import { MemoryRouter } from 'react-router-dom';
import { mount, ReactWrapper, shallow } from 'enzyme';
import configureMockStore from 'redux-mock-store';

import ActionBanner from 'components/shared/ActionBanner';
import { businessCaseInitialData } from 'data/businessCase';
import { initialSystemIntakeForm } from 'data/systemIntake';

import Home from './index';

jest.mock('@okta/okta-react', () => ({
  useOktaAuth: () => {
    return {
      authState: {
        isAuthenticated: true
      },
      authService: {
        getAccessToken: () => Promise.resolve('test-access-token'),
        getUser: () =>
          Promise.resolve({
            name: 'John Doe'
          })
      }
    };
  }
}));

describe('The home page', () => {
  it('renders without crashing', () => {
    const mockStore = configureMockStore();
    const store = mockStore({});
    shallow(
      <MemoryRouter initialEntries={['/']} initialIndex={0}>
        <Provider store={store}>
          <Home />
        </Provider>
      </MemoryRouter>
    );
  });

  describe('User is logged in', () => {
    it('displays login button', async () => {
      const mockStore = configureMockStore();
      const store = mockStore({
        systemIntakes: {
          systemIntakes: []
        },
        businessCases: {
          businessCases: []
        }
      });
      let component: ReactWrapper;

      await act(async () => {
        component = mount(
          <MemoryRouter initialEntries={['/']} initialIndex={0}>
            <Provider store={store}>
              <Home />
            </Provider>
          </MemoryRouter>
        );

        component.update();

        expect(component.find('a[children="Start now"]').exists()).toEqual(
          true
        );
      });
    });

    it('displays banners for intakes and biz cases', async () => {
      const mockStore = configureMockStore();
      const store = mockStore({
        systemIntakes: {
          systemIntakes: [
            {
              ...initialSystemIntakeForm,
              id: '1'
            },
            {
              ...initialSystemIntakeForm,
              id: '2',
              status: 'INTAKE_SUBMITTED'
            },
            {
              ...initialSystemIntakeForm,
              id: '3'
            },
            {
              ...initialSystemIntakeForm,
              id: '4',
              status: 'INTAKE_SUBMITTED',
              businessCaseId: '1'
            }
          ]
        },
        businessCases: {
          businessCases: [
            {
              ...businessCaseInitialData,
              id: '1',
              systemIntakeId: '4'
            }
          ]
        }
      });
      let component: ReactWrapper;

      await act(async () => {
        component = mount(
          <MemoryRouter initialEntries={['/']} initialIndex={0}>
            <Provider store={store}>
              <Home />
            </Provider>
          </MemoryRouter>
        );

        component.update();
        expect(component.find(ActionBanner).length).toEqual(4);
      });
    });
  });
});
