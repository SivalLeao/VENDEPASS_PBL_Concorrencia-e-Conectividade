
import { Box, Dialog } from '@mui/material';
import { NavBar, CadastroLogin, CardList } from './components';
import { AppThemeProvider } from './contexts/ThemeContext';
import { useState } from 'react';

function App() {
  const items = [
    { id: '1', title: 'Porto-Alegre' },
    { id: '2', title: 'Rio-de-Janeiro' },
    { id: '3', title: 'Belo-Horizonte' },
    { id: '4', title: 'Cachoeira' },
    { id: '9', title: 'Sao-Paulo' },
    { id: '10', title: 'Serrinha' },
    { id: '11', title: 'Salvador' },
    { id: '12', title: 'Feira-de-Santana' },
    { id: '12', title: 'Xique-Xique' },
    { id: '12', title: 'Aracaju' },
    { id: '12', title: 'Maceio' },
    { id: '12', title: 'Recife' }, 
    { id: '12', title: 'Fortaleza' },
    { id: '12', title: 'Acre' },
    { id: '12', title: 'Manaus' },
  ];

  const [tickets, setTickets] = useState([
    { id: '5', title: 'Natal' },
    { id: '6', title: 'Maceió' },
    { id: '7', title: 'Fortaleza' },
  ]);

  const handleCancel = (id) => {
    console.log(`Você cancelou o item com id: ${id}`);
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

  const showLoginModal = credentials.name === '' && credentials.password === '';

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
        />
        <Box marginTop={2} marginLeft={1}>
          <CardList items={items} onBuy={handleBuy} />
        </Box>
        
        {/* Modal de login */}
        <Dialog open={showLoginModal} onClose={() => {}} fullWidth>
          <CadastroLogin open={showLoginModal} onClose={() => {}} onLogin={handleLogin} />
        </Dialog>
        </Box>
      </AppThemeProvider>
    </>
  );
}

export default App;
