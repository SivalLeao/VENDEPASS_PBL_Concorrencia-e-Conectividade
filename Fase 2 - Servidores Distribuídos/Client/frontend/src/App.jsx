import { Box, Dialog } from '@mui/material';
import { NavBar, CadastroLogin, CardList, SideMenu } from './components';
import { AppThemeProvider } from './contexts/ThemeContext';
import { useState, useEffect, useRef } from 'react';
import { getRotas } from './func/userServices/UserServices';

function App() {
  const [items, setItems] = useState([]);
  const [tickets, setTickets] = useState([]);
  const [endpoint, setEndpoint] = useState('');
  const [clientId, setClientId] = useState(null);
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [credentials, setCredentials] = useState({ name: '', password: '' });
  
  // Usamos um ref para armazenar o último estado de rotas, evitando re-renderizações desnecessárias
  const lastRotasRef = useRef([]);

  const handleCancel = (id) => {
    setTickets((prevTickets) => prevTickets.filter(ticket => ticket.id !== id));
  };

  const handleBuy = (id) => {
    const item = items.find((item) => item.id === id);
    if (item) {
      // Não adiciona a compra ao estado tickets
      console.log(`Você comprou o item com id: ${id}`);
    }
  };

  const handleLogin = (name, password) => {
    setCredentials({ name, password });
  };

  const handleCloseLoginModal = () => {
    setCredentials({ name: '', password: '' });
  };

  const showLoginModal = credentials.name === '' && credentials.password === '';

  const fetchRotas = async () => {
    if (!endpoint) return;
    try {
      const response = await getRotas(endpoint);
      const filteredItems = Object.entries(response.rotas)
        .filter(([key, value]) => value === 0)
        .map(([key]) => ({ id: key, title: key }));

      // Só atualiza o estado se as rotas forem diferentes das anteriores
      if (JSON.stringify(filteredItems) !== JSON.stringify(lastRotasRef.current)) {
        setItems(filteredItems);
        lastRotasRef.current = filteredItems; // Atualiza o cache de referência
      }
    } catch (error) {
      console.error('Erro ao buscar rotas:', error);
    }
  };

  useEffect(() => {
    const intervalId = setInterval(fetchRotas, 5000); // Consulta o servidor a cada 5 segundos

    return () => clearInterval(intervalId); // Limpa o intervalo ao desmontar o componente
  }, [endpoint]);

  return (
    <> 
      <AppThemeProvider>
        <Box sx={{ height: '100vh', width: '100vw', overflowY: 'auto', overflowX: 'hidden' }}>
          <NavBar
            tickets={tickets}
            handleCancel={handleCancel}
            onToggleMenu={() => setIsMenuOpen(!isMenuOpen)}
            open={isMenuOpen}
            onClose={() => setIsMenuOpen(false)}
            credentials={credentials}
            setEndpoint={setEndpoint}
            setClientId={setClientId}
            clientId={clientId}
            endpoint={endpoint}
            setTickets={setTickets}
          />
          <Box marginTop={2} marginLeft={1}>
            <CardList items={items} onBuy={handleBuy} endpoint={endpoint} clientId={clientId} />
          </Box>
          
          <SideMenu
            open={isMenuOpen}
            onClose={() => setIsMenuOpen(false)}
            purchasedItems={tickets}
            onCancel={handleCancel}
            endpoint={endpoint}
            clientId={clientId}
          />

          <Dialog open={showLoginModal} onClose={handleCloseLoginModal} fullWidth>
            <CadastroLogin open={showLoginModal} onClose={handleCloseLoginModal} onLogin={handleLogin} />
          </Dialog>
        </Box>
      </AppThemeProvider>
    </>
  );
}

export default App;