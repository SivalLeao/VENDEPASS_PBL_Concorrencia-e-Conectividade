import axios from "axios";

// Get rotas
async function getRotas(ENDPOINT) {
  try {
    const response = await axios.get(`${ENDPOINT}/rotas`);
    return response.data; // Retorna apenas os dados
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;
  }
}

// Post cadastro
async function create(ENDPOINT, userData) {
  try {
    const response = await axios.post(`${ENDPOINT}/cadastro`, userData);
    return response.data;
  } catch (error) {
    console.error('Error creating user:', error);
    throw error;
  }
}

// Patch comparRotas
async function comparRotas(ENDPOINT, userData) {
  try {
    const response = await axios.patch(`${ENDPOINT}/comprar_rota`, userData);
    return response.data;
  } catch (error) {
    console.error('Error updating user data:', error.response ? error.response.data : error.message);
    throw error;
  }
}

export { getRotas, create, comparRotas };