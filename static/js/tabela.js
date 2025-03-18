const root = document.getElementById('root');

let limit = 10;
let offset = 0;

async function buscaArquivos(limit, offset) {
    const url = `http://localhost:5000/teste?limit=${limit}&offset=${offset}`;
    try {
        const resp = await fetch(url, { method: 'GET' });
        const dados = await resp.json();
        return dados.documentos;
    } catch (erro) {
        console.log('Erro ao buscar arquivos:', erro);
    }
}

async function carregarDados() {
    const documentos = await buscaArquivos(limit, offset);
    
    if (!documentos) {
        console.error("Nenhum dado encontrado");
        return;
    }

    // Criar a tabela se ainda não existir
    let tabela = document.getElementById('tabela');
    if (!tabela) {
        tabela = document.createElement('table');
        tabela.id = 'tabela';
        tabela.innerHTML = `
            <thead>
                <tr>
                    <th>Nome</th>
                    <th>Categoria</th>
                    <th>Última Alteração</th>
                    <th><th>
                </tr>
            </thead>
            <tbody></tbody>
        `;
        root.appendChild(tabela);
    }

    // Preencher o corpo da tabela
    const tbody = tabela.querySelector('tbody');
    tbody.innerHTML = ''; // Limpar dados anteriores

    documentos.forEach(doc => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${doc.Nome}</td>
            <td>${doc.Categoria || 'N/A'}</td>
            <td>${new Date(doc.UpdatedAt).toLocaleString()}</td>
            <td>
                <button value="${doc.ID}" class="btnDownload" onClick="baixaArquivo(${doc.ID})">
                    <span class="material-symbols-outlined">download</span>
                </button>
                <a href="/index/documentos/${doc.ID}">
                    <button>
                        <span class="material-symbols-outlined">edit</span>
                    </button>
                </a>
                <button value="${doc.ID}" class="btnDelete" onClick="deletaArquivo(${doc.ID})">
                    <span class="material-symbols-outlined">delete</span>
                </button>
            </td>
        `;
        tbody.appendChild(row);
    });

    // Criar botões de paginação se ainda não existirem
    let pagination = document.getElementById('pagination');
    if (!pagination) {
        pagination = document.createElement('div');
        pagination.id = 'pagination';
        pagination.innerHTML = `
            <button id="prevPage">Anterior</button>
            <button id="nextPage">Próximo</button>
        `;
        root.appendChild(pagination);
        
        document.getElementById('prevPage').addEventListener('click', () => {
            if (offset > 0) {
                offset -= limit;
                carregarDados();
            }
        });

        document.getElementById('nextPage').addEventListener('click', () => {
            offset += limit;
            carregarDados();
        });
    }
}

carregarDados();
