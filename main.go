package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez la longueur souhaitée du mot de passe: ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // enlever le '\n'
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Longueur du mot de passe non valide. Utilisation de 8 caractères par défaut.")
		n = 8
	}
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?@#-_=&;:><"
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, n)
	for i := range password {
		password[i] = characters[rand.Intn(len(characters))]
	}

	fmt.Println("Mot de passe généré:", string(password))

	err = ioutil.WriteFile("password.txt", password, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
