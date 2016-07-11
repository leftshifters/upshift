// Code generated by go-bindata.
// sources:
// scripts.bash
// DO NOT EDIT!

package basher

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x59\xef\x72\xda\x48\x12\xff\x6c\x9e\xa2\x17\xbb\xe2\x38\x8b\x90\xc1\xc9\xe6\xaf\xf7\x0a\x83\x62\xeb\x16\x03\x05\x78\x53\x5b\x49\x8a\x1a\xa4\x01\xe9\x2c\x34\x3a\x69\x64\xcc\x65\x7d\x6f\x70\xcf\x70\xef\x71\x4f\x75\x8f\x70\xdd\x33\x12\x48\x36\x71\x72\x6b\xdf\x87\x4b\xa5\x6c\x23\xf5\xf4\xff\xfe\x75\xf7\x30\xe6\x89\x1c\x39\xb1\x1f\xc9\xa7\x07\xf0\xa5\xb2\xc3\x1d\x4f\xc0\x5e\xa3\xb2\x93\x04\x9c\x47\x80\x7f\x04\x09\x18\x01\xab\xdc\x54\x2a\x1b\xda\xa6\x26\x3e\xb7\x46\xa3\xd6\xa9\x75\x4c\xf4\xdd\xfe\xe9\x7b\xbb\x6b\x8d\x8e\xf7\x9a\x77\xb8\x34\x73\x2e\xd0\xfc\xf9\x49\x03\x7e\x07\xc9\x39\x20\x1d\xf2\x3c\x8d\x99\x1b\xf0\x0f\x31\x8b\x22\x1e\x6b\xae\x73\xf5\x08\x96\xfa\xd9\x86\x68\x39\x66\xc9\xa5\x26\x19\xb4\x86\xad\xf3\x51\x2e\x77\x32\x68\x8d\xcf\x94\xdc\xba\xa9\x0f\x2f\xa1\xba\xd7\xa8\xde\x95\x36\xe2\x32\x8d\x4e\xf9\x42\x73\x39\xb5\xce\x15\x8b\xb6\xad\x0e\x0f\x85\x90\x03\x96\x24\x4b\x11\xbb\xc7\x7b\x47\x95\xca\x8e\x3f\x83\x8f\xc8\xa9\x59\x85\xe3\x63\xa8\xca\x38\xe5\x55\xf8\xfc\x16\xa4\xc7\xc3\xca\xce\xce\x2e\xb4\x6d\xf0\x13\xa0\xe7\x35\x58\x72\x08\xc5\x12\x42\xce\x5d\x88\x32\x2e\x48\x94\xb1\x38\xaa\xc2\x0f\xc8\xa2\x78\x5c\x3b\xc9\xe0\xfa\xed\xef\x90\xa4\x2e\x7e\xbc\x04\x63\x04\x73\xbe\x00\x3f\x4c\x24\x0b\x02\xe5\x45\xa4\xbd\xf6\x25\x1c\xe2\x5f\x3c\x48\xf8\xfa\x01\xbd\x9a\xf9\x95\xfc\xe1\x2e\x5c\x24\x3c\x86\xc4\x13\x69\xe0\x82\x5c\x45\x1c\xb9\x90\xb8\xa2\x42\xb9\x98\x3b\x32\x72\x11\xc8\x10\x3d\x75\x11\x66\x2f\xff\x3f\xbc\x95\xe6\xea\x82\xc1\xec\xeb\xff\xa9\xd3\xb6\x89\x2a\xfb\x6e\x20\x5c\x5b\x93\x68\xcf\x65\xa5\xa1\xbc\x17\x09\x77\xed\xf6\x62\x7a\x36\xb2\x83\x43\x1e\x89\x8b\xc8\x65\x92\x6f\x3f\x1b\xe3\x7b\x48\x15\x01\x18\xc6\x15\x8f\xa7\x22\xe1\x77\x39\x9d\xfa\x72\x94\x4e\x17\xc2\x4d\x03\x6e\x87\xbe\xbc\xcb\x6c\x8e\x1a\x27\x39\x09\xaa\x84\x1f\xef\xe5\xf2\x35\xa5\xca\x7c\x32\xcd\xb6\x71\x1a\xa4\xb9\x3f\x86\xd6\x79\x7f\xac\x4f\x9f\x0c\x5b\xbd\xb6\xae\xdd\x35\xd3\x23\xcd\x34\x4a\x55\x66\x62\xe5\x96\xd8\x1d\x11\xbb\x93\x98\x2f\x4b\x2e\x1e\xf7\xfb\x5d\xc5\x70\x8a\x6f\x4a\x79\x9d\x85\x26\x3b\x74\x11\x11\x40\xf0\x6d\x87\x52\xfd\x6a\xdb\xa1\x3c\xe2\x5b\x8f\x85\xdb\xa5\x65\x92\x8a\xd8\x7a\x62\x8d\x5b\xea\xa0\xdd\x1b\x8d\x5b\xdd\xee\xe4\x62\xd8\x3d\xae\x7a\x52\x46\xc9\x1b\xd3\x8c\xd9\xb2\x8e\x76\x7b\xe9\x34\xc5\x84\x74\x44\x28\x79\x28\xeb\x8e\x58\x98\x01\x9f\xc9\xc4\xf3\x67\x92\xc7\x89\x99\x46\xea\x4f\x73\xc1\x12\xfc\x6c\x66\xc2\xeb\x89\x57\x5d\x57\xe0\x17\x12\x74\xb3\xbd\x0c\x1f\x43\xf4\x5c\x94\xc5\x52\x35\x39\x69\x8c\x05\x31\x4b\x46\x5d\x52\xa0\x20\x05\xf5\x80\x9f\x61\x43\x0f\x4f\x9e\x80\xe3\x61\xba\xc0\x8f\xd7\xb7\x1e\xd7\xcd\xf2\xe7\x78\x51\x7a\xa4\x00\x5c\xb2\x58\x8e\xfc\x45\x1a\x30\x29\xb2\x7e\xd1\xb1\x7e\xb5\xdb\x3a\x9d\xae\x9d\x38\x0d\x15\xd7\x38\x5d\xa0\x0d\xd8\x74\xb2\x56\xd0\xf8\xd9\x74\xf9\x95\x19\xa6\x59\xdd\x11\x33\xbb\x3f\xb2\x43\x27\xe6\x44\x79\x92\xfa\x81\xdb\x4b\x17\xd3\xbc\x0b\x0d\x86\xfd\x3f\x5b\xed\xf1\xa4\xd7\x3a\x57\xbc\xd1\x44\xb1\x58\xf8\xb2\x2d\xd2\x50\x1e\xef\x3d\xa5\x14\x8d\xf9\x95\x11\xf8\x89\xc4\x4a\x54\x78\x60\x38\xf4\xf2\x40\x79\x23\x2e\x33\xc5\x13\x66\x9a\xc4\x66\xe0\x4f\xf9\x35\x77\xcc\x01\x9d\x3b\x49\x5d\x77\x05\x86\x03\xd5\x41\xec\x87\x12\xda\xef\x4f\xd2\x10\xdb\xd7\xaf\xe8\x6d\x5f\x84\x55\x4c\x29\xd3\x0e\x67\xa2\x1e\x11\xf5\x01\xea\x10\xf2\x65\x91\x69\x75\x13\xf4\x82\x76\x55\x30\xe6\x92\x1e\xdd\xd1\xa2\x98\x09\xb7\x58\x15\x19\xac\xc1\xf1\x36\xcd\xd3\xa7\x5b\x98\xfe\xd8\x38\x38\x50\x49\x50\xd9\xb9\xcf\x44\xec\xbd\xf0\xe6\x96\x85\xb0\x57\x96\x50\xa5\x60\x15\x4c\xae\x22\xcf\x5d\x68\xb9\x2e\xd8\xe3\x51\x2b\x8a\x10\xaa\x93\x9e\x08\xad\x6b\xbe\x88\xa4\x85\xb1\x5b\x45\x92\xd8\xb0\x04\x66\x0c\x55\x06\x74\x06\xc6\xc5\x4f\xc2\x7d\x44\x25\x8e\x86\xb0\x30\x59\xf2\x78\x2c\x36\xc4\xdf\x15\x89\x7b\xc5\xdd\xd1\xf2\x40\x47\xe1\x23\xec\xdd\x15\x47\x55\xf8\xac\xda\x11\x3c\x81\x9e\x90\x60\x5d\xd3\x81\x67\xf0\x79\x13\x88\xfb\xb4\x21\xd3\xdf\xdc\x6f\xfb\x54\x88\x40\x5b\x7f\xd7\x7b\x59\x43\xc2\x44\x57\x4e\x2e\x67\xf6\xf8\xb7\x81\xae\x9a\xfc\xc1\x7a\x8c\x1a\xb5\xcf\xac\x73\x8d\xc4\x79\x71\x3d\x2f\x4e\x5a\x2f\x70\xb0\xc3\x68\x1a\x02\x22\x3f\xe2\x33\xe6\x07\x54\xae\xd7\x8e\x70\xf9\x94\xe4\x80\xa1\x6a\x4e\xcd\x02\x46\xe2\x78\x58\x61\xba\x6f\x1b\x9e\x8f\x90\xe8\xf1\x20\xd0\xb8\x68\x85\x57\x7e\x2c\x42\x2a\x40\x24\x74\x2f\xc1\x8f\x3c\x11\xf2\x24\x2f\x70\x30\x5c\x9c\x39\xfd\x90\x29\x4b\xab\x11\x3e\x9c\x89\x78\x71\x7c\x8b\xac\x16\xb2\x05\x47\x1d\xab\x44\x1f\xfb\x57\xdc\xed\x30\xc9\x06\x4c\x7a\x50\xcf\x41\x4b\x2b\xa6\x9b\x48\x75\xef\x05\xcd\x10\xd7\x4e\x14\x73\x29\x57\x99\x8b\x68\xbc\x7d\x44\x0f\x5d\x3b\x92\x22\xf3\x55\x57\x3c\x86\xb9\x12\xc9\x0b\x36\x65\x86\xb4\x62\xc7\x43\x27\xfc\x01\x5b\xca\xa0\xf7\x48\x31\xff\x46\x48\x0c\xa6\xd5\x2d\xbf\xdb\x7b\x5e\xbf\x76\xb2\x37\x90\xff\xbe\x2f\x7a\xd6\x75\x24\x62\xb9\x1d\xbb\xcb\x4b\xc2\x37\xcc\xe0\x8a\x51\xe6\xc3\xfc\x63\x5f\xd5\x5a\xa2\xaa\x13\xea\x11\x1a\x84\x63\x8e\xa9\xdf\xe9\x62\xfb\x9a\x1d\x8d\x82\x1d\x19\xb3\x32\x45\x71\xba\x51\xde\x2b\x9b\xf6\x9e\x4b\xc7\x6b\x85\x34\x18\x32\x3f\x1e\xc4\xe2\xca\x27\xe4\xf4\xc3\x39\xfe\x3d\xf3\x03\x9e\x68\x9b\x77\x61\x77\xdc\xef\xf4\xa1\xd7\x1f\xc3\xc5\xc8\xee\x9d\xc2\xf8\xcc\x1e\x41\xab\xf7\xdb\x87\x33\x6b\x68\xd5\xa0\x67\x59\x1d\x18\xf7\xa1\x63\xb5\xed\x8e\x05\xf6\x7b\xf8\x60\x6d\x1e\xf6\x15\x39\xb1\x91\x1e\x4e\xe9\x38\x07\x24\x90\xf8\x73\x8f\x9e\xb8\x62\x19\x06\x82\xb9\x40\x4d\x0e\xd1\x51\xac\x3d\x00\x33\x11\x60\x78\x89\x88\x85\xae\x02\x34\xa0\x0e\x4c\xe3\xf3\x2c\x0d\x1d\x95\xc7\x03\x11\x51\xde\xf2\x6d\xba\x57\x76\x5a\xed\x76\xff\xa2\x37\x9e\x58\xe7\x2d\x5b\x8f\x55\x24\x96\xe6\x5c\x34\x17\x8c\x14\xd6\x8f\x72\x35\x26\xaa\xd7\xea\x17\x8b\x2b\x78\x56\x5f\x88\x29\xf2\x8a\x72\xf6\x6b\xed\x30\x11\xee\x13\x8d\xce\x45\x78\x1d\x2d\x31\x08\x23\x91\xc6\x4e\xee\xc9\xd6\xb0\x7d\x66\xff\x6a\x6d\xb2\xc7\x1e\xb4\x36\x1f\x54\xbb\x35\x5c\xa8\x6e\x0d\xb1\xa9\xd9\xa5\x11\xc5\xd9\x2c\xb6\xdb\x88\xfa\xca\x4c\x21\x3a\x2a\x01\x45\xb1\x80\x1e\x25\x8f\xa1\x1c\x74\x68\x8c\x86\x7d\x0a\x11\xbb\x77\xd2\xf0\x6f\x7e\x04\xc6\x5f\x0b\xd9\xd4\xac\xfb\x11\x23\xf1\xeb\x47\x72\x11\x21\x29\x8d\x4a\x65\x22\x7c\xb8\xb8\x74\xc9\x85\x71\x89\xb8\xac\xa1\x46\x16\x81\x71\xd8\x71\x50\x54\x3c\xdb\x9e\xb9\xdb\x0f\x99\xcf\xea\xee\x0a\x1b\xd7\x77\xf0\xcf\x49\x49\xce\x1d\xe5\x95\x95\xa8\x65\x3d\xb7\xaf\x9e\x91\xd5\xcd\x7a\x5d\x1b\x57\xd2\x4c\x9d\xca\xd6\xd4\x88\x32\x02\x1d\x87\xa1\x1d\xa7\x61\x1e\x42\xc4\x64\xab\xdb\x1f\x58\xc3\x49\x96\x5e\xeb\x38\xae\x51\x20\xf2\x03\x6c\xc7\xa9\x3a\x8f\x93\x1b\x0d\xbe\x04\xab\xb4\x71\x18\x06\x29\xa1\x91\xcc\x48\x2e\xfd\x68\xb2\x64\x3e\xc2\xf2\x7c\x82\xd1\x99\x28\x98\x98\x60\xb6\x61\xe4\x12\x0a\x25\x8d\xd8\x9b\x5d\x4c\xd5\xac\x1f\xba\x5b\x52\xee\x3e\xdd\x4e\x2e\x7a\x9d\xae\x35\xc1\xca\xec\x8d\xed\xf7\xb6\x35\xd4\x50\x45\x79\xaf\x52\x1d\xe1\x25\xbb\xc3\x68\xc7\x1c\x33\x1a\x07\x82\x7e\x68\xcb\x6f\x99\xbc\x9d\x6d\x19\x1f\x11\xf7\xd1\x1a\x37\x75\x78\x51\x14\xcd\xb4\x51\x34\x51\x3e\x21\x1c\xd7\xfb\xe9\xd7\xcb\xe9\x5b\x28\x74\xde\xff\x2a\x08\x7d\xb0\xbb\xdd\x22\x02\x9d\x22\x40\x53\x3d\x5c\x5c\xd8\x1d\x98\xc5\x62\x51\xa8\xe8\x5d\xc8\x97\x96\x39\x82\x6e\xb6\xb5\xa8\x55\x65\x71\x1d\xc5\xe6\xab\xe6\xe1\xab\xe6\xab\xd7\xec\xa7\x23\x87\x3d\xe7\x47\xec\xe8\x05\x7b\x4e\x87\xba\x42\x44\xc8\x33\x16\x29\x3a\x94\xf0\x43\xe9\x5c\xa3\x79\x71\x25\x52\xdc\xeb\x25\x30\xb8\xe8\xd8\x9d\x1a\x30\x57\x01\xd9\x22\xaf\x4a\x05\xee\x62\x06\x51\x8e\x1b\x45\x0c\x30\xd7\xbd\xa0\x7c\xd7\x31\xc6\x83\xb7\x40\x12\x38\x8d\x7e\x54\x69\x54\xe3\x19\xb7\x1e\xb9\xd7\xdf\x40\x96\x79\x07\xcd\xde\x22\xee\xd1\xdd\x45\x9a\xfa\xee\xfd\xa3\xeb\xbe\x1e\x5d\xc9\x6b\xfb\xa0\xb6\x9d\x44\x22\xd6\xc0\xbb\x77\xef\x60\xef\x69\xc2\x71\x72\xf7\xe5\x0a\x9c\x05\x6e\x45\x1d\x30\x7c\x5a\xd4\x0a\x6a\xdc\xd0\x7d\xd9\x7a\x49\x3a\xa0\x6d\x83\x2c\xb1\x67\xe4\x19\x8a\x84\xd2\xbf\xa6\x41\xde\x11\xd1\x0a\x27\xed\x5a\x36\x70\x7b\x4c\x4d\xdc\x2c\xc0\xdc\x44\x65\xa6\x5c\x93\xf8\x9c\x2e\x4f\xd6\x6b\x29\x99\x70\xb3\xe5\x6a\x47\x13\xfc\x80\xeb\x23\xfc\xdd\xec\xfa\xd3\x98\xc5\x2b\xf3\x5c\xb9\xa1\xc3\xaf\x7c\x87\x9b\xc5\x84\xfb\x04\x79\xca\x99\x19\xcb\x3b\x0d\xa0\xc8\x9b\x2e\x9e\x48\x59\x3a\x41\x21\xd5\x49\x95\xb5\x2d\xfa\x47\xb8\x37\xbb\xeb\x8b\x47\xd1\x44\x4b\xc8\xa1\xff\x1c\x9f\x22\x5e\x94\x25\xc1\xc7\xec\xec\x67\xd2\x2e\x93\xa9\xf1\x5f\xdf\x53\x65\x3f\x5d\xc4\x51\xbd\x6a\x6f\xee\x18\x70\x2c\x88\x85\xef\x76\x19\xb6\x5a\xcf\x2a\x2e\xc5\x95\x9d\x98\xbb\x6d\x11\x60\x83\x39\xde\xff\x74\x78\x74\xf4\xf1\xf0\xed\x51\x63\xb1\x4f\xb7\xab\x18\x9c\xdb\x6f\x9a\xf4\x66\x1a\xa4\xfc\xf6\x8b\xe7\xf4\x22\x14\xe5\xc7\xf8\xac\xb2\x63\x9d\x5f\x74\x5b\xe3\xfe\xf0\x2b\x43\xd6\xfa\xf5\xd0\x1a\x5d\x74\xc7\x23\x4c\xdd\x50\x78\x69\x84\x8e\x6e\xf5\x3a\xc3\xbe\xdd\x99\x9c\xf5\xcf\x2d\x93\x66\xe4\xc4\xe4\x99\xee\x08\xb8\xec\xca\xcd\x2f\x6f\x11\x84\x1a\xf4\xe3\xc9\x81\x5a\x03\x15\xb0\xbc\xc1\xd2\xc2\xf9\x04\xff\x33\x98\xfa\x73\xd8\x3d\x6b\xb5\x7f\xa9\x81\x08\x83\x15\xf0\x38\x16\x31\xbe\x88\x39\x8e\x0f\x32\x8d\x43\xee\xe6\x77\x79\x33\x3f\xc6\x2a\x96\x4b\x81\x5b\xa1\x23\x42\x17\x4b\xdf\x86\x24\x75\x2e\xd5\xdc\x62\x63\x8d\x51\xfe\x5e\xd2\x9d\x24\x83\x25\x5b\x81\x48\x65\xe1\xf6\x7a\x8b\x35\xef\xf6\x9a\x5a\xad\xb6\xc7\x91\x0d\xa6\x30\xca\x41\xc9\x4b\x46\xaa\x0d\x5a\x3d\xbb\x0d\x1f\x09\x42\x68\x44\xa7\x99\xaa\x06\x51\x2a\x73\x7d\x96\xb8\xf5\xcc\x21\x37\x1b\x08\x65\x3f\xd3\xb5\x36\x9e\x9a\x64\xe8\xfd\x54\xdd\x78\x62\x5e\xde\x16\x7e\x43\x23\x22\x86\x11\x7d\xa9\x0e\xa0\xd3\x1c\xad\x8a\x5d\x56\x22\x62\xa1\xef\xd4\x14\xec\x2d\xa9\x4a\x99\x86\x55\x6a\x75\xaa\x8c\x57\xb0\xc0\x16\x86\x4e\x4a\xa3\x1a\x84\xd4\x0d\x93\xcd\x1d\x52\x41\x97\x1b\x7d\xa1\x70\xb8\x65\x9a\x21\x9c\x5b\x5b\xb1\x54\x4e\x54\x0d\x35\x8d\xea\x70\xce\x56\x53\xae\x24\xee\x7d\x59\x67\xe3\x4d\x4e\x4d\xc9\xbf\xf7\x25\x4f\xad\x1b\xb8\x44\x75\xf4\xde\xee\x88\x38\xe6\x8e\xd4\x35\x90\x0b\x3a\x43\xb3\xf6\x13\x58\x7a\x99\x11\x81\x98\x43\xc2\x56\xc9\x9b\x4f\x61\x99\x70\x8b\xbf\x72\x12\x7d\xfd\x7b\x08\xbb\xf4\x83\x40\x78\x40\x0e\x2a\xdc\x04\x8f\x38\x47\x70\x0c\xfc\x4b\x9e\xf9\x11\xd3\x2c\x14\xb9\x1f\x03\x2e\x51\x03\x47\x85\x9b\x4e\xeb\x74\xc3\x83\xd6\x70\x88\xd2\xfe\x8b\xb0\xa9\x03\x3a\x6c\x1b\x74\x2c\x70\xd9\xe2\xf0\xef\xb1\x2f\x33\xb0\x41\x06\x36\x94\x8a\x16\xa9\x58\xb8\x00\x5f\x0f\xa1\xd8\x01\x69\x0d\x9a\xb2\xe9\xaa\x86\xe1\x29\x80\xc2\x4d\x42\x57\x6c\x88\x53\xc5\xe0\x60\x1e\x67\x61\x55\xce\xcf\x63\x98\x89\xdd\x85\x56\xe7\x04\xe6\x38\x27\x26\x7a\x7b\x10\xfa\x66\x02\x4f\xd1\xe0\xa4\x34\xa1\x63\xae\x82\x50\x90\x3e\xad\x11\x14\x67\x4c\x3f\x95\x7b\xda\xb3\x7f\x49\xb1\x52\x5c\x1e\x71\xac\x48\x3c\x8d\xdd\x45\x33\x97\x59\xc1\x63\x6f\x09\x02\x9d\xb3\xb0\xc4\xbe\x8f\xe9\x8a\xf9\xc0\xc3\x4d\xac\x18\xb5\xd5\x69\x80\xbd\x5b\xbd\x2f\xaa\x0a\xb4\xea\x25\x6b\x43\xa6\xdc\x61\x58\x07\xd4\xd7\x7d\x92\x8d\xab\x05\x4c\xd3\x39\x45\x73\x87\xb9\x53\xa5\xb7\x81\x7a\x1b\x5a\xe7\xac\x17\xf6\xb0\x8e\x64\x9e\x81\x99\x35\x24\xf6\x0a\x79\x33\x94\x5b\xd3\xb5\x2d\x08\x6b\xc8\x0d\x57\xba\xc7\x94\x31\x62\xa3\x92\x47\xd7\x56\x7e\xe8\x27\x1e\xd6\xe0\x54\x48\xf2\xba\xea\xa8\x58\x8b\x35\x20\xf0\x71\x23\xe5\xbc\x66\x8e\x5a\x7a\xcb\x8a\x57\xda\x27\x6c\xce\xfc\xb0\xb6\xf6\x0f\x46\x8a\xd8\xb1\x00\x29\xfc\xf9\x5c\xcb\xce\x23\x5e\x4e\xed\x82\xf2\x04\x79\xe5\x04\x58\x9b\x53\xc8\x00\xf5\x8d\x0d\x01\x2b\x4e\x4a\xa4\x26\x38\x01\xce\xba\x59\xfc\x47\xed\xa1\x65\xf5\x26\xdd\x7e\xab\x83\x13\x1f\x96\x00\x79\x30\xa1\x5b\x1e\x22\xc7\x90\x44\x90\xac\x92\xfa\x54\x08\x39\xc1\xa0\x47\x18\x6e\xae\xae\x62\xe8\xc2\x03\xf6\x3f\xc5\xfb\x54\x07\x3b\x4b\x8f\x5a\x34\xd5\x42\x99\xa1\x9e\x18\x1a\x6a\x64\xd0\xb3\x50\x06\xcc\xcf\x2b\xc5\xee\xfa\x5d\x3e\x16\xda\xc7\x4b\x6f\xa5\xbe\x8b\xf2\xf4\x0f\xb4\x27\x11\xe8\x85\x4d\x07\xbc\x71\xd9\x22\x2c\x56\x40\x12\x88\xe5\x9f\xf2\xb6\xfc\x48\x06\xab\x96\x9e\x15\x50\x80\x21\xf3\x30\xea\x1c\x73\x66\xa5\x55\xca\x82\x43\x54\xf5\x75\xd7\x53\x45\x92\x12\x68\x63\x2a\xa7\xb8\x0f\xa3\xcd\x9b\x80\xd6\xcb\x45\xbe\x10\x48\x5c\x0e\x2e\x71\x2b\xc5\x55\x8a\xd4\xf1\x8a\x39\x21\x42\x27\x0f\xec\xc6\x2c\x3f\xa4\xe6\x85\x18\x8d\xfa\xe1\x8c\xf9\xaa\x59\x86\xa4\xa2\x04\x5c\xfd\x78\xfc\x43\xd1\x77\xa5\x26\x91\x59\x15\xa7\x21\x8d\x50\x75\xf8\x4d\xa4\xfb\x31\x49\x45\x4c\xc7\x7e\xbc\x02\x3c\x9d\x60\xf9\xff\xfb\x9f\xff\xf8\x57\x11\xd6\x0e\xb3\xaf\xed\xe8\x9e\x19\x47\x9f\xdb\xff\xb2\x51\x28\xff\xd6\xa5\xf3\x4b\xbe\x8f\xb4\xd5\x18\xaa\xd7\x08\x85\x13\x86\xda\x20\x70\x81\x40\xa0\x73\x2e\xa9\x46\x67\x18\x5b\xb5\x42\x30\xf3\xa8\xf1\xfa\xf0\xf0\x79\xf3\xa5\xf9\xe2\xe5\xeb\x86\xda\x1c\x70\x4c\x50\x80\x55\x98\x10\x3c\x15\x95\x1c\xec\xb0\xdc\x2f\x2b\x3b\x4f\x41\xe7\xe4\x0b\xba\x4e\xca\x33\xb9\xa1\x53\x36\x7b\xd5\x78\x0b\xaa\x37\xac\xde\xaa\x98\xc2\x01\x66\x04\xd3\x6a\xe7\xdf\x9b\xd1\x9d\xa0\x61\x84\xc2\x48\xfd\xc2\x7c\x97\x7d\xd3\xb5\x36\x4a\x73\xf9\xda\x69\xfd\x4d\x84\xe2\x81\xbf\x71\xce\x94\xb8\x7b\x54\x33\x52\xa3\x89\x9b\xdc\xc3\x18\x34\x1f\xca\xa0\xf1\x50\x06\x87\x0f\x64\xd0\x78\xfd\x50\x06\xaf\x1e\xca\xe0\xe5\x43\x19\xfc\xf4\x07\x19\xa8\x79\xfb\x0f\x9e\xcd\xef\xa3\x8d\x87\x30\xe1\xd7\x32\x66\x46\x6e\xc7\xa2\x49\x5f\x68\x27\x3e\x22\xc3\xea\x41\x0c\xe7\x42\xcc\x03\xfe\x18\xfc\x10\xb9\x0d\x7f\x31\x37\x58\xbc\xe0\x6c\xea\x1b\x57\x2f\x37\xea\x52\xf2\xdf\x54\xfe\x13\x00\x00\xff\xff\xc7\xe8\xdc\xab\x34\x23\x00\x00")

func scriptsBashBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBash,
		"scripts.bash",
	)
}

func scriptsBash() (*asset, error) {
	bytes, err := scriptsBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts.bash", size: 9012, mode: os.FileMode(420), modTime: time.Unix(1468256562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"scripts.bash": scriptsBash,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"scripts.bash": &bintree{scriptsBash, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

