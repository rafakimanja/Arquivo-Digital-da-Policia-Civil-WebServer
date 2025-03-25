function copiarEmail() {
    const email = "tidposorio@gmail.com";
    
    navigator.clipboard.writeText(email)
        .then(() => alert("E-mail copiado para a área de transferência!"))
        .catch(err => err);
}