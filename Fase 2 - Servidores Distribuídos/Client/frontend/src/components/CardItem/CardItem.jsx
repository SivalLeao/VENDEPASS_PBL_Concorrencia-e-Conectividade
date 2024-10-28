import PropTypes from 'prop-types';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import img_default from '/image/rota-de-voo.png';

export const CardItem = ({ title, id, onBuy }) => {
  const handleBuyClick = () => {
    onBuy(id); // Passa o id para a função de compra
  };

  return (
    <Card sx={{ width: 150, height: 250, display: 'flex', flexDirection: 'column', justifyContent: 'space-between' }}>
      <CardMedia
        sx={{ height: 140 }}
        image={img_default}
        title="img_voos"
      />
      <CardContent sx={{ paddingBottom: 0, paddingTop: 1 }}>
        <Typography gutterBottom variant="h5" component="div" sx={{ marginBottom: 0, textAlign: 'center' }}>
          {title}
        </Typography>
      </CardContent>
      <CardActions sx={{ paddingTop: 0, justifyContent: 'center' }}>
        <Button size="small" onClick={handleBuyClick}>Comprar</Button>
      </CardActions>
    </Card>
  );
};

// Validação de props
CardItem.propTypes = {
  title: PropTypes.string.isRequired,
  id: PropTypes.string.isRequired, // Adicione validação para o id
  onBuy: PropTypes.func.isRequired,
};
