import { Box, Drawer, Typography, useTheme } from "@mui/material";
import PropTypes from "prop-types";
import { TicketsList } from '../TicketsList/TicketsList'; // Ajuste o caminho conforme necessário
import LocalActivityIcon from '@mui/icons-material/LocalActivity';

export const SideMenu = ({ open, onClose, purchasedItems, onCancel, endpoint, clientId }) => {
  const theme = useTheme(); // Obtém o tema para usar theme.spacing()

  return (
    <Drawer 
      anchor="right" 
      open={open} 
      onClose={onClose} 
      sx={{ width: theme.spacing(44), flexShrink: 0 }} // Ajusta a largura usando theme.spacing
      PaperProps={{
        sx: { width: theme.spacing(44) } // Define também a largura do conteúdo do Drawer
      }}
    >
      <Box 
        sx={{   
          width: '100%', // Mantém o Box com a largura total do Drawer
          height: '100%', 
          display: 'flex', 
          flexDirection: 'column',
        }}
      >
        <Box 
          sx={{
            display: 'flex',
            alignItems: 'center', // Alinha o ícone e o texto verticalmente
            padding: theme.spacing(2),
            position: 'sticky', // Define o título como fixo
            top: 0, // Posiciona no topo
            backgroundColor: theme.palette.background.paper, // Mantém o fundo consistente
            zIndex: 1, // Garante que o título fique acima dos itens roláveis
          }}
        >
          <LocalActivityIcon sx={{ fontSize: theme.spacing(4), marginRight: theme.spacing(1) }} />
          <Typography variant="h6">Minhas Compras</Typography>
        </Box>
        <TicketsList 
          items={purchasedItems} 
          onCancel={onCancel} 
          endpoint={endpoint} // Passa o endpoint
          clientId={clientId} // Passa o ID do cliente
        /> {/* Usa TicketsList */}
      </Box>
    </Drawer>
  );
};

// Validação de props
SideMenu.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  purchasedItems: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string.isRequired,
      id: PropTypes.string.isRequired,
    })
  ).isRequired,
  onCancel: PropTypes.func.isRequired,
  endpoint: PropTypes.string.isRequired, // Adicione validação para o endpoint
  clientId: PropTypes.number.isRequired, // Adicione validação para o ID do cliente
};