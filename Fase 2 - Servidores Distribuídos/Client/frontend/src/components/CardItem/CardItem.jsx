import PropTypes from 'prop-types';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import { useState } from 'react';
import img_default from '/image/icon/rota-de-voo.png';
import { comparRotas } from '../../func/userServices/UserServices'; // Atualize o caminho da importação

export const CardItem = ({ title, id, onBuy, endpoint, clientId }) => {
  const imageSrc = `/image/city/${title}.png`;
  console.log("Caminho da imagem:", imageSrc);

  const [currentImage, setCurrentImage] = useState(imageSrc);
  const [notification, setNotification] = useState({ open: false, message: '', severity: '' });

  const handleImageError = () => {
    setCurrentImage(img_default);
  };

  const handleBuyClick = async () => {
    try {
      const response = await comparRotas(endpoint, { Id: clientId, Rota: title });
      console.log('Resposta da comparação de rotas:', response);
      if (response.status) {
        console.log('Compra realizada com sucesso:', response.status);
        onBuy(id);
        setNotification({ open: true, message: 'Compra realizada com sucesso!', severity: 'success' });
      } else if (response.error) {
        console.error('Erro ao realizar a compra:', response.error);
        setNotification({ open: true, message: 'Erro ao realizar a compra.', severity: 'error' });
      }
    } catch (error) {
      console.error('Erro ao fazer a requisição de compra:', error.message);
      if (error.response) {
        console.error('Detalhes do erro:', error.response.data);
      }
      setNotification({ open: true, message: 'Erro ao fazer a requisição de compra.', severity: 'error' });
    }
  };

  const handleCloseNotification = () => {
    setNotification({ ...notification, open: false });
  };

  return (
    <>
      <Card sx={{ width: 150, height: 250, display: 'flex', flexDirection: 'column', justifyContent: 'space-between' }}>
        <CardMedia
          sx={{ height: 140 }}
          component="img"
          image={currentImage}
          title={title}
          onError={handleImageError}
        />
        <CardContent sx={{ paddingBottom: 0, paddingTop: 1 }}>
          <Typography gutterBottom variant="h5" component="div" sx={{ marginBottom: 0, textAlign: 'center' }}>
            {title.replace(/-/g, ' ')}
          </Typography>
        </CardContent>
        <CardActions sx={{ paddingTop: 0, justifyContent: 'center' }}>
          <Button size="small" onClick={handleBuyClick}>Comprar</Button>
        </CardActions>
      </Card>

      {/* Notificação */}
      <Snackbar
        open={notification.open}
        autoHideDuration={4000}
        onClose={handleCloseNotification}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={handleCloseNotification} severity={notification.severity} sx={{ width: '100%' }}>
          {notification.message}
        </Alert>
      </Snackbar>
    </>
  );
};

// Validação de props
CardItem.propTypes = {
  title: PropTypes.string.isRequired,
  id: PropTypes.string.isRequired,
  onBuy: PropTypes.func.isRequired,
  endpoint: PropTypes.string.isRequired,
  clientId: PropTypes.number.isRequired,
};
