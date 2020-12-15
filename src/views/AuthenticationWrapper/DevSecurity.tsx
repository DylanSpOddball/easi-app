import React, { ReactEventHandler, useEffect, useState } from 'react';
import { OktaContext } from '@okta/okta-react';

const storageKey = 'dev-user-config';

const initialAuthState = {
  isAuthenticated: false,
  isPending: false,
  name: '',
  euaId: '',
  groups: { EASI_D_GOVTEAM: true }
};

type ParentComponentProps = {
  children: React.ReactNode;
};

const DevSecurity = ({ children }: ParentComponentProps) => {
  const [authState, setAuthState] = useState(initialAuthState);
  const [euaId, setEuaId] = useState('');
  const [checkboxValues, setCheckboxValues] = useState({
    EASI_D_GOVTEAM: true
  });

  const authService = {
    login: () => {},
    logout: () => {
      window.localStorage.removeItem(storageKey);
      window.location.href = '/';
    },
    getUser: () => {
      return Promise.resolve({ name: authState.name });
    },
    getTokenManager: () => {
      return {
        off: () => {},
        on: () => {},
        renew: () => {}
      };
    }
  };

  useEffect(() => {
    if (window.localStorage[storageKey]) {
      const state = JSON.parse(window.localStorage[storageKey]);
      setAuthState(as => {
        return {
          ...as,
          name: `User ${state.eua}`,
          isAuthenticated: true,
          euaId: state.eua,
          groups: state.jobCodes
        };
      });
    }
  }, []);

  const handleSubmit: ReactEventHandler = event => {
    event.preventDefault();
    const value = {
      eua: euaId,
      jobCodes: checkboxValues
    };
    localStorage.setItem(storageKey, JSON.stringify(value));
    localStorage.removeItem('okta-token-storage'); // ensure that the dev token is used
    setAuthState({
      ...authState,
      isAuthenticated: true,
      name: `User ${euaId}`,
      euaId,
      groups: checkboxValues
    });
  };

  const checkboxChange: ReactEventHandler<HTMLInputElement> = event => {
    const newCheckboxValues = checkboxValues;
    newCheckboxValues[event.currentTarget.value] = event.currentTarget.checked;
    setCheckboxValues(newCheckboxValues);
  };

  return authState.isAuthenticated ? (
    <OktaContext.Provider value={{ authService, authState }}>
      {children}
    </OktaContext.Provider>
  ) : (
    <form onSubmit={handleSubmit} style={{ padding: '1rem 3rem' }}>
      <h1
        style={{
          backgroundImage: 'linear-gradient(to left, orange, red)',
          WebkitBackgroundClip: 'text',
          WebkitTextFillColor: 'transparent',
          display: 'inline-block'
        }}
      >
        Dev auth
      </h1>
      <p>
        <label>
          Enter an EUA
          <br />
          <input
            type="text"
            maxLength={4}
            minLength={4}
            required
            value={euaId}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
              setEuaId(e.target.value.toUpperCase())
            }
            style={{ border: 'solid 1px orangered' }}
          />
        </label>
      </p>
      <fieldset
        style={{ display: 'inline-block', border: 'solid 1px orangered' }}
      >
        <legend>Select job codes</legend>
        {Object.keys(authState.groups).map(code => (
          <p key={code}>
            <label>
              <input
                type="checkbox"
                value={code}
                onChange={e => checkboxChange(e)}
                checked
              />
              {code}
            </label>
          </p>
        ))}
      </fieldset>
      <p>
        <button
          type="submit"
          style={{
            backgroundImage: 'linear-gradient(to left, orange, red)',
            color: 'white',
            fontWeight: 'bold',
            border: 0,
            padding: '.3rem 1rem',
            borderRadius: '3px',
            cursor: 'pointer'
          }}
        >
          Login
        </button>
      </p>
    </form>
  );
};

export default DevSecurity;
