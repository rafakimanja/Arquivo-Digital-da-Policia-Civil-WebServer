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

async function carregaDados(){
    dados = await buscaArquivos()
    console.log(dados)
    criaTabela(dados)
}

// variaveis
let tabela = document.querySelector('#tabela')

let paginaAtual = 0
const itensPorPagina = 10
// =-=-=-=-=-=-=-=-=-=-=-=-=-=

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

function criaTabela(array){

    let totalPaginasSpan = document.querySelector('#total-paginas')
    let paginaAtualSpan = document.querySelector('#pagina-atual')

    let qtdPaginas = Math.ceil(array.length / itensPorPagina)

    let inicio = paginaAtual * itensPorPagina
    let fim = inicio + itensPorPagina

    let dadosPaginaAtual = array.slice(inicio, fim)

    preencheTabela(dadosPaginaAtual)
    paginaAtualSpan.innerHTML = `${paginaAtual+1}`
    totalPaginasSpan.innerHTML = `${qtdPaginas}`
}

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

carregaDados()