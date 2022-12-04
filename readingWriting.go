package main

import (
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"io"
	"log"
	"os"
)

func FileReading(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	data := make([]byte, 64)
	text := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		text += string(data[:n])
	}
	return text
}

func FileWriting(filename string, text string) {
	data := []byte(text)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func AudioReading(filename string) (*audio.IntBuffer, uint16) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Decode the original audio file
	// and collect audio content and information.
	d := wav.NewDecoder(f)
	buf, err := d.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File: ", d)
	return buf, d.WavAudioFormat
}

func AudioWriting(filename string, buf *audio.IntBuffer, WavAudioFormat uint16) {
	// Destination file
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// setup the encoder and write all the frames
	e := wav.NewEncoder(out, buf.Format.SampleRate, buf.SourceBitDepth,
		buf.Format.NumChannels, int(WavAudioFormat))
	if err = e.Write(buf); err != nil {
		log.Fatal(err)
	}
	// close the encoder to make sure the headers are properly
	// set and the data is flushed.
	if err = e.Close(); err != nil {
		log.Fatal(err)
	}
	err = out.Close()
	if err != nil {
		log.Fatal(err)
	}
}
