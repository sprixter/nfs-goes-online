package compressfilecontent

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

// compressFileContent compresses the contents of a source file and returns the compressed data as a byte slice.
func compressFileContent(src string) ([]byte, error) {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Create a buffer to hold the compressed data
	var buf bytes.Buffer

	// Create a zlib writer with the buffer as the underlying writer
	zlibWriter := zlib.NewWriter(&buf)
	defer zlibWriter.Close()

	// Copy data from the source file to the zlib writer, which compresses the data and writes it to the buffer
	if _, err := io.Copy(zlibWriter, srcFile); err != nil {
		return nil, fmt.Errorf("failed to compress file: %v", err)
	}

	// Close the zlib writer to flush any remaining data to the buffer
	if err := zlibWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close zlib writer: %v", err)
	}

	return buf.Bytes(), nil
}
