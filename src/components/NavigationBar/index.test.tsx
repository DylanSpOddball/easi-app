import React from 'react';
import { useTranslation } from 'react-i18next';
import { MemoryRouter } from 'react-router-dom';
import { render } from '@testing-library/react';

import NavigationBar, { navLinks } from './index';

jest.mock('react-i18next', () => ({
  // this mock makes sure any components using the translate hook can use it without a warning being shown
  useTranslation: () => {
    return {
      t: (str: String) => str,
      i18n: {
        changeLanguage: () => new Promise(() => {})
      }
    };
  }
}));

describe('The NavigationBar component', () => {
  it('renders without errors', done => {
    const { getByTestId } = render(
      <MemoryRouter initialEntries={['/']}>
        <NavigationBar />
      </MemoryRouter>
    );

    expect(getByTestId('navigation-bar')).toBeInTheDocument();
    done();
  });

  it('displays every navigation element', done => {
    const { getByText } = render(
      <MemoryRouter initialEntries={['/system/making-a-request']}>
        <NavigationBar />
      </MemoryRouter>
    );

    const { t } = useTranslation();

    navLinks.map(link => {
      const linkTitle = t(`header:${link.label}`);
      return expect(getByText(linkTitle)).toBeInTheDocument();
    });
    done();
  });
});
