import { CardItem } from '../CardItem/CardItem';
// import { CardItem } from './CardItem';
import PropTypes from 'prop-types';
import { Grid2 } from '@mui/material';

export const CardList = ({ items, onBuy }) => {
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
