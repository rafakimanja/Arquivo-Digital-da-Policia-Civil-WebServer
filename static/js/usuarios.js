async function deletaUsuario(id) {
    let url = `http://localhost:5000/index/usuarios/${id}`
    try{
        const resp = await fetch(url, {
            method: 'DELETE'
        })
        if(resp.status === 200) alert('Usuario deletado com sucesso!')
        if(resp.status === 404) alert('Usuario nao encontrado!')
    } catch (error){
        alert('Ocorreu um erro ao excluir o usuario!')
        return
    } finally {
        window.location.href = "/index/usuarios"
    }
}