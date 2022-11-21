package main

import (
	"bufio"
	"gRPCcalculatorClient/pkg/answerfromserv"
	"gRPCcalculatorClient/pkg/requestparameters"
	"log"
	"os"
)

func main() {
	for {

		//читаем построчно команды
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Printf("can not read: %s\n", err)
		}
		//ищем параметры
		ipAddres, par1, par2, err := requestparameters.GetRequestParameters(reader)
		if err != nil {
			log.Printf("can not read: %s\n", err)
			continue
		}
		//Делаепм запрос на gRPC сервер
		answerfromserv.GetAnswerFromServer(ipAddres, par1, par2)

	}

}
