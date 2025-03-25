
async function updateForm(event) {
    event.preventDefault()

    const form = document.querySelector('.form-update')
    const id = form.getAttribute('id')
    const nome = document.querySelector('#nome').value
    const nickname = document.querySelector('#nickname').value
    const senha = document.querySelector('#senha').value
    const senhaConfirm = document.querySelector('#senha-confirm').value
    const admin = document.querySelector('#admin')

    const resp = validaForm(nome, nickname, senha, senhaConfirm)

    if(!resp.resp){
        alert(resp.message)
        return
    }

    const updateUser = {nome, rg: nickname, senha, admin: admin.checked}

    try{
        const url = `http://localhost:5000/index/usuarios/${id}`
        const response = await fetch(url, {
            method: 'PATCH',
            headers: {
                "Content-Type": "application/json"
              },
            body: JSON.stringify(updateUser)
        })

        if(response.status === 200){
            alert('Usuario atualizado com sucesso!')
        }
    } catch (error) {
        alert('Erro ao atualizar usuario!')
        return
    } finally {
        window.location.href = "/index/usuarios"
    }
}

async function postForm(event) {
    event.preventDefault()

    const nome = document.querySelector('#nome').value
    const nickname = document.querySelector('#nickname').value
    const senha = document.querySelector('#senha').value
    const senhaConfirm = document.querySelector('#senha-confirm').value

    const resp = validaForm(nome, nickname, senha, senhaConfirm)

    if(!resp.resp){
        alert(resp.message)
        return
    }

    const postUser = {nome, rg: nickname, senha, admin: false}

    try{
        const url = `http://localhost:5000/index/usuarios`
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                "Content-Type": "application/json"
              },
            body: JSON.stringify(postUser)
        })

        if(response.status === 200){
            alert('Usuario atualizado com sucesso!')
        }
    } catch (error) {
        alert('Erro ao atualizar usuario!')
        return
    } finally {
        window.location.href = "/index/usuarios"
    }
}

function validaForm(nome, nickname, senha, senhaConfirm){
    const regexNickname = new RegExp("^[A-Za-z0-9_-]{4,}$") 
    const regexNome = new RegExp("^[A-Za-zÀ-ÖØ-öø-ÿ ]{4,}$");

    if(!testeNome(nome, regexNome)) return {resp: false, message: "Nome invalido!"}

    if(!testeNickname(nickname, regexNickname)) return {resp: false, message: "Nome de usuario invalido!"}

    if(senha.trim() === "" || senha.length < 4) return {resp: false, message: "Senha invalida!"}

    if(senhaConfirm.trim() === "" || senhaConfirm !== senha) return {resp: false, message: "Senhas nao coincidem!"}

    return {resp: true, message: ""}
}

function testeNome(nome, regex){
    if(nome.trim() === "") return false
    return regex.test(nome)
}

function testeNickname(nickname, regex){
    if(nickname.trim() === "") return false
    return regex.test(nickname)
}