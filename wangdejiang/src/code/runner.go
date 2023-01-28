package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func main() {
	// 运行用户代码
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdPipe, err := cmd.StdinPipe()
	// 根据测试的输入案例运行拿到输出结果和标准输出结果对比
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdPipe, "23 1\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	println(out.String() == "24\n")
}
