// src/contexts/EndpointContext.js
import { createContext, useContext, useState } from 'react';

// Cria o contexto para o endpoint e clientId
const EndpointContext = createContext();

export const EndpointProvider = ({ children }) => {
  const [endpoint, setEndpoint] = useState('');
  const [clientId, setClientId] = useState(null);

  return (
    <EndpointContext.Provider value={{ endpoint, setEndpoint, clientId, setClientId }}>
      {children}
    </EndpointContext.Provider>
  );
};

// Hook para facilitar o uso do contexto
export const useEndpoint = () => useContext(EndpointContext);
