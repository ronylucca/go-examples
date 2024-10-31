package main

import fmt "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Empresa struct {
	Nome  string
	CNPJ  string
	Ativo bool
}

func (e Empresa) Desativar() Empresa {
	e.Ativo = false
	fmt.Printf("Desativando empresa %s\n", e.Nome)
	return e
}

func (c Cliente) Desativar() Cliente {
	c.Ativo = false
	fmt.Printf("Desativando cliente %s\n", c.Nome)
	return c
}

func main() {

	rony := Cliente{
		Nome:  "Rony",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 1",
			Numero:     123,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}
	NovoCliente := rony.Desativar()

	fmt.Printf("Cliente %s está ativo? %t\n", NovoCliente.Nome, NovoCliente.Ativo)

}
