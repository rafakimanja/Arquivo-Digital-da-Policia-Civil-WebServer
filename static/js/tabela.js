let dados = []

async function buscaArquivos() {
    const url = `http://localhost:5000/index/documentos/json`;
    try {
        const resp = await fetch(url, { method: 'GET' });
        const dados = await resp.json();
        return dados.documentos;
    } catch (erro) {
        alert('Erro ao buscar os arquivos!')
        console.log(erro)        
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
    let ano = document.querySelector('#ano').value;
    let categoria = document.querySelector('#categoria').value

    let anoAtual = new Date().getFullYear()

    console.log(`Nome=${nome} | ano=${ano} | categoria=${categoria}`)

    if (nome.trim() !== '') {
        filtros_ativos.push({type: 'nome', value: nome})
    }

    if ( ano.trim() !== '' && Number(ano) <= anoAtual) {
        filtros_ativos.push({type: 'ano', value: ano})
    }

    if (categoria.trim() !== ''){
        filtros_ativos.push({type: 'categoria', value: categoria})
    }

    document.querySelector('#nome').value = ''
    document.querySelector('#ano').value = ''
    document.querySelector('#categoria').value = ''

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

        if(element.type === 'ano'){
            dadosFiltrados = dadosFiltrados.filter(item => {return item.Ano == element.value})
        }

        if(element.type === 'categoria'){
            dadosFiltrados = dadosFiltrados.filter(item => {return item.Categoria.toLocaleLowerCase() === element.value.toLocaleLowerCase()})
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
            <td>${element.Ano}</td>
            <td>${element.Categoria}</td>
            <td>${new Date(element.UpdatedAt).toLocaleString()}</td>
            <td>
                <button value="${element.ID}" class="btnDownload" onClick="baixaArquivo(${element.ID})">
                    <span class="material-symbols-outlined">download</span>
                </button>
                <a href="/index/documentos/${element.ID}">
                    <button>
                        <span class="material-symbols-outlined">edit</span>
                    </button>
                </a>
                <button value="${element.ID}" class="btnDelete" onClick="deletaArquivo(${element.ID})">
                    <span class="material-symbols-outlined">delete</span>
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
    dados = await buscaArquivos()
    atualizaTabela()
}

carregaDados()