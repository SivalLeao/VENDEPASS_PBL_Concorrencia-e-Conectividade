<h1 align="center">VENDEPASS: Venda de Passagens</h1>
<h3 align="center">
    Este projeto foi desenvolvido como parte do primeiro problema da disciplina MI - Concorrência e Conectividade do curso de graduação em Engenharia de Computação na UEFS
</h3>

<div id="sobre">
    <h2>Sobre o projeto</h2>
    <div align="justify">
        O projeto desenvolvido consiste em um sistema de compra de passagens aéreas para diversas localidades, incluindo a opção de cancelamento das compras. O sistema é composto por dois principais componentes: os clientes, responsáveis por solicitar a compra e obter informações sobre as passagens, e o servidor, que realiza o processamento e o armazenamento das passagens adquiridas, bem como a vinculação dessas passagens aos respectivos compradores. Tanto o cliente quanto o servidor foram desenvolvidos na linguagem de programação Go, recomendada por sua eficiência em projetos que envolvem comunicação em redes e tratamento adequado de concorrência.
    </div>
</div>

<h2>Equipe:<br></h2>

<ul>
    <li><a href="https://github.com/avcsilva">Antonio Vitor Costa da Silva</a></li>
    <li><a href="https://github.com/SivalLeao">Sival Leão de Jesus</a></li>
</ul>

<h1 align="center">Sumário</h1>
<div id="sumario">
    <ul>
        <li><a href="#arquitetura">Arquitetura do sistema</a></li>
        <li><a href="#comunicacao">Protocolo de comunicação</a></li>
        <li><a href="#formatação">Formatação e tratamento de dados</a></li>
        <li><a href="#resultados">Resultados</a></li>
        <li><a href="#execucao">Execução do Projeto</a></li>
        <li><a href="#conclusao">Conclusão</a></li>
    </ul>
</div>

<div id="arquitetura">
    <h2>Arquitetura do sistema</h2>
    <div align="justify">
        <p>
            O sistema opera com uma arquitetura de comunicação cliente-servidor, na qual ocorre a troca de mensagens utilizando o protocolo TCP/IP. Esse protocolo assegura que as mensagens sejam entregues ao destino com integridade, garantindo a confiabilidade e a segurança do sistema.
        </p>
        <p>
            Todos os dados recorrentes de procesamento de comprar e cancelamento são atribuido ao servidor, assim quando um cliente se desconectar suas infomacoes ficaram salvas desde que o servidor esteja funcionando, garantido assim uma seguranca nos dados. 
        </p>
        <h3>Servidor</h3>
        <p>
            O servidor é responsável pelo processamento e armazenamento de todas as informações referentes ao funcionamento do sistema. Sendo elas, as rotas de voo dispostas em seu sistema e suas disponibilidades para compra. 
        </p>
        <p>
            Os clientes são cadastrados automaticamente assim que realizam seu primeiro acesso, o sistema armazenará seus dados de compra e permite cancelar voos já comprados. Os usuários não precisam se cadastrar, porém, deverão acessar sua conta pelo mesmo dispositivo.
        </p>
        <p>
            Os voos podem ser comprados por qualquer cliente desde que o voo esteja disponível, ou seja, que nenhum outro cliente tenha a posse.
        </p>
        <p>
            As ações do servidor incluem:
        </p>
        <ol>  
            <li>
            O servidor exibe no terminal o endereço IP e a porta em que está em funcionamento, permitindo que o cliente saiba exatamente onde deve estabelecer a conexão. 
            </li> 
            <li>
                Listar os clientes já previamente conectados e cadastrados, tais como os registrar a partir de ID’s, que lhes são atribuídos no momento de suas conexões.
            </li>
            <li>
                Enviar para o cliente uma listagem das localidades disponíveis para compra e oferecer que uma delas possa ser adquirida.
                <ol type="a">
                    <li>
                        Caso a passagem escolhida pelo cliente esteja disponível, esta logo lhe será atribuída a partir de seu ID.
                    </li>
                    <li>
                        Caso o servidor verifique que a passagem selecionada já foi adquirida por outro cliente, será devolvida uma mensagem de alerta para o cliente envolvendo o ocorrido.
                    </b>
                </ol>
            </li>
            <li>
                Enviar para o cliente a listagem de suas passagens atualmente adquiridas.
                <ol type="a">
                    <li>
                        Caso o cliente não possua nenhuma aquisição, será devolvida uma mensagem indicando esse fato.
                    </li>
                    <li>
                        Caso possua, o cliente pode indicar uma de suas atuais passagens para o devido cancelamento. Com isso, essa passagem será removida da sua lista de aquisições e estará aberta para compra por parte de um novo cliente.
                    </li>
                </ol>
            </li>
            <li>
                Encerrar a conexão de forma segura com um cliente.
            </li>
        </ol>
        É utilizado o protocolo <em>stateful</em>, salvando as informações em variaveis no sistema do servidor, porém é importante frisar que tais informações armazenadas estarão disponíveis apenas enquanto o servidor estiver funcionando. No momento de seu desligamento, todos os registros serão retornados a seus valores padrões.
        <h3>Cliente</h3>
        É a parte do sistema com o qual o usuário irá interagir para realizar suas solicitações, como comprar voos, ver voos comprados e até mesmo cancelá-los. É responsável por oferecer uma interface baseada em terminal para possibilitar que os usuários possam visualizar as informações e inserirem as ações que desejam realizar. Por meio dessa parte do sistema será possível:
        <ol>
            <li>
                Indicar com qual endereço IP e porta se deseja conectar para interação.
            </li>
            <li>
                Solicitar a lista de localidades disponíveis.
                <ol type="a">
                    <li>
                        Selecionar uma das localidades mostradas para aquisição.
                    </li>
                    <li>
                        Retornar para o menu principal.
                    </li>
                </ol>
            </li>
            <li>
                Consultar a lista de passagens já adquiridas.
                <ol type="a">
                    <li>
                        Selecionar uma das passagens para cancelamento.
                    </li>
                    <li>
                        Retornar para o menu principal.
                    </li>
                </ol>
            </li>
            <li>
                Encerrar a conexão de forma segura com o servidor.
            </li>
        </ol>
        O cliente utiliza o protocolo <em>stateless</em>, não possui nenhum armazenamento de dados e realiza processamento apenas para o envio e recebimento de mensagens, tal como processa a exibição da lista de passagens disponíveis, representando com cores quais estão liberadas para compra e quais estão atualmente ocupadas, respectivamente as cores verde e vermelho.
    </div>
