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
		reader := bufio.NewReader(os.Stdin)
		if reader != nil {
			newstring, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("can not read: %s\n", err)

				//ищем параметры
				ipAddres, par1, par2, err := requestparameters.GetRequestParameters(newstring)
				if err != nil {
					log.Printf("can not read: %s\n", err)
					continue
				}
				//Делаепм запрос на gRPC сервер
				answerfromserv.GetAnswerFromServer(ipAddres, par1, par2)

			}

		}

	}
}
