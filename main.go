package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// reading files
	textName, err := inputStringPath("\nВведите полное название скрываемого файла.",
		"\nExample: C:/Users/User1/Downloads/text.txt")
	if err != nil {
		log.Println(err)
		return
	}
	audioName, err := inputStringPath("\nВведите полное название wav-аудиофайла.",
		"\nExample: C:/Users/User1/Downloads/audio.wav")
	if err != nil {
		log.Println(err)
		return
	}
	text := FileReading(textName)
	buf, f := AudioReading(audioName)
	// hiding
	b := intsByRunes([]rune(text))
	buf.Data = encrypt(buf.Data, b)
	AudioWriting(audioName[:len(audioName)-4]+"_encoded.wav", buf, f)
	// releasing
	buff, f := AudioReading(audioName[:len(audioName)-4] + "_encoded.wav")
	fmt.Println(decrypt(buff.Data))
	FileWriting(textName[:len(audioName)-5]+"_secret.txt", string(runesByInts(decrypt(buff.Data))))
}

func inputStringPath(args ...any) (string, error) {
	for _, v := range args {
		fmt.Print(v)
	}
	fmt.Print("\nВвод: ")
	filename := ""
	_, err := fmt.Scanf("%s\n", &filename)
	if err != nil {
		return "", fmt.Errorf("input error")
	}
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("input error")
	}
	err = file.Close()
	if err != nil {
		return "", err
	}
	return filename, nil
}

func intsByRunes(data []rune) []int {
	arr := make([]int, len(data))
	for i, v := range data {
		arr[i] = int(v)
	}
	return arr
}

func runesByInts(data []int) []rune {
	arr := make([]rune, len(data))
	for i := range arr {
		arr[i] = rune(data[i])
	}
	return arr
}
