package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 开发者 GitHub token
const token = ""

// 定义一个访问者结构体
type Visitor struct{}

func (self *Visitor) md2html(arg map[string]string) error {
	from := arg["from"]
	to := arg["to"]
	s := `<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
`
	err := filepath.Walk(from+"/", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if (f.Mode() & os.ModeSymlink) > 0 {
			return nil
		}
		if !strings.HasSuffix(f.Name(), ".md") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		input_byte, _ := ioutil.ReadAll(file)
		input := string(input_byte)
		input = regexp.MustCompile(`\[(.*?)\]\(<?(.*?)\.md>?\)`).ReplaceAllString(input, "[$1](<$2.html>)")

		if f.Name() == "README.md" {
			input = regexp.MustCompile(`https:\/\/github\.com\/astaxie\/build-web-application-with-golang\/blob\/master\/`).ReplaceAllString(input, "")
		}

		// 以#开头的行，在#后增加空格
		// 以#开头的行, 删除多余的空格
		input = FixHeader(input)

		// 删除页面链接
		input = RemoveFooterLink(input)

		// remove image suffix
		input = RemoveImageLinkSuffix(input)

		var out *os.File
		filename := strings.Replace(f.Name(), ".md", ".html", -1)
		fmt.Println(to + "/" + filename)
		if out, err = os.Create(to + "/" + filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v", f.Name(), err)
			os.Exit(-1)
		}
		defer out.Close()
		client := &http.Client{}

		req, err := http.NewRequest("POST", "https://api.github.com/markdown/raw", strings.NewReader(input))
		if err != nil {
			// handle error
		}

		req.Header.Set("Content-Type", "text/plain")
		req.Header.Set("charset", "utf-8")
		req.Header.Set("Authorization", "token "+token)
		//
		resp, err := client.Do(req)
		if err!=nil {
			fmt.Println("err:",err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		w := bufio.NewWriter(out)
		n4, err := w.WriteString(s + string(body)) //m.Render()
		fmt.Printf("wrote %d bytes\n", n4)
		// fmt.Printf("wrote %d bytes\n", n4)
		//使用 Flush 来确保所有缓存的操作已写入底层写入器。
		w.Flush()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Parsing Error", err)
			os.Exit(-1)
		}

		return nil
	})
	return err
}

func FixHeader(input string) string {
	re_header := regexp.MustCompile(`(?m)^#.+$`)
	re_sub := regexp.MustCompile(`^(#+)\s*(.+)$`)
	fixer := func(header string) string {
		s := re_sub.FindStringSubmatch(header)
		return s[1] + " " + s[2]
	}
	return re_header.ReplaceAllStringFunc(input, fixer)
}

func RemoveFooterLink(input string) string {
	re_footer := regexp.MustCompile(`(?m)^#{2,} links.*?\n(.+\n)*`)
	return re_footer.ReplaceAllString(input, "")
}

func RemoveImageLinkSuffix(input string) string {
	re_footer := regexp.MustCompile(`png\?raw=true`)
	return re_footer.ReplaceAllString(input, "png")
}

func main() {
	tmp := os.Getenv("TMP")
	if tmp == "" {
		tmp = "."
	}

	workdir := os.Getenv("WORKDIR")
	if workdir == "" {
		workdir = "."
	}

	arg := map[string]string{
		"from": workdir,
		"to":   tmp,
	}

	v := &Visitor{}
	err := v.md2html(arg)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
