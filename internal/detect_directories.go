package internal

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"strings"
)

func mysqlBinaryDirectory() []string {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\WOW6432Node\\MySQL AB", "/s", "/v", "/f", "Location", "/k")
	var outb bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	var datas []string
	var inArray = false
	lines := strings.Split(outb.String(), "\n")
	for i := 0; i < len(lines); i++ {

		if strings.Contains(lines[i], "Location") {
			l := strings.Split(lines[i], "    ")
			data := strings.TrimSpace(l[3])
			for _, v := range datas {
				if data == v {
					inArray = true
				}
			}
			if !inArray {
				datas = append(datas, data)
			}
			inArray = false
		}
	}
	return datas
}

func postgresqlBinaryDirectories() []string {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\PostgreSQL", "/s", "/v", "/f", "Base Directory", "/k")
	var outb bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	var datas []string
	var inArray = false
	lines := strings.Split(outb.String(), "\n")
	for i := 0; i < len(lines); i++ {

		if strings.Contains(lines[i], "Base Directory") {
			l := strings.Split(lines[i], "    ")
			data := strings.TrimSpace(l[3])
			for _, v := range datas {
				if data == v {
					inArray = true
				}
			}
			if !inArray {
				datas = append(datas, data)
			}
			inArray = false
		}
	}
	return datas
}

func oraclelBinaryDirectories() []string {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Oracle", "/s", "/v", "/f", "dllpath", "/k")
	var outb bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	var datas []string
	var inArray = false
	lines := strings.Split(outb.String(), "\n")
	for i := 0; i < len(lines); i++ {

		if strings.Contains(lines[i], "DllPath") {
			l := strings.Split(lines[i], "    ")
			data := strings.TrimSpace(l[3])
			for _, v := range datas {
				if data == v {
					inArray = true
				}
			}
			if !inArray {
				datas = append(datas, data)
			}
			inArray = false
		}
	}
	return datas
}
