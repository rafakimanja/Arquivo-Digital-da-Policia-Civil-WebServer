
let btnLogin = document.querySelector('#btnLogin')

btnLogin.addEventListener('click', Login)

function Login(){
    let usuario = document.querySelector('#rg').value
    let senha = document.querySelector('#senha').value

    alert(`Login User: ${usuario} | ${senha}`)
}