</div>

<div id="comunicacao">
    <h2>Protocolo de comunicação</h2>
    <div align="justify">
    <p>
        Toda a comunicação do sistema foi projetada sobre o modelo TCP/IP, tratando-se de uma comunicação orientada a conexão, no qual deve haver a garantia de conexão estabelecida antes de qualquer comunicação, e que toda informação deve ser devidamente entregue e em sua ordem proposta. Além disso, como o servidor armazena informações de quais usuários já se conectaram e cadastraram, tal como quais deles já efetuaram compras de passagens, diz-se que é aplicado o paradigma de serviço stateful, caracterizado por um servidor  que mantém o estado das interações com clientes. Esse método garante que o usuário não perca seus dados, mesmo que o programa seja excluído ou desligado. 
    </p>
    <p>
        O sistema desenvolvido tem como proposto o seguinte protocolo de comunicação, iniciando-se a partir do momento da conexão de um cliente com o servidor:
    </p>
        <ol>
            <li>
                O servidor inicia enviando ao cliente um número de ID, sendo este um número inteiramente novo ou o número já previamente cadastrado, caso seja um cliente em reconexão.
            </li>
            <li>
                O cliente verifica se a primeira mensagem recebida na conexão é um número. Caso seja, é enviada uma resposta de confirmação de reconhecimento para o servidor ("ID_ok").
            </li>
            <li>
                Com a resposta de confirmação ("ID_ok") sendo devidamente validada no lado do servidor, ambos poderão finalmente iniciar a interação com base em solicitações e comandos do usuário.
            </li>
        </ol>
    </p>
    <p>
        Após a realização dessa comunicação inicial, tanto o servidor quanto o cliente estarão em sua etapa da realização de transações de informações sobre as passagens aéreas. A comunicação ocorrerá da seguinte forma, explicitando-se cada uma de suas possíveis etapas:
        <ol>
            <li>
                Todas as mensagens do cliente com destino ao servidor serão compostas pelo seu ID atribuído e o comando que se deseja realizar no momento. Tendo como um exemplo de uma mensagem “1:1” no momento de um usuário na tela de menu principal, significando que o cliente com ID 1 deseja visualizar a lista de passagens para possivelmente realizar uma compra. Caso o servidor verifique que foi recebida uma mensagem num formato diferente desse, a conexão é encerrada automaticamente.
            </li>
            <li>
                Caso seja solicitada uma operação de compra:
                <ol type="a">
                    <li>
                        O servidor inicia a etapa enviando, em formato JSON, um <em>map</em> para o cliente com de todos os possíveis destinos e se estes estão ocupados ou disponíveis.
                    </li>
                    <li>
                        Tendo o cliente recebido a lista, é esperado que responda com um comando de retorno ou com o nome de um dos destinos para compra.
                        <ol type="i">
                            <li>
                                Caso o servidor receba um comando de retorno (por exemplo “1:3”), este retornará para a etapa de menu principal, tal qual fará o cliente.
                            </li>
                            <li>
                                Caso o servidor receba uma informação diferente do comando de retorno, como um possível destino (exemplo: “1:Fortaleza”), este verificará se é possível realizar a operação.
                            </li>
                            <ol>
                                <li>
                                    Caso o destino recebido não exista ou já esteja ocupado, será enviada ao cliente uma mensagem de erro ("Rota inválida!") e ambos cliente e servidor retornarão à etapa de menu principal.
                                </li>
                                <li>
                                    Caso o destino exista e esteja passível de compra, o servidor realizará a operação de compra e responderá ao cliente com uma mensagem de confirmação ("ok"), que por sua vez será validada de forma a informar ao usuário que a operação foi bem sucedida. Após isso, tanto o cliente quanto o servidor retornarão à etapa de menu principal.
                                </li>
                            </ol>
                        </ol>
                    </li>
                </ol>
            </li>
            <li>
                Caso seja solicitada uma operação de consulta:
                <ol type="a">
                    <li>
                        O servidor inicia verificando se o referido cliente possui já registrada alguma compra. Caso não haja compras, será respondido ao cliente com uma mensagem que indique o ocorrido ("Sem passagens compradas"). Entretanto, caso o cliente possua passagens registradas, o servidor enviará uma mensagem de confirmação de posse ("ok").
                    </li>
                    <li>
                        Tendo o cliente recebido a mensagem de confirmação, este estará esperando, por parte do servidor, novamente em JSON, uma lista de passagens registradas para o ID do referido cliente. Com a lista de passagens recebida, desserializada e devidamente tratada para exibição em terminal, é esperado que o cliente responda com um comando de retorno ou com o nome de uma das possíveis passagens.
                        <ol type="i">
                            <li>
                                Caso o servidor receba comando de retorno (exemplo: "1:3"), ambas as partes do sistema retornarão para a etapa de menu principal.
                            </li>
                            <li>
                                Caso o servidor receba uma informação diferente do comando de retorno (exemplo: "1:Maceio"), este verificará se foi possível realizar a operação de cancelamento de passagem com o nome recebido.
                                <ol>
                                    <li>
                                        Caso o nome recebido não exista na lista ou pertença a algum outro cliente, será enviada ao cliente uma mensagem indicando o erro ("Rota inválida!"), e ambos retornarão à etapa de menu principal.
                                    </li>
                                    <li>
                                        Caso o nome exista e pertença ao cliente em questão, a operação será realizada e o servidor responderá com uma mensagem de confirmação ("ok"). Após isso, ambos irão retornar à etapa de menu principal.
                                    </li>
                                </ol>
                            </li>
                        </ol>
                    </li>
                </ol>
            </li>
            <li>
                Caso seja solicitado o encerramento da conexão:
                <ol type="a">
                    <li>
                        O servidor inicia enviando uma mensagem de confirmação para encerramento de conexão ("exit_ok") e, em seu lado, encerra a conexão.
                    </li>
                    <li>
                        O cliente, por sua vez, tendo recebido e validado corretamente a mensagem de confirmação, exibirá em sua interface tal confirmação e encerrará a execução do programa. Caso, por algum motivo, receba uma mensagem diferente da confirmação, exibirá em sua interface uma mensagem de erro e continuará com a execução do serviço, retornando à etapa de menu principal.
                    </li>
                </ol>
            </li>
            <li>
                Em caso de um comando inválido (exemplo: "1:5", "1:6", etc) o servidor enviará uma mensagem informando o problema ("Operação inválida!"), de forma que o cliente a receba e exiba em seu terminal para indicar ao usuário o ocorrido.
            </li>
        </ol>
        </li>
    </p>
    <p>
        Para a realização da conexão com diversos clientes simultaneamente, foram utilizadas as chamadas goroutines, disponíveis nativamente na linguagem Go, que possibilitam a utilização de threads para execução de processos em paralelo. Com isso, é possível que mais de um cliente se conecte e comunique com o servidor simultaneamente e, graças às funcionalidades oferecidas pela tecnologia da linguagem, sem haver problemas de choques de comunicação ou problemas de concorrência. 
    </p>
    <p>
        Dessa forma, se dois usuários tentarem comprar a mesma passagem simultaneamente, apenas um deles conseguirá realizar a compra, devido ao tratamento adequado de concorrência proporcionado pela linguagem.
    </p>
