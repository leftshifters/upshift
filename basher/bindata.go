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

var _scriptsBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3a\xeb\x72\xda\x48\xd6\xbf\xcd\x53\xf4\x60\x57\x1c\x67\x2c\x30\x38\x99\x5c\x3d\x5f\x11\x20\x0e\xdf\xd8\x40\x01\xce\xd6\x54\x92\xa2\x1a\xa9\x81\x5e\x4b\x6a\x6d\x77\xcb\x98\xcd\x78\xdf\x60\x9f\x61\xdf\x63\x9f\x6a\x1f\x61\xcf\xe9\x56\x83\x04\xe4\x32\xe3\xcc\x8f\x75\xcd\xd8\xa0\x3e\x7d\xee\x77\x65\xc4\x94\x1e\xfa\x92\x27\xfa\xe1\x11\xf9\x54\xda\x63\xfe\x5c\x90\x83\x5a\x69\x4f\x85\x8c\x25\x04\x3e\x84\x8a\x78\x21\x2d\xdd\x95\x4a\x6b\xd8\xba\x05\xbe\x6c\x0f\x87\x8d\xf3\xf6\x19\xc2\x5f\xf4\xce\xdf\x74\x2e\xda\xc3\xb3\x83\xfa\x16\x96\xba\xc3\x42\xea\x3f\x3f\xa8\x91\xdf\x88\x66\x8c\x00\x1c\xe0\x3c\x97\x34\x08\xd9\x5f\x24\x4d\x12\x26\x2d\xd6\x99\x79\x44\x16\xf6\xd9\x1a\x68\x31\xa2\xea\xda\x82\xf4\x1b\x83\xc6\xe5\xd0\xd1\x1d\xf7\x1b\xa3\xb7\x86\xee\x25\xbd\x66\x6f\x44\x18\x30\xa9\x4a\x7b\x95\xaa\xc5\xb4\x20\xe5\x83\x5a\x79\x9b\xf4\x90\xe9\x34\x39\x67\x91\x45\x79\xde\xbe\x34\xf8\x9a\x1d\x83\x69\x20\x84\xee\x53\xa5\x16\x42\x06\x67\x07\xa7\xa5\xd2\x1e\x9f\x92\xf7\x80\xa9\x5e\x26\x67\x67\xa4\xac\x65\xca\xca\xe4\xe3\x4b\xa2\xe7\x2c\x2e\xed\xed\xed\x93\x66\x87\x70\x45\xf0\xf9\x31\x59\x30\x12\x8b\x05\x89\x19\x0b\x48\x92\x61\x01\xa0\x0c\xc5\x69\x99\xfc\x00\x28\xf2\xd7\xad\xc6\x3c\x66\x4f\x7f\x23\x2a\x0d\xe0\xeb\x35\xf1\x86\x64\xc6\x22\xc2\x63\xa5\x69\x18\x1a\x95\x02\xec\x2d\xd7\xe4\x04\x3e\xb1\x50\xb1\xd5\x03\x3c\x9a\xf2\x92\x7b\xb8\x4f\xae\x14\x93\x44\xcd\x45\x1a\x06\x44\x2f\x13\x06\x58\x90\x5c\x9e\x21\x47\x66\x8b\x86\x23\x01\x08\x41\x53\x57\x71\x76\xf8\xbf\xa1\xad\xd4\xb1\x4b\x3c\xda\xb9\xfd\x53\x95\xb6\x8b\x54\x51\x77\x7d\x11\x74\x2c\x88\xd5\x5c\x16\x27\x46\x7b\x05\x77\x4d\x44\xb0\xb2\x41\xde\x57\x6b\x19\x96\x01\x4b\xc4\x55\x12\x50\xcd\xb6\x11\xe1\x5d\x09\xe7\x24\x35\x00\xc4\xf3\x6e\x98\x9c\x08\xc5\xb6\x31\x9d\x73\x3d\x4c\x27\x91\x08\xd2\x90\x75\x62\xae\xb7\x91\xcd\x80\x7d\xe5\x40\x80\x25\xf8\xba\x0b\x4b\x3f\x75\x22\x0d\xda\x97\xbd\x91\xbd\xfb\x7a\xd0\xe8\x36\x6d\x2c\xae\x50\x9e\x5a\x94\x49\x6a\x9c\x0b\x82\xaf\x80\xee\x14\xd1\x65\x3f\x57\x09\x46\x2c\xcb\xe7\x23\x3f\x95\xa0\xdb\xa9\x1a\x5e\x90\xb9\xd6\x89\x7a\x51\xad\x4a\xba\xa8\x00\xc2\x79\x3a\x49\xc1\x58\xbe\x88\x35\x8b\x75\xc5\x17\x51\x35\x64\x53\xad\xe6\x7c\xaa\x41\xa1\xd5\x34\x31\x1f\xab\x11\x55\xf0\xbd\x9a\xa9\xb6\xa2\xe6\xe4\x67\x92\x9d\x55\x34\x8b\x12\xf2\xe0\x01\xf1\xe7\x20\x2d\xf9\xf1\x76\xeb\xa0\x52\xdd\x7c\x22\xa3\x02\xd0\x2a\x91\xbc\x96\x90\xa0\x84\xc8\x54\x32\xea\xf5\x2e\x8c\x42\x26\xf0\xb8\x10\x5a\x99\x77\x6c\x58\xe2\x73\x86\x2d\xda\x22\xb3\xee\x96\x35\x86\x9a\x4a\x3d\xe4\x51\x1a\x52\x2d\xb2\x2c\xda\x6a\xbf\xeb\x34\x2d\x92\x5b\x5f\xa6\xb1\x61\x42\xa6\x11\xe8\x0a\x52\x71\x96\x13\x6b\x3f\x57\x03\x76\x53\x8d\xd3\xcc\xe7\x10\x59\x53\x44\x09\x07\xdf\xe8\x0d\xb3\x5c\x3b\xe8\xfd\x7f\xbb\x39\x1a\x8f\x7e\xed\x5b\x74\xee\xc1\x2a\xeb\x0e\x9b\x6f\xdb\x97\xd6\xd0\x8e\xea\xe3\x7c\x62\x7e\x02\x75\x80\x69\xe2\x09\x92\xf0\x84\x4d\x29\x0f\x51\x91\xb7\xbe\x08\xd8\x24\xe5\x10\x68\x9e\x61\xc6\x64\x0b\x4f\xf9\x73\x16\x65\x91\xed\xcd\x39\x78\xc3\x9c\x85\xa1\x75\x89\x76\x7c\xc3\xa5\x88\x51\x06\x00\x0c\xae\x09\x4f\xe6\x22\x66\xca\x49\x4e\xbc\x00\x4a\x14\x8f\xa9\xe6\x22\x26\xe5\x04\x1e\x4e\x85\x8c\xce\x36\xc0\x8e\x63\x1a\x31\xe0\xb1\x8c\xf0\x92\xdf\xb0\xa0\x45\x35\xed\x53\x3d\x27\x15\xe7\x35\x96\x31\xab\xe4\xf2\xc1\x13\xcc\x32\xb7\x7e\x22\x99\xd6\x4b\x57\x0a\xbf\xaf\x86\x6e\x7d\x0d\xde\xf3\x79\x55\x7c\x0f\x71\x35\x80\xe7\x64\x42\x41\x1a\xd2\x9f\x83\x06\xfe\x98\x2c\xee\xac\xdb\xb8\xfc\x7e\x36\xff\x8a\x49\x3c\x6a\x39\x2e\x9e\x1d\x3c\xae\xdc\xfa\xd9\x09\x71\x7f\x3f\x6b\xbd\xf6\x6d\x22\xe4\xb6\xfd\xac\x18\x9b\x3d\xc5\x57\xc4\x60\x06\x57\xa6\x46\xf7\xb5\x97\xa0\x49\x54\x3f\xe4\xa0\xf0\x4a\x02\x02\x41\xdc\x56\xed\x59\x25\x31\x4f\x3f\x23\x47\x2d\x27\x47\x86\xac\x08\x91\x8f\x7e\xa3\xbd\xa2\x68\x6f\x98\xf6\xe7\x8d\x18\xab\x05\xe5\xb2\x2f\xc5\x0d\x57\xc0\x0a\x8f\x67\xf0\x79\x0a\x71\xad\xac\xcc\xfb\x50\xd5\xa0\x00\x43\x06\x55\x44\xf1\xd9\x1c\x9f\x04\x62\x11\x87\x82\x06\x04\x73\x15\x8f\xb5\x58\xb1\x4e\xa6\xa6\x4a\x21\x10\x8d\x03\x53\x86\x09\xe6\x14\xac\x8c\xd3\x34\xf6\x8d\x03\xf6\x45\x82\x0e\xc7\x76\x11\x2d\xed\x35\x9a\xcd\xde\x55\x77\x34\x6e\x5f\x36\x3a\x36\x35\x22\x59\xac\x5a\xc0\x27\xf1\x52\xb2\x7a\xe4\xd8\x18\x9b\xb2\x6a\x0f\xa2\x1b\xf2\xa8\x12\x89\x09\xe0\x4a\x1c\xfa\x15\x77\x60\xc1\x2f\x91\x46\xad\xf0\x38\xd8\x71\xb6\xca\x94\xed\x8b\x5e\xbf\x3d\x18\x67\x3c\xda\x4a\x76\xd5\x6d\x5d\xb4\xc7\x9d\x56\xbb\x3b\xea\xbc\xe9\xb4\x07\xd6\x19\x90\x41\xc3\x13\x18\x30\xeb\x23\xbf\x44\xdc\x29\xfb\x1c\x9c\x08\x95\x75\x75\xd5\x69\x91\xa9\x14\x51\x8e\xf9\xfd\x55\x65\x9b\x81\x63\x64\xa5\xcd\xd4\xb3\xe8\x36\x91\xd5\x67\xf5\x93\x67\xf5\x67\xcf\xe9\x4f\xa7\x3e\x7d\xcc\x4e\xe9\xe9\x13\xfa\x18\x2f\x5d\x08\x91\x00\x4e\x29\x52\x60\x09\x55\x65\x08\x1e\x13\x68\x9b\x96\x22\x85\xee\x44\x13\x4a\xae\x5a\x9d\xd6\x31\xa1\x81\xb1\x59\x44\xc0\xa4\xc8\x85\x71\x40\x31\x25\x89\x53\x51\xd6\xaf\x79\x01\x29\x57\xaa\x2b\x7f\x2d\x76\x6c\x23\xb8\xb8\xe1\x0f\x04\x8a\x99\xd2\x0a\x7b\x28\x91\x1a\x0d\x1b\x74\x67\x53\x6a\xdb\x29\xc8\x46\x8e\x46\x17\x72\x10\xf6\x52\x2b\xe4\x5b\xe6\x7c\x09\x86\xc7\xbe\x2c\x4d\x39\xf4\x90\x0f\xab\xa9\x92\xd5\x90\x4f\xd8\x2d\xf3\xab\x26\x90\x5e\xa7\x41\xb0\x24\x9e\x4f\x0e\xfb\x12\xbc\xd3\xe8\xf2\x90\x98\x02\xa6\x74\x00\xb8\x5f\xbd\x7a\x45\x0e\x1e\x2a\x06\x3d\x03\xd7\x4b\xe2\x47\x50\xe8\x5a\xc4\xe3\x10\x27\x9f\x72\x6c\xdc\xe1\x2c\xb0\xaa\x7b\x47\x47\x25\xa4\xba\x4f\x3a\x53\xd4\x17\xda\xc7\x48\x75\x6c\xbd\xdc\x17\xc9\x92\x70\x6d\xd4\x0a\xb5\x78\x4e\x55\x7c\x08\x8a\x0d\x25\xa3\xc0\xcc\x84\x59\x10\xce\xb0\x31\x74\x1d\xeb\x27\x14\xe1\x6e\x47\xdb\x9a\x69\x79\x4a\xfe\x51\xbd\xe0\x13\x49\xe5\xb2\x7a\x69\x94\xd0\x62\x37\xdc\x67\xd5\xbc\x07\x7d\x20\x4e\x9d\xd5\x0c\xe1\x96\xff\xe7\x31\x83\x00\x43\x38\x8c\x75\x08\xec\xce\x62\x21\x99\x7d\x9c\xa0\xaa\xa6\x9b\x1a\xc8\x0c\xe7\x9a\xdb\x8c\x99\x0f\x71\xd9\x5c\x72\x1d\xb2\xe9\xd3\x51\x7e\xbc\x87\xbe\x63\xbd\x37\x4b\x05\xf8\xe3\x27\x28\xcd\x96\x7a\xbf\x8b\x78\x45\xfe\x2f\xe1\x69\x3c\x23\x1b\x62\xbc\xcf\xee\x7e\x44\xee\x36\xa4\xc0\xb6\x7e\x6f\xd3\x33\x71\xe0\x28\xb9\xc3\x00\xaa\x64\x69\x3d\x68\x14\x20\xcd\xdc\x82\xd0\x85\xb1\x63\x73\x86\x70\xdc\xb5\xd2\x00\xe6\x98\x5f\x21\xec\xcc\x10\x03\xcc\x60\xc8\x29\x01\x3e\x9f\xe4\x84\x5e\x05\x5c\x3e\x12\x32\x7e\x77\x8d\x23\x0e\xfd\x5b\xb6\x3c\x36\x41\x9d\xc7\x0e\xdf\xe5\xb7\x62\x77\xc8\xed\x48\x02\xe5\x41\x0a\x1e\x98\xae\xb1\x01\xb9\xfb\x06\xa2\xc5\xcd\xd5\xcd\x5f\xdc\x40\x0f\x79\xba\x0b\x19\x71\xd4\x79\xd7\x19\xfd\x6a\x92\x1f\x0d\x26\x30\x1a\x41\x37\x46\x68\x44\x14\xde\x26\x5e\x0c\xa9\x10\xab\x96\x4d\x85\x19\xe6\x0b\x0a\x35\x61\xde\xce\xf7\xa3\xa5\x3d\xc9\x82\xa6\x08\x81\xe9\xb3\xc3\x0f\x27\xa7\xa7\xef\x4f\x5e\x9e\xd6\xa2\x43\x1c\xf7\x21\x88\x36\x4f\xea\x78\x32\x09\x53\xb6\x79\xf0\x18\x0f\x62\x51\x7c\x0c\xcf\x4a\x7b\xed\xcb\xab\x8b\xc6\xa8\x37\xf8\x4c\x19\x5f\x1d\x0f\xda\xc3\xab\x8b\xd1\x10\x52\x4c\x2c\xe6\x69\x02\x86\x6f\x74\x5b\x83\x5e\xa7\x35\x7e\xdb\xbb\x6c\x57\xb1\x0b\x53\x55\x96\xf1\x0e\xfd\x08\xbd\x09\xdc\x02\x01\x46\x97\x1a\xfe\x7a\x80\x39\x03\x32\x62\xaf\xd5\x23\x2f\x20\x31\x42\x21\x85\xff\x28\x99\xf0\x19\xd9\x7f\x0b\x3a\x3c\x26\x22\x86\x40\x64\x52\x0a\x09\x07\x92\x41\x9d\xd3\xa9\x8c\x59\xe0\x42\x6e\xca\x25\xe4\x60\xbd\x10\x04\xf2\x95\x88\x03\x48\xdc\x1d\xe8\xf4\xfd\x6b\x53\x60\x3b\x90\x0b\x31\xcf\x5c\xe3\x5c\x4c\xc9\x82\x2e\x89\x48\x75\x6e\x9d\xb2\x43\x9a\x57\x07\x75\xcb\x56\x73\xce\x00\x0d\x38\x35\xd0\x01\xca\x0b\x8a\xac\xf5\x1b\xdd\x4e\x93\xbc\xc7\x02\x80\x4d\x20\x16\xff\x63\x18\xc9\xb4\xe3\x67\x01\x7d\xf5\x8c\x38\xb1\x09\xb6\x8c\x1f\xd1\x1f\xe0\xd6\x38\x2b\x8a\x0f\xcd\xd4\x0d\xc1\xbe\x49\xfc\x0e\x9b\x10\x30\x23\xe8\xd2\x5c\x00\xa5\xf9\x96\x95\x4e\x91\x89\x84\xc6\xdc\x3f\x36\x45\x6b\x81\xd9\x94\xda\xa2\x88\x63\x9c\x49\xb7\x4b\x12\x31\xa5\x40\x49\x69\x72\x4c\x62\xa1\xe1\x64\xb5\x49\xf8\x94\xe3\x05\x08\x7a\x33\x88\xc4\x5c\x6c\xba\x50\xc1\x2a\xb5\x92\x62\x61\x94\x68\x5a\x9a\x34\xa9\x90\x4b\xba\x9c\x30\x43\xf1\xe0\xd3\xca\x1b\xef\x1c\x34\x66\x94\x83\x4f\xce\xb5\xee\xc8\x35\xb0\xc3\x4d\xb6\xf7\x85\x94\xcc\xd7\x36\x94\xd6\x31\x29\xd9\xa1\x22\x8b\x79\x26\x44\x28\x66\x44\xd1\xa5\x7a\xf1\x21\x2e\x02\xee\xd0\x97\x03\xb1\x61\x79\x42\xf6\xf1\x17\x16\xcb\x3e\x2a\x28\xb7\x8d\x18\x32\x06\x45\x2c\xe4\xd7\x2c\xd3\x23\xb8\x59\x2c\x9c\x1e\x43\xa6\x81\x03\xdf\x98\x1b\x6f\x5b\x77\x83\x8b\xed\xc1\x00\xa8\xfd\x0e\xb3\x99\x0b\xd6\x6c\xeb\x2a\x96\xc3\xb2\x43\xe1\xdf\x22\x5f\x26\x60\x0d\x05\xac\x19\x16\xdb\xc8\xe2\x8e\x04\xda\x80\xa4\x82\x8d\xf6\x84\x4e\x20\xd5\x1d\x7c\xca\x25\x85\x3b\x93\x69\x20\xbf\xe5\x8d\x03\x7e\x9c\x99\xd5\x28\xdf\xd9\x30\x23\xbb\x4f\x1a\xad\xd7\x64\x06\x3d\xb4\xb2\x6d\xae\x30\xfd\x38\xde\x5a\x50\xe0\x08\x39\xc1\x6b\x81\xa9\x4b\x44\x73\xec\x77\xd1\xce\xe0\x7e\xc6\xf7\xac\x66\xff\x9a\x42\xa4\x04\x2c\x61\x10\x91\x70\x1b\xba\x00\x8b\x5c\x67\x01\x0f\x3d\x40\x18\x5a\x9f\x25\x0b\xe8\xda\xc0\x5d\xc1\x1f\x58\xbc\xb6\x15\xc5\x9c\x3c\x09\xa1\xf3\x32\xe7\x79\x56\x09\x0e\x13\x6a\x25\xc8\x84\xf9\x14\xe2\x00\xbb\x32\x8e\xb4\xa1\x07\x26\x93\x74\x86\xd6\x34\x79\x17\xf9\xf6\x80\x6f\xcf\xf2\x9c\xf5\x2c\x5d\x88\x23\xed\x3c\x30\x93\x06\xc9\xde\x00\x6e\x0a\x74\x8f\x6d\x6c\x0b\xcc\x35\xa8\x86\x1b\x5b\xb8\x8b\x39\x62\xcd\x12\x74\x37\x90\x98\x62\x0e\x49\x1e\x38\x12\x1a\xb5\x6e\x3a\x1f\x88\xc5\x63\x82\xc9\x27\x48\x8c\xf2\xea\x2e\x6b\xd9\x71\x40\x2e\xad\x4e\xe8\x8c\xf2\xf8\x78\xa5\x1f\x2c\x50\x80\x8e\x86\x00\xc1\x67\x33\x4b\xdb\x59\xbc\xe8\xda\x39\xe6\x31\xe5\x15\x1d\x60\x25\x4e\xce\x03\xcc\xd6\x10\x13\x2b\xf4\xb9\xc8\x26\xf1\x43\xa1\x5c\x2d\x1d\x36\x07\xed\x76\x77\x7c\xd1\x6b\xb4\x3a\xdd\x73\x08\x81\x75\xe5\x02\x70\x30\x49\x42\xd4\x52\x55\x26\x42\xe8\x31\x18\x3d\x01\x73\x33\x33\xec\xe3\x48\x4d\x0e\x3f\xc8\x43\x8c\x83\xbd\xc5\x1c\xfb\x1e\x8c\x85\x22\x42\xdb\xd9\xd5\x4c\x6b\x67\x7b\xd6\x2c\x31\x3f\x2e\xe5\x5b\x96\x6f\xd2\xb1\xb0\x3a\x5e\xcc\x97\x66\x1f\x3a\xb7\xbf\x40\x1e\x25\x40\x0b\xeb\x0a\x78\x17\xd0\x28\xce\x47\x80\x0a\xc5\xe2\xff\x5c\xaf\xf3\x9d\x04\x76\x0d\x11\x06\x50\x08\x26\x9b\x83\xd5\x19\xf8\xcc\xd2\xb2\x94\x19\x07\xa1\x2a\xab\xaa\x67\x82\x24\x55\xb6\x2f\x49\x61\x70\x03\x99\xd7\x06\xad\x14\x83\x3c\x12\x00\x5c\x34\x2e\x62\x2b\xd8\x55\x8b\xd4\x9f\xe7\x7d\x42\xc4\xbe\x33\xec\x5a\x2c\x1e\x63\xf1\x82\x1c\x0d\xfc\xc1\x2c\xf0\xac\x5e\x4c\x49\x79\x0a\xc3\x34\x61\xf2\x87\xbc\xee\x0a\x45\x22\x93\x0a\xc6\x59\x6c\xa2\x2a\xd8\xc5\x1d\x4a\xa4\x0a\x39\x1d\xea\xf1\x92\xc0\x6d\x05\xe1\xff\x9f\x7f\xfd\xf3\xdf\xf9\xb4\x76\x92\xf5\x6a\xf0\x7f\xae\xf5\x71\x2b\xcc\xd6\x2f\x6e\xf4\x6b\x9a\xf1\xc0\x0e\x7d\x26\x2f\x78\x66\xde\x83\x71\x0f\x12\x9b\x7f\x8d\x31\x39\x05\x5b\x9a\x81\x8f\x56\x4f\x6b\xcf\x4f\x4e\x1e\xd7\x9f\x56\x9f\x3c\x7d\x5e\x33\x73\x1e\xb4\x05\x26\x41\xe5\x3a\x82\xb9\xb1\x82\x4b\x6e\x10\xde\xd7\xa5\xbd\x87\xc4\xfa\xe0\x13\x5c\x50\x38\xcf\xad\x59\x17\xcd\x8e\x6a\x2f\x89\xa9\x05\xcb\x97\xc6\x86\xe4\x08\x3c\x80\x5a\xb6\xdd\x6a\x11\xb7\x4c\x9e\x17\x0b\x2f\xcd\x77\x8a\xd9\x02\x7b\x25\x94\xc5\xf2\xb9\xdb\x66\x6c\xb7\x38\xe0\x2f\x74\xa4\x1a\x26\xc5\x72\x06\xea\xd5\x4f\xcb\xf7\x44\x50\xbf\x2f\x82\xda\x7d\x11\x9c\xdc\x13\x41\xed\xf9\x7d\x11\x3c\xbb\x2f\x82\xa7\xf7\x45\xf0\xd3\x1f\x44\x60\xfa\xeb\x3f\x78\xd7\x6d\x38\xbd\xfb\x20\x61\xb7\x5a\x52\xcf\xc9\x11\xd5\xf1\xbd\x89\xe2\x90\x09\x96\xf7\x42\x38\x13\x62\x16\xb2\xef\x81\x0f\x32\xb5\xc7\xa3\x99\x47\x65\xc4\xe8\x84\x7b\x37\x4f\xd7\xec\xa2\xf3\x43\x5c\x66\x01\xd9\x64\xd0\x18\x4d\xb9\x0f\x18\xb3\x4d\xd3\xeb\xc6\xb0\x9d\xcd\x3c\xb8\x4e\x73\xab\x10\x1e\xe1\x22\x11\x47\x35\x9a\x40\xf2\xaf\xf8\x40\xc7\xbb\xce\xcd\xe9\xbf\xb0\xa5\x3f\x87\xe2\xad\xaa\xd0\xc7\xf2\xb8\x72\x9d\x7d\x27\xde\x88\x98\x4d\xcc\x84\xc7\x55\xdc\x7a\x2a\x3e\x2b\x3e\x74\x34\x76\x52\x0b\xb8\x82\xd2\x3f\x49\xb1\xfb\xaa\x24\xb5\xfa\xf7\x26\x4a\xbc\x3e\x29\x97\xbf\x4e\xfa\xcf\x90\x17\x5f\x95\x40\x07\xa8\x59\x23\x49\x7a\x71\x47\xa7\xb1\x33\xc2\xef\x59\x28\x16\x77\xcf\xa7\x25\xa8\x61\x22\x48\xa1\xea\xe5\x96\x8c\xe8\x29\x49\x32\xc6\x99\xcc\xec\xc8\xcd\x2b\xd9\x04\xdb\xc6\x4e\xbf\x01\xa3\xff\xe8\x6b\xa4\x01\x6c\x3d\x08\x27\x3c\xc4\xc9\xca\xdc\x07\xcc\x38\x7c\x19\xcc\x48\xcd\xe3\x09\xcd\x76\xf2\x9e\xba\xe6\xc9\x18\xdb\x4e\xa8\x8a\x63\x88\xba\xb1\x59\x78\x8f\x81\x3f\x1f\x1a\x5e\xec\x07\xcc\x7a\x64\xf5\xaa\xd1\xd4\x8b\x20\x18\x2e\xf8\x54\x0f\xa1\xcc\xfa\x8e\xa7\xc6\xa0\xf9\xb6\xf3\xae\xbd\x9e\xcb\x91\x9d\xd5\x97\xf5\xe2\x71\xd7\x0e\xbc\x6a\xd1\xa5\x09\x9a\xb5\xb0\x90\x5c\x35\x15\x41\x80\xbc\xe4\xc9\xba\x35\x27\xd0\x31\xdd\xea\x41\xcd\xd6\xed\x34\xfe\x3b\x4f\x88\xf7\xb7\xdc\xba\xbd\x5e\x41\x89\x81\xfc\xea\x91\x8e\x12\x00\x95\xd1\x26\x10\x3c\x8c\xae\x03\x5c\x55\xcb\x02\x70\x91\x43\xfb\xea\x45\x60\xdb\x8e\xbb\x31\x39\xdd\xbd\xda\xdf\x7d\xa9\xfa\xa8\x12\x2c\x43\x3e\xf9\x06\xfc\x0e\x14\xe9\x6c\x31\x6f\xa4\x04\x2e\x2b\x4e\xbe\x4a\x06\x56\xa9\x56\x2a\x56\xb8\x02\x67\xe6\x96\x5d\x0d\x75\x7a\xc3\x4e\xec\x4b\x86\xaf\xd8\x5e\xa3\xbd\xbb\x69\x34\x71\xff\x30\x63\xf3\x3d\x49\x69\x0f\x7a\x96\x88\xeb\xa6\x48\x63\x0d\x4d\x27\xbe\xac\x94\xec\xc6\xb3\xef\x37\xb2\xec\xe6\xe3\xe1\x91\x79\x91\x2b\x8b\x48\xbf\xbc\xe4\x2d\xdb\x25\x6f\xf3\xcd\xeb\x34\x0e\x42\xf6\x0e\x9a\x30\x88\xe4\x32\x86\x76\x27\x9e\x0a\xfb\x16\x05\xb7\x0c\x31\x5b\xe4\x91\x62\x3a\xc8\x26\xd8\x1c\x77\x76\x7e\x85\x47\x5b\x5c\xe4\x7d\x6a\x03\x55\x1e\xc1\x6a\x28\xdf\x84\x79\xf8\x70\x07\xd2\x1f\x6b\x47\x47\xb6\x33\xdc\xfb\x92\x88\x43\x98\x78\x5f\x6c\x48\x48\x0e\x8a\x14\xf0\x1d\x59\x5e\xe4\xb2\x59\xac\x80\xd7\x93\xce\x68\x08\xa9\xe7\x4a\x31\xd5\x15\x71\xfb\x96\x45\xf8\x82\xd4\x97\x4b\x3b\xeb\xe2\x98\x81\x6b\xf8\x6c\x71\x6d\x37\x19\x30\x62\x97\xf6\x68\xac\x16\x4c\x8e\xc4\x1a\xf8\x9b\x2c\xf1\x45\x72\x5b\x5c\x1e\x59\x2b\xbc\x27\x07\xdb\xe4\x70\xb5\xfa\xa8\xdc\xc2\x06\xb6\x0b\xa9\xa8\x8d\xbb\xe8\xf2\x23\xf2\x71\x6d\x88\x2f\x71\x83\xa2\xbf\xf8\xb2\xec\x13\x7c\x93\x6a\xa4\xdf\xd6\x9e\x75\xf4\xff\x06\x00\x00\xff\xff\x13\x35\x36\xf9\xf7\x24\x00\x00")

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

	info := bindataFileInfo{name: "scripts.bash", size: 9463, mode: os.FileMode(420), modTime: time.Unix(1468083986, 0)}
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

