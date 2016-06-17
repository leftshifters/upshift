// Code generated by go-bindata.
// sources:
// scripts/main.bash
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

var _scriptsMainBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x59\x7f\x52\xe3\x46\xf6\xff\xdb\x3e\xc5\x8b\xa1\x86\x99\xef\xd7\xb2\xb0\x99\x49\x26\x99\x61\xb6\x1c\xdb\x61\xbc\x01\xec\xc2\x90\xad\x54\x26\x35\x69\x4b\x6d\xab\x17\x49\xad\xea\x6e\x61\x5c\x2c\xb9\xc1\x9e\x61\xef\xb1\xa7\xda\x23\xec\x7b\xdd\x92\x91\x0c\x54\x48\xc8\x56\xcd\x80\xad\xee\xfe\xbc\xdf\xef\x7d\x5a\x9c\xb0\x4b\xfe\x9d\x8c\x43\xae\xf4\xcb\x57\x70\xd3\x6c\x24\x97\xa1\x50\xe0\x65\xd0\xc9\x33\x1d\x89\x85\xf1\x63\xb9\xd4\xfe\x43\x0b\xf3\x5c\xc4\xa1\xdf\xbc\x6d\x36\x2f\xb2\xa5\x62\x21\x9f\x05\x4a\x64\xc6\xe1\x04\xb9\x8a\xc1\x5b\xe8\xd9\x31\x44\xc6\x64\xfa\x1b\xdf\x57\x6c\xd5\x59\x0a\x13\xe5\xf3\x5c\x73\x15\xc8\xd4\xf0\xd4\x74\x02\x99\xf8\x31\x5f\x18\x8b\x89\x6a\xf8\x25\x7c\xc2\x34\x7e\xf7\x45\xaa\x0d\x8b\xe3\x8e\x8e\xe0\x03\x14\x6b\x1d\xc3\x93\x0c\x5e\xbc\x80\x20\x4a\x64\x08\xff\x7f\x7d\x6f\xa1\xe3\x6f\x3f\x51\x49\x6d\x13\xa9\x3d\xe3\x26\xcf\x8e\x78\xe2\x34\x3e\x1a\x9d\x1c\xee\x76\x9b\x8d\xc1\xf8\x70\xb7\xd7\x6c\x9c\x49\x69\xa6\x4c\xeb\x95\x54\xe1\xe1\xee\x41\xb3\xd9\x10\x0b\xf8\x09\x5a\xbb\xbd\x16\x1c\x1e\x42\xcb\xa8\x9c\xb7\xe0\xe7\x77\x60\x22\x9e\x36\x1b\x8d\x1d\x18\x8c\x41\x68\xa0\xe7\x6d\x58\x71\x48\xe5\x0a\x52\xce\x43\xc8\x0a\x14\xdc\x54\x40\x1c\xb4\xe0\x0b\x84\xa8\x1e\x6f\xf0\x20\x92\xe0\x71\xb7\xfa\x0f\xd0\x79\x88\x5f\x67\xb0\xe4\x09\x14\x2e\x00\xd2\x0e\x37\x5e\x0b\x03\xfb\xf8\x89\xc7\x9a\x6f\x1e\xd0\xd2\x42\x34\xcb\x87\x3b\x70\x81\x3e\x06\x1d\xc9\x3c\x0e\xc1\xac\x33\x8e\x28\x24\xab\xaa\xcd\xa3\x32\x4a\x11\x08\x88\x6e\x3a\x12\x66\x9a\xc7\xb1\xf3\xd2\xd9\xe8\x64\x72\x3e\xb2\x8e\xfa\xf6\xac\x7f\x3a\xf8\x68\x9d\x75\x3c\x39\xfa\x6e\x7c\x3c\x22\x3f\x35\x2a\x29\xd5\x6c\x60\xbc\x21\xcb\x2d\x2e\xec\xf6\xa0\xf7\xe1\x45\x17\x8d\x33\x9c\x03\x6e\x75\xd8\xb3\x7c\x8e\x41\xcc\x63\x3e\x4e\x45\x91\x3c\x1b\xb8\xee\x03\x70\xba\xdc\x8f\x3a\xe3\xd7\x2a\x64\x77\x1b\xf2\x22\x0b\x99\xe1\xbf\x0b\x34\xb7\x47\xee\xc3\x4e\x65\x38\x76\x3e\xfa\x0d\xb8\x0c\x33\xb2\xf4\x66\x0d\xa5\x51\xc0\x9c\xf1\x4c\x3e\xa6\xd8\x85\x4b\xd1\x81\x4c\x17\x62\xe9\xb0\x14\x6e\x2f\xb5\xf2\xbc\x2b\xae\xe6\x52\xf3\x07\x90\x5d\x36\x63\x1d\xc6\xfc\x6f\x0e\xb8\x6e\xa5\x5d\x81\x95\x62\x59\xc6\x95\x3d\x60\x98\x32\x33\x91\xe4\x31\x33\x52\xb9\x23\xc3\xd1\x0f\xe3\xc1\x03\x46\x5d\x07\x2a\x4f\xad\x59\x2a\x4f\xb0\x6c\x35\x78\x2b\x4c\xd5\x6e\x0b\xba\x1f\xfc\x90\x5f\xf9\x69\x5e\x98\x4b\xc8\x03\x99\x64\x02\x03\x3a\x99\x39\xd4\xe9\xd9\xe4\xaf\xa3\xc1\xf9\xe7\xf3\x1f\xa7\x0e\xbb\x7c\x30\xed\x9f\xbb\x04\x9a\x0d\x3e\x8e\x4e\x5c\xfe\x94\x2a\xbc\xb6\x9e\x29\xb6\xbc\xd9\xd2\x47\x73\x03\x9e\x84\x4c\x64\x7c\xc1\x44\x4c\x05\x7e\x1d\xc8\x90\xdb\x9e\x04\x9e\xd5\xcc\x96\xaa\xa7\x83\x88\x27\x45\x59\x79\x91\xc0\x2e\x15\xf1\x38\x76\xad\x6a\x94\x5e\x09\x25\x53\x32\x08\x37\x86\x97\x20\xb2\x48\xa6\x5c\x97\x3e\x01\x2f\xe4\xda\x88\x94\x19\x21\x53\x68\x65\xf8\x70\x21\x55\x72\xb8\xb5\xad\x9d\xb2\x84\xa3\xc2\x2d\xda\xaf\xc4\x15\x0f\x87\xcc\xb0\x29\x33\xd1\x56\xb3\x2c\x02\xd6\xda\x7d\x43\x25\x7e\x1d\x64\x8a\x1b\xb3\x26\x8f\xf5\x55\x10\xe1\xc1\x3f\xe6\xb1\x72\xed\xb4\x7f\xf2\x3f\xf2\xdb\x6f\x98\xe5\x31\xa7\x7e\x7d\x6d\xf7\x75\xe7\x3a\x28\x56\xa0\xfc\xfd\xa8\x07\x46\xd7\x99\x54\xe6\x9e\x03\x9c\x4d\xdd\xaa\x4d\xbd\xdf\x67\x13\xb7\xc0\x85\x83\xcb\xaf\x93\x8c\x42\xaa\xa7\xb1\xd0\x06\x3a\x19\x5a\x87\xe5\xe5\xbb\xb5\x4e\x66\x9f\x3e\x62\x54\xb7\x62\x54\x01\x76\x6f\x87\xc8\x58\xb5\x40\xad\x43\xeb\xd6\x4e\x65\x46\xb9\xc3\xa7\x4a\x5e\x09\x8d\xaa\x88\x74\x89\x9f\x17\x58\x35\xc5\x20\xde\x81\x23\xb4\x8b\x3a\xf6\xc5\xc5\x78\x08\x0b\x25\x93\x8d\xa2\xb4\x5a\x4e\xd5\x25\xea\x5a\x8c\x55\x3b\x4b\x93\xeb\x4c\xf9\x6f\x7b\xfb\x6f\x7b\x6f\xbf\x66\x5f\x1e\x04\xec\x35\x3f\x60\x07\x6f\xd8\x6b\x3a\x74\x2c\x65\x86\x98\x4a\xe6\xcb\x08\xa8\x41\x59\x81\x6d\xc0\xb9\xb4\x96\x39\x0e\x02\x03\x0c\x2e\x86\xe3\x61\x1b\x58\x18\x92\xf4\x04\x8c\xb4\x5a\x58\x9f\xc8\x05\x64\x85\x96\x5b\x41\xb0\x93\xcd\x0b\xa1\xd5\xf1\x37\xee\xac\x8f\xc7\x73\x04\x29\x97\x60\x61\xcf\x01\x0e\x1a\x6d\x34\x8d\x2e\x99\xa7\x61\xe9\x80\xc3\x05\x73\x53\x0c\x8b\xad\x94\x77\x8a\x25\x46\x23\x6c\x03\xfe\x7f\x9d\x44\xce\x71\x21\x2b\x3d\xf8\x0e\x42\x49\xe3\x30\xcf\x05\x0e\xec\x97\x7e\xae\x95\x1f\x8b\x39\xbf\xe6\x81\x6f\xe3\xfc\x6d\x1e\x86\x6b\xf0\x02\xd8\x9b\x2a\x81\x15\x4f\x7e\xdd\x03\xdb\xbd\xb4\x09\x11\xfb\xfd\xfb\xf7\xb0\xfb\x52\x73\xe4\x2e\xc2\xac\x21\x48\xb0\xcb\x0d\xc1\x13\x18\xc0\x9b\x8a\x1a\xb7\x2d\x0c\xee\xa6\xe9\xbd\x7a\xd5\x24\xa9\x3b\x30\x5e\x90\xef\x28\x56\xd6\xaa\xb6\xb5\x1c\x02\x99\xad\x41\x18\xeb\x62\x9c\x33\x11\xd3\xe9\x1e\x3a\x39\x56\x9c\xa1\x32\x73\xee\xb6\x08\x4e\xf3\xb8\xa4\x07\x37\x64\xc2\xed\x03\x1c\x81\x48\x06\xe1\x91\x26\x14\x17\x97\x19\xd6\x95\x76\xf9\x97\x20\x43\xce\x05\x75\x6d\xe1\x57\xff\x58\xcc\x15\x53\x6b\xff\xc4\x7a\x6c\xc8\xaf\x44\xc0\xfd\x6a\xea\x7d\x82\xd2\xf7\x7e\x21\x7d\xdb\xbb\xbf\x58\x01\x19\x39\x6e\x01\xad\x73\xc5\x52\x4d\xa9\x8f\x67\x9d\x36\x8f\x1c\xab\xe7\x2d\xe9\x5c\xe8\xf2\x29\x6d\x59\xc4\x7a\xe0\x89\x3c\xd1\x63\x62\x32\x8d\x10\x7b\x6c\xf3\x8e\x34\xd5\x76\x5a\x0e\x46\xbb\x6b\x14\x6a\x9b\x19\x95\xea\x0e\xf3\x10\x39\xd9\x8f\x98\xe1\x96\x90\xa1\x16\x94\xdd\x5a\x62\x4a\x65\x15\x2f\x6c\x72\xbb\x9a\x68\x85\xa2\x0f\x91\xac\x12\xfe\x23\x5f\xb7\x6d\xfd\x54\xd1\xf1\xbb\x7a\x2a\x7a\x09\xee\xf8\x56\x3f\x0d\x95\x14\xe1\x20\xe6\x2c\xdd\x30\x84\xa2\xf7\x6d\xcf\xe5\x8e\xef\xe6\xfa\x0a\x02\xda\x7e\x9f\xb6\x14\x60\x17\xa9\xd8\x22\x2f\xbf\x09\x98\x97\x47\xfa\xdb\x44\xa6\x82\x5b\x50\xa2\x21\x36\xdb\xe5\xd3\xa1\x45\xe5\x14\x72\x1a\xfc\x12\x5c\x1a\xc5\x82\x07\x68\x57\x21\xa7\xaf\x35\x4f\xe6\x31\x7f\xba\x0c\x56\x9c\x78\x1a\xfe\x31\x46\xf2\xe9\xd8\x31\xb5\x8f\xc7\xa0\x2c\xad\xea\x07\x46\x5c\x61\x13\x29\x86\x59\x7f\xf0\x7d\xff\xa8\xa0\x55\xfd\xf1\xe9\xe7\xfe\xe0\x7c\xfc\xc3\xf8\xfc\xc7\xfb\xc3\x8c\x85\x73\x64\xed\xc8\x50\x80\x25\xa0\x09\x0a\xbc\x14\x05\xd0\x58\xc1\xcd\x15\x8d\x59\x9e\x06\xd1\xa8\xca\xde\xea\x50\xcd\x86\xe2\xe1\x40\xc6\x98\x87\x87\x7b\x9f\xf6\x0f\x0e\x7e\xda\x7f\x77\xd0\x4d\xf6\x88\x0a\x62\xdb\xd9\x5e\xe9\xd1\xca\x3c\xce\xf9\xf6\xc2\x6b\x5a\x48\x65\xfd\x31\x3e\x6b\x36\x46\x27\x17\xc7\xfd\xf3\xc9\xd9\x23\x43\x7a\xb3\x7c\x36\x9a\x5d\x1c\x9f\xcf\xb0\x29\xa7\x32\xca\x33\xac\xe5\xfe\xe9\xf0\x6c\x32\x1e\x7e\xfe\x38\x39\x19\xf9\x46\xca\x58\xfb\xbc\xb0\x04\xd9\x06\xbb\x0a\x1d\xb7\xec\x7d\xc0\x3b\x43\x97\x7e\xbc\xa0\x2e\x8b\x33\x64\x32\x9c\xc0\x37\x38\x4a\xf0\xa2\x85\xff\x18\xcc\xc5\x12\x76\x3e\xa2\x7b\xdb\x20\xd3\x78\x0d\x5c\x29\xa9\x70\x41\x71\xe4\xcd\x26\x57\x29\x0f\xcb\x8b\xcf\x42\x28\x9c\x60\x66\x25\x01\x3b\xbc\x4c\x43\x1c\x7b\x63\xe4\xfd\xc1\x25\xb0\x34\xc4\x8f\xd8\x6e\xb0\x33\x5f\xd2\xb5\x8d\xc1\x8a\xad\x41\xe6\x06\xc9\x45\xcc\x79\x06\x0f\x5b\xf3\x7e\xb7\xe7\xd4\x1a\x44\x1c\x61\xb0\x4f\xa1\x1c\x94\xbc\x62\xa4\xda\xb4\x7f\x3a\x1e\xc0\x4f\x34\x3e\x39\x09\x46\x9d\xdb\x78\x17\x32\xa5\x3e\x2b\x64\x9e\x4b\x28\xcd\x06\xe2\x90\x3f\x53\xaa\xe0\xa9\xcf\x83\xc9\xc5\xe9\x39\x4a\xb0\x97\x42\x1c\x08\xdb\xc2\x6f\x89\x4f\x60\x18\xd1\x97\xf6\x00\x3a\x2d\x70\xaa\x8c\xeb\x4a\x64\x2c\x15\x41\xdb\x8e\xfc\x15\xcd\x1f\xe6\x28\x05\x5d\xc0\xed\x80\x5a\x43\xc2\xb1\x50\x42\xbc\x60\xb4\xf1\xc6\x6a\x70\x65\x73\xd1\xbd\xa9\xe8\x82\x02\xbd\x25\x36\xd7\x4a\xbb\xdd\xcc\x02\xc4\xdb\x58\xb1\xb2\x4e\x8c\x25\x23\xc4\x0e\x9c\xb0\xf5\x9c\x5b\x89\xbb\x37\x9b\x6c\xbc\x2d\x77\xd3\x78\xda\xbd\x29\x53\xeb\x16\x2e\x51\x1d\x61\xe7\x63\x20\x95\xe2\x81\x71\xdd\xf1\xae\xcd\x2a\xbe\xa7\x61\x15\x15\x46\xc4\x72\x09\x9a\xad\xf5\x37\x9f\xd2\xfa\xc6\x07\xfc\x55\x6e\x71\x9d\x76\x1f\x76\xe8\x07\xd1\x8b\x29\x39\xa8\x72\x6d\x9e\x71\x8e\x63\x3f\x16\x97\xbc\xf0\x23\xa6\x59\x2a\x4b\x3f\xc6\xdc\xa0\x06\x81\x0d\x37\x9d\x76\xe9\x86\x07\x47\x67\x67\x28\xed\x77\x84\xcd\x1e\x70\x61\xbb\x9b\xfb\x15\x94\x07\x1c\xfe\x14\xfb\x0a\x03\xbb\x64\x60\xd7\xaa\x38\x22\x15\x1f\x98\x89\xd4\xd5\x89\x39\xcf\xd9\x1c\xa7\xd7\xee\x4d\xa5\x29\xdc\xda\xbe\x83\x23\xab\x1a\x1c\xcc\xe3\x22\xac\xd6\xf9\x65\x0c\x0b\xb1\x3b\xd0\x1f\x7e\x0b\x4b\x24\xc5\xda\x66\x3a\x48\x4b\xb0\xe9\xd4\x8a\xa1\x46\xa4\x09\x1d\x0b\x2d\xf7\x00\x23\x50\xba\xa0\x38\x63\xfa\xd9\xdc\x73\x9e\xfd\x7b\x8e\x95\x12\xf2\x8c\x63\x45\xe2\x69\xe4\x4d\x0e\xdc\x14\x05\x8f\xac\x29\x8e\x5d\xce\xc2\x0a\x39\x2f\xa6\x2b\xe6\x03\x4f\xef\x62\xc5\x68\xcc\x62\xdb\x4f\xdc\x7a\x55\x55\xa0\xdb\x81\xde\x18\x32\xe7\x01\xc3\x3a\x20\x4e\x2b\x48\xb6\x5c\xa5\x80\xf3\x88\xa2\x69\xbb\x30\xe9\xed\xa1\xde\x9e\xd3\xb9\x60\x79\xa7\x58\x47\xa6\xcc\xc0\xc2\x1a\x12\x7b\x85\xd8\x0c\xe5\xb6\x5d\x6d\x4b\xea\x35\xe4\x86\x2b\x47\xcd\xea\x3d\xe2\x4e\x25\xe4\x83\xd8\x98\x52\x81\x2d\x1f\x35\x92\x86\xbc\x6e\xb9\x22\xd6\x62\x1b\xa8\xf9\x84\x99\x75\x5e\xaf\xec\x5a\xb6\x5b\x19\xb5\x76\x3e\x61\x4b\x26\xd2\xf6\xc6\x3f\xc4\x39\x10\x8e\xc5\xb8\x43\x2c\x97\x4e\x76\x19\xf1\x7a\x6a\x57\x94\xa7\x96\x57\x4f\x80\x8d\x39\x95\x0c\xb0\x2f\xb5\xa8\xb1\xe2\x2d\xc1\xf2\xbe\x20\x96\xba\xa4\x47\xb3\xc1\xd9\x68\x74\xfa\xf9\x78\xd2\x1f\x8e\x4f\x8f\xb0\x04\xee\xe6\x18\x6e\xc7\x90\x64\xa0\xd7\xba\x33\x97\xd2\x7c\xc6\xa0\x67\x18\x6e\x6e\xaf\xc3\x74\xc7\x86\xbd\x4f\x6a\x8f\xea\xa0\xb1\x8a\x88\x4b\x52\x2d\xd4\x01\x1d\x17\xee\x5a\x32\xec\x58\x7e\xd1\x98\x5f\xd7\x68\xe9\x93\x7c\x2c\x9d\x8f\x57\xd1\xda\xbe\xae\x8b\xdc\x0f\xb4\x47\x4b\xf4\xc2\xdd\x04\xbc\x0d\x59\x92\x56\x2b\x40\xc7\x72\xf5\x97\x92\xb7\xfe\x49\x06\x97\x1c\x97\x0a\x28\xc6\x90\x45\x18\x75\x8e\x39\xb3\x76\x2a\x15\xc1\xa1\x5d\x9d\xcd\xd4\xb3\x45\x92\x6b\x47\x35\xf3\x34\x96\x68\xf3\x5d\x40\x3b\xf5\x22\x4f\x24\x6e\xae\x07\x97\xd0\x6a\x71\x35\x32\x0f\xa2\x6a\x4e\xc8\x34\x28\x03\x7b\x67\x96\x48\x69\x78\x61\x8f\x46\xfd\x90\xfe\xbc\xed\xd5\x5b\x52\x55\xc2\x2c\xcf\xb8\xfa\xa2\xea\xbb\xda\x90\x28\xac\x52\x79\x4a\xbc\xb8\x43\xc4\x7c\x4f\x91\x54\xec\xe9\x38\x8f\xd7\x80\xa7\x35\x96\xff\x7f\xfe\xf5\xcf\x7f\x57\xdb\xda\x7e\x41\xbf\xf1\x7f\x95\xda\x16\x2f\x9f\x87\xdf\x3b\xb2\xf5\x12\x5c\x6a\xbc\xa1\x17\x01\x65\x42\x75\x5d\xe6\x14\x4b\xdd\x77\x60\x5b\xf4\xfa\x9d\x75\x2d\xbc\xc2\xc0\x30\x87\x56\xbe\x69\xa3\xb7\x41\x9e\x97\x4a\x2f\x17\xf7\xe9\xee\x46\x96\x43\x79\xec\x34\x5d\xb0\x0b\x0c\xfc\x8d\xdc\xdf\xe0\x95\xb7\x55\x6c\xf5\x7a\x07\xad\x67\x02\xf4\x9e\x0b\xd0\x7d\x2e\xc0\xfe\x33\x01\xba\x5f\x3f\x17\xe0\xed\x73\x01\xbe\x7a\x2e\xc0\x97\x7f\x10\xc0\xd2\xde\x3f\x78\xb6\x7c\x13\xe9\x3d\x07\x84\x5f\xe3\x9d\xc8\x2b\xed\x48\x7a\xf4\xa2\x59\x0b\x2c\xd0\xf5\xb3\x00\x97\x52\x2e\x63\xfe\x67\xe0\x61\x03\xf5\x44\xb2\xf4\x98\x4a\x38\x9b\x0b\xef\xea\xab\x3b\x75\x29\xf9\xed\x5f\x9e\x2a\xef\xcb\xb7\xfe\x82\xf5\xab\x5f\xbe\x93\x6b\x36\x5c\x8b\xbb\x7b\xe2\x07\xf6\x44\xc7\xc8\x24\x6e\xde\xfe\x37\x00\x00\xff\xff\x39\xb3\xc7\xce\x06\x1b\x00\x00")

func scriptsMainBashBytes() ([]byte, error) {
	return bindataRead(
		_scriptsMainBash,
		"scripts/main.bash",
	)
}

func scriptsMainBash() (*asset, error) {
	bytes, err := scriptsMainBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/main.bash", size: 6918, mode: os.FileMode(420), modTime: time.Unix(1466007084, 0)}
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
	"scripts/main.bash": scriptsMainBash,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"main.bash": &bintree{scriptsMainBash, map[string]*bintree{}},
	}},
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

