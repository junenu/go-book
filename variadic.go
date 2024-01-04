package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type holiday struct{
	Date   string
	Year   int
	Month  int
	Day    int
	Name   string
}

const (
	SJIS_CSV = "./csv/syukujitsu.csv"
	UTF8_CSV = "./csv/syukujitsu_utf8.csv"
	JSON_FILE = "./holidays.json"
)

func main() {
	convert_sjis_utf8(SJIS_CSV, UTF8_CSV)
	file, err := os.Open(UTF8_CSV) 
	if err != nil {
	log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
	log.Fatal(err)
	}

	var holidays []holiday

	for i, row := range rows {
	if i == 0 {
		continue //ヘッダをスキップ
	}

	d, err := time.Parse("2006/1/2", row[0])
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	// 年、月、日を取得
	year, month, day := d.Year(), int(d.Month()), d.Day()

	date := d.Format("20060102") //yyyyMMddフォーマットに変換する
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	holiday := holiday {
		Date:  date,
		Year:  year,
		Month: month,
		Day:   day,
		Name:  row[1],
	}
	holidays = append(holidays, holiday)
	}

	f, err := os.Create(JSON_FILE) //jsonファイル作成

	for _, h := range holidays {
		j, err := json.MarshalIndent(h, "", " ")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		f.Write(j)
	}	
}

func convert_sjis_utf8(from_file string,to_file string) {
	// ShiftJISファイルを開く
	sjisFile, err := os.Open(from_file)
	if err != nil {
		log.Fatal(err)
	}
	defer sjisFile.Close()

	// ShiftJISのデコーダーを噛ませたReaderを作成する
	reader := transform.NewReader(sjisFile, japanese.ShiftJIS.NewDecoder())

	// 書き込み先ファイルを用意
	utf8File, err := os.Create(to_file)
	if err != nil {
		log.Fatal(err)
	}
	defer utf8File.Close()

	// 書き込み
	tee := io.TeeReader(reader, utf8File)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}