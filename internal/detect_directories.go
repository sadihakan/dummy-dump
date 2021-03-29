package internal

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"strings"
)

func mysqlBinaryDirectory() string {
	connstr := fmt.Sprintf("%s:%s@/%s", "root", "deneme332", "deneme")
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	r, err := db.Query("select @@basedir")
	if err != nil {
		log.Fatalln(err)
	}
	var p string
	for r.Next() {
		err = r.Scan(&p)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return p
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
