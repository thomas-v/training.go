package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceLine(line, old, new string) (found bool, res string, occ int) {
	old = old + " "
	new = new + " "
	if strings.Contains(line, old) || strings.Contains(line, strings.ToLower(old)) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, strings.ToLower(old))
		res = strings.ReplaceAll(line, old, new)
		res = strings.ReplaceAll(res, strings.ToLower(old), strings.ToLower(new))
	}
	return found, res, occ
}

func replaceFile(src, dst, old, new string) (occ int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, err
	}
	defer dstFile.Close()

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := replaceLine(scanner.Text(), old, new)
		if found {
			occ += o
		}
		fmt.Fprintf(writer, res)
	}
	return occ, err
}

func main() {
	occ, err := replaceFile("wikigo.txt", "test.txt", "Go", "Php")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Nombre d'occurences remplacées : %d\n", occ)

}
