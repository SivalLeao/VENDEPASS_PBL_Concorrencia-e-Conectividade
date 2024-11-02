import PropTypes from 'prop-types';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { useState } from 'react';
import img_default from '/image/icon/rota-de-voo.png';

export const CardItem = ({ title, id, onBuy }) => {
  // Monta o caminho da imagem a partir do título
  const imageSrc = `/image/city/${title}.png`;
  console.log("Caminho da imagem:", imageSrc);

  
  // Estado para armazenar a fonte da imagem atual
  const [currentImage, setCurrentImage] = useState(imageSrc);

  const handleImageError = () => {
    // Se ocorrer um erro ao carregar a imagem, usa a imagem padrão
    setCurrentImage(img_default);
  };

  const handleBuyClick = () => {
    onBuy(id); // Passa o id para a função de compra
  };

  return (
    <Card sx={{ width: 150, height: 250, display: 'flex', flexDirection: 'column', justifyContent: 'space-between' }}>
      <CardMedia
        sx={{ height: 140 }}
        component="img" // Usa o componente img para permitir o evento onError
        image={currentImage} // Usa a imagem atual, que pode ser a do título ou a padrão
        title={title} // Define o título da imagem
        onError={handleImageError} // Chama a função se a imagem não carregar
      />
      <CardContent sx={{ paddingBottom: 0, paddingTop: 1 }}>
        <Typography gutterBottom variant="h5" component="div" sx={{ marginBottom: 0, textAlign: 'center' }}>
          {title.replace(/-/g, ' ')} {/* Substitui "-" por " " apenas na exibição */}
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
