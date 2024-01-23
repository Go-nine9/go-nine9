import { createContext, useEffect, useMemo, useState } from 'react';

export const AuthContext = createContext();

// eslint-disable-next-line react/prop-types
export function AuthProvider({ children }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = getCookie('authToken');
    setIsAuthenticated(!!token);
  }, []);

  const login = async (email, password) => {
    try {
      const response = await fetch('http://localhost:8097/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        mode: 'cors',
        body: JSON.stringify({
          email,
          password,
        }),
      });

      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête de connexion');
      }

      const { jwt: authToken } = await response.json();

      if (authToken) {
        setIsAuthenticated(true);
        setCookie('authToken', authToken, 1); 
      }
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  };

  const register = async (firstname, lastname, email, password) => {
    try {
      const response = await fetch('http://localhost:8097/api/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        mode: 'cors',
        body: JSON.stringify({
          firstname,
          lastname,
          email,
          password,
        }),
      });
  
      if (!response.ok) {
        const errorResponse = await response.json();
        throw new Error(errorResponse.message || 'Échec de la requête d\'inscription');
      }
  
      
      const { jwt: authToken } = await response.json();

      if (authToken) {
        setIsAuthenticated(true);
        setCookie('authToken', authToken, 1);
      }
    } catch (err) {
      throw new Error(err.message || 'Une erreur inattendue s\'est produite');
    }
  };

  const logout = () => {
    setIsAuthenticated(false);
    deleteCookie('authToken');
  };

  const authContextValue = useMemo(
    () => ({
      isAuthenticated,
      login,
      logout,
      register,
    }),
    [isAuthenticated],
  );

  return <AuthContext.Provider value={authContextValue}>{children}</AuthContext.Provider>;
}

function setCookie(name, value, days) {
  const expires = new Date(Date.now() + days * 24 * 60 * 60 * 1000).toUTCString();
  document.cookie = `${name}=${value}; expires=${expires}; path=/; secure;`;
}

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}

function deleteCookie(name) {
  document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/; secure;`;
}