/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/23 12:46:30 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/25 13:38:50 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var nm = flag.Bool("n", false, "Remove file before conversion")

func myGetNextLine(file io.Reader, write io.Writer, nm *bool) {
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if *nm {
			fmt.Printf("    %v: ", i)
		}
		line := scanner.Text()
		write.Write([]byte(line + "\n"))
	}
}

func main() {
	// 標準入力から読み込む
	var filenames string
	var r io.Reader
	var w io.Writer
	flag.Parse()
	w = os.Stdout
	if args := flag.Args(); len(args) > 0 {
		filenames = args[0]
	}
	switch filenames {
	case "":
		r = os.Stdin
		myGetNextLine(r, w, nm)
	default:
		for i := 0; i < flag.NArg(); i++ {
			r, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprint(os.Stderr, "opne file error!")
			}
			defer r.Close()
			myGetNextLine(r, w, nm)
		}
	}
}
