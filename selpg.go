package main

import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "os/exec"
)
//pu the arguments in a struct
type args struct {
  sPage int
  ePage int
  pageLen int
  pageType bool
  inputFile string
  dest string
}

func main()  {
// initial the struct
  info := args{
    sPage : -1,
    ePage : -1,
    pageLen : 72,
    pageType : true,
    inputFile : "",
    dest : "",
  }
//get all the arguments by flag
  flag.IntVar(&info.sPage,"s",-1,"the starting page")
  flag.IntVar(&info.ePage,"e",-1,"the ending page")
  flag.IntVar(&info.pageLen,"l",72,"the length of one page")
  flag.BoolVar(&info.pageType,"f",false,"if exists,then pages are separated by page changer")
  flag.StringVar(&info.dest,"d","","the printing position")
  flag.Parse()

//judge the wrong input
  if  info.sPage < 1 || info.ePage < 1 || info.sPage > info.ePage {
      flag.Usage()
      return
  }

  if info.pageLen != 72 && info.pageType == true {
      flag.Usage()
      return
  }

//deal with right input
  if len(flag.Args()) > 1 {
      flag.Usage()
      return
  }

  if len(flag.Args()) == 1 {
        info.inputFile = flag.Args()[0]
  }

//judge which kind of pages
  if info.pageType == false {
      type_1(info,info.inputFile != "",info.dest != "")
  }else {
    type_2(info,info.inputFile != "",info.dest != "")
  }

}

//reading and writing files
func type_1(info args,file bool,pipe bool) {
    cmd := exec.Command("cat", "-n")
    stdin, err:= cmd.StdinPipe()
    if err != nil {
        panic(err)
    }
    cur_page := 1
    cur_lines := 0
    if file {
        file_in, err := os.OpenFile(info.inputFile,os.O_RDWR,os.ModeType)
        defer file_in.Close()
        if err != nil {
            panic(err)
            return
        }
    line := bufio.NewScanner(file_in)
    for line.Scan() {
            if cur_page >= info.sPage && cur_page <= info.ePage {
                os.Stdout.Write([]byte(line.Text()+"\n"))
                stdin.Write([]byte(line.Text()+"\n"))
            }
            cur_lines++;
            if cur_lines %= info.pageLen; cur_lines == 0 {
                cur_page++;
            }
    }
    } else {
        tmp_s := bufio.NewScanner(os.Stdin)
        for tmp_s.Scan() {
            if cur_page >= info.sPage && cur_page <= info.ePage {
                os.Stdout.Write([]byte(tmp_s.Text()+"\n"))
                stdin.Write([]byte(tmp_s.Text()+"\n"))
            }
            cur_lines++;
            if cur_lines %= info.pageLen; cur_lines == 0 {
                cur_page++;
            }
        }
    }
    if cur_page < info.ePage {
        fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
    }
    if pipe {
        stdin.Close()
        cmd.Stdout = os.Stdout;
        cmd.Start()
    }
}

func type_2(info args,file bool,pipe bool) {
    cmd := exec.Command("cat", "-n")
    stdin, err:= cmd.StdinPipe()
    if err != nil {
        panic(err)
    }
    cur_page := 1
    if file {
        file_in, err := os.OpenFile(info.inputFile,os.O_RDWR,os.ModeType)
        defer file_in.Close()
        if err != nil {
            panic(err)
            return
        }
        line := bufio.NewScanner(file_in)
    for line.Scan() {
            flag := false
            for _,c := range line.Text() {
                if c == '\f' {
                    if cur_page >= info.sPage && cur_page <= info.ePage {
                        flag = true
                        os.Stdout.Write([]byte("\n"))
                        stdin.Write([]byte("\n"))
                    }
                    cur_page++;
                } else {
                    if cur_page >= info.sPage && cur_page <= info.ePage {
                        os.Stdout.Write([]byte(string(c)))
                        stdin.Write([]byte(string(c)))
                    }
                }
            }
            if flag != true && cur_page >= info.sPage && cur_page <= info.ePage {
                os.Stdout.Write([]byte("\n"))
                stdin.Write([]byte("\n"))
            }
            flag = false
    }
    } else {
        tmp_s := bufio.NewScanner(os.Stdin)
        for tmp_s.Scan() {
            flag := false
            for _,c := range tmp_s.Text() {
                if c == '\f' {
                    if cur_page >= info.sPage && cur_page <= info.ePage {
                        flag = true
                        os.Stdout.Write([]byte("\n"))
                        stdin.Write([]byte("\n"))
                    }
                    cur_page++;
                } else {
                    if cur_page >= info.sPage && cur_page <= info.ePage {
                        os.Stdout.Write([]byte(string(c)))
                        stdin.Write([]byte(string(c)))
                    }
                }
            }
            if flag != true && cur_page >= info.sPage && cur_page <= info.ePage {
                os.Stdout.Write([]byte("\n"))
                stdin.Write([]byte("\n"))
            }
            flag = false
        }
    }
    if cur_page < info.ePage {
        fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
    }
    if pipe {

        stdin.Close()
        cmd.Stdout = os.Stdout
        cmd.Start()
    }
}

func Usage() {
    fmt.Fprintf(os.Stderr,"Usage : [-s the starting page] [-e the ending page] [-l length of one page(default 72)] [-f type of file(default false)] [-d dest] [input file]")
}