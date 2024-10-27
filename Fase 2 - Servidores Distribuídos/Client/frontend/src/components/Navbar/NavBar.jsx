import { Box, IconButton, InputBase, Paper, useTheme } from '@mui/material';
import Divider from '@mui/material/Divider';
import SearchIcon from '@mui/icons-material/Search';
import AirplaneTicketIcon from '@mui/icons-material/AirplaneTicket';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import logo  from '/image/aviao-de-papel.png';

export const NavBar = () => {
    const theme = useTheme();
    return (
        // Box é um componente de layout do Material-UI 
        <Box
            marginX={1}
            paddingY={1}
            paddingX={2}
            height={theme.spacing(5)}
            display='flex'
            gap={1}
            alignItems='center'
            //width='100%' // Ocupa toda a largura disponível


        >
            {/* Logo da Passcom */}
            <Box
                display='flex'
                alignItems='center'
                gap={1}
            >
                <img src={logo} alt="Logo Passcom" width="35" height="35" />
                <span style={{fontFamily: 'Roboto, sans-serif', fontSize: '2rem', fontWeight: 700 }}>PASSCOM</span> {/* Texto menor */}
            </Box>

            {/* Barra de pesquisa */}
            <Paper
                component="form"
                sx={{
                    p: theme.spacing(1),
                    display: 'flex',
                    alignItems: 'center',
                    width: theme.spacing(50),
                    height: theme.spacing(2),
                    borderRadius: theme.spacing(5), // Ajuste para bordas arredondadas
                }}
            >
                <InputBase
                    sx={{ ml: theme.spacing(1), flex: 1 }}
                    placeholder="Buscar Rotas de Voos"
                    inputProps={{ 'aria-label': 'buscar rotas de voos' }}
                />
                <IconButton type="button"
                    sx={{
                        p: theme.spacing(1),
                        color: 'primary',
                        '&:focus': {
                            outline: 'none', // Remove a caixa branca de foco ao clicar
                        },
                        '&:hover': {
                            color: '#1565c0', // Cor de fundo ao passar o mouse
                        },
                    }}
                    aria-label="search"
                >
                    <SearchIcon />
                </IconButton>
                <Divider sx={{ height: theme.spacing(4), m: theme.spacing(0.5) }} orientation="vertical" />
            </Paper>

            {/* Minhas compras / Tema */}
            <Box
                display="flex"
                flex={1}
                justifyContent='end'
            >
                <IconButton
                    aria-label="Minha Passagens"
                    sx={{
                        color: 'black', // Cor do ícone
                        display: 'flex', // Para alinhar o ícone e o texto
                        flexDirection: 'column', // Empilha o ícone e o texto verticalmente
                        alignItems: 'center', // Alinha horizontalmente o ícone e o texto
                        gap: 0, // Espaçamento entre o ícone e o texto
                        '&:hover': {
                            color: '#1565c0', // Cor de fundo ao passar o mouse
                        },
                        '&:focus': {
                            outline: 'none', // Remove a caixa branca de foco ao clicar
                        },
                        borderRadius: theme.spacing(0), // Borda arredondada
                    }}
                >
                    <AirplaneTicketIcon sx={{ fontSize: '2rem' }} /> {/* Tamanho do ícone ajustado */}
                    <span style={{ fontSize: '0.70rem' }}>Passagens</span> {/* Texto menor */}
                </IconButton>

                <IconButton
                    aria-label="Minha Passagens"
                    sx={{
                        color: 'black', // Cor do ícone
                        display: 'flex', // Para alinhar o ícone e o texto
                        flexDirection: 'column', // Empilha o ícone e o texto verticalmente
                        alignItems: 'center', // Alinha horizontalmente o ícone e o texto
                        gap: 0, // Espaçamento entre o ícone e o texto
                        '&:hover': {
                            color: '#1565c0', // Cor de fundo ao passar o mouse
                        },
                        '&:focus': {
                            outline: 'none', // Remove a caixa branca de foco ao clicar
                        },
                        borderRadius: theme.spacing(0), // Borda arredondada
                    }}
                >
                    <Brightness4Icon sx={{ fontSize: '2rem' }} /> {/* Tamanho do ícone ajustado */}
                    <span style={{ fontSize: '0.70rem' }}>Tema</span> {/* Texto menor */}
                </IconButton>
            </Box>

        </Box>
    );
}
