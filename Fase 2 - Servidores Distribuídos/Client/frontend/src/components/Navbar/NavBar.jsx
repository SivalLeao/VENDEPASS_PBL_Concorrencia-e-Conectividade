import { useState, useRef } from 'react';
import { Box, IconButton, InputBase, Paper, useTheme, ClickAwayListener } from '@mui/material';
import Divider from '@mui/material/Divider';
import SearchIcon from '@mui/icons-material/Search';
import AirplaneTicketIcon from '@mui/icons-material/AirplaneTicket';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import logo from '/image/icon/aviao-de-papel.png';
import { useAppThemeContext } from '../../contexts/ThemeContext';
import { SideMenu } from '../SideMenu/SideMenu';
import PropTypes from 'prop-types';
import { Cadastrar } from '../../func/Cadastrar/Cadastrar';

export const NavBar = ({ tickets, handleCancel, credentials }) => {
    const theme = useTheme();
    const { toggleTheme } = useAppThemeContext();

    const [menuOpen, setMenuOpen] = useState(false);
    const [searchUrl, setSearchUrl] = useState('');

    const menuRef = useRef(null);

    const handleMenuToggle = () => {
        setMenuOpen((prev) => !prev);
    };

    const handleMenuClose = () => {
        setMenuOpen(false);
    };

    const handleSearch = async () => {
        if (searchUrl) {
            const id = await Cadastrar(searchUrl, credentials);
            if (id !== -1) {
                console.log(`Cadastro realizado com ID: ${id}`);
            } else {
                console.error('Erro no cadastro.');
            }
        } else {
            console.warn('URL de busca não pode estar vazia.');
        }
    };

    return (
        <ClickAwayListener onClickAway={handleMenuClose}>
            <Box
                marginX={1}
                paddingY={1}
                paddingX={2}
                height={theme.spacing(5)}
                display='flex'
                alignItems='center'
                component={Paper}
                justifyContent="space-between"
                ref={menuRef}
            >
                {/* Logo */}
                <Box display='flex' alignItems='center' gap={1}>
                    <img src={logo} alt="Logo Passcom" width="35" height="35" />
                    <span style={{ fontFamily: 'Roboto, sans-serif', fontSize: '2rem', fontWeight: 700 }}>PASSCOM</span>
                </Box>

                {/* Barra de pesquisa */}
                <Paper
                    component="form"
                    sx={{
                        p: theme.spacing(1),
                        display: 'flex',
                        alignItems: 'center',
                        width: { xs: '100%', sm: theme.spacing(50) },
                        maxWidth: theme.spacing(60),
                        height: theme.spacing(2),
                        borderRadius: theme.spacing(5),
                        flexGrow: 1,
                        mx: { xs: 2, sm: 4 },
                    }}
                >
                    <InputBase
                        sx={{ ml: theme.spacing(1), flex: 1 }}
                        placeholder="Buscar Rotas de Voos"
                        inputProps={{ 'aria-label': 'buscar rotas de voos' }}
                        value={searchUrl}
                        onChange={(e) => setSearchUrl(e.target.value)}
                    />
                    <IconButton
                        type="button"
                        onClick={handleSearch}
                        sx={{
                            p: theme.spacing(1),
                            color: 'primary',
                            '&:focus': { outline: 'none' },
                            '&:hover': { color: '#1565c0' },
                        }}
                        aria-label="search"
                    >
                        <SearchIcon />
                    </IconButton>
                    <Divider sx={{ height: theme.spacing(4), m: theme.spacing(0.5) }} orientation="vertical" />
                </Paper>

                {/* Ícones de Minhas compras / Tema */}
                <Box display="flex" gap={1} alignItems="center">
                    <IconButton
                        aria-label="Minha Passagens"
                        onClick={handleMenuToggle}
                        sx={{
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            gap: 0,
                            '&:hover': { color: '#1565c0' },
                            '&:focus': { outline: 'none' },
                            borderRadius: theme.spacing(0),
                        }}
                    >
                        <AirplaneTicketIcon sx={{ fontSize: '2rem' }} />
                        <span style={{ fontSize: '0.70rem' }}>Passagens</span>
                    </IconButton>

                    <IconButton
                        aria-label="Tema"
                        onClick={toggleTheme}
                        sx={{
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            gap: 0,
                            '&:hover': { color: '#1565c0' },
                            '&:focus': { outline: 'none' },
                            borderRadius: theme.spacing(0),
                        }}
                    >
                        <Brightness4Icon sx={{ fontSize: '2rem' }} />
                        <span style={{ fontSize: '0.70rem' }}>Tema</span>
                    </IconButton>
                </Box>

                {/* Menu Lateral */}
                <SideMenu 
                    open={menuOpen} 
                    onClose={handleMenuClose} 
                    purchasedItems={tickets} 
                    onCancel={handleCancel} 
                />
            </Box>
        </ClickAwayListener>
    );
};

// Validação de props
NavBar.propTypes = {
    tickets: PropTypes.arrayOf(
        PropTypes.shape({
            title: PropTypes.string.isRequired,
            id: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
        })
    ).isRequired,
    handleCancel: PropTypes.func.isRequired,
    credentials: PropTypes.shape({
        name: PropTypes.string.isRequired,
        password: PropTypes.string.isRequired,
    }).isRequired,
};
