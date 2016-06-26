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

var _scriptsMainBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x59\x7f\x72\xdb\x36\xf6\xff\x5b\x3a\xc5\xab\xec\x89\x93\xef\xd7\x14\x2d\x39\x69\x92\x26\xce\x8e\x22\x29\x89\xb6\xb6\xa5\xb1\xec\xee\x74\x92\x8e\x07\x22\x21\x11\x6b\x92\xe0\x00\xa0\x65\x4d\xd6\xbd\xc1\x9e\x61\xef\xb1\xa7\xda\x23\xec\x7b\x00\x29\x51\xb2\xbc\x49\xea\xb4\xd3\xd8\x26\x01\x7c\xf0\x7e\xbf\x0f\xc0\x13\x76\xc5\xdf\xc9\x38\xe4\x4a\x3f\x7e\x02\x9f\xeb\xb5\xe4\x2a\x14\x0a\xbc\x0c\x9a\x79\xa6\x23\x31\x35\x7e\x2c\x67\xda\xdf\x36\x30\xc9\x45\x1c\xfa\xf5\xdb\x7a\xfd\x22\x9b\x29\x16\xf2\x71\xa0\x44\x66\x1c\x4e\x90\xab\x18\xbc\xa9\x1e\x1f\x43\x64\x4c\xa6\x7f\xf2\x7d\xc5\xe6\xcd\x99\x30\x51\x3e\xc9\x35\x57\x81\x4c\x0d\x4f\x4d\x33\x90\x89\x1f\xf3\xa9\xb1\x98\x28\x86\x5f\xc2\x27\x4c\xe3\xb3\x2f\x52\x6d\x58\x1c\x37\x75\x04\x6f\xa0\x18\x6b\x1a\x9e\x64\xf0\xe8\x11\x04\x51\x22\x43\xf8\xff\x9b\x3b\x03\x4d\x7f\xf3\x8d\x4a\xd6\x26\x91\xd8\x63\x6e\xf2\xec\x3d\x4f\x9c\xc4\xef\xfb\x27\x47\xbb\xad\x7a\xad\x3b\x38\xda\x6d\xd7\x6b\x67\x52\x9a\x11\xd3\x7a\x2e\x55\x78\xb4\x7b\x58\xaf\xd7\xc4\x14\x3e\x42\x63\xb7\xdd\x80\xa3\x23\x68\x18\x95\xf3\x06\xfc\xf6\x0a\x4c\xc4\xd3\x7a\xad\xb6\x03\xdd\x01\x08\x0d\xf4\x7e\x1f\xe6\x1c\x52\x39\x87\x94\xf3\x10\xb2\x02\x05\x27\x15\x10\x87\x0d\xf8\x01\x21\xaa\xcb\x6b\x3c\x88\x24\x78\xdc\x8d\xfe\x03\x74\x1e\xe2\xe3\x18\x66\x3c\x81\xc2\x04\x40\xd2\xe1\xc4\x1b\x61\xe0\x00\xff\xe2\xb1\xe6\xcb\x17\x34\x34\x15\xf5\xf2\xe5\x0e\x5c\xa0\x8d\x41\x47\x32\x8f\x43\x30\x8b\x8c\x23\x0a\xed\x55\x95\xe6\xde\x3d\xca\x2d\x10\x10\xcd\xf4\x5e\x98\x51\x1e\xc7\xce\x4a\x67\xfd\x93\xe1\x79\xdf\x1a\xea\xed\x59\xe7\xb4\xfb\xc1\x1a\xeb\x78\xf8\xfe\xdd\xe0\xb8\x4f\x76\xaa\x55\x42\xaa\x5e\x43\x7f\x43\x96\x5b\x5c\xd8\x6d\x43\xfb\xcd\xa3\x16\x2a\x67\x38\x07\x9c\xea\xb0\xc7\xf9\x04\x9d\x98\xc7\x7c\x90\x8a\x22\x78\x96\x70\xad\x2d\x70\xba\x9c\x8f\x32\xe3\x63\x15\xb2\xb5\x09\x79\x91\x85\xcc\xf0\x6f\x02\xcd\xed\x92\xbb\xb0\x23\x19\x0e\x9c\x8d\xbe\x00\x97\x61\x44\x96\xd6\x5c\x43\xa9\x15\x30\x67\x3c\x93\xf7\x09\x76\xe1\x42\xb4\x2b\xd3\xa9\x98\x39\x2c\x85\xd3\x4b\xa9\x3c\xef\x9a\xab\x89\xd4\x7c\x0b\xb2\x8b\x66\xcc\xc3\x98\xff\xcd\x01\xaf\x6b\x69\x47\x60\xae\x58\x96\x71\x65\x17\x18\xa6\xcc\x58\x24\x79\xcc\x8c\x54\x6e\x49\xaf\xff\xcb\xa0\xbb\x45\xa9\x9b\x40\xe5\xa9\x55\x4b\xe5\x09\xa6\xad\x06\x6f\x8e\xa1\xda\x6a\x40\xeb\x8d\x1f\xf2\x6b\x3f\xcd\x0b\x75\x09\xb9\x2b\x93\x4c\xa0\x43\x87\x63\x87\x3a\x3a\x1b\xfe\xb5\xdf\x3d\xbf\x3c\xff\x75\xe4\xb0\xcb\x17\xa3\xce\xb9\x0b\xa0\x71\xf7\x43\xff\xc4\xc5\x4f\x29\xc2\x53\x6b\x99\x62\xca\xb3\x0d\x79\x34\x37\xe0\x49\xc8\x44\xc6\xa7\x4c\xc4\x94\xe0\x37\x81\x0c\xb9\xad\x49\xe0\x59\xc9\x6c\xaa\x7a\x3a\x88\x78\x52\xa4\x95\x17\x09\xac\x52\x11\x8f\x63\x57\xaa\xfa\xe9\xb5\x50\x32\x25\x85\x70\x62\x78\x05\x22\x8b\x64\xca\x75\x69\x13\xf0\x42\xae\x8d\x48\x99\x11\x32\x85\x46\x86\x2f\xa7\x52\x25\x47\x1b\xd3\xf6\x53\x96\x70\x14\xb8\x41\xf3\x95\xb8\xe6\x61\x8f\x19\x36\x62\x26\xda\x28\x96\x85\xc3\x1a\xbb\xcf\x28\xc5\x6f\x82\x4c\x71\x63\x16\x64\xb1\x8e\x0a\x22\x5c\xf8\xc7\x2c\x56\x8e\x9d\x76\x4e\xfe\x24\xbb\x7d\x41\x2d\x8f\x39\xf1\xd7\xc7\x76\x9f\x36\x6f\x82\x62\x04\xca\xdf\xf7\x5a\xa0\x7f\x93\x49\x65\xee\x18\xc0\xe9\xd4\xaa\xea\xd4\xfe\x36\x9d\xb8\x05\x2e\x0c\x5c\x3e\x0e\x33\x72\xa9\x1e\xc5\x42\x1b\x68\x66\xa8\x1d\xa6\x97\xef\xc6\x9a\x99\x7d\x7b\x8f\x52\xad\x8a\x52\x05\xd8\xfa\x8c\x6a\x6a\x5a\x53\xae\xeb\xf9\x8e\x9b\x20\xea\xa4\x54\x05\x98\x50\x23\x25\xaf\x85\x46\x51\x44\x3a\xc3\xbf\xa7\x98\x35\x45\x23\xde\xc1\x6a\x8d\xed\x04\x5b\xa5\x06\x2d\x66\x11\xbd\x09\xe5\x3c\x8d\x25\x0b\x81\x6a\x8b\x48\x8d\x5c\x8a\x0e\x53\x6b\x0d\x9a\xc4\xd2\xd0\x36\x15\xa0\x8c\xa5\x8a\x3f\xcd\xd3\xc0\x06\xf0\x48\x66\x14\xb0\x7c\xdb\xa6\xf5\x5a\xa7\xdb\x1d\x5e\x9c\x9e\x5f\xf6\x4f\x3a\x83\x63\x6b\x72\xda\x96\xca\x0f\xa3\xc6\x9f\xc3\xf2\x55\x29\xc6\x25\x89\x51\x0c\xac\xe8\x81\x5f\xca\x84\x2f\xaf\xe1\xff\x9a\x89\x9c\xe0\x06\x59\xb9\x27\xac\x86\xff\xa7\x3c\x64\x2a\x91\x86\x5b\xc6\x96\x95\xaa\x7f\x3c\x1c\xf5\xcf\x2e\x0b\xc1\x5d\x4f\xba\x38\xed\x1d\xf7\x2f\x07\xbd\xfe\xe9\xf9\xe0\xdd\xa0\x7f\x66\xc3\xc5\x4a\x6d\x05\x45\xaf\x52\x1f\xf2\x3c\x16\x46\x32\x70\x05\xf9\x7e\x21\x4a\x4f\xbc\xc7\x08\x23\x4b\x5e\x5c\x0c\x7a\x30\x55\x32\xa9\x28\xb1\xb3\xe4\x37\x33\x8c\x9a\x82\xe0\x58\x56\x93\xdc\x64\xca\x7f\xd1\x3e\x78\xd1\x7e\xf1\x92\xfd\x78\x18\xb0\xa7\xfc\x90\x1d\x3e\x63\x4f\x69\xd1\xb1\x94\x19\x62\x2a\x99\xa3\x68\x64\x47\xbb\xe1\x3e\x20\x43\x58\xc8\x1c\x5b\xb2\x01\x06\x17\xbd\x41\x6f\x1f\x58\x68\x1d\x9a\x00\xfa\x9b\xa4\xb0\xd1\x29\xa7\x90\x2d\x5d\xb7\x96\x0e\x96\x63\x78\x21\x34\x56\x9e\xf0\xd7\x89\xca\x39\x82\x6c\x04\x0e\x60\xcb\xd7\x46\x13\x89\x90\xb9\xb5\xba\x85\x3e\x9a\x32\xc7\x27\xb0\xec\x95\xfb\x9d\x62\xb1\x23\x32\xb1\x04\xbf\xe3\xe2\x57\x18\x21\x44\x4c\xf2\x5c\x20\x75\x7a\xec\xe7\x5a\xf9\xb1\x98\xf0\x1b\x1e\xf8\x36\xe3\xde\xe6\x61\xb8\x00\x2f\x80\xbd\x91\xc2\x30\xb6\x76\xdd\x03\xdb\x47\xb4\x09\x11\xfb\xf5\xeb\xd7\xb0\xfb\x58\x73\x64\x91\xc2\x2c\x20\x48\xb0\xdf\xf4\xc0\x13\x98\x50\x9f\x2b\x62\xdc\x36\x30\xd9\x96\xed\xe7\xc9\x93\x3a\xed\xba\x03\x83\x29\xd9\x8e\x7c\x65\xb5\xda\x77\xe9\x10\xc8\x6c\x01\xc2\x58\x13\x63\xc7\x8f\x98\x4e\xf7\xd0\xc8\xb1\xe2\x0c\x85\x99\x70\x37\x45\x70\x62\x46\x25\x51\xfb\x4c\x2a\xdc\x6e\x61\x6b\x85\x95\xa7\xf0\xbb\x7f\x2c\x26\x8a\xa9\x85\x7f\x62\x8d\xd0\xe3\xd7\x22\xe0\x7e\x35\x9a\x3e\x41\x69\x4e\xbf\x00\xbc\x93\x13\x55\x64\x54\x60\x8c\x83\xa9\x89\x51\xdc\x59\x2a\x15\x77\xaf\x33\x32\xd5\x74\xd3\x02\x85\xe3\x4a\x76\x57\x08\xf3\x29\x6d\xd8\x45\x25\x45\xb4\xf4\x94\xf4\xa7\x75\x14\x47\x2e\x92\x8b\x9a\x41\xff\x05\x19\x69\x73\xc7\xbc\xdf\x45\xbd\x75\xf9\x4f\xf0\x6d\x3a\x83\x0d\x35\x3e\x16\x6b\x7f\x23\xe9\x36\xb4\x20\x5e\x5b\xdb\x8c\x4c\xe2\xd9\xf5\x72\x30\xc4\x76\x5c\x5f\xf1\xeb\xb5\x99\x96\xae\xd3\xec\x35\xb6\xbd\x49\xa2\x4b\xe9\x7a\x79\x88\xf4\xfd\x57\x4c\x41\xcb\xdd\x51\x18\x4a\x3f\x2d\x31\xe6\xb3\x8a\xd2\xcb\xe4\xab\x66\x42\x21\xef\x36\x3e\x5e\xc2\x7f\xe0\x8b\x7d\x9b\xe0\x55\x74\x7c\x56\x5f\x8b\x5e\x82\x3b\x6a\x8e\x7d\x44\x49\x11\x76\x63\xce\xd2\x25\x99\x2c\xda\xe4\x26\x85\x6b\xfa\x8e\x02\xce\x21\xa0\xe9\x77\x19\x6e\x01\x76\x91\x8a\x0d\x9e\xfb\x45\xc0\xbc\x5c\xd2\xd9\xe4\xbc\x15\xdc\x82\x3d\xf7\xb0\x2f\xcf\xbe\x1e\x5a\x54\x56\x61\xd9\xc6\x87\xe0\xca\x28\x16\x6c\x61\xe8\xc5\x3e\x1d\xad\x79\x32\x89\xf9\xd7\xef\xc1\x8a\x15\x5f\x87\x7f\x8c\x9e\xfc\x7a\xec\x98\xea\xdb\x7d\x50\x96\x81\x77\xb0\x39\x5f\x63\x95\x2b\x78\x4f\xa7\xfb\x73\xe7\x7d\xc1\xc0\x3b\x83\x53\xec\x6e\xe7\x83\x5f\x06\xe7\xbf\xde\xe5\x3d\x2c\x9c\xe0\x01\x0f\xc9\x2c\xb0\x04\x34\x41\x81\x97\xe2\x06\xc4\x51\x70\x72\x45\x62\x86\x0c\x20\xea\x57\x89\xfe\x3a\x54\xbd\xa6\x78\xd8\x95\x31\xc6\xe1\xd1\xde\xa7\x83\xc3\xc3\x8f\x07\xaf\x0e\x5b\xc9\x1e\x9d\x1a\xb0\x2e\x6e\x8e\xb4\x69\x64\x12\xe7\x7c\x73\xe0\x29\x0d\xa4\x72\xfd\x35\xbe\xab\xd7\xfa\x27\x17\xc7\x9d\xf3\xe1\xd9\x3d\x7c\x6e\x39\x7c\xd6\x1f\x5f\x1c\x9f\x8f\xb1\x6b\xa4\x32\xca\x33\xcc\xe5\xce\x69\xef\x6c\x38\xe8\x5d\x7e\x18\x9e\xf4\x7d\x23\x65\xac\x7d\x5e\x68\x82\xc4\x94\x5d\x87\xee\x18\xd2\x7e\x83\x6d\xbd\x45\x3f\x1e\x51\x1b\xc0\x26\x37\xec\x0d\xe1\x27\xec\x75\x48\xa2\xf0\x7f\x06\x13\x31\x83\x9d\x0f\x68\xde\x7d\x90\x29\xd6\x56\xae\x94\x54\x38\xa0\x38\x72\x1c\x93\xab\x94\x87\x65\x15\x9d\x0a\x85\x2d\xd6\xcc\x25\x60\x0b\x92\x69\x88\x7d\x79\x80\x47\xc4\xe0\xca\x92\xab\x01\xb6\x37\x6a\x1d\x57\x74\xc2\x67\x30\x67\x0b\x90\xb9\x41\x96\x11\x73\x9e\xc1\x76\x6d\x5e\xef\xb6\x9d\x58\xdd\x88\x23\x0c\xd6\x29\xdc\x07\x77\x9e\x33\x12\x6d\xd4\x39\x1d\x74\xe1\x23\xf5\x77\x4e\x1b\xa3\xcc\xfb\x78\x6c\x36\xa5\x3c\x73\x3c\xa4\xcc\xa0\x54\x1b\xe8\xb8\xf1\x1b\x85\x0a\xae\xba\x2c\xb8\xcf\x63\x7b\x7f\x80\xf5\x7b\x73\xf3\x5b\x22\xa0\xe8\x46\xb4\xa5\x5d\x80\x46\x0b\x9c\x28\x83\x75\x21\x32\x96\x8a\x60\xdf\x72\x92\x39\x35\x48\xe6\x38\x0f\xdd\xd5\xd8\x0e\xba\x80\x84\x63\xa2\x84\x78\x16\xdd\x87\x54\x1a\x1c\x59\xde\x89\x7c\xae\xc8\x82\x1b\x7a\x33\x2c\xae\x95\x72\x5b\x56\x3f\x22\x1e\x4b\x2d\xe6\xd6\x88\x96\xce\xe6\x59\x13\x4e\xd8\x62\xc2\xed\x8e\xbb\x9f\x97\xd1\x78\x5b\xce\xa6\x26\xb1\xfb\xb9\x0c\xad\x5b\xb8\x42\x71\x84\x6d\xe0\x81\x54\x8a\x07\xc6\x55\xc7\x55\x99\x55\x7c\x4f\xc3\x3c\x2a\x94\x88\xe5\x0c\x34\x5b\xe8\x9f\x3e\xa5\xeb\x13\xb7\xd8\xab\x9c\xe2\x2a\xed\x01\xec\xd0\x0f\xe2\x3f\x23\x32\x50\xe5\x86\x65\xcc\x39\xf2\x92\x58\x5c\xf1\xc2\x8e\x18\x66\xa9\x2c\xed\x18\x73\x83\x12\x04\xd6\xdd\xb4\xda\x85\x1b\x2e\xec\x9f\x9d\xe1\x6e\xdf\xe0\x36\xbb\xc0\xb9\x6d\x45\x4c\x2a\x28\x5b\x0c\xfe\x35\xfa\x15\x0a\xb6\x48\xc1\x96\x15\xb1\x4f\x22\x6e\xe9\x89\x54\xd5\xe9\x90\x35\x61\x13\xec\x5e\xbb\x9f\x2b\x45\xe1\xd6\xd6\x1d\x6c\x59\x55\xe7\x60\x1c\x17\x6e\xb5\xc6\x2f\x7d\x58\x6c\xbb\x03\x9d\xde\x5b\x98\xe1\xf9\x49\xbb\x23\x8e\xb4\x67\x31\x5a\x35\x67\x28\x11\x49\x42\xcb\x42\x4b\x35\xc0\x08\x3a\xeb\x90\x9f\x31\xfc\x6c\xec\x39\xcb\xfe\x3d\xc7\x4c\x09\x79\xc6\x31\x23\x71\x35\x12\x3b\x07\x6e\x8a\x84\x47\x5a\x17\xc7\x2e\x66\x61\x8e\xa4\x1c\xc3\x15\xe3\x81\xa7\x2b\x5f\x31\x6a\xb3\x58\xf6\x13\x37\x5e\x15\x15\xe8\x20\xa9\x97\x8a\x4c\x78\xc0\x30\x0f\x88\x74\x0b\xda\x1b\xcf\x3f\x80\xfd\x88\xbc\x69\xab\x30\xc9\xed\xa1\xdc\x9e\x93\xb9\xa0\xa1\xa7\x98\x47\xa6\x8c\xc0\x42\x1b\xda\xf6\x1a\xb1\x19\xee\xbb\xef\x72\x5b\x52\xad\x21\x33\x5c\x3b\x2e\xb6\x5e\x23\x56\x22\x21\x61\xc5\xc2\x94\x0a\x2c\xf9\x28\x91\x34\x64\x75\x4b\x66\x31\x17\xf7\x81\x8a\x4f\x98\x59\xe3\xb5\xcb\xaa\xe5\x8e\x82\x6a\xe1\x6c\xc2\x66\x4c\xa4\xfb\x4b\xfb\x10\xe7\x40\x38\x16\xe3\x0c\x31\x9b\xb9\xbd\x4b\x8f\xaf\x87\x76\x45\x78\x2a\x79\xeb\x01\xb0\x54\xa7\x12\x01\xf6\xfe\x93\x0a\x2b\x1e\x63\x48\x4c\x64\x1c\x52\x97\xf4\x68\xdc\x3d\xeb\xf7\x4f\x2f\x8f\x87\x9d\xde\xe0\xf4\x3d\xa6\xc0\xaa\x8f\xe1\x74\x74\x49\x06\x7a\xa1\x9b\x13\x29\xcd\x25\x3a\x3d\x43\x77\x73\x7b\x73\x42\xd7\x31\xb0\xf7\x49\xed\x51\x1e\xd4\xe6\x11\x51\x59\xca\x85\x75\x40\x47\xd6\x5b\x96\xad\xbb\x63\x48\x51\x98\x9f\xd6\xab\x2c\xf4\xab\x6c\x2c\x9d\x8d\xe7\xd1\xc2\xde\xec\x46\xee\x07\xea\xa3\x25\x5a\x61\xd5\x01\x6f\x43\x96\xa4\xd5\x0c\xd0\xb1\x9c\xff\xa5\xa4\xaf\xdf\x49\xe1\x92\xe3\x52\x02\xc5\xe8\xb2\x08\xbd\xce\x31\x66\x16\x4e\xa4\xc2\x39\x34\xab\xb9\xec\x7a\x36\x49\x72\xed\xa8\x66\x8e\x87\x76\xd4\x79\xe5\xd0\xe6\x7a\x92\x27\x12\x27\xaf\x3b\x97\xd0\xd6\xfc\x6a\x64\x1e\x44\xd5\x98\x90\x69\x50\x3a\x76\xa5\x96\x48\xa9\x79\x61\x8d\x46\xf9\x90\xfe\xbc\x68\xaf\x97\xa4\xea\x0e\xe3\x3c\xe3\xea\x87\xaa\xed\xd6\x9a\x44\xa1\x95\xca\x53\xe2\xc5\x4d\x22\xe6\x7b\x8a\x76\xc5\x9a\x8e\xfd\x78\x01\xb8\x5a\x63\xfa\xff\xe7\x5f\xff\xfc\x77\xb5\xac\x1d\x14\xf4\x1b\xff\x55\xa9\x6d\xf1\x9d\xa2\xf7\x73\x79\xb2\xef\xda\x13\x9f\x3b\xd3\xdb\xba\xe0\xd9\xe3\x3c\x9e\xe6\x2d\x13\xa4\x9c\x9c\xa2\x2f\xed\x79\x9e\xf9\x87\xad\x97\x07\x07\x4f\xdb\xcf\xfd\x67\xcf\x5f\xb6\xec\x31\x1e\x69\x81\x2d\x50\x15\x46\x10\x59\x2f\x94\xc5\x0d\xd3\xfb\xaa\x5e\x7b\x0c\x2e\x06\x9f\xd1\xe5\x54\x19\xb9\x2d\x17\xa2\xc5\x50\xeb\x15\xd8\x5e\xb0\x78\x65\x7d\x08\x4f\x30\x02\x98\x13\xbb\xbc\xfd\xa5\x1b\x4a\xcf\x4b\xa5\x97\x8b\xbb\xbc\x7a\xa9\x94\x43\xb9\x6f\xb5\xbd\xb2\x71\x18\xf8\x1b\x0f\x19\x06\x0f\xff\x8d\x62\xaa\xd7\x3e\x6c\x3c\x10\xa0\xfd\x50\x80\xd6\x43\x01\x0e\x1e\x08\xd0\x7a\xf9\x50\x80\x17\x0f\x05\x78\xfe\x50\x80\x1f\xff\x20\x80\xe5\xd7\x7f\x70\x6d\x79\x3b\xee\x3d\x04\x84\xdf\xe0\xe1\xcb\x2b\xf5\x48\xda\xf4\xf1\x43\x0b\xac\x04\x8b\x07\x01\xce\xa4\x9c\xc5\xfc\x7b\xe0\x61\xa5\xf6\x44\x32\xf3\x98\x4a\x38\x9b\x08\xef\xfa\xf9\x4a\x5c\x0a\x7e\xfb\x35\xb4\xf2\x0d\x67\xe3\xab\xea\xef\x7e\x79\x4f\x5c\xaf\xb9\x5a\xba\x7a\xe3\x07\xee\xab\x0f\x42\x14\x39\xdd\xe5\xc8\xad\xa6\x22\x40\xa1\x8a\xbb\xc8\xb7\x9d\x71\x7f\x75\xe6\x5c\x5e\x90\x89\x84\xee\xa1\xe9\xec\xc7\x32\xec\x1f\xcd\x00\x45\xf5\xae\x2a\xb7\x37\x3f\xf3\x45\x10\x61\xff\xd7\xf4\x39\x57\xa4\xcd\xab\xe2\x19\xbc\x73\xb0\xf7\x73\x13\x91\xfa\x74\x69\xae\xc5\x6c\xfd\x65\xb9\xc7\xd6\xdd\x42\xa1\x91\x3d\x4c\x72\x22\x70\xcd\xac\xd5\xfe\xde\x9b\x82\x37\x82\x46\xe3\xcb\x5b\xff\x19\xfa\xd2\x77\x2c\x24\x91\x86\x77\xb2\x6c\x98\x0e\x4c\x9e\x96\x4e\xf8\x96\xab\xe7\xf5\xef\x18\x87\x75\x6c\x83\x32\xcc\xb1\x71\x6e\x5c\x47\x67\xd9\x25\x1d\xeb\xec\xf7\x96\xfa\xed\x7f\x03\x00\x00\xff\xff\xad\x56\xbd\xb7\x97\x1f\x00\x00")

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

	info := bindataFileInfo{name: "scripts/main.bash", size: 8087, mode: os.FileMode(420), modTime: time.Unix(1466943555, 0)}
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

