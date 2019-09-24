package mainpackage

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)
cmd

import (
"bufio"
"fmt"
"github.com/spf13/cobra"
"os"
"regexp"
"sort"
"strings"
)

const (
	InitiatorRegion = 1
	ResponderRegion = 2
)

type info struct {
	Src map[string]uint32
	Dst map[string]uint32
}

type result struct {
	ip string
	count uint32
}

var number uint32


var rootCmd = &cobra.Command{
	Use:   "filter",
	Short: `filter info`,
	Example: `
main.exe fw.log -n 10`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		f, err := os.Open(args[0])
		//f, err := os.Open(`C:\Users\guanzhang\Desktop\1.txt`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open err: %v\n", err)
			return
		}
		defer f.Close()
		run(f)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Uint32VarP(&number, "number","n",3, "result line number")
}


func run(f *os.File) {

	Initiator := &info{
		Src: make(map[string]uint32,200),
		Dst: make(map[string]uint32,200),
	}
	Responder := &info{
		Src: make(map[string]uint32,200),
		Dst: make(map[string]uint32,200),
	}

	line := bufio.NewScanner(f)

	var flag uint8

	re, _ := regexp.Compile(`(\d{1,3}\.){3}\d+`)

	for line.Scan() {
		if strings.Contains(line.Text(),"Initiator:") {
			flag = InitiatorRegion
		} else if strings.Contains(line.Text(),"Responder:") {
			flag = ResponderRegion

		} else if  strings.Contains(line.Text(),"  Source      IP/port") ||
			strings.Contains(line.Text(),"  Destination IP/port") {
			if flag == InitiatorRegion{
				filter(line.Text(), re, Initiator)
			} else {
				filter(line.Text(), re, Responder)
			}
		} else if strings.Contains(line.Text(), "DS-Lite"){
			flag = 0
		} else {
			continue
		}
	}

	var ISrcSlice, IDstSlice, RSrcSlice, RDstSlice  []result
	for k, v := range Initiator.Src {
		ISrcSlice = append(ISrcSlice, result{ip:k, count:v})
	}
	for k, v := range Initiator.Dst {
		IDstSlice = append(IDstSlice, result{ip:k, count:v})
	}
	for k, v := range Responder.Src {
		RSrcSlice = append(RSrcSlice, result{ip:k, count:v})
	}
	for k, v := range Responder.Dst {
		RDstSlice = append(RDstSlice, result{ip:k, count:v})
	}

	sort.Slice(ISrcSlice, func(i, j int) bool {
		return ISrcSlice[i].count > ISrcSlice[j].count
	})
	sort.Slice(IDstSlice, func(i, j int) bool {
		return IDstSlice[i].count > IDstSlice[j].count
	})
	sort.Slice(RSrcSlice, func(i, j int) bool {
		return RSrcSlice[i].count > RSrcSlice[j].count
	})
	sort.Slice(RDstSlice, func(i, j int) bool {
		return RDstSlice[i].count > RDstSlice[j].count
	})

	fmt.Printf("%-16s\t%-16s\t%-16s\t%-16s\n","Initiator Src","Initiator Dst", "Responder Src", "Responder Dst")
	var i uint32
	for ; i < number; i++ {
		fmt.Printf("%-17s%d\t%-17s%d\t%-17s%d\t%-17s%d\n",
			getIP(i, ISrcSlice), getCount(i, ISrcSlice),
			getIP(i, IDstSlice), getCount(i, IDstSlice),
			getIP(i, IDstSlice), getCount(i, RSrcSlice),
			getIP(i, RDstSlice), getCount(i, RDstSlice))
	}
}

func filter(line string, re *regexp.Regexp, container *info ){
	if strings.Contains(line, "Source") {
		ip := string(re.Find([]byte(line)))
		if _, ok := container.Src[ip]; ok {
			container.Src[ip]++
		} else {
			container.Src[ip] = 1
		}
	} else {
		ip := string(re.Find([]byte(line)))
		if _, ok := container.Dst[ip]; ok {
			container.Dst[ip]++
		} else {
			container.Dst[ip] = 1
		}
	}
}


func getIP(i uint32, slice []result ) string {

	if uint32(len(slice)) <= i {
		return "-"
	}
	return slice[i].ip
}

func getCount(i uint32, slice []result ) uint32 {

	if uint32(len(slice)) <= i {
		return 0
	}
	return slice[i].count
}