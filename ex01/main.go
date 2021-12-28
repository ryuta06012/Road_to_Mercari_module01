/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/25 13:48:03 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 11:43:49 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	imgconv "test/imgconv"
)

var from = flag.String("f", ".jpg", "Extension before conversion")
var to = flag.String("t", ".png", "Extension after conversion")
var rm = flag.Bool("r", false, "Remove file before conversion")

func dirWalk(dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *from {
				dst := path[:len(path)-len(*from)] + *to
				println(path)
				println(dst)
				if err := imgconv.Convert(path, dst, *rm); err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return err
}

func checkDir(dir string) {
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		fmt.Fprint(os.Stderr, "The directory does not exist!\n")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, "Not enough arguments\n")
		os.Exit(1)
	} else if flag.NArg() > 1 {
		fmt.Fprint(os.Stderr, "There is only one argument\n")
		os.Exit(1)
	}
	checkDir(flag.Arg(0))
	if err := dirWalk(flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
}