</div>

<div id="formatação">
    <h2>Formatação e tratamento de dados</h2>
    <div align="justify">
        <p>
            Para o correto funcionamento da comunicação cliente-servidor, é essencial definir o formato dos dados que serão enviados e recebidos por ambos. Para isso, foram analisadas as estruturas disponíveis na linguagem, com o objetivo de transmitir apenas os dados necessários, minimizando o volume de envio para atender ao problema proposto. Optou-se por utilizar um map tanto na comunicação do servidor para o cliente quanto do cliente para o servidor, pois essa estrutura permite o envio de dados associados, como o nome das rotas, disponibilidade, e as requisições do usuário vinculadas ao seu ID.
        </p>
        <p>
            Antes de qualquer envio, os dados são serializados e convertidos em um arquivo JSON, seguindo o formato de um <em>map</em> em Go. O destinatário, por sua vez, desserializa a mensagem para tratá-la adequadamente. Devido à utilização de uma estrutura de dados específica da linguagem, tanto o servidor quanto o cliente devem estar implementados em Go para garantir a compatibilidade na comunicação.
        </p>
    </div>
</div>

<div id="resultados">
    <h2>Resultados</h2>
    <div align="justify">
        <p>
            Tendo sido testado em laboratório com uso de diversos computadores para simular a conexão simultânea de múltiplos clientes, tal como com um <em>script</em> que realiza uma execução teste programada de múltiplos clientes ao mesmo tempo, foi possível averiguar que o sistema consegue lidar correta e eficientemente com as diversas comunicações ocorrendo simultaneamente, não apresentando nenhum tipo de atraso ou travamento, tal como se era esperado em teoria de acordo com as tecnologias oferecidas pela linguagem Go. Além disso, foi possível comprovar que o servidor foi capaz de reconhecer corretamente cada cliente que se conectou e reconectou, sendo possível a recuperação dos dados e compras de cada usuário simulado.
        </p>
        <p>
            Em laboratório, foi testado também o que ocorria com o funcionamento do sistema caso um dos clientes conectados perdesse sua conexão de maneira forçada, como a remoção de um cabo de rede. Com isso, foi averiguado que o servidor mantém seu funcionamento normalmente, podendo ainda se comunicar com outros clientes, e encerrando automaticamente a comunicação com o cliente que perdeu sua conexão, graças à funcionalidade de <em>timeout</em> embutida na biblioteca utilizada para a realização das conexões e comunicações.
        </p>
        <p>
            Entretanto, da maneira como projeto foi concebido, o cliente em questão que teve sua conexão perdida não consegue reconhecer o erro relatado em tempo real, mantendo a execução do programa na etapa em que parou, até que se tente enviar algum comando. Somente após a tentativa de enviar alguma mensagem o programa reconhece a perda da conexão e exibe uma mensagem de erro, solicitando em seguida um endereço alvo para realizar uma nova conexão. Caso o cliente receba de volta sua conexão com a rede, como tendo seu cabo de rede posto de volta, após o servidor ter encerrado sua conexão, esta não será iniciada novamente de forma automática. O usuário do cliente deverá indicar novamente o endereço alvo para poder se reconectar ao servidor e recuperar seus dados.
        </p>
        <p>
            Uma considerável porção do código fonte do projeto possui documentação sobre suas operações, indicando o que cada parte ou linha de código deve estar realizando para o funcionamento do sistema.
        </p>
    </div>
