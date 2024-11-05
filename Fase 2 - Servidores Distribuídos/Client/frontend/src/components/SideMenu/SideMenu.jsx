import { Box, Drawer, Typography, useTheme } from "@mui/material";
import PropTypes from "prop-types";
import { TicketsList } from '../TicketsList/TicketsList'; 
import LocalActivityIcon from '@mui/icons-material/LocalActivity';

export const SideMenu = ({ open, onClose, purchasedItems, onCancel, endpoint, clientId }) => {
  const theme = useTheme();

  console.log("Client ID in SideMenu:", clientId); // Verifica o valor de clientId

  return (
    console.log('SM endpoint:', endpoint),
    <Drawer 
      anchor="right" 
      open={open} 
      onClose={onClose} 
      sx={{ width: theme.spacing(44), flexShrink: 0 }}
      PaperProps={{
        sx: { width: theme.spacing(44) }
      }}
    >
      <Box 
        sx={{   
          width: '100%', 
          height: '100%', 
          display: 'flex', 
          flexDirection: 'column',
        }}
      >
        <Box 
          sx={{
            display: 'flex',
            alignItems: 'center',
            padding: theme.spacing(2),
            position: 'sticky',
            top: 0,
            backgroundColor: theme.palette.background.paper,
            zIndex: 1,
          }}
        >
          <LocalActivityIcon sx={{ fontSize: theme.spacing(4), marginRight: theme.spacing(1) }} />
          <Typography variant="h6">Minhas Compras</Typography>
        </Box>
        <TicketsList 
          items={purchasedItems} 
          onCancel={onCancel} 
          endpoint={endpoint} 
          clientId={clientId} // Passa o clientId diretamente
        />
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
  endpoint: PropTypes.string.isRequired,
  clientId: PropTypes.number.isRequired, // Adicione validação para o ID do cliente
};