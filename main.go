package main

import (
	"github.com/IacopoMelani/vortex/cmd"
)

// func aggregateChunks() {

// 	dirEntries, err := os.ReadDir(".")
// 	if err != nil {
// 		panic(err)
// 	}

// 	recreatedFileName := "recreated_file.exe"

// 	file, err := os.Create(recreatedFileName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, dirEntry := range dirEntries {

// 		if dirEntry.IsDir() {
// 			continue
// 		}

// 		if !strings.Contains(dirEntry.Name(), "chunk_") {
// 			continue
// 		}

// 		b, err := os.ReadFile(dirEntry.Name())
// 		if err != nil {
// 			panic(err)
// 		}

// 		file.Write(b)
// 	}
// }

// func chunkFile(fileToBeChunked string) {

// 	file, err := os.Open(fileToBeChunked)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	fileInfo, _ := file.Stat()

// 	var fileSize int64 = fileInfo.Size()

// 	const fileChunk = 1 * (1 << 18) // chunk sized

// 	// calculate total number of parts the file will be chunked into

// 	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

// 	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

// 	for i := uint64(0); i < totalPartsNum; i++ {

// 		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
// 		partBuffer := make([]byte, partSize)

// 		file.Read(partBuffer)

// 		// write/save buffer to disk
// 		fileName := "chunk_" + strconv.FormatUint(i, 10)
// 		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

// 		fmt.Println("Split to : ", fileName)
// 	}
// }

func main() {
	cmd.Parse()
}
