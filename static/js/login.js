function validaLogin(status){
    if(!status){
        alert('Usuario ou senha incorretos. Tente novamente!')
    }
}

function validaForm(event){
    event.preventDefault()

    const nickname = document.querySelector('#nickname').value
    const senha = document.querySelector('#password').value

    const resp = validaCampos(nickname, senha)

    if(!resp.resp){
        alert(resp.message)
        return
    } else {
        event.target.submit()
    }
}

function validaCampos(nickname, senha){
    const regexNickname = new RegExp("^[A-Za-z0-9_-]{4,}$") 

    if(!testeNickname(nickname, regexNickname)) return {resp: false, message: "Nome de usuario invalido!"}

    if(senha.trim() === "" || senha.length < 4) return {resp: false, message: "Senha invalida!"}

}

function testeNickname(nickname, regex){
    if(nickname.trim() === "") return false
    return regex.test(nickname)
}