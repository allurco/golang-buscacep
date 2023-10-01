package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		request, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao buscar o cep: %v\n", err)
		}

		defer request.Body.Close()

		res, err := io.ReadAll(request.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o resultado: %v\n", err)
		}

		var viaCep ViaCep

		err = json.Unmarshal(res, &viaCep)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		file, err := os.Create(cep + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		}
		defer file.Close()

		bytes, err := file.Write([]byte("CEP: " + viaCep.Cep + "\n Logradouro " + viaCep.Logradouro + "\n Complemento: " + viaCep.Complemento + "\n Bairro: " + viaCep.Bairro + "\n Localidade: " + viaCep.Localidade + "\n UF: " + viaCep.Uf + "\n IBGE: " + viaCep.Ibge + "\n GIA: " + viaCep.Gia + "\n DDD: " + viaCep.Ddd + "\n SIAFI: " + viaCep.Siafi + "\n"))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}

		fmt.Println("Arquivo criado com sucesso! Bytes escritos: ", bytes, "")

	}
}
