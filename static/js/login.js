function mostraSenha() {
    const inputSenha = document.querySelector('#password')
    const iconOlho = document.querySelector('#icon-olho');

    if (inputSenha.type === "password") {
        inputSenha.type = "text";
        iconOlho.classList.remove("bi-eye");
        iconOlho.classList.add("bi-eye-slash");
    } else {
        inputSenha.type = "password";
        iconOlho.classList.remove("bi-eye-slash");
        iconOlho.classList.add("bi-eye");
    }
}

function validaForm(event){
    event.preventDefault()

    const nickname = document.querySelector('#nickname').value
    const senha = document.querySelector('#password').value

    const resp = validaCampos(nickname, senha)

    if(!resp.resp){
        event.preventDefault()
        document.querySelector('#mensagem-erro').textContent = resp.message
    } else {
        event.target.submit()
    }
}

function validaCampos(nickname, senha){
    const regexNickname = new RegExp("^[A-Za-z0-9_-]{4,}$") 

    if(!testeNickname(nickname, regexNickname)) return {resp: false, message: "Nome de usuario invalido!"}

    if(senha.trim() === "" || senha.length < 4) return {resp: false, message: "Senha invalida!"}

    return {resp: true, message: ""}
}

function testeNickname(nickname, regex){
    if(nickname.trim() === "") return false
    return regex.test(nickname)
}