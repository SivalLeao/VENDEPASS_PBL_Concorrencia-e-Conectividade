import { Box, Dialog } from '@mui/material';
import { NavBar, CadastroLogin, CardList } from './components';
import { AppThemeProvider } from './contexts/ThemeContext';
import { useState, useEffect } from 'react';
import { getRotas } from './func/userServices/UserServices';

const ENDPOINT = 'http://localhost:8080'; // Defina seu endpoint aqui

function App() {
  const [items, setItems] = useState([]); // Inicializa o estado dos itens como um array vazio
  const [tickets, setTickets] = useState([ 
    { id: '5', title: 'Natal' },
    { id: '6', title: 'Maceió' },
    { id: '7', title: 'Fortaleza' },
  ]);

  const handleCancel = (id) => {
    console.log(`Você cancelou o item com id: ${id}`);
    // Remova o ticket cancelado da lista de tickets
    setTickets((prevTickets) => prevTickets.filter(ticket => ticket.id !== id));
  };

  const handleBuy = (id) => {
    const item = items.find((item) => item.id === id);
    if (item) {
      setTickets((prevTickets) => [...prevTickets, item]);
    }
  };

  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [credentials, setCredentials] = useState({ name: '', password: '' });

  const handleLogin = (name, password) => {
    setCredentials({ name, password });
    console.log("Nome:", name, "Senha:", password);
  };

  const handleCloseLoginModal = () => {
    setCredentials({ name: '', password: '' }); // Limpa as credenciais ao fechar
  };

  const showLoginModal = credentials.name === '' && credentials.password === '';

  // useEffect para buscar os dados das rotas
  useEffect(() => {
    const fetchRotas = async () => {
      try {
        const response = await getRotas(ENDPOINT); // Chama a função getRotas com o ENDPOINT
        const filteredItems = Object.entries(response.rotas)
          .filter(([key, value]) => value === 0)
          .map(([key, value]) => ({ id: key, title: key })); // eslint-disable-line no-unused-vars
        setItems(filteredItems); // Atualiza o estado com os itens filtrados
      } catch (error) {
        console.error('Erro ao buscar rotas:', error);
      }
    };

    fetchRotas(); // Executa a função de busca
  }, []); // Executa apenas uma vez ao montar o componente

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
          />
          <Box marginTop={2} marginLeft={1}>
            <CardList items={items} onBuy={handleBuy} />
          </Box>
          
          {/* Modal de login */}
          <Dialog open={showLoginModal} onClose={handleCloseLoginModal} fullWidth>
            <CadastroLogin open={showLoginModal} onClose={handleCloseLoginModal} onLogin={handleLogin} />
          </Dialog>
        </Box>
      </AppThemeProvider>
    </>
  );
}

export default App;