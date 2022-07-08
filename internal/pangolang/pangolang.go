package pangolang

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func ExecuteUnixCmd(c string, arg ...string) {
	cmd := exec.Command(c, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		CheckError(err)
	}
}

func ProgressBarbasic() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}

func ProgressBarCustom(msg string) {
	doneCh := make(chan struct{})

	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset]"+msg),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionOnCompletion(func() {
			doneCh <- struct{}{}
		}),
	)

	go func() {
		for i := 0; i < 1000; i++ {
			bar.Add(1)
			time.Sleep(5 * time.Millisecond)
		}
	}()

	// got notified that progress bar is complete.
	<-doneCh
	fmt.Printf("\n ======= progress bar completed ==========\n")
}

func ProgressBarDownload(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	CheckError(err)
	defer resp.Body.Close()

	filename := (path.Base(req.URL.Path))
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}

func ProgressBarDownloadUnknown(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	CheckError(err)
	defer resp.Body.Close()

	filename := (path.Base(req.URL.Path))
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		-1,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}

func GetUserInput(req string) string {
	var a string
	fmt.Printf("%s ", req)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a := scanner.Text()
		if len(a) < 1 {
			a = "null"
			return a
		} else {
			return a
		}
	}
	return a
}

func DirExist(d string) bool {
	_, err := os.Stat(d)
	var b bool
	if err == nil {
		b = true
	} else if os.IsNotExist(err) {
		b = false
	} else {
		CheckError(err)
	}
	return b
}

func DirIsEmpty(d string) (bool, error) {
	f, err := os.Open(d)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

///////////
// 	dir := "/path/to/dir"
//
// 	if !dirExist(dir) {
// 		fmt.Println("Directory doesn't exist")
// 	}
// 	e, err := dirIsEmpty(dir)
// 	checkError(err)
// 	if !e {
// 		fmt.Println("Directory insn't empty")
// 	}
// }
//////////

func Tcp4Client(proto string, ip string, data string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
	CheckError(err)

	conn, err := net.DialTCP(proto, nil, tcpAddr)
	CheckError(err)

	_, err = conn.Write([]byte(data))
	CheckError(err)

	result, err := ioutil.ReadAll(conn)
	CheckError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func HttpServer() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	fmt.Printf("URL : http://localhost:%s \nPress CTRL+C for interrupt...\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func Govalidator() {
	str := "https://www.example.com"
	validURL := govalidator.IsURL(str)
	fmt.Printf("%s is a valid URL : %v \n", str, validURL)
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// usage : fmt.Println(pangolang.ArrayToString(c, ""))

func ListProcess() {
	matches, _ := filepath.Glob("/proc/*/exe")
	for _, file := range matches {
		target, _ := os.Readlink(file)
		if len(target) > 0 {
			fmt.Printf("%+v\n", target)
		}
	}
}
