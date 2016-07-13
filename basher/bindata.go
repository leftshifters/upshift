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

var _scriptsBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5a\xef\x72\xe2\xca\x72\xff\x6c\x9e\xa2\x0f\xa6\x8e\xed\x73\x2d\x30\xd8\xfb\x7f\x7d\x52\x18\xb4\xb6\x72\x31\x50\x80\x77\xeb\xd4\xee\x16\x35\x48\x03\x28\x16\x1a\x65\x24\x19\x93\xbd\xce\x1b\xe4\x19\xf2\x1e\x79\xaa\x3c\x42\xba\x67\x46\x20\xd9\xd8\xbb\x59\x6f\x3e\xe4\xd4\x29\x16\x34\x33\xfd\xbf\x7f\xdd\x3d\xf2\x88\xc7\xc9\xd0\x95\x7e\x94\xec\x1f\xc0\xb7\xd2\x0e\x77\xe7\x02\x2a\xf5\xd2\x4e\x1c\x70\x1e\x01\x7e\x09\x62\xb0\x02\x56\xba\x2b\x95\x36\x7b\x1b\x7a\xf3\xa5\x3d\x1c\x36\xcf\xed\x53\xda\xdf\xe9\x9d\x7f\x70\x3a\xf6\xf0\xb4\xd2\x78\x40\xa5\x91\x51\x81\xc6\x9f\xbf\xd7\xe1\x1f\x90\x70\x0e\xb8\x0f\x69\x9e\x4b\xe6\x05\xdc\x09\xe3\x84\x05\x81\xa6\xea\x74\x87\xa3\x66\xa7\x33\xbe\x1a\x74\x4e\xcb\xf3\x24\x89\xe2\xb7\xb5\x5a\xcc\xe5\x8d\xef\xf2\xb8\x3a\x53\x07\xaa\x42\xce\x6a\x9e\x1f\x27\xd2\x9f\xa4\x89\x2f\xc2\xb8\xa6\x17\xac\x46\xb5\x7e\x62\x4d\xfc\xb0\xfa\x6f\x7e\x54\x2e\xed\xf8\x53\xf8\x0c\xd6\x14\xca\xdb\x96\xe1\xeb\x3b\x48\xe6\x3c\x2c\xed\xec\xb8\xa9\x0c\xc0\xb2\xe2\xb9\x58\x5a\x5c\x4a\x21\xf1\x47\x24\xc5\x4c\xf2\x38\xb6\x26\x8c\x7e\x4e\x99\x4f\x5b\x02\xe1\x32\xe2\x08\xe5\xca\xb7\x9c\xa8\x77\x65\xf8\x13\xb6\x30\x29\xed\x4c\xfd\xd2\x4e\x1a\xe2\xd7\xed\xcb\x6d\x67\x70\x5a\xae\xec\x83\xeb\x01\xfd\xe3\xf9\x32\x64\x0b\x4e\xd4\xcf\x9a\xc3\x8b\xf1\xb0\x77\x35\x68\xd9\x9f\x8f\xbe\x22\x83\x83\x32\xfc\xfe\x3b\x44\x4b\x0f\xbf\xa1\x91\x6f\x23\x21\x13\x38\x1f\x34\xdb\x1d\x7b\x7c\xd1\xbb\xb4\x91\xce\x37\x24\x77\x97\xb7\xc5\x66\x63\xbf\x39\xba\x38\xad\xd0\xe7\xdb\x4a\xee\x50\x0d\x25\xd9\x78\xe2\x93\x64\x51\xc4\xa5\xf6\x84\x26\x03\x4b\xfd\x6c\xb3\x69\x39\x62\xf1\xb5\xde\xd2\x6f\x0e\x9a\x97\xc3\x2c\x02\xc6\x9a\x07\xfa\xbb\x6a\x64\x58\xa2\x22\xf5\xf2\x43\xbf\x0f\x79\x92\x46\xe7\x7c\xa1\xa9\x9c\xdb\x97\x8a\x44\xcb\x51\x87\x07\x42\x24\x7d\x16\xc7\x4b\x21\xbd\xd3\xca\x71\xc9\xb8\xb1\x5c\x69\x94\xe1\xf4\x14\xca\x89\x4c\x79\xde\x79\xbb\xd0\x72\xc0\x8f\x81\x9e\x1f\xc2\x92\x43\x28\x96\x10\x72\xee\x41\x64\xa8\xe0\x26\x43\xe2\xb8\x0c\xbf\x21\x89\xfc\x71\x1d\xae\x16\xd7\xab\xff\x80\x38\xf5\xf0\xe7\x35\x58\x43\x98\xf1\x05\xf8\x3a\x38\x55\x3c\xe3\xde\x5b\x3f\x81\x23\xfc\xc6\x83\x98\xaf\x1f\xd0\x12\xf9\xd9\x3c\xdc\x85\x2b\x0c\x58\xc0\x68\x4a\x03\x0f\x92\x55\xc4\x91\x0a\xb1\xcb\x0b\x94\xb1\x79\xc0\x23\x63\x81\x04\xd1\x52\x57\xa1\x59\xfc\xff\x61\xad\x34\x13\x17\x2c\xe6\xdc\xfe\x9f\x1a\x6d\x1b\xab\xa2\xed\xfa\xc2\x2b\x40\x8b\x01\x29\x65\xbd\x48\x78\x6b\xb3\xe7\xc3\xb3\x6e\x0e\x0e\x78\x24\xae\x22\x8f\x25\x7c\xfb\x59\x89\xeb\x90\xaa\x0d\x88\x0a\x37\x5c\x4e\x44\xcc\x1f\x52\x3a\xf7\x93\x61\x3a\x59\x08\x2f\x25\x98\xf3\x93\x87\xc4\x66\x28\x71\x9c\x6d\x41\x91\xf0\xe7\x93\x54\x1e\x13\xaa\x48\xc7\x48\xb6\x8d\x52\x3f\xcd\xec\x31\xb0\x2f\x7b\x23\x7d\xfa\x6c\xd0\xec\xb6\x74\xee\xae\x89\x1e\x6b\xa2\x51\xaa\x22\x13\x33\xb7\x40\xee\x98\xc8\x9d\x49\xbe\x2c\x98\x78\xd4\xeb\x75\x14\xc1\x09\xae\x14\xe2\xda\xb8\xc6\x1c\xba\x8a\x08\x20\xf8\xb6\x43\xa9\x5e\xda\x76\x28\xf3\xf8\xd6\x63\xe1\x76\x6e\x86\x53\xbe\xca\x9d\xd9\xa3\xa6\x3a\xb8\xb5\xd6\x48\xb6\xac\xa2\xde\xf3\x74\x92\x62\x40\xba\x22\x4c\x78\x98\x54\x5d\xb1\xa8\x05\x7c\x9a\xc4\x73\x7f\x9a\x70\x19\xd7\xd2\x48\x7d\xad\x2d\x58\x8c\xbf\x6b\x86\x79\x35\x9e\x97\xd7\x19\xf8\x8d\x18\xdd\x6d\x4f\xc3\x5f\xc1\x7a\x26\x8a\x6c\x29\x9b\x74\x25\x9b\xc6\xc3\xce\xc3\x0a\x85\x25\x6a\xb3\x9f\x6a\x89\x3b\xc7\x70\x81\xbf\xdd\xde\x7b\x5c\xad\x15\x7f\xcb\x45\xe1\x91\x02\xf0\x84\xc9\x64\xe8\x2f\xd2\x80\x25\xc2\xd4\x8b\xb6\xfd\xd1\x69\xe9\x70\xba\x75\x65\x1a\x2a\xaa\x32\x5d\xa0\x0e\x58\xfe\x4d\x29\xa8\xff\x59\xf3\xf8\x4d\x2d\x4c\x4d\xde\x11\x31\xa7\x37\x74\x42\x57\x72\xda\x79\x96\xfa\x81\xd7\x4d\x17\x93\xac\x0a\xf5\x07\xbd\x7f\xb6\x5b\xa3\x71\xb7\x79\xa9\x68\xa3\x8a\x62\xb1\xf0\x93\x96\x48\xc3\xe4\xb4\xb2\x4f\x21\x2a\xf9\x8d\x15\x60\x3b\x80\x99\xa8\xf0\xc0\x72\x69\xf1\x40\x59\x43\x16\x89\xe2\x89\x5a\x1a\xcb\x5a\xe0\x4f\xf8\x2d\x77\x6b\x7d\x3a\x77\x96\x7a\xde\x0a\x2c\x17\xca\x7d\xe9\x87\x09\xb4\x3e\x9c\xa5\x21\x96\xaf\x8f\x68\x6d\xac\xf5\x65\x0c\xa9\x9a\x13\x4e\x45\x35\xa2\xdd\x07\x28\x43\xc8\x97\x79\xa2\xe5\x8d\xd3\x73\xd2\x95\xc1\x9a\x25\xf4\xe8\x81\x14\xf9\x48\xb8\x47\x2a\x4f\x60\x0d\x8e\xf7\xf7\xec\xef\x6f\x21\xfa\xb7\xfa\xc1\x81\x0a\x82\xd2\xce\x53\x2a\x62\xed\x85\xb7\xf7\x34\x84\x4a\x91\x43\x99\x9c\x95\x53\xb9\x8c\x34\x77\xa1\xe9\x79\xe0\x8c\x86\xcd\x28\x42\xa8\x8e\xbb\x22\xb4\x6f\xf9\x22\x4a\x6c\xf4\xdd\x2a\x52\x4d\x11\x8b\x61\xca\x50\x64\x40\x63\xa0\x5f\xfc\x38\xdc\x43\x54\xe2\xa8\x08\x0b\xe3\x25\x97\x23\xb1\xd9\xfc\x43\x9e\x78\x92\xdd\x03\x29\x0f\xb4\x17\x3e\x43\xe5\x21\x3b\xca\xc2\x3f\xca\x6d\xc1\x63\xe8\x8a\x04\xec\x5b\x3a\xf0\x07\x7c\xdd\x38\xe2\x29\x69\x48\xf5\xb7\x4f\xeb\x3e\x11\x22\xd0\xda\x3f\xb4\x9e\x29\x48\x18\xe8\xca\xc8\xc5\xc8\x1e\xfd\xd5\xd7\x59\x93\x3d\x58\xb7\x51\xc3\xd6\x85\x7d\xa9\x91\x38\x4b\xae\x93\x7c\xa7\xf5\x02\x5b\x6c\xf4\xa6\x25\x20\xf2\x23\xae\x3a\x54\x4c\xd7\x5b\x57\x78\x7c\x42\x7c\xc0\x52\x39\xa7\x7a\x01\x2b\x76\xe7\x7c\x61\xea\xb6\x35\xf7\x11\x12\xe7\x3c\x08\x34\x2e\xda\xe1\x8d\x2f\x45\x48\x09\x88\x1b\xbd\x6b\xf0\xa3\xb9\x08\x79\x9c\x25\x38\x58\x1e\x76\xff\x7e\x68\x5a\xdf\x08\x1f\x4e\x85\x5c\x9c\xde\xdb\x76\x48\xbd\x2b\xca\x58\xa6\xfd\xd2\xbf\xe1\x5e\x9b\x25\xac\xcf\x92\x39\x54\x33\xd0\xd2\x82\xe9\x22\x52\xae\xbc\xa0\x1e\xe2\xd6\x8d\x24\x4f\x92\x95\x31\x11\x0d\x1a\xbf\xd0\x42\xb7\x6e\x42\x9e\x79\xd4\x14\xbf\x42\xdd\x04\xb7\xe7\x74\x32\x8a\x34\xa5\x3b\x47\x23\xfc\x84\x2e\x45\xd0\xfb\x45\x3e\xff\x8e\x4b\x2c\xa6\xc5\x2d\xae\x55\x4e\xaa\xb7\xae\x59\x81\xec\xdf\xa7\xbc\x67\xab\x71\x63\x3b\x76\x17\x87\x84\xef\xa8\xa1\xe7\x16\x63\xc3\xec\x67\x4f\xe5\x5a\xac\xb2\x13\xaa\x11\x2a\x84\x6d\x4e\x4d\xaf\xe9\x64\x7b\x4c\x8f\x7a\x4e\x0f\x43\xac\xb8\x23\xdf\xdd\x28\xeb\x15\x55\xfb\xc0\x13\x77\xde\x0c\xa9\x31\x64\xbe\xec\x4b\x71\xe3\x13\x72\xfa\xe1\x0c\xbf\x4f\xfd\x80\xc7\x5a\xe7\x5d\xd8\x1d\xf5\xda\x3d\xe8\xf6\x46\x70\x35\x74\xba\xe7\x30\xba\x70\x86\xd0\xec\xfe\xf5\xe9\xc2\x1e\xd8\x87\xd0\xb5\xed\x36\x8c\x7a\xd0\xb6\x5b\x4e\xdb\x06\xe7\x03\x7c\xb2\x37\x0f\x7b\x6a\x3b\x91\x49\xe6\xd8\xa5\x63\x1f\x10\x43\xec\xcf\xe6\xf4\xc4\x13\xcb\x30\x10\xcc\x03\x2a\x72\x88\x8e\x62\x6d\x01\x98\x8a\x00\xdd\x4b\x9b\x58\xe8\x29\x40\x03\xaa\xc0\xd4\x3e\x4f\xd3\xd0\x55\x71\xdc\x17\x11\xc5\x2d\xdf\x26\x7b\x69\xa7\xd9\x6a\xf5\xae\xba\xa3\xb1\x7d\xd9\x74\x74\x5b\x45\x6c\xa9\xcf\x45\x75\xc1\x4a\x61\xfd\x28\x13\x63\xac\x6a\xad\x5e\x58\xdc\xc0\x1f\xd5\x85\x98\x20\xad\x28\x23\xbf\x96\x0e\x03\xe1\x29\xd6\x68\x5c\x84\xd7\xe1\x12\x9d\x30\x14\xa9\x74\x33\x4b\x36\x07\xad\x0b\xe7\xa3\xbd\x89\x1e\xa7\xdf\xdc\xfc\xd0\xa3\x3d\x4e\xcd\x5b\x5d\x5c\xd3\xe4\xd2\x88\xfc\x5c\xcb\x97\xdb\x88\xea\xca\x54\x21\x3a\x0a\x01\x79\xb6\x80\x16\x25\x8b\x21\x1f\x34\xa8\x44\xc5\xbe\x84\x88\xdd\x66\x7c\xb7\xfe\x35\x17\x4d\x8d\xaa\x1f\x31\x62\xbf\x7e\x94\x2c\x70\x9a\xdf\xa1\x56\xa9\xb8\x09\x1f\x2e\xae\x3d\x32\xa1\x2c\x6c\x2e\x4a\xa8\x91\x45\xc4\x74\x17\x81\xac\xe4\x74\x7b\xe4\x6e\x3f\x54\xfb\xa3\xea\xad\xb0\x70\xfd\x00\xfd\x6c\x2b\xf1\x79\x20\xbc\xd2\x12\xa5\xac\x66\xfa\x55\xcd\xb6\x6a\xad\x5a\xd5\xca\x15\x24\x53\xa7\xcc\x98\x1a\x51\x44\xa0\xe1\xd0\xb5\xa3\x34\xcc\x5c\x88\x98\x6c\x77\x7a\x7d\x7b\x30\x36\xe1\xb5\xf6\xe3\x1a\x05\x22\x3f\xc0\x72\x9c\xaa\xf3\xd8\xb9\x51\xe3\xab\x6e\x40\x70\xe2\xb0\x2c\x12\x42\x23\x99\x15\x5f\xfb\xd1\x78\xc9\x7c\x84\xe5\xd9\x18\xbd\x33\x56\x30\x31\xc6\x68\x43\xcf\xc5\xe4\x4a\x6a\xb1\x37\xb3\x98\xca\x59\x3f\xf4\xb6\x84\xdc\x53\xb2\x9d\x5d\x75\xe9\x6a\x04\x33\xb3\x3b\x72\x3e\x38\xf6\x40\x43\x15\xc5\xbd\x0a\x75\x84\x17\x73\x87\xd1\x92\x1c\x23\x1a\x1b\x82\x5e\xe8\x24\xdf\x53\x79\x3b\xd9\x22\x3e\x22\xee\xa3\x36\x5e\xea\xf2\x3c\x2b\xea\x69\xa3\x68\x6c\x6e\x85\x8e\xcb\x7a\x3e\x7d\x3c\x9d\xbe\x87\x42\x97\xbd\x47\x41\xe8\x93\xd3\xe9\xe4\x11\xe8\x1c\x01\x9a\xf2\xe1\xea\xca\x69\xc3\x54\x8a\x45\x2e\xa3\x77\x21\x1b\x5a\x66\x08\xba\x66\x6a\x51\xa3\xca\xe2\x36\x92\xb5\xd7\x8d\xa3\xd7\x8d\xd7\x6f\xd8\xcb\x63\x97\x9d\xf0\x63\x76\xfc\x82\x9d\xd0\xa1\x8e\x10\x11\xd2\x94\x22\x45\x83\x12\x7e\x28\x99\x0f\xa9\x5f\x5c\x89\x14\xe7\xfa\x04\x18\x5c\xb5\x9d\xf6\x21\x30\x4f\x01\xd9\x22\xcb\x4a\x05\xee\x62\x0a\x51\x86\x1b\x79\x0c\xa8\xad\x6b\x41\xf1\xae\x63\x84\x07\xef\x81\x24\x70\x6a\xfd\x28\xd3\x28\xc7\x0d\xb5\x2e\x99\xd7\xdf\x40\x56\xed\x01\x9a\xbd\x43\xdc\xa3\xbb\x8b\x34\xf5\xbd\xa7\x5b\xd7\x3d\xdd\xba\x92\xd5\xf6\x40\x4d\x3b\x71\x82\x58\x03\xef\xdf\xbf\x87\xca\x7e\xcc\xb1\x73\xf7\x93\x15\xb8\x0b\x9c\x8a\xda\x60\xf9\x34\xa8\xe5\xc4\xb8\xa3\xfb\xb2\xf5\x90\x74\x40\xd3\x06\x69\xe2\x4c\xc9\x32\xe4\x09\x25\xff\xa1\x06\x79\x57\x44\x2b\xec\xb4\x0f\x4d\xc3\x3d\x67\xaa\xe3\x66\x01\xc6\x26\x0a\x33\xe1\x7a\x8b\xcf\xe9\xf2\x64\x3d\x96\x92\x0a\x77\x5b\xae\x76\xf4\x86\xdf\xe8\xc2\xf4\xdf\x6b\x1d\x7f\x22\x99\x5c\xd5\x2e\x95\x19\xda\x9c\xee\x5f\x6b\xf9\x80\xfb\x02\x59\xc8\xd5\x0c\xc9\x07\x05\x20\x4f\x9b\x2e\x9e\x48\x58\x3a\x41\x2e\xd5\x41\x65\xca\x16\xfd\x47\xb8\x37\x7d\x68\x8b\x5f\x22\x89\xe6\x90\x41\xff\x25\x3e\x45\xbc\x28\x72\x82\xcf\xe6\xec\x57\x92\xce\xf0\xd4\xf8\xaf\xef\xa9\xcc\xa7\x87\x38\xaa\x47\xed\xcd\x1d\x03\xb6\x05\x52\xf8\x5e\x87\x61\xa9\x9d\xdb\xf9\xa1\xb8\xb4\x23\xb9\xd7\x12\x01\x16\x98\xd3\xbd\x2f\x47\xc7\xc7\x9f\x8f\xde\x1d\xd7\x17\x7b\x74\xbb\x8a\xce\xb9\xbf\xd2\xa0\x95\x49\x90\xf2\xfb\x0b\x27\xb4\x10\x8a\xe2\x63\x7c\x56\xda\xb1\x2f\xaf\x3a\xcd\x51\x6f\xf0\x48\x93\xb5\x5e\x1e\xd8\xc3\xab\xce\x68\x88\xa1\x1b\x8a\x79\x1a\xa1\xa1\x9b\xdd\xf6\xa0\xe7\xb4\xf5\x35\x30\xf5\xc8\x71\x8d\x1b\xd9\x11\x70\xd9\x8d\x97\x5d\xde\x22\x08\xd5\xe9\xe3\xf7\x03\x35\x06\x2a\x60\x79\x8b\xa9\x85\xfd\x09\xfe\xcf\x60\xe2\xcf\x60\xf7\xa2\xd9\xfa\xfb\x21\x88\x30\x58\x81\xba\x41\xc7\x05\xc9\xb1\x7d\x48\x52\x19\x72\x2f\xbb\xcb\x9b\xfa\x12\xb3\x38\x59\x0a\x9c\x0a\x5d\x11\x7a\x98\xfa\x0e\xc4\xa9\x7b\xad\xfa\x16\x07\x73\x8c\xe2\xf7\x9a\xee\x24\x19\x2c\xd9\x0a\x44\x9a\xe4\xde\x23\x6c\xd1\xe6\x7d\xa5\xa1\xc5\x6a\xcd\x39\x92\xc1\x10\x46\x3e\xc8\x79\xc9\x48\xb4\x7e\xb3\xeb\xb4\xe0\x33\x41\x08\xb5\xe8\xd4\x53\x1d\x42\x94\x26\x99\x3c\x4b\x9c\x7a\x66\x90\xa9\x0d\x84\xb2\x5f\xe9\x5a\x1b\x4f\x8d\x0d\x7a\xef\xab\x1b\x4f\x8c\xcb\xfb\xcc\xef\xa8\x45\x44\x37\xa2\x2d\xd5\x01\x34\x9a\xab\x45\x71\x8a\x42\x44\x2c\xf4\xdd\x43\x05\x7b\x4b\xca\x52\xa6\x61\x95\x4a\x9d\x4a\xe3\x15\x2c\xb0\x84\xa1\x91\xd2\xe8\x10\x42\xaa\x86\xf1\xe6\x0e\x29\x27\xcb\x9d\xbe\x50\x38\xda\xd2\xcd\x10\xce\xad\xb5\x58\x2a\x23\xaa\x82\x9a\x46\x55\xb8\x64\xab\x09\x57\x1c\x2b\xdf\xd6\xd1\x78\x97\xed\xa6\xe0\xaf\x7c\xcb\x42\xeb\x0e\xae\x51\x1c\x3d\xb7\xbb\x42\x4a\xee\x26\x3a\x07\x32\x46\x17\xa8\xd6\x5e\x0c\xcb\xb9\x51\x22\x10\x33\x88\xd9\x2a\x7e\xfb\x25\x2c\x6e\xdc\x62\xaf\x6c\x8b\xbe\xfe\x3d\x82\x5d\xfa\x20\x10\xee\x93\x81\x72\x37\xc1\x43\xce\x11\x1c\x03\xff\x9a\x1b\x3b\x62\x98\x85\x22\xb3\x63\xc0\x13\x94\xc0\x55\xee\xa6\xd3\x3a\xdc\xf0\xa0\x3d\x18\x20\xb7\xff\x85\xdb\xd4\x01\xed\xb6\x0d\x3a\xe6\xa8\x6c\x31\xf8\x8f\xe8\x67\x14\xac\x93\x82\x75\x25\xa2\x4d\x22\xe6\x2e\xc0\xd7\x4d\x28\x56\x40\x1a\x83\x26\x6c\xb2\x3a\x44\xf7\xe4\x40\xe1\x2e\xa6\x2b\x36\xc4\xa9\xbc\x73\x30\x8e\x8d\x5b\x95\xf1\x33\x1f\x1a\xb6\xbb\xd0\x6c\x9f\xc1\x0c\xfb\xc4\x58\x4f\x0f\x42\xdf\x4c\xe0\x29\x6a\x9c\x94\x24\x74\xcc\x53\x10\x0a\x89\x4f\x63\x04\xf9\x19\xc3\x4f\xc5\x9e\xb6\xec\xbf\xa4\x98\x29\x1e\x8f\x38\x66\x24\x9e\xc6\xea\xa2\x89\x27\x26\xe1\xb1\xb6\x04\x81\x8e\x59\x58\x62\xdd\xc7\x70\xc5\x78\xe0\xe1\xc6\x57\x8c\xca\xea\x24\xc0\xda\xad\xd6\xf3\xa2\x02\x8d\x7a\xf1\x5a\x91\x09\x77\x19\xe6\x01\xd5\x75\x9f\x78\xe3\x68\x01\x93\x74\x46\xde\xdc\x61\xde\x44\xc9\x6d\xa1\xdc\x96\x96\xd9\xd4\xc2\x2e\xe6\x51\x92\x45\xa0\xd1\x86\xd8\xde\x20\x6d\x86\x7c\x0f\x75\x6e\x0b\xc2\x1a\x32\xc3\x8d\xae\x31\x45\x8c\xd8\x88\x34\xa7\x6b\x2b\x3f\xf4\xe3\x39\xe6\xe0\x44\x24\x64\x75\x55\x51\x31\x17\x0f\x81\xc0\xc7\x8b\x94\xf1\x1a\x19\x6a\xe9\x29\x4b\xae\xb4\x4d\xd8\x8c\xf9\xe1\xe1\xda\x3e\xe8\x29\x22\xc7\x02\xdc\xe1\xcf\x66\x9a\x77\xe6\xf1\x62\x68\xe7\x84\x27\xc8\x2b\x06\xc0\x5a\x9d\x5c\x04\xa8\x37\x36\x04\xac\xd8\x29\x91\x98\xe0\x06\xd8\xeb\x1a\xff\x0f\x5b\x03\xdb\xee\x8e\x3b\xbd\x66\x1b\x3b\x3e\x4c\x01\xb2\x60\x4c\xb7\x3c\xb4\x1d\x5d\x12\x41\xbc\x8a\xab\x13\x21\x92\x31\x3a\x3d\x42\x77\x73\x75\x15\x43\x17\x1e\xb0\xf7\x45\xee\x51\x1e\xec\x2c\xe7\x54\xa2\x29\x17\x8a\x04\x75\xc7\x50\x57\x2d\x83\xee\x85\x0c\x30\x9f\x94\xf2\xd5\xf5\x87\x6c\x2c\xb4\x8d\x97\xf3\x95\x7a\x17\x35\xd7\x1f\xa8\x4f\x2c\xd0\x0a\x9b\x0a\x78\xe7\xb1\x45\x98\xcf\x80\x38\x10\xcb\x7f\xca\xca\xf2\x2f\x52\x58\x95\x74\x93\x40\x01\xba\x6c\x8e\x5e\xe7\x18\x33\x2b\x2d\x92\x71\x0e\xed\xaa\xae\xab\x9e\x4a\x92\x94\x40\x1b\x43\x39\xc5\x79\x18\x75\xde\x38\xb4\x5a\x4c\xf2\x85\xc0\xcd\x45\xe7\x12\xb5\x82\x5f\x13\x91\xba\xf3\x7c\x4c\x88\xd0\xcd\x1c\xbb\x51\xcb\x0f\xa9\x78\x21\x46\xa3\x7c\xd8\x63\xbe\x6e\x14\x21\x29\xcf\x01\x47\x3f\x2e\x7f\xcb\xdb\xae\x50\x24\x8c\x56\x32\x0d\xa9\x85\xaa\xc2\x5f\x22\xdd\x93\xc4\x15\x31\x1d\xeb\xf1\x0a\xf0\x74\x8c\xe9\xff\xdf\xff\xf9\x1f\xff\x95\x87\xb5\x23\xf3\xda\x8e\xee\x99\xb1\xf5\xb9\xff\x9f\x69\x85\xb2\xb7\x2e\xed\xbf\x67\xf3\x48\x4b\xb5\xa1\x7a\x8c\x50\x38\x61\xa9\x09\x82\x5e\xee\x27\xcc\xbd\xa6\x1c\x9d\xa2\x6f\xd5\x08\xc1\x6a\xc7\xf5\x37\x47\x47\x27\x8d\x57\xb5\x17\xaf\xde\xd4\xd5\xe4\x80\x6d\x82\x02\xac\x5c\x87\x30\x57\x5e\xc9\xc0\x0e\xd3\xfd\xba\xb4\xb3\x0f\x3a\x26\x5f\xd0\x75\x52\x16\xc9\x75\x1d\xb2\x66\xa9\xfe\x0e\x54\x6d\x58\xbd\x53\x3e\x85\x03\x8c\x08\xa6\xc5\xce\xde\x9b\xd1\x9d\xa0\x65\x85\xc2\x4a\xfd\x5c\x7f\x67\xde\x74\x6d\x94\xd2\x64\x1e\x3b\xae\x5f\x45\x28\x22\xf4\xf7\x03\x7e\x90\xe0\xf0\x51\x36\x5b\xad\xc6\x71\xd9\xfc\xa5\xc4\xcf\x12\x68\x94\x9f\x2b\x42\xfd\xd9\x14\x8e\x9e\x4b\xa1\xfe\xe6\xd9\x14\x5e\x3f\x9b\xc2\xab\x67\x53\x78\xf9\x93\xce\x54\x6d\xf7\x4f\x9e\xcd\xae\xa5\xad\xe7\x10\xe1\xb7\x89\x64\x56\xa6\xc7\xa2\x41\xef\xb5\x63\x1f\x01\x62\xf5\x2c\x82\x33\x21\x66\x01\xff\x15\xf4\x10\xc0\x2d\x7f\x31\xb3\x6e\x5f\xbf\xb4\xf2\xa1\xff\x20\x2d\x9b\x67\xce\x33\xd2\x32\x63\xc3\xe4\x82\xb3\x89\x6f\xdd\xbc\x62\xd6\xf3\x33\xed\x31\xe1\xd7\x90\xdd\x0f\x38\xa3\xd7\x65\xe6\xd5\x35\xc1\xbf\x36\x1d\xd0\x91\x97\x27\x80\xb2\x28\xf4\x35\x7f\x32\xb1\xd1\xda\x5c\x3f\x7d\x6c\x6b\xa5\xd7\xd3\xdf\xa8\x39\x38\xb7\x47\x6a\xf6\x43\x8b\xdc\x7f\x33\xa2\x94\xc0\x56\x7a\xa3\x85\x75\x03\xae\x22\x05\x34\xf1\x59\xa1\x1e\xfa\xac\x24\xbb\x77\x43\x01\xcc\xfb\x03\xd3\x83\xe1\x2f\x7a\xc1\x43\x5d\x99\x4b\xb7\x6c\xff\x13\x00\x00\xff\xff\xb6\x5a\xad\xbf\x44\x26\x00\x00")

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

	info := bindataFileInfo{name: "scripts.bash", size: 9796, mode: os.FileMode(420), modTime: time.Unix(1468387098, 0)}
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

