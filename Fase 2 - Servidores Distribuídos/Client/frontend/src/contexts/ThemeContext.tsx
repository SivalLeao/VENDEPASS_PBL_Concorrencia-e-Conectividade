import { createContext, useCallback, useContext, useMemo, useState } from 'react';
import { ThemeProvider } from '@emotion/react';
import { LightTheme, DarkTheme } from "./../theme";
import { Box } from '@mui/material';
import * as React from 'react';

interface IThemeContextData {
    themeName: 'Light' | 'Dark';
    toggleTheme: () => void;
}

const ThemeContext = createContext({} as IThemeContextData);

export const useAppThemeContext = () => {
    return useContext(ThemeContext);
}

interface AppThemeProviderProps {
    children: React.ReactNode;
}
export const AppThemeProvider: React.FC<AppThemeProviderProps> = ({ children }) => {
    const [themeName, setThemeName] = useState<'Light' | 'Dark'>('Light');
    const toggleTheme = useCallback(() => {
        setThemeName(oldTheme => oldTheme === 'Light' ? 'Dark' : 'Light');       
    }, []);

    const theme = useMemo(() => {
        if (themeName === 'Dark') return LightTheme;
        return DarkTheme;
    }, [themeName])

    return (
        <ThemeContext.Provider value={{ themeName, toggleTheme }}>
            <ThemeProvider theme={theme}>
                <Box width="100vw" height="100vh" bgcolor={theme.palette.background.default}>
                    {children}
                </Box>
            </ThemeProvider>
        </ThemeContext.Provider>
    );
}