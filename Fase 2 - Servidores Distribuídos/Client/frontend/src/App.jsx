import { NavBar } from './components'
import { AppThemeProvider } from './contexts/ThemeContext'

function App() {
  // const [count, setCount] = useState(0)

  return (
    <>
      <AppThemeProvider>
        <NavBar/>
      </AppThemeProvider>
    </>
  )
}

export default App
