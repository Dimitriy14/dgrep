package translate

import (
	"bufio"
	"os"
	"strings"

	"github.com/astaxie/beego"
)

func New() *Translate {
	return &Translate{}
}

type Translate struct {
	beego.Controller
	Path string
	Lang string
}

func (b *Translate) Set() {
}

func (r *Translate) readFile() ([]string, error) {
	file, err := os.Open(r.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	in := bufio.NewScanner(file)
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return lines, in.Err()
}

func (r *Translate) funcMakeMapTrans() (map[string]string, error) {
	strngs, err := r.readFile()
	if err != nil {
		return nil, err
	}
	tr := make(map[string]string)

	for _, str := range strngs {
		tmp := strings.Split(str, "=")
		tr[tmp[0]] = tmp[1]
	}
	return tr, nil
}

func (r *Translate) Tr(s string) string {
	tmp, _ := r.funcMakeMapTrans()
	if _, ok := tmp[s]; ok {
		return tmp[s]
	}
	return s
}
