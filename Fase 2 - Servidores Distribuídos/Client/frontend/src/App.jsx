import { Box } from '@mui/material';
import { CardList, NavBar } from './components';
import { AppThemeProvider } from './contexts/ThemeContext';
import { useState } from 'react';

function App() {
  // Lista de destinos disponíveis para compra
  const items = [
    { id: '1', title: 'Xique-Xique BA' },
    { id: '2', title: 'São Paulo' },
    { id: '3', title: 'Aracaju' },
    { id: '4', title: 'Feira de Santana' },
  ];

  // Lista de tickets comprados simulados
  const [tickets, setTickets] = useState([
    { id: '5', title: 'Natal' },
    { id: '6', title: 'Maceió' },
    { id: '7', title: 'Fortaleza' },
  ]);

  // Função para lidar com o cancelamento de um ticket
  const handleCancel = (id) => {
    console.log(`Você cancelou o item com id: ${id}`);
    
  };

  // Função para lidar com a compra de um item
  const handleBuy = (id) => {
    const item = items.find((item) => item.id === id);
    if (item) {
      setTickets((prevTickets) => [...prevTickets, item]);
    }
  };

  // Estado para controle de abertura do menu lateral
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <> 
      <AppThemeProvider>
        <NavBar 
          tickets={tickets} // Passando a lista de tickets
          handleCancel={handleCancel} // Passando a função de cancelamento
          onToggleMenu={() => setIsMenuOpen(!isMenuOpen)} 
          open={isMenuOpen} 
          onClose={() => setIsMenuOpen(false)} 
        />
        <Box marginTop={2} marginLeft={1}>
          <CardList 
            items={items} 
            onBuy={handleBuy} />
        </Box>
      </AppThemeProvider>
    </>
  );
}

export default App;
