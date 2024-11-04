import { CardItem } from '../CardItem/CardItem';
import PropTypes from 'prop-types';
import { Grid2 } from '@mui/material';

export const CardList = ({ items, onBuy }) => {
  // Verifica se items é um array antes de chamar map
  if (!Array.isArray(items)) {
    return null; // Ou exiba uma mensagem de erro ou um componente de fallback
  }

  return (
    <Grid2 container spacing={2}>
      {items.map((item) => (
        <Grid2 item xs={12} sm={6} md={4} lg={3} key={item.id}>
          <CardItem title={item.title} id={item.id} onBuy={onBuy} />
        </Grid2>
      ))}
    </Grid2>
  );
};

// Validação de props
CardList.propTypes = {
    items: PropTypes.arrayOf(
      PropTypes.shape({
        title: PropTypes.string.isRequired,
        id: PropTypes.string.isRequired, // Validação para id
      })
    ).isRequired,
    onBuy: PropTypes.func.isRequired,
};