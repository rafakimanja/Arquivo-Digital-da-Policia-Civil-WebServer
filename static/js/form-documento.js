
async function updateForm(event) {
    event.preventDefault()

    const form = document.querySelector('.form-update')
    const id = form.getAttribute('id')
    const nome = document.querySelector('#nome').value
    const ano = document.querySelector('#ano').value
    const categoria = document.querySelector('#categoria').value
    const arquivo = document.querySelector('#arquivo').files[0]

    let resp

    if(arquivo){resp = validaForm(nome, ano, categoria, arquivo, true)} 
    else {resp = validaForm(nome, ano, categoria, arquivo, false)}

    if(!resp.resp){
        alert(resp.message)
        return
    }

    const formData = new FormData()
    formData.append("nome", nome)
    formData.append("ano", ano)
    formData.append("categoria", categoria)
    formData.append("arquivo", arquivo)

    try{
        const url = `http://localhost:5000/index/documentos/${id}`
        const response = await fetch(url, {
            method: 'PUT',
            body: formData
        })

        if(response.status === 200){
            alert('Arquivo atualizado com sucesso!')
        }
    } catch (error) {
        alert('Erro ao atualizar arquivo!')
        return
    } finally {
        window.location.href = "/index/documentos"
    }
}

async function postForm(event) {
    event.preventDefault()

    const nome = document.querySelector('#nome').value
    const ano = document.querySelector('#ano').value
    const categoria = document.querySelector('#categoria').value
    const arquivo = document.querySelector('#arquivo').files[0]

    const resp = validaForm(nome, ano, categoria, arquivo, true)

    if(!resp.resp){
        alert(resp.message)
        return
    }

    const formData = new FormData()
    formData.append("nome", nome)
    formData.append("ano", ano)
    formData.append("categoria", categoria)
    formData.append("arquivo", arquivo)

    try{
        const url = `http://localhost:5000/index/documentos`
        const response = await fetch(url, {
            method: 'POST',
            body: formData
        })

        if(response.status === 201){
            alert('Arquivo criado com sucesso!')
        }
    } catch (error) {
        alert('Erro ao criar novo arquivo!')
        return
    } finally {
        window.location.href = "/index/documentos"
    }
}

function validaForm(nome, ano, categoria, arquivo, updateArq){
    const regexNome = new RegExp("^[A-Za-z0-9_.-]+$") 
    const regexAno = new RegExp("^[0-9]{4}$")

    if(!testeNome(nome, regexNome)) return {resp: false, message: "Nome invalido!"}

    if(!testeAno(ano, regexAno)) return {resp: false, message: "Ano invalido!"}

    if(categoria === "") return {resp: false, message: "Categoria invalida!"}

    if(updateArq){
        if(!testeArquivo(arquivo)) return {resp: false, message: "Arquivo invalido!"}
    }

    return {resp: true, message: ""}
}

function testeNome(nome, regex){
    if(nome.trim() === "") return false
    return regex.test(nome)
}

function testeAno(ano, regex){
    const anoAtual = new Date().getFullYear()

    if(ano === "") return false
    if(ano < "2000" || ano > anoAtual ) return false
    return regex.test(ano)
}

function testeArquivo(arq){
    if(!arq) return false
    if(!(arq instanceof File)) return false
    if(arq.size > 20_971_520) return false
    return true
}