</div>

<div id="execucao">
    <h2>Execução do Projeto</h3>
    <div align="justify">
    <h3>Abrir o Terminal</h3>
    <p>
        Este projeto deve ser executado no terminal do sistema operacional ou em IDEs Ambientes de Desenvolvimento Integrado (Integrated Development Environments).
    </p>
    <p>
    Para abrir o terminal: 
    <li>
        No Windows, pressione as teclas <code>Windows + R</code>, digite cmd na janela que abrir e confirme.
    </li>
    <li>
        No Linux, pressione as teclas <code>Ctrl + Alt + T</code> para abrir o terminal. 
    </li>
    Com o terminal aberto, navegue até o diretório onde os arquivos foram baixados utilizando o comando <code>cd</code>, por exemplo,
    </p>
    <p> 
    <code>cd C:\VENDEPASS_PBL_Concorrencia-e-Conectividade\Servidor</code>
    </p>
    <h3>Sem docker</h3>
    <p>
        Para executar o projeto sem Docker, primeiramente, é necessário configurar o ambiente de trabalho instalando a linguagem de programação <a href="https://go.dev/doc/install">Go</a>. Em seguida, faça o download dos arquivos disponibilizados neste repositório.
    </p>
    <p>
        Deve ser aberto um terminal para cada código, e cada um possui um diretório diferente.
    </p>
    <p>
        O primeiro arquivo a ser executado deve ser o servidor. Embora o cliente possa ser iniciado primeiro, o servidor é quem comunica o endereço da conexão.
    </p>
    <p> 
    Para iniciar o servidor, insira o seguinte comando no terminal:

    go run server.go 

