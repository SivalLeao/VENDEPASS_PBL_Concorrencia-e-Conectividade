import axios from 'axios';

export const Cadastrar = async (url, nome) => {
  const cadastro = { nome }; // Define o objeto a ser enviado

  // Adicionando logs para verificar os valores
  console.log('URL:', url);
  console.log('Nome:', nome);
  console.log('Cadastro:', cadastro); // Verifica a estrutura do objeto cadastro

  try {
    // Faz a requisição POST usando axios
    const resposta = await axios.post(`${url}/cadastro`, cadastro, {
      headers: { 'Content-Type': 'application/json' },
    
    });

    console.log('Resposta do servidor:', resposta); // Exibe a resposta completa

    // Verifica se o ID está presente na resposta e o retorna
    if (resposta.data && resposta.data.id) {
      console.log('Resposta do servidor:', resposta.data); // Exibe a resposta completa
      return resposta.data.id;
    } else {
      console.error('Erro ao obter o ID do cadastro');
      return -1;
    }
  } catch (error) {
    console.error('Erro ao fazer a requisição POST:', error);
    return -1;
  }
};

  