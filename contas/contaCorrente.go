package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!"
	} else {
		return "Valor do deposito incorreto!"
	}
}

func (c *ContaCorrente) Transferencia(valorDaTransferencia float64, contaQueRecebe *ContaCorrente) string {
	podeTransferir := valorDaTransferencia <= c.saldo && valorDaTransferencia > 0

	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaQueRecebe.saldo += valorDaTransferencia
		return "Transferencia realizada com sucesso!"
	} else {
		return "erro na Transferencia"
	}
}
