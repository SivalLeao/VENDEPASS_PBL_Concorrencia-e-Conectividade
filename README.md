<h1 align="center">Redes de Computadores: Cliente Servidor</h1>
<h3 align="center">
    Este projeto foi realizado como parte da disciplina MI - Concorrência e Conectividade, no âmbito do curso de graduação em Engenharia de Computação da Universidade Estadual de Feira de Santana (UEFS).
</h3>

<div id="sobre">
    <h2>Sobre o projeto</h2>
    <div align="justify">
    Neste projeto, um sistema de venda de passagens é desenvolvido em duas fases, utilizando a arquitetura cliente-servidor para permitir uma comunicação eficiente entre múltiplos usuários. Na primeira fase, o sistema é centralizado, com um único servidor responsável por gerenciar as solicitações de compra e cancelamento de passagens, garantindo a integridade dos dados e evitando conflitos durante transações simultâneas. Na segunda fase, o sistema evolui para uma arquitetura com servidores distribuídos, aumentando a escalabilidade e a resiliência do sistema. O cliente, por sua vez, fornece uma interface simples, onde os usuários podem interagir com o sistema de maneira intuitiva.
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
        <li><a href="#projeto1">Projeto 1: Servidor Centralizado</a></li>
        <li><a href="#projeto2">Projeto 2: Servidores Distribuidos</a></li>
    </ul>
</div>

<div id="projeto1">
    <h2>Projeto 1: Servidor Centralizado - Sistema Vendepass: Venda de Passagens</h2>
    <div align="justify">
    <h3>Sobre o projeto</h3>
        <p>
            Neste projeto, um sistema de venda de passagens é desenvolvido utilizando a arquitetura cliente-servidor, permitindo uma comunicação eficiente entre múltiplos usuários. O servidor é responsável por gerenciar as solicitações de compra e cancelamento de passagens, garantindo a integridade dos dados e evitando conflitos durante transações simultâneas. O cliente, por sua vez, fornece uma interface simples, onde os usuários podem interagir com o sistema de maneira intuitiva.
        </p>
        <h3>Requisitos</h3>
        <ul>
            <li>Desenvolver utilizando contêineres Docker para garantir isolamento e portabilidade.</li>
             <li>Nenhum framework de troca de mensagens deve ser usado para implementar a comunicação entre os clientes e o servidor;</li>
            <li>A comunicação do sistema deve implementada usando a interface de socket nativa do TCP/IP</li>
        </ul>
        <ul>
        </ul>
    <a href="[(link)](https://github.com/SivalLeao/VENDEPASS_PBL_Concorrencia-e-Conectividade/blob/main/Fase%201%20-%20Servidor%20Centralizado/README.md)">README do Projeto 1</a>
</div>

<div id="projeto2">
    <h2>Projeto 2: Servidores Distribuidos - Sistema PASSCOM: Venda Compartilhada de Passagens</h2>
    <div align="justify">
    <h3>Sobre o projeto</h3>
        <p>
            Neste projeto, um cliente e um servidor são desenvolvidos para permitir a interação em tempo real, utilizando a linguagem Go e arquiteturas de servidores distribuídos. O servidor, por meio de uma API REST, escuta as requisições de conexão e processa as transações de venda e cancelamento de passagens. O cliente envia comandos e recebe feedback instantâneo sobre suas ações, garantindo uma experiência fluida e eficiente.
        </p>
        <h3>Requisitos</h3>
        <ul>
            <li>Desenvolver utilizando contêineres Docker para garantir isolamento e portabilidade.</li>
            <li>Implementar comunicação via API REST para interação eficiente com servidores distribuídos.</li>
            <li>Evitar o uso de soluções centralizadas, promovendo um sistema totalmente distribuído.</li>
        </ul>
        </ul>
        <ul>
        </ul>
    <a href="(link)">README do Projeto 2 (Não disponíveis)</a>
</div>