function mostraSenha() {
    const inputSenha = document.getElementById("input-senha");
    const iconOlho = document.getElementById("icon-olho");

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
