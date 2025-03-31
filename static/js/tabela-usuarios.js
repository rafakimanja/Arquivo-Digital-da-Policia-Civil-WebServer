let dados = []

async function buscaUsuarios() {
    const url = `http://${env.IP_SERVIDOR}:${env.PORTA_SERVIDOR}/index/usuarios/json`;
    try {
        const resp = await fetch(url, { method: 'GET' });
        const dados = await resp.json();
        return dados.usuarios;
    } catch (erro) {
        alert('Erro ao buscar os usuarios!')
    }
}

// variaveis
let tabela = document.querySelector('#tabela')
let divFiltrosAtivos = document.querySelector('#filtros-ativos')


let filtros_ativos = []

let paginaAtual = 0
const itensPorPagina = 10
// =-=-=-=-=-=-=-=-=-=-=-=-=-=

//filtros
function adicionaFiltro() {
    let nome = document.querySelector('#nome').value;

    if (nome.trim() !== '') {
        filtros_ativos.push({type: 'nome', value: nome})
    }

    document.querySelector('#nome').value = ''

    exibeFiltrosAtivos()
    atualizaTabela()
}


function deletaFiltro(id){
    filtros_ativos.splice(id, 1)
    exibeFiltrosAtivos()
    atualizaTabela()
}


function exibeFiltrosAtivos() {
    
    if (filtros_ativos && filtros_ativos.length > 0) {

        divFiltrosAtivos.innerHTML = '';

        filtros_ativos.forEach((itemFiltro, index) => {
            const campoFiltro = document.createElement('div');
            campoFiltro.classList.add("campo-filtro");
            campoFiltro.innerHTML = `<p>${itemFiltro.type}</p> <button onclick="deletaFiltro(${index})">X</button>`;
            divFiltrosAtivos.appendChild(campoFiltro);
        })
    } else {
        divFiltrosAtivos.innerHTML = '<p>Nenhum filtro ativo</p>';
    }
}


function filtraDados() {
        
    let dadosFiltrados = [...dados]
        
    filtros_ativos.forEach(element => {
        if(element.type === 'nome'){
            dadosFiltrados = dadosFiltrados.filter(item => {return item.Nome.toLocaleLowerCase().includes(element.value.toLocaleLowerCase())})
        }
    });

    return dadosFiltrados
}
// =-=-=-=-=-=-=-=-=-=-=-=-=-=

// tabela
function criaTabela(dados){

    let totalPaginasSpan = document.querySelector('#total-paginas')
    let paginaAtualSpan = document.querySelector('#pagina-atual')

    let qtdPaginas = Math.ceil(dados.length / itensPorPagina)

    let inicio = paginaAtual * itensPorPagina
    let fim = inicio + itensPorPagina

    let dadosPaginaAtual = dados.slice(inicio, fim)

    preencheTabela(dadosPaginaAtual)
    paginaAtualSpan.innerHTML = `${paginaAtual+1}`
    totalPaginasSpan.innerHTML = `${qtdPaginas}`
}

function preencheTabela(dados){
    
    let tbody = document.querySelector('#content')
    tbody.innerHTML = ''

    if(dados.length < 1) {
        tabela.innerHTML = '<h3>Sem dados cadastrados!</h3>'
        return
    }

    dados.forEach(element => {
        const linha = document.createElement('tr')
        linha.innerHTML = `
            <td>${element.Nome}</td>
            <td>${element.Admin ? 'Sim' : 'Não'}</td>
            <td>
                <a href="/index/usuarios/${element.ID}">
                    <button class="btn btn-warning">
                        <i class="bi bi-pencil"></i>
                    </button>
                </a>
                <button value="${element.ID}" class="btn btn-danger" onClick="deletaUsuario(${element.ID})">
                    <i class="bi bi-trash"></i>
                </button>
            </td>
        `
        tbody.appendChild(linha)
    });
}

function atualizaTabela(){
    if(filtros_ativos.length > 0){
        const dadosFiltrados = filtraDados()
        criaTabela(dadosFiltrados)
    } else {
        criaTabela(dados)
    }
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=


// paginacao
function voltaDados(){
    if(paginaAtual > 0){
        paginaAtual--;
        criaTabela(dados)
    }
}

function avancaDados(){
    let maxPaginas = Math.ceil(dados.length / itensPorPagina)
    if(paginaAtual < maxPaginas-1){
        paginaAtual++
        criaTabela(dados)
    }
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-


async function carregaDados(){
    dados = await buscaUsuarios()
    atualizaTabela()
}

carregaDados()
