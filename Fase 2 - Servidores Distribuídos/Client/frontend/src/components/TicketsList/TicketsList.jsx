import { TicketsItem } from "../TicketsItem/TicketsItem";
import PropTypes from "prop-types";
import { Box, Typography } from "@mui/material";

export const TicketsList = ({ items, onCancel, endpoint, clientId }) => {
  return (
    <Box display="flex" flexDirection="column" gap={2}>
      {items.length > 0 ? (
        items.map((item) => (
          <Box key={item.id} width="100%">
            <TicketsItem title={item.title} id={item.id} onCancel={onCancel} endpoint={endpoint} clientId={clientId} />
          </Box>
        ))
      ) : (
        <Box width="100%">
          <Typography>Nenhum item comprado.</Typography>
        </Box>
      )}
    </Box>
  );
};

// Validação de props
TicketsList.propTypes = {
  items: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string.isRequired,
      id: PropTypes.string.isRequired,
    })
  ).isRequired,
  onCancel: PropTypes.func.isRequired,
  endpoint: PropTypes.string.isRequired, // Adicione validação para o endpoint
  clientId: PropTypes.number.isRequired, // Adicione validação para o ID do cliente
};