O servidor estará funcionando e exibirá o IP e a porta da conexão. Após o servidor ser iniciado, não será possível interagir diretamente com ele, apenas visualizar suas saídas.
    </p>
    <p align="center">
      <img src="img/Tela inicial do servidor.jpeg" width = "400" />
    </p>
    <p align="center"><strong>Tela inicial do servidor</strong></p>
    <h3>Cliente</h3>
    <p>
    Para iniciar o cliente, insira o comando no terminal:

    go run client.go

Logo após, será solicitado que insira o endereço da conexão exatamente como foi informado pelo servidor, incluindo todos os pontos separadores.
    <p align="center">
    <img src="img/solicitacao de endereco.png" width = "400"/>
    </p>
    <p align="center"><strong>Tela de solicitação para se conectar ao servidor</strong></p>


O menu do cliente será exibido, permitindo que o usuário interaja com o sistema utilizando os números do teclado para selecionar as opções desejadas.
    <p align="center">
    <img src="img/menu do usuario.png" width = "400"/>
    </p>
    <p align="center"><strong>Menu do usuário</strong></p>
</p>
    <h3>Com Docker</h3>
    <p>
        Para executar o projeto, com Docker é necessário ter o docker instalado na sua máquina, em seguida baixar os arquivos dispostos neste repositório.
    </p>
    <h3>Servidor</h3>
    <p>
        Para utilizar os arquivos em contêiner é necessário primeiro criar a imagem docker.

Utilize o comando para gerar a imagem:

    docker build -t server .

Para executar a imagem, roda a aplicação em container, utilize:  

    docker run -it -p 8088:8088 server

O código será executado e exibirá o endereço e porta, similar ao funcionamento sem docker, e os mesmo procedimentos deverão ser seguido
</p>
    <h3>Cliente</h3>
    <p>
        Para iniciar o cliente, crie a imagem utilizando o comando a seguir:
        
    docker build -t client .
Para executar a imagem: 
    
    docker run -it --rm client

Logo após, será solicitado que você insira o endereço da conexão exatamente como foi informado pelo servidor, incluindo todos os pontos separadores. 

O menu do cliente será exibido, permitindo que o usuário interaja com o sistema utilizando os números do teclado para selecionar as opções desejadas.
</p>
    <h3>Comprar/Cancelar Compra</h3>
    <p>
        Na tela que apresenta os nomes das cidades disponíveis para compra ou cancelamento de passagens, é importante que o nome da cidade seja digitado exatamente como está exibido, respeitando letras maiusculas e/ou minúsculas e eventuais assentos.
        <p align="center">
    <img src="img/comprar passagem.png" width = "400"/>
    </p>
    <p align="center"><strong>Comprando passagem</strong></p>
    <p align="center">
    <img src="img/cancelando compra.png" width = "400"/>
    </p>
    <p align="center"><strong>Cancelando compra de passagem</strong></p>
    </p>
</div>

<div id="conclusao">
    <h2>Conclusão</h2>
    <div align="justify">
        <p>
            De acordo com os resultados obtidos em testes em laboratório, é possível afirmar que o produto cumpre com o que se propõe inicialmente. Com a execução correta do servidor e do cliente, é possível realizar e cancelar compras de passagens mesmo que haja a presença de diversos usuários simultâneos, com o servidor encarregado de realizar todo o processamento e tratamento de concorrência para o caso de requisições coincidentes de múltiplos usuários.
        </p>
        <p>
            Ainda é possível aprimorar o sistema, como implementando uma reconexão automática para o cliente e servidor em caso de perda de rede. Porém, o projeto ainda consegue lidar adequadamente com suas outras propostas, sendo assim bem favorável para a sua utilização.
        </p>
    </div>
</div>
