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

// Patch cancelarRota
async function cancelarRota(ENDPOINT, userData) {
  console.log('Cancelando rota pp:', userData);
  console.log('Cancelando rota endpoint:', ENDPOINT);
  try {
    const response = await axios.patch(`${ENDPOINT}/cancelar_rota`, userData);
    return response.data;
  } catch (error) {
    console.error('Error cancelling route:', error.response ? error.response.data : error.message);
    throw error;
  }
}

// Get rotas do cliente
async function getRotasCliente(ENDPOINT, clientId) {
  try {
    const response = await axios.get(`${ENDPOINT}/rotas_cliente`, {
      params: { id: clientId }
    });
    return response.data.rotas; // Retorna a lista de passagens do cliente
  } catch (error) {
    console.error('Error fetching client routes:', error);
    throw error;
  }
}

export { getRotas, create, comparRotas, cancelarRota, getRotasCliente };