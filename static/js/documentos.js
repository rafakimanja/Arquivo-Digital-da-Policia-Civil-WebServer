
async function deletaArquivo(id) {
    let url = `http://localhost:5000/index/documentos/${id}`
    try{
        const resp = await fetch(url, {
            method: 'DELETE'
        })
        if(resp.status === 200) alert('Arquivo deletado com sucesso!')
        if(resp.status === 404) alert('Arquivo nao encontrado!')
    } catch (error){
        alert('Erro ao excluir o documento!')
        return
    } finally {
        window.location.href = "/index/documentos"
    }
}

async function baixaArquivo(id) {
    let url = `http://localhost:5000/index/documentos/download/${id}`

    try {
        const response = await fetch(url);
        
        if (response.status === 404) {
            alert('Arquivo nao encontrado!')
            return;
        }

        //extrai o nome do cabealho
        let contentDisposition = response.headers.get("Content-Disposition")
        let filename = "arquivo.pdf"
        if(contentDisposition){
            let indexIgual = contentDisposition.indexOf("=")
            filename = contentDisposition.slice(indexIgual+1, contentDisposition.length)
        }

        // Converte a resposta para um blob
        const blob = await response.blob();

        // Cria um link temporário para o download
        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = filename
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    } catch (error) {
        return
    }
}