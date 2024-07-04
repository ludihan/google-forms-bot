package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type ResponseRow struct {
	resp    *http.Response
	row     int
	success bool
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage:", os.Args[0], "<csv> <url>")
		os.Exit(1)
	}
	targetUrl := os.Args[2]

	_, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csv := csv.NewReader(file)

	records, err := csv.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keys := records[0]
	rows := records[1:]

	fmt.Printf("Keys:\n%v\n\nRows:\n", keys)
	for _, v := range rows {
		fmt.Println(v)
	}
	fmt.Printf("\nTarget: %v\n\n", targetUrl)

	for {
		fmt.Print("Do want to send? [y/n] ")
		stdin := bufio.NewReader(os.Stdin)
		var char string
		fmt.Fscan(stdin, &char)
		if err != nil || char != "y" && char != "n" {
			continue
		}
		if char == "n" {
			os.Exit(0)
		} else {
			break
		}
	}

	ch := make(chan ResponseRow)

	for i := 0; i < len(rows); i++ {
		go func() {
			row := rows[i]
			v := url.Values{}
			for i := 0; i < len(keys); i++ {
				v.Set(keys[i], row[i])
			}
			resp, err := http.PostForm(targetUrl, v)
			if err != nil {
				ch <- ResponseRow{
					&http.Response{},
					i,
					false,
				}
			}
			ch <- ResponseRow{
				resp,
				i,
				true,
			}
		}()
	}

	for i := 0; i < len(rows); i++ {
		response := <-ch
		fmt.Printf("Row:    %v\n", response.row)
		fmt.Printf("Status: %v\n", response.resp.Status)
		fmt.Printf("Sucess: %v\n", response.success)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("\n\n")
	}

	fmt.Println("All rows done!")
}
