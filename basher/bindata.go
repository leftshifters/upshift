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

var _scriptsBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x5a\xfb\x72\x1a\x39\xd6\xff\x1b\x9e\x42\x83\x5d\x71\x9c\x71\xd3\x06\x27\x93\xab\xf3\x15\x01\x92\xf0\x8d\x6d\x28\x83\xb3\x35\x95\xa4\x5c\xa2\x5b\x80\xd6\xdd\x52\xaf\x5a\x6d\xcc\x66\x3c\x6f\xb0\xcf\xb0\xef\xb1\x4f\xb5\x8f\xb0\xe7\x48\x2d\xe8\xc6\x38\x71\xc6\x99\xa9\xdd\x4c\x68\x49\xe7\x7e\xf9\x1d\x69\x46\x2c\xd5\xc3\x40\xf1\x44\x3f\xdc\x25\x5f\xaa\x15\x16\xcc\x24\xd9\x6e\x54\x2b\x69\xc4\x58\x42\xe0\x2f\x51\x4a\xbc\x88\x56\xaf\xab\xd5\xd5\xde\xa6\xdd\x7c\xdc\x1d\x0e\x5b\xef\xba\x87\xb8\xff\xa8\xff\xee\x6d\xef\xa8\x3b\x3c\xdc\x6e\xde\xa0\xd2\x74\x54\x48\xf3\xf5\x83\x06\xf9\x9d\x68\xc6\x08\xec\x03\x9a\xef\x14\x0d\x23\xf6\x37\x45\x93\x84\x29\x4b\x75\x6a\x3e\x91\xb9\xfd\xb6\xda\x34\x6f\x47\x8c\x0a\xbb\x07\xb8\x9d\x0f\x5a\xa3\xf7\x86\xf5\x31\xbd\x60\x6f\x65\x14\x32\x95\x56\x2b\x75\xdf\x9e\x9f\x93\x00\xb7\x97\x38\x36\x90\x18\xfe\x53\x38\x61\xe9\xc5\x17\x21\x57\xc4\x4b\x48\x3d\x4b\xd2\x19\x9f\x68\x3f\x92\xd3\xd4\xdf\xb4\x30\xce\x78\x14\x96\x56\x12\xc5\x2f\xa9\x66\x48\xfc\x2c\x41\xee\xac\x68\xd1\x20\x53\x11\xf1\x26\xe9\xf0\x88\xcc\xb4\x4e\xd2\x17\xbe\xaf\xe8\xbc\x3e\xe5\x7a\x96\x8d\xb3\x94\xa9\x40\x0a\xcd\x84\xae\x07\x32\xf6\x23\x36\xd1\x86\x0d\x48\xe6\x3b\x8e\x31\x4d\xe1\xb7\xcf\x45\xaa\x69\x14\xd5\xd3\x19\x79\x4d\xf2\xb5\xba\x66\x71\x42\x1e\x3c\x20\xc1\x2c\x96\x21\xf9\xf9\xea\xc6\x42\xdd\x5f\xff\xa2\xe2\xd2\x26\x14\x7b\xc8\x74\x96\xbc\x51\x6c\x3e\x92\x32\xb2\x62\x8f\xfa\xfd\x23\x63\xdd\x31\x7c\x26\x39\x6f\xe3\x53\x76\xc5\x35\xd9\x5f\x1e\x7b\xc7\x62\x7b\xe2\x5d\xf7\xd8\x1c\x68\xf7\x4c\x0c\x9c\x4a\xa9\x07\x34\x4d\xe7\x52\x85\x87\xdb\x07\xd5\x6a\x85\x4f\xc8\x47\x52\xdb\x6e\xd6\xc8\xe1\x21\xa9\x69\x95\xb1\x1a\xf9\xfc\x92\xe8\x19\x13\xd5\x4a\x65\x8b\xb4\x7b\x84\xa7\x04\xbf\xef\x91\x39\x23\x42\xce\x89\x60\x2c\x24\x49\x4e\x05\x36\xe5\x24\x0e\x6a\xe4\x27\x20\x51\x3c\x6e\x63\xce\x63\x76\xf5\x77\x92\x66\x21\xfc\x1c\x92\x29\x8b\x4b\xd2\x57\x9c\xfc\xf0\x97\x28\x65\xcb\x0f\xb8\x34\xe1\x55\xf7\x71\x8b\x9c\x81\x6b\x48\x3a\x93\x59\x14\x12\xbd\x48\x18\x50\x41\x5e\x45\x69\x6e\xe5\xe1\x58\x00\x41\x0c\x5f\xae\x07\x59\x94\xdb\xf5\xb4\x7b\xdc\x1f\xd9\x94\x79\x73\xda\x3a\x69\xbf\x37\xc6\xca\xb3\x07\xed\x54\x0e\x67\x08\x13\x92\x64\x86\x2e\x64\x4c\x29\x9a\x0f\x72\xda\xc3\x6c\x0c\xbe\xcf\x22\xd6\x13\x5c\x2f\xd3\xc3\x92\x5b\xcf\x0e\x24\x97\xba\xfd\x20\x33\xfc\xbc\x91\x20\x45\x92\x67\x49\x08\x91\xfd\x5d\x44\x33\x73\xe4\x26\xd9\x81\x0c\x7b\xd6\x46\xdf\x20\x97\x40\x20\x3b\x6b\x96\xa8\x54\x72\x32\xa7\x2c\x91\xb7\x09\x76\x66\x23\xbb\x2d\xc5\x84\x4f\x2d\x2d\x05\xdb\x9d\x54\x9e\x77\xc9\xd4\x58\xa6\x6c\x03\xe5\xa1\xa6\x4a\x0f\x79\x9c\x45\x54\xcb\xbc\x16\x75\xba\x1f\x7a\xed\x0d\x32\x5e\x05\x2a\x13\x46\x4a\x95\xc5\x90\xbc\x50\xdd\xe6\x10\x79\x8d\x1a\x69\xbc\xf6\x43\x76\xe9\x8b\x2c\x97\x1e\x29\xb7\x65\x9c\x70\xf0\x4f\x7f\x68\xa9\x0e\x4e\xfb\xff\xdf\x6d\x8f\xce\x47\xbf\x0d\x2c\x6d\xf7\xc1\x96\x34\x88\x87\x61\xfb\x7d\xf7\xd8\x86\x83\x13\xe1\x71\xb1\xea\x3d\x59\x93\x27\x65\x9a\x78\x92\x24\x3c\x61\x13\xca\x23\x4c\xf3\xab\x40\x86\xcc\x14\x2b\xe2\x19\xc9\x4c\xe6\x79\x69\x30\x63\x71\x9e\x25\xde\x8c\x43\xad\x9a\xb1\x28\xb2\x05\xab\x2b\x2e\xb9\x92\x02\x15\x82\x8d\xe1\x05\xe1\xc9\x4c\x0a\x96\x3a\x9b\x10\x2f\x84\x16\xc0\x05\xd5\x5c\x0a\x52\x4b\xe0\xe3\x44\xaa\xf8\x70\x6d\xdb\x9e\xa0\x31\x03\x81\x6b\xb8\x1f\x6a\x23\x0b\x3b\x54\xd3\x01\xd5\xb3\xb5\x2a\x9a\xdb\xbf\xb6\xfd\x04\x33\xf6\x2a\x48\x14\xd3\x7a\xe1\x5a\xcd\x5f\x68\xae\xab\x40\x43\xa1\xbb\xdd\x2e\x3f\x42\x77\x0d\xdb\x0b\x0a\xa2\x56\x2d\x15\xcc\xc0\x1c\x7f\x4e\x31\xb7\x76\xd2\x3a\xfe\x8b\xa2\xe1\x1b\xce\xf2\xa8\x15\xbf\xbc\xb6\xfd\xb8\x7e\x15\xe4\x2b\xc4\xfd\xfb\x56\xbf\x76\xaf\x12\xa9\x6e\x7a\xd6\xea\xd4\x28\xea\xd4\xfc\x3e\x9d\x98\x21\x9c\x1b\xd8\xfd\xec\x27\xe8\xac\x74\x10\x71\x70\x85\x6b\xd3\xbe\x5d\xab\x27\xe6\xeb\x2d\x4a\x35\x0a\x4a\xe5\xc4\xca\x3b\x8a\xf5\xc3\x98\xb2\xac\xe7\x5b\xa6\x83\x59\x4b\x60\xa9\xa2\x5c\x0d\x94\xbc\xe4\x29\x88\xc2\xc5\x14\xfe\x3e\x81\x5a\x90\xe3\x8e\x2d\x68\x29\xd0\xf3\x00\x06\xa4\x24\xe5\xd3\x19\x7e\x09\xe5\x5c\x44\x92\x86\x04\x0b\x20\x17\x5a\x2e\x45\x27\x13\x63\x0d\xdc\x44\x45\x68\x3a\x1f\xc1\x3a\x84\x6d\x69\x92\x89\xc0\x84\xe6\x40\x26\x18\x8a\x6c\x13\xd3\x6a\xa5\xd5\x6e\xf7\xcf\x4e\x46\xe7\xdd\xe3\x56\xcf\xf6\x77\x64\x8b\x35\x92\x22\x9a\xc9\xc8\xf2\x93\x13\xe3\x1c\xc5\xc8\x17\x56\x98\xc7\x5f\xa2\x9e\x4a\x7c\x49\x1e\xd5\x63\x39\x06\x06\x89\xe3\xb9\x02\x45\x95\xaf\xca\x83\xa6\xe2\x22\xdc\xb0\xb6\xac\xbf\xdd\xa3\xfe\xa0\x7b\x7a\x9e\x0b\x6e\x1b\xe7\xd9\x49\xe7\xa8\x7b\xde\xeb\x74\x4f\x46\xbd\xb7\xbd\xee\xa9\x09\x17\x23\xb5\x11\x14\xbc\x9a\xc3\xcb\xaf\x31\x77\x1e\x78\x07\x91\x85\x16\x3c\x3b\xeb\x75\xc8\x44\xc9\xb8\x20\xfc\xd6\x12\xb3\x4d\x21\x5a\x72\xd0\x66\x90\x5a\x7c\x95\x28\xff\x59\x73\xff\x59\xf3\xd9\x73\xfa\xcb\x41\x40\x1f\xb3\x03\x7a\xf0\x84\x3e\xc6\x43\x47\x52\x26\x40\x53\xc9\x0c\x44\x42\xfb\x19\x86\x7b\x04\xe0\xcb\x42\x66\x80\x17\x34\xa1\xe4\xac\xd3\xeb\xec\x11\x1a\x1a\x47\xc6\x04\xfc\x8c\x52\x98\xa8\x94\x13\x92\x2c\x5d\x56\x4a\x03\x03\x80\xbc\x90\xd4\x56\x1e\xf0\xcb\x28\x6a\x04\x44\xd6\x02\x86\x00\x1e\x49\x75\x8a\x08\x47\x66\xc6\xda\x86\xf4\xe1\x84\x5a\xb0\x03\x85\xcc\xf1\x3b\x81\xf2\x85\x48\x67\x49\xfc\x86\x6b\x5f\x42\x64\x20\x6a\xca\x32\x0e\xb8\xee\xa1\x9f\xa5\xca\x8f\xf8\x98\x5d\xb1\xc0\x37\x99\xf6\x26\x0b\xc3\x05\xf1\x02\xb2\x33\x50\x10\xbe\xc6\xae\x3b\xc4\x74\xc5\x54\x87\x40\xfb\xd5\xab\x57\x64\xfb\x61\xca\x00\x19\x73\xbd\x20\x41\x0c\xdd\xb3\x43\x3c\x0e\x89\xf4\xa5\x20\xc6\x75\x0d\x92\x6c\xd9\x4c\x77\x77\xab\xc8\x75\x8b\xf4\x26\x68\x3b\xf4\x95\xd1\x6a\xcf\xa6\x41\x20\x93\x05\xe1\xda\x98\x18\xe0\xc8\x8c\xa6\x62\x07\x8c\x1c\x29\x46\x41\x98\x31\xb3\x5b\x38\x43\xd8\xe6\x50\xe4\x17\x54\xe1\x7a\x03\x94\xcc\xad\x3c\x21\x7f\xf8\x47\x7c\xac\xa8\x5a\xf8\xc7\xc6\x08\x1d\x76\xc9\x03\xe6\x17\xa3\xe9\x13\x71\xe6\xf4\x73\x82\x37\x72\xa1\x48\x19\x14\x18\xc2\xa2\xd0\x11\x88\x3b\x15\x52\x31\xfb\x39\x41\x53\x4d\xd6\x2d\x90\x3b\xce\x41\xcf\x5c\x98\x4f\xa2\x66\x0e\x39\xfc\x6a\xb0\x33\xea\x8f\xe7\x30\x8e\x6c\x24\xe7\xb5\x02\xff\x09\x12\xd4\xe6\x86\x79\x7f\x88\x7a\x65\xf9\x8f\xe1\xab\x98\x92\x35\x35\x3e\xe6\x67\x3f\xa3\x74\x6b\x5a\x20\xe8\xae\xac\x47\x26\x0e\x01\x55\xb7\x18\x42\x83\xad\xae\xc0\x7f\x69\xa7\x99\x25\x70\x77\x69\x14\x58\x47\xf8\x4e\xba\x4e\x16\xc2\x6c\xf1\x1b\xa4\xa0\x19\x2c\x40\x18\x4c\xbf\x54\x42\xcc\x27\x05\xa5\x97\xc9\x57\xcc\x84\x5c\xde\x4d\xc3\x82\x23\xff\x9e\x2d\xf6\x4c\x82\x17\xa9\xc3\x6f\x75\x57\xea\x8e\xb8\x9d\x1b\xa0\x7f\x28\xc9\xc3\x33\xc1\xd7\x70\xf3\x37\x47\xdf\xcc\x1d\x69\xad\x63\xe8\x02\xdd\x1c\x8d\x77\xa0\x85\x4e\xef\x4e\x9a\x17\x4e\x01\x9c\x86\x1f\xc1\x85\x56\x34\xd8\x80\xf8\x73\x3e\xad\x34\x65\xf1\x38\x62\x77\xe7\x41\xf3\x13\x77\xa3\x7f\x04\xc6\xbf\x3b\xed\x08\x4b\xd2\x6d\xa4\xcc\x08\xd0\x82\x3e\x7a\x09\x85\x29\x87\x28\xad\xf6\xaf\xee\x8e\x03\x7a\xe6\x09\x34\xa2\x51\xef\x43\x6f\xf4\xdb\x4d\x88\x42\xc3\x31\x0c\x8c\x80\xa6\x09\x8d\x49\x8a\xa4\x88\x27\x80\x01\xc2\x09\xdb\x8e\x9c\xc4\x14\x9a\xf5\xac\x5b\x9c\x34\xca\xa4\xaa\x15\xc5\xc2\xb6\x8c\x20\x74\x0e\x77\x3e\xed\x1f\x1c\x7c\xdc\x7f\x79\xd0\x88\x77\xf0\x72\x04\x4a\xd9\xfa\x4a\x13\x57\xc6\x51\xc6\xd6\x17\x1e\xe3\x82\x90\xe5\xcf\xf0\xad\x5a\xe9\x1e\x9f\x1d\xb5\x46\xfd\xd3\x5b\xa0\xd7\x72\xf9\xb4\x3b\x3c\x3b\x1a\x0d\xa1\xd0\x0b\x39\xcb\x12\x48\xbf\xd6\x49\xe7\xb4\xdf\xeb\x9c\xbf\xef\x1f\x77\x7d\x84\xd1\xa9\xcf\x72\x4d\x00\x43\xd2\xcb\xd0\xce\x41\xcd\xd7\x30\xae\x36\xf0\x8f\x07\x58\xb9\xa1\x2f\xf5\x3b\x7d\xf2\x02\xda\x13\xe0\x1d\xf8\x1f\x25\x63\x3e\x25\x5b\xef\xc1\xbc\x7b\x44\x0a\x28\x87\x4c\x29\xa9\x60\x41\x31\x80\x23\x3a\x53\x82\x85\xae\xf0\x4d\xb8\x82\xae\xa8\xe7\x92\x40\xd7\x90\x22\x84\x56\xda\x83\x91\x33\xb8\x30\x38\xa8\x07\x1d\x09\xab\xfd\x05\xde\x18\x50\x32\xa7\x0b\x22\x33\x5d\xb8\x7c\xda\xa0\xcd\xab\xed\xa6\x15\xab\x3d\x63\x40\x06\x4a\x0b\xf0\x01\xce\x73\x8a\xa2\x0d\x5a\x27\xbd\x36\xf9\x88\x2d\x19\x51\x3c\x62\xb4\x3d\x18\xc3\xb5\x93\x67\x0e\x53\xd2\x94\x38\xb5\x09\x62\xfe\xcf\x18\x2a\x70\xea\x3c\x87\x29\x0f\xcd\x7d\x04\x94\xdc\x75\xe6\xd7\x88\x15\xc1\x8d\x60\x4b\x73\x00\x8c\x16\x58\x51\x7a\x65\x21\x12\x2a\x78\xb0\x67\x60\xc4\x1c\x7b\x1a\xb5\x30\x05\xaf\x8c\x4c\xd3\x5b\x90\x98\x41\xa2\x84\x30\xdb\xee\x11\x21\x35\xac\x2c\xef\x58\xbe\x14\x64\x01\x86\xde\x14\xea\x61\xa1\x42\xba\x82\x85\x58\x61\xa9\xc5\xdc\x18\xd1\x20\xcf\x2c\xa9\x93\x63\xba\x18\x33\xc3\x71\xfb\xcb\x32\x1a\xaf\xdd\x6e\xac\xeb\xdb\x5f\x5c\x68\x5d\x93\x0b\x10\x87\x9b\x9e\x1b\x48\xa5\x58\xa0\x6d\x41\x5b\x55\x46\xc5\x76\x52\x32\x9f\xe5\x4a\x44\x72\x4a\x52\xba\x48\x5f\x7c\x12\xe5\x8d\x1b\xec\xe5\xb6\xd8\xe2\xb8\x4f\xb6\xf0\x0f\x84\x2c\x03\x34\x50\xe1\xc6\x66\xc8\x18\x40\x89\x88\x5f\xb0\xdc\x8e\x10\x66\x42\x3a\x3b\x46\x4c\x83\x04\x81\x71\x37\x9e\xb6\xe1\x06\x07\xbb\xa7\xa7\xc0\xed\x3b\xdc\x66\x0e\x58\xb7\xad\xb0\x44\x81\xca\x06\x83\xdf\x45\xbf\x5c\xc1\x06\x2a\xd8\x30\x22\x76\x51\xc4\x0d\x6d\x0c\xab\x3a\xce\x43\x63\x3a\x86\x86\xb3\xfd\xa5\x50\x14\xae\x4d\xdd\x81\x2e\x53\x74\x0e\xc4\x71\xee\x56\x63\x7c\xe7\xc3\x9c\xed\x16\x69\x75\xde\x90\x29\x8c\x3a\xa9\x9d\x46\xa4\x19\x9b\xf0\xd4\x9c\x82\x44\x28\x09\x1e\x0b\x0d\x3a\x20\x9a\xe3\x58\x82\x7e\x86\xf0\x33\xb1\x67\x2d\xfb\xf7\x0c\x32\x25\x64\x09\x83\x8c\x84\xd3\x80\xc5\x2c\x71\x9d\x27\x3c\x20\xb1\x28\xb2\x31\x4b\xe6\x80\xa3\x21\x5c\x21\x1e\x98\x58\xf9\x8a\x62\x67\x84\xb2\x1f\xdb\xf5\xa2\xa8\x04\x67\xbe\x74\xa9\xc8\x98\x05\x14\xf2\x00\x71\x32\x47\xde\x30\xaa\x10\xe8\x47\xe8\x4d\x53\x85\x51\x6e\x0f\xe4\xf6\xac\xcc\x39\x72\x3c\x81\x3c\xd2\x2e\x02\x73\x6d\x90\xed\x25\xd0\xa6\xc0\x77\xcf\xe6\xb6\xc4\x5a\x83\x66\xb8\xb4\xf0\xa9\x5c\x23\x56\x22\x01\xc6\x84\xc2\x24\x38\x94\x7c\x90\x48\x6a\xb4\xba\xc1\x9f\x90\x8b\x7b\x04\x8b\x4f\x98\x18\xe3\x35\x5d\xd5\xb2\x53\x9b\x5a\x58\x9b\xd0\x29\xe5\x62\x6f\x69\x1f\x84\x09\x40\x8e\x46\xb0\x83\x4f\xa7\x96\xb7\xf3\x78\x39\xb4\x0b\xc2\x63\xc9\x2b\x07\xc0\x52\x9d\x42\x04\x98\xfb\x54\x2c\xac\x30\x79\xa0\x98\x24\x88\x64\xea\x10\xcd\xb0\x7d\xda\xed\x9e\x9c\x1f\xf5\x5b\x9d\xde\xc9\x3b\x48\x81\x55\x1f\x83\xed\xe0\x92\x84\xa4\x8b\xb4\x3e\x96\x52\x9f\x83\xd3\x13\x70\x37\x33\x57\x37\x78\x27\x42\x76\x3e\xa9\x1d\xcc\x83\xca\x7c\x86\xe8\x13\x73\xa1\x4c\xd0\xe2\xeb\x86\x01\xd8\x76\x72\xc8\x0b\xf3\xe3\x6a\x11\x38\xde\xc9\xc6\xd2\xda\x78\x3e\x5b\x98\x9b\xe2\x99\xfd\x03\xf4\x49\x25\x58\x61\xd5\x01\xaf\x43\x1a\x8b\x62\x06\xa4\x91\x9c\xff\x9f\x43\x9c\x3f\x48\x61\x07\x4b\x31\x81\x22\x70\xd9\x0c\xbc\xce\x20\x66\x16\x56\xa4\xdc\x39\xb8\xab\xbe\xec\x7a\x26\x49\xb2\xd4\xa2\xc3\x0c\xe6\x6b\xd0\x79\xe5\xd0\x7a\x39\xc9\x63\x09\x9b\xcb\xce\x45\x6a\x25\xbf\x6a\x99\x05\xb3\x62\x4c\x48\x11\x38\xc7\xae\xd4\xe2\x02\x9b\x17\xd4\x68\x90\x0f\xe0\xcf\xb3\x66\xb9\x24\x15\x39\x0c\xb3\x84\xa9\x9f\x8a\xb6\x2b\x35\x89\x5c\x2b\x95\x09\x84\xb2\x75\xc4\xd2\x3b\x0a\xb9\x42\x4d\x87\x7e\xbc\x20\x70\x3a\x85\xf4\xff\xef\xbf\xff\xf5\x9f\x62\x59\xdb\xcf\x11\x33\xfc\xbf\x08\x6d\xf3\xe7\x92\xce\xaf\x6e\x18\x6f\x9b\x21\xcd\x8e\xe1\xa6\x2e\x78\x66\x02\x87\x01\xdc\x20\x41\xcc\xc9\x09\xf8\xd2\x8c\xe0\xd4\x3f\x68\x3c\xdf\xdf\x7f\xdc\x7c\xea\x3f\x79\xfa\xbc\x61\x26\x6f\x80\x05\xa6\x40\x15\x10\xc1\xcc\x78\xc1\x15\x37\x48\xef\x8b\x6a\xe5\x21\xb1\x31\xf8\x04\xef\x91\x5c\xe4\x36\x6c\x88\xe6\x4b\x8d\x97\xc4\xf4\x82\xc5\x4b\xe3\x43\xb2\x0b\x11\x40\xad\xd8\xee\x36\x19\xaf\x09\x3d\x4f\x48\x2f\xe3\x37\x71\xf5\x52\x29\x4b\xe5\xb6\xd3\xe6\x76\xc5\xd2\x80\x7f\xc3\x5c\xa0\x61\x5e\xaf\xe5\x5b\xbd\xe6\x41\xed\x9e\x04\x9a\xf7\x25\xd0\xb8\x2f\x81\xfd\x7b\x12\x68\x3c\xbf\x2f\x81\x67\xf7\x25\xf0\xf4\xbe\x04\x7e\xf9\x93\x04\x0c\xbe\xfe\x93\x67\xdd\x15\xb5\x77\x1f\x22\xec\x0a\x86\x2f\xcf\xe9\x11\x37\xf1\x31\x25\xe5\x50\x09\x16\xf7\x22\x38\x95\x72\x1a\xb1\x1f\x41\x0f\x2a\xb5\xc7\xe3\xa9\x47\x55\xcc\xe8\x98\x7b\x97\x4f\x57\xe2\x62\xf0\x9b\x47\xd9\xc2\x9b\xd0\xda\x7b\xef\x1f\xbe\xbb\xd2\xad\x56\x6c\x2d\x5d\x7d\xf1\x03\xfb\x8a\x04\x24\xf2\x9c\x6e\x33\xc0\x56\x13\x1e\x80\x50\xf9\xf5\xe1\x9b\xd6\xb0\xbb\x9a\x39\x97\x77\x5a\x3c\xc6\x2b\x63\x9c\xfd\x68\x02\xfd\xa3\x1e\x80\xa8\xde\x45\xe1\xc2\xe5\x57\xb6\x08\x66\xd0\xff\x53\x7c\x68\xe6\xa2\x7e\x91\xff\x26\xde\x88\x98\x2b\xb5\x31\x17\x3e\xde\x6f\xa7\x7c\x5a\xfe\xe8\x78\x6c\xe4\x16\xf2\x14\xd0\xc3\x38\x43\x00\x57\x4f\x1a\xcd\x1f\xcd\x94\x78\x03\x52\xab\x7d\x9b\xf5\x5f\xa1\x2f\x3e\xa4\x01\x88\xd4\xac\x95\x24\x7d\xd1\xd3\x99\x70\x4e\xf8\x9e\x5b\xe2\xf2\x93\xc3\x41\x15\xda\xa0\x0c\x33\x68\x9c\x85\x9b\x63\x0c\xb6\x24\x39\xc7\xb1\xce\x3c\x8d\xd8\x20\x42\xe4\xd9\x1b\xb4\xa4\xe0\xa3\x6f\xb1\x86\x6d\xab\x59\x3a\xe1\x11\x0e\x67\xe6\x3c\x50\xc6\xf9\xcd\x50\x46\x6e\x1e\x4f\x68\xfe\x14\xe3\xa5\x17\x3c\x39\x47\xe4\x0a\x8d\xf5\x1c\x12\xf7\xdc\x3c\x6d\x9c\x83\x7c\x01\x60\x66\x84\x14\xe6\x9e\x6b\xf5\x84\x69\x7a\x4e\x18\x0e\xe7\x10\xab\x43\x68\xd5\x81\x13\xaa\x75\xda\x7e\xdf\xfb\xd0\x5d\xcd\xf6\x28\xcf\xf2\xc7\xea\x0a\x79\xd3\x73\x87\x6f\xc9\x65\x09\xfa\xb5\x74\xb5\xbc\x04\x26\x61\x88\xc2\x14\xd9\xba\xcb\x6b\xe0\x63\x10\xef\x76\xc3\xf6\xfe\x4c\xfc\x93\x27\xc4\xfb\x47\xe1\x65\xa5\x59\x47\x95\x81\xfd\xf2\x93\x8e\x13\xd8\xaa\xe2\xf5\x4d\xf0\x31\x4f\x53\x55\xda\x5c\x96\xd0\xbe\xbf\x49\x84\xfe\x78\xcb\xa9\x26\x9b\x5f\x71\x36\x1f\xf2\x1f\xd5\xc3\x45\xc4\xc7\x77\xa0\xef\xb6\x22\x9f\x1b\xc2\x1b\x2d\x41\xca\xba\xd3\xaf\x9e\x6f\xab\xfb\xf5\xba\x55\xae\x24\x99\x39\x65\x2f\xf9\x7a\xfd\x61\x4f\x04\x8a\xe1\xa3\xeb\x1b\x74\xf8\x49\x16\x8f\xdd\x7f\x0a\xb3\xfe\x3e\x56\xad\x00\xee\x89\x39\x94\xb2\x4c\x68\x00\xae\xf8\xf2\xae\xd8\xa5\x67\x9f\xb2\xf2\x0a\x19\xe0\xe2\xae\xf9\x0f\x4f\x54\x99\xe8\xd7\xaf\xeb\x6b\xf6\xba\xbe\xfd\xf6\x4d\x26\xc2\x88\x7d\x00\x20\x07\xa9\x5c\xc3\xdc\xee\x89\x89\xb4\x0f\x66\x78\x53\x21\xd8\xbc\x48\x14\xeb\x41\x3e\x05\x17\xa4\xb3\x33\x30\x7c\xba\x21\x45\x31\xa6\xd6\x48\x15\x09\x2c\x07\xfb\xf5\x3d\x0f\x1f\x6e\x20\xfa\x73\x63\x77\xd7\xa2\xcb\xca\xd7\x54\x1c\xc2\xd4\xfc\x62\x4d\x43\xb2\x5d\xe6\x80\x6f\xa3\x45\x95\x6b\xe6\x72\x06\xa2\x9e\xf4\x46\x43\xa8\x3d\x67\x29\x4b\x4f\xa4\xe8\x5e\xb1\x18\x9f\xcc\x03\xb5\xb0\xf3\x32\x8e\x2a\xf8\xa0\x92\x3f\x41\xd8\xdb\x10\x18\xd3\xab\x15\x2a\xd2\x39\x53\x23\xb9\xda\x7c\x27\x4f\x7c\x95\xdd\x0d\x29\x77\xad\x17\x3e\x92\xed\x9b\xec\xf0\x92\xfc\x51\xad\x83\x20\xf8\x04\x6a\x51\x17\x5f\x15\x6a\x8f\xc8\xe7\x95\x23\xbe\x26\x0d\xaa\xfe\xe2\xeb\xba\x8f\xf1\x39\xdd\x68\x7f\xd3\x7a\x26\xd0\xff\x17\x00\x00\xff\xff\xb3\x07\x66\x3e\x68\x26\x00\x00")

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

	info := bindataFileInfo{name: "scripts.bash", size: 9832, mode: os.FileMode(420), modTime: time.Unix(1467952751, 0)}
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

