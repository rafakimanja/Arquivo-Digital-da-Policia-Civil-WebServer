package services

import (
	"adpc-webserver/src/models"
	"fmt"
	"os"
	"strconv"
)

type Gerenciador struct {
	diretorio string
}

func Construtor() Gerenciador {
	return Gerenciador{diretorio: "./arquivos"}
}

func (g *Gerenciador) GetDiretorio() string {
	return g.diretorio
}

func (g *Gerenciador) criaDiretorio(path string) bool {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (g *Gerenciador) validaDiretorio(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (g *Gerenciador) excluiDiretorio(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (g *Gerenciador) qtdArquivosDiretorio(path string) (int, error) {
	arquivos, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	return len(arquivos), nil
}

func (g *Gerenciador) moveArquivo(origem, destino string) bool {
	err := os.Rename(origem, destino)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}

func (g *Gerenciador) criaCaminho(ano int, categoria string) string {
	anoStr := strconv.Itoa(ano)
	return fmt.Sprintf(g.diretorio + "/" + anoStr + "/" + categoria)
}

func (g *Gerenciador) SalvaArquivo(arq models.Documento) (bool, string) {

	path := g.criaCaminho(arq.Ano, arq.Categoria)

	if g.validaDiretorio(path) {
		return false, path
	} else {
		if g.criaDiretorio(path) {
			return true, path
		} else {
			return false, path
		}
	}
}

func (g *Gerenciador) AtualizarArquivo(arqOri models.Documento, arqDes models.Documento) bool {

	var flag bool

	pathArqOri := g.criaCaminho(arqOri.Ano, arqOri.Categoria) + "/" + arqOri.Arquivo
	pathArqDes := g.criaCaminho(arqDes.Ano, arqDes.Categoria) + "/" + arqDes.Arquivo

	diretorioOrigem := g.criaCaminho(arqOri.Ano, arqOri.Categoria)
	diretorioDestino := g.criaCaminho(arqDes.Ano, arqDes.Categoria)

	if g.validaDiretorio(diretorioDestino) {

		if g.moveArquivo(pathArqOri, pathArqDes) {
			flag = true
		} else {
			flag = false
		}

	} else {

		if g.criaDiretorio(diretorioDestino) {

			if g.moveArquivo(pathArqOri, pathArqDes) {
				flag = true
			} else {
				flag = false
			}
		} else {
			flag = false
		}
	}

	qtd, err := g.qtdArquivosDiretorio(diretorioOrigem)
	if err != nil {
		fmt.Println(err.Error())
		return flag
	}

	if qtd == 0 {
		g.excluiDiretorio(diretorioOrigem)
	}
	return flag
}

func (g *Gerenciador) DeletaArquivo(arquivo models.Documento, temp bool) bool {

	pathDiretorio := g.criaCaminho(arquivo.Ano, arquivo.Categoria)
	pathArquivo := pathDiretorio + "/" + arquivo.Arquivo
	tempPath := g.GetDiretorio() + "/temp"

	if !temp {
		if g.criaDiretorio(tempPath) {
			if g.moveArquivo(pathArquivo, tempPath+"/"+arquivo.Arquivo) {
				return true
			}
		}
		return false
	} else {
		if !g.excluiDiretorio(tempPath) {
			g.moveArquivo(tempPath+"/"+arquivo.Arquivo, pathArquivo)
			return false
		} else {
			qtd, err := g.qtdArquivosDiretorio(pathDiretorio)
			if err == nil {
				if qtd == 0 {
					g.excluiDiretorio(pathDiretorio)
				}
			}
			return true
		}
	}
}
