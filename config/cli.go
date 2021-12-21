package config

import (
	"bufio"
	"fmt"
	"github.com/bykovme/goconfig"
	"log"
	"os"
	"strconv"
	"strings"
)

const welcomeText = "controller.conf is not found when starting the application," +
	" trying create new one"

func (c *Config) CreateConfByUser() error {
	fmt.Println(welcomeText)
	if err := c.inputServerPort(); err != nil {
		log.Println(err)
		return err
	}
	if err := c.CreateNodes(); err != nil {
		log.Println(err)
		return err
	}

	if err := c.SaveConfig(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func isInputEmpty(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	return len(input) == 0
}

func (c *Config) inputServerPort() (err error) {
	for isInputEmpty(c.ServerPort) {
		fmt.Println("input serverPort: ")
		_, err = fmt.Scan(&c.ServerPort)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (c *Config) RemoveNodes() (err error) {
	fmt.Println("list of nodes:")
	for _, node := range c.Nodes {
		fmt.Println("node uuid:", node.UUID)
	}

	fmt.Println("input uuids for delete")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	isExist := func(str string) bool {
		for _, s := range line {
			if s == str {
				return true
			}
		}
		return false
	}

	var nodeArr []Node
	for _, node := range c.Nodes {
		if !isExist(node.UUID) {
			nodeArr = append(nodeArr, node)
		}
	}

	c.Nodes = nodeArr

	if err := c.SaveConfig(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Config) CreateNodes() (err error) {
	var strInput string
	fmt.Println("input list of nodes uuid\n", "how many nodes?")
	fmt.Scan(&strInput)

	intInput, err := strconv.ParseInt(strInput, 10, 64)
	if err != nil {
		log.Println(err)
		intInput = 0
	}

	for i := int64(0); i < intInput; i++ {
		var node Node
		for isInputEmpty(node.UUID) {
			fmt.Println("input uuid: ")
			_, err := fmt.Scan(&node.UUID)
			if err != nil {
				return err
			}
			fmt.Println("input blockchain: ")
			_, err = fmt.Scan(&node.Blockchain)
			if err != nil {
				return err
			}
		}

		c.Nodes = append(c.Nodes, node)
	}

	if err := c.SaveConfig(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Config) SaveConfig() error {
	err := goconfig.SaveConfig(c.usrHomePath+cConfigPath, c)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("config path:", c.usrHomePath+cConfigPath)
	return nil
}
