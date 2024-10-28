import { Box } from '@mui/material';
import { CardList, NavBar } from './components';
import { AppThemeProvider } from './contexts/ThemeContext';

function App() {
  const items = [
    { id: '1', title: 'Xique-Xique BA' },
    { id: '2', title: 'São Paulo' },
    { id: '3', title: 'Aracaju' },
    { id: '4', title: 'Feira de Santana' },
    // Adicione mais destinos conforme necessário
  ];

  // Função para lidar com a compra de um item
  const handleBuy = (id) => {
    console.log(`Você comprou o item com id: ${id}`);
    // Aqui você pode adicionar a lógica para processar a compra
  };

  return (
    <>
      <AppThemeProvider>
        <NavBar />
        <Box marginTop={2} marginLeft={1}>
          <CardList items={items} onBuy={handleBuy} />
        </Box>
      </AppThemeProvider>
    </>
  );
}

export default App;
