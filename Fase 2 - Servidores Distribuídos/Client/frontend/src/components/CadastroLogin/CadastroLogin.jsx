import * as React from 'react';
import PropTypes from 'prop-types';
import {
  Button,
  FormControl,
  InputLabel,
  OutlinedInput,
  TextField,
  InputAdornment,
  IconButton,
  Box,
  Typography,
  Dialog,
  DialogContent,
} from '@mui/material';
import AccountCircle from '@mui/icons-material/AccountCircle';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { AppProvider } from '@toolpad/core/AppProvider';
import { useTheme } from '@mui/material/styles';

export const CadastroLogin = ({ open, onClose, onLogin }) => {
  const theme = useTheme();
  const [formData, setFormData] = React.useState({ name: '', password: '' });
  const [showPassword, setShowPassword] = React.useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSignIn = (e) => {
    e.preventDefault();
    const { name, password } = formData;
    if (name && password) {
      onLogin(name, password);
    }
  };

  const handleClickShowPassword = () => setShowPassword((show) => !show);
  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };

  return (
    <AppProvider theme={theme}>
      <Dialog open={open} onClose={onClose}>
        <DialogContent>
          <Box
            component="form"
            onSubmit={handleSignIn}
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              justifyContent: 'center',
              bgcolor: theme.palette.background.default,
              padding: 3,
              width: '100%',
              maxWidth: 400,
              margin: 'auto',
            }}
          >
            <AccountCircle sx={{ fontSize: 60, mb: 2, color: '#1565c0' }} />
            
            <Typography variant="h5" component="h1" sx={{ mb: 1 }}>
              Entrar
            </Typography>

            <Typography variant="body2" component="h2" sx={{ mb: 2, textAlign: 'center' }}>
              Aviso: Ao inserir os dados, o sistema criará um novo usuário automaticamente. Se os dados já existirem, o usuário será conectado.
            </Typography>

            <TextField
              id="input-with-icon-textfield"
              label="Nome"
              name="name"
              type="text"
              size="small"
              required
              fullWidth
              onChange={handleChange}
              variant="outlined"
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <AccountCircle fontSize="inherit" />
                  </InputAdornment>
                ),
              }}
            />

            <FormControl sx={{ my: 2 }} fullWidth variant="outlined">
              <InputLabel size="small" htmlFor="outlined-adornment-password">
                Senha
              </InputLabel>
              <OutlinedInput
                id="outlined-adornment-password"
                type={showPassword ? 'text' : 'password'}
                name="password"
                size="small"
                required // Torna o campo de senha obrigatório
                onChange={handleChange}
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton
                      aria-label="toggle password visibility"
                      onClick={handleClickShowPassword}
                      onMouseDown={handleMouseDownPassword}
                      edge="end"
                      size="small"
                    >
                      {showPassword ? <VisibilityOff fontSize="inherit" /> : <Visibility fontSize="inherit" />}
                    </IconButton>
                  </InputAdornment>
                }
                label="Senha"
              />
            </FormControl>

            <Button
              type="submit"
              variant="outlined"
              color="info"
              size="small"
              disableElevation
              fullWidth
              sx={{ my: 2 }}
            >
              Entrar
            </Button>
          </Box>
        </DialogContent>
      </Dialog>
    </AppProvider>
  );
};

// Validação das props com PropTypes
CadastroLogin.propTypes = {
  open: PropTypes.bool.isRequired,  // Estado de abertura do modal
  onClose: PropTypes.func.isRequired,  // Função para fechar o modal (não usada aqui)
  onLogin: PropTypes.func.isRequired,  // Define `onLogin` como uma função obrigatória
};
