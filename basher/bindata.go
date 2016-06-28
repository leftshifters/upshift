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

var _scriptsMainBash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x5a\x7d\x72\xdb\xb8\x15\xff\x5b\x3a\xc5\x5b\xd9\x13\xc7\x59\x53\xb4\xe4\x64\xf3\xe9\x74\x14\x49\x49\xd4\xb5\x2d\x8d\x25\xa7\xb3\x93\x64\x3c\x10\x09\x49\xa8\x49\x80\x05\x49\xcb\x6a\xea\xbd\x41\xcf\xd0\x7b\xf4\x54\x3d\x42\xdf\x03\x48\x89\x94\xe5\xc4\x59\x67\x3b\x6d\x12\x11\xc0\xfb\xfe\xf8\x3d\xa0\xc7\xec\x82\xbf\x55\x81\xcf\x75\xfc\x70\x17\xbe\x54\x2b\xe1\x85\x2f\x34\x38\x11\xd4\xd3\x28\x9e\x89\x49\xe2\x06\x6a\x1a\xbb\x9b\x16\xc6\xa9\x08\xfc\xd2\x4a\xa4\xc5\x25\x4b\x78\xf5\xba\x5a\x3d\x8b\xa6\x9a\xf9\x7c\xe8\x69\x11\x25\x96\xb4\x97\xea\x00\x9c\x49\x3c\x3c\x82\x59\x92\x44\xf1\x0b\xd7\xd5\x6c\x5e\x9f\x8a\x64\x96\x8e\xd3\x98\x6b\x4f\xc9\x84\xcb\xa4\xee\xa9\xd0\x0d\xf8\x24\x31\x6c\x50\x32\x37\xe7\x18\xb2\x18\x7f\xbb\x42\xc6\x09\x0b\x82\x7a\x3c\x83\xd7\x90\xad\xd5\x13\x1e\x46\xf0\xe0\x01\x78\xb3\x50\xf9\xf0\xf3\xd5\x8d\x85\xba\xbb\xfe\x45\x87\xa5\x4d\x24\xf6\x90\x27\x69\xf4\x8e\x87\x56\xe2\x77\xdd\xe3\xc3\xed\x46\xb5\xd2\xee\x1d\x6e\x37\xab\x95\x53\xa5\x92\x01\x8b\xe3\xb9\xd2\xfe\xe1\xf6\x41\xb5\x5a\x11\x13\xf8\x08\xb5\xed\x66\x0d\x0e\x0f\xa1\x96\xe8\x94\xd7\xe0\xf3\x4b\x48\x66\x5c\x56\x2b\x95\x2d\x68\xf7\x40\xc4\x40\xdf\xf7\x60\xce\x41\xaa\x39\x48\xce\x7d\x88\x32\x2a\xb8\x29\x23\x71\x50\x83\x9f\x90\x44\xf1\x78\x85\x7b\x33\x05\x0e\xb7\xab\xff\x82\x38\xf5\xf1\xe7\x10\xa6\x3c\x84\xcc\x04\x40\xd2\xe1\xc6\x2b\x91\xc0\x3e\xfe\x8b\x07\x31\x5f\x7e\xa0\xa5\x89\xa8\xe6\x1f\xb7\xe0\x0c\x6d\x0c\xf1\x4c\xa5\x81\x0f\xc9\x22\xe2\x48\x85\x78\x15\xa5\xb9\x95\x47\xce\x02\x09\xa2\x99\xde\x89\x64\x90\x06\x81\xb5\xd2\x69\xf7\xb8\x3f\xea\x1a\x43\xbd\x39\x6d\x9d\xb4\xdf\x1b\x63\x1d\xf5\xdf\xbd\xed\x1d\x75\xc9\x4e\x95\x42\x94\x55\x2b\xe8\x6f\x88\x52\x43\x17\xb6\x9b\xd0\x7c\xfd\xa0\x81\xca\x25\x9c\x03\x6e\xb5\xb4\x87\xe9\x18\x9d\x98\x06\xbc\x27\x45\x16\x3c\x4b\x72\x8d\x0d\xe4\xe2\x7c\x3f\xca\x8c\x3f\x8b\x24\x1b\xeb\x24\xcf\x22\x1f\x43\xf4\xbb\x88\xa6\xe6\xc8\x4d\xb2\x03\xe5\xf7\xac\x8d\xbe\x41\x2e\xc2\x88\xcc\xad\x59\xa2\x52\xc9\xc8\x9c\xf2\x48\xdd\x26\xd8\x99\x0d\xd1\xb6\x92\x13\x31\xb5\xb4\x34\x6e\xcf\xa5\x72\x9c\x4b\xae\xc7\x2a\xe6\x1b\x28\xdb\x68\xc6\x3c\x0c\xf8\xdf\x2c\xe1\xb2\x96\x66\x05\xe6\x9a\x45\x11\xd7\xe6\x40\xc2\x74\x32\x14\x61\x1a\xb0\x44\x69\x7b\xa4\xd3\xfd\xd0\x6b\x6f\x50\xea\xca\xd3\xa9\x34\x6a\xe9\x34\xc4\xb4\x8d\xc1\x99\x63\xa8\x36\x6a\xd0\x78\xed\xfa\xfc\xd2\x95\x69\xa6\x2e\x51\x6e\xab\x30\x12\xe8\xd0\xfe\xd0\x52\x1d\x9c\xf6\xff\xda\x6d\x8f\xce\x47\xbf\x0d\x2c\xed\xfc\xc3\xa0\x35\xb2\x01\x34\x6c\xbf\xef\x1e\xdb\xf8\xc9\x45\x78\x6c\x2c\x93\x6d\x79\xb2\x26\x4f\xcc\x13\x70\x14\x44\x22\xe2\x13\x26\x02\x4a\xf0\x2b\x4f\xf9\xdc\x94\x29\x70\x8c\x64\x26\x55\x9d\xd8\x9b\xf1\x30\x4b\x2b\x67\x26\xb0\x4a\xcd\x78\x10\xd8\x52\xd5\x95\x97\x42\x2b\x49\x0a\xe1\x46\xff\x02\x44\x34\x53\x92\xc7\xb9\x4d\xc0\xf1\x79\x9c\x08\xc9\x12\xa1\x24\xd4\x22\xfc\x38\x51\x3a\x3c\x5c\xdb\xb6\x27\x59\xc8\x51\xe0\x1a\xed\xc7\xaa\xc8\xfd\x0e\x4b\xd8\x80\x25\xb3\xb5\xfa\x99\x39\xac\xb6\xfd\x84\x52\xfc\xca\x8b\x34\x4f\x92\x05\x59\x6c\x84\x7c\xfe\x44\x73\x5d\x79\x89\x52\xc1\xed\x76\xf9\x11\xba\x27\xb8\xbd\xa0\x20\x69\xd5\xd2\xde\x0c\xcd\xf1\xc7\x14\xcb\xd7\x4e\x5a\xc7\x7f\x52\x34\x7c\xc3\x59\x0e\xb3\xe2\x97\xd7\xb6\x1f\xd7\xaf\xbc\x6c\x05\xf2\xbf\x6f\xf5\x6b\xf7\x2a\x52\xfa\xa6\x67\xad\x4e\x8d\xa2\x4e\xcd\xef\xd3\x89\x1b\xc2\x99\x81\xf3\x9f\xfd\x88\x9c\x15\x0f\x02\x81\xae\xc8\x1b\xb4\x6b\xd7\xea\x91\xf9\x7a\x8b\x52\x8d\x82\x52\x19\xb1\xf2\x8e\x62\xc1\x31\xa6\x2c\xeb\xf9\x96\x27\xde\xac\x25\xa9\xb6\x31\xa1\x07\x5a\x5d\x8a\x18\x45\x11\x72\x8a\xff\x9e\x60\x2d\xc8\x10\xc7\x16\xf6\x20\x6c\x92\x08\x00\x62\x88\xc5\x74\x46\x5f\x7c\x35\x97\x81\x62\x3e\x50\xc5\x14\x32\x51\x4b\xd1\x61\x62\xac\x41\x9b\x98\xf4\x4d\xab\x04\xaa\x43\xd4\xc7\x26\xa9\xf4\x4c\x68\x0e\x54\x44\xa1\xc8\x37\x31\xad\x56\x5a\xed\x76\xff\xec\x64\x74\xde\x3d\x6e\xf5\x8e\x8c\xc9\x89\x2d\x15\x55\x46\x38\x26\x85\xe5\xa7\x5c\x8c\x73\x12\x23\x5b\x58\xa1\x1d\x77\x89\x77\x2a\xe1\x25\x3c\xaa\x87\x6a\x8c\x0c\xa2\x9c\xe7\x0a\x0e\x55\xbe\x2a\x0f\x99\x4a\x48\x7f\xc3\xda\xb2\xfe\x76\x8f\xfa\x83\xee\xe9\x79\x26\xb8\xed\xb4\x67\x27\x9d\xa3\xee\x79\xaf\xd3\x3d\x19\xf5\xde\xf6\xba\xa7\x26\x5c\x8c\xd4\x46\x50\xf4\x2a\x76\x57\xdb\x5e\x6e\x67\x9e\x7b\xe0\x1d\x46\x16\x59\xf0\xec\xac\xd7\x81\x89\x56\x61\x41\xf8\xad\x25\x5a\x9b\x62\xb4\x64\x70\xcd\x60\xb4\xf0\x2a\xd2\xee\xb3\xe6\xfe\xb3\xe6\xb3\xe7\xec\x97\x03\x8f\x3d\xe6\x07\xec\xe0\x09\x7b\x4c\x87\x8e\x94\x8a\x90\xa6\x56\x29\x8a\x44\xf6\x33\x0c\xf7\x00\xf1\xce\x42\xa5\x08\x30\x12\x60\x70\xd6\xe9\x75\xf6\x80\xf9\xc6\x91\x21\xa0\x9f\x49\x0a\x13\x95\x6a\x02\xd1\xd2\x65\xa5\x34\x30\x88\xc9\xf1\xa1\xb6\xf2\x80\x5b\x86\x5d\x23\x24\xb2\x16\x30\x80\x00\x26\x4e\x62\x82\x44\x2a\x35\xd6\x36\xa4\x0f\x27\xcc\xa2\x23\x2c\x64\x39\xbf\x13\x2c\x5f\x04\x8d\x96\xc4\x6f\xb8\xf6\x25\x46\x06\xc1\xac\x34\x15\x08\x04\x1f\xba\x69\xac\xdd\x40\x8c\xf9\x15\xf7\x5c\x93\x69\x6f\x52\xdf\x5f\x80\xe3\xc1\xce\x40\x63\xf8\x1a\xbb\xee\x80\xe9\x8a\x71\xe2\x23\xed\x57\xaf\x5e\xc1\xf6\xc3\x98\x23\x26\x16\xc9\x02\xbc\x10\xbb\x67\x07\x1c\x81\x89\xf4\xa5\x20\xc6\x75\x0d\x93\x6c\xd9\x4c\x77\x77\xab\xc4\x75\x0b\x7a\x13\xb2\x1d\xf9\xca\x68\xb5\x67\xd3\xc0\x53\xd1\x02\x44\x62\x4c\x8c\xf8\x65\xc6\x62\xb9\x83\x46\x0e\x34\x67\x28\xcc\x98\xdb\x2d\x82\x13\xce\xcb\x61\xe7\x17\x52\xe1\x7a\x03\xf6\xcc\xac\x3c\x81\xdf\xdd\x23\x31\xd6\x4c\x2f\xdc\x63\x63\x84\x0e\xbf\x14\x1e\x77\x8b\xd1\xf4\x09\x72\x73\xba\x19\xc1\x1b\xb9\x50\xa4\x8c\x0a\x0c\x71\x51\x26\x01\x8a\x3b\x95\x4a\x73\xfb\x39\x22\x53\x4d\xd6\x2d\x90\x39\x2e\xc7\xaa\x99\x30\x9f\x64\xcd\x1c\xca\x01\xaf\x01\xdb\xa4\x3f\x9d\xa3\x38\xb2\x91\x9c\xd5\x0a\xfa\x8f\x17\x91\x36\x37\xcc\xfb\x43\xd4\x2b\xcb\x7f\x8c\x5f\xe5\x14\xd6\xd4\xf8\x98\x9d\xfd\x4c\xd2\xad\x69\x41\x28\xbd\xb2\x1e\x99\x34\x35\x54\xf3\x45\x1f\x1b\x6c\x75\x35\x2d\x94\x76\x9a\xe1\x83\x76\x97\x66\x87\xf5\x91\x20\x97\xae\x93\xfa\x38\x8c\xfc\x86\x29\x68\x26\x11\x14\x86\xd2\x2f\x56\x18\xf3\x51\x41\xe9\x65\xf2\x15\x33\x21\x93\x77\xd3\x74\x91\x93\x7f\xcf\x17\x7b\x26\xc1\x8b\xd4\xf1\xb7\xbe\x2b\xf5\x9c\xb8\x1d\x34\xb0\x7f\x68\x25\xfc\x76\xc0\x99\x5c\x42\xe3\xac\x3d\xae\x03\xd2\xba\x6b\x01\xed\x1c\x3c\xda\x7e\x13\xaf\x67\xc4\xce\xa4\x58\x43\xed\xdf\x24\x98\xe6\x47\x5a\xeb\x08\xbe\x40\x37\x9b\x05\x3a\xd8\x8f\xa7\x77\x27\x2d\x0a\xa7\x10\xcc\xe3\x0f\xef\x22\xd1\xcc\xdb\x30\x6f\x64\x7c\x5a\x71\xcc\xc3\x71\xc0\xef\xce\x83\x65\x27\xee\x46\xff\x08\x3d\x79\x77\xda\x01\xd5\xb7\xdb\x48\x99\x79\xa2\x85\x4d\xf9\x12\xab\x5c\x86\x77\x5a\xed\x5f\x5b\xef\xb2\x79\xa2\xd5\x3b\xc1\xae\x36\xea\x7d\xe8\x8d\x7e\xbb\x89\x77\x98\x3f\xc6\x71\x15\xa1\x39\xb0\x10\x62\x22\x05\x8e\x44\x06\x84\x4d\x6c\x6f\xcb\x25\x66\xd8\xf9\x67\xdd\xe2\xd8\x52\x26\x55\xad\x68\xee\xb7\x55\x80\x71\x78\xb8\xf3\x69\xff\xe0\xe0\xe3\xfe\xcb\x83\x46\xb8\x43\x33\x10\xd6\xc5\xf5\x95\x26\xad\x8c\x83\x94\xaf\x2f\x3c\xa6\x05\xa9\xca\x9f\xf1\x5b\xb5\xd2\x3d\x3e\x3b\x6a\x8d\xfa\xa7\xb7\xe0\xb8\xe5\xf2\x69\x77\x78\x76\x34\x1a\x62\xd7\x90\x6a\x96\x46\x98\xcb\xad\x93\xce\x69\xbf\xd7\x39\x7f\xdf\x3f\xee\xba\x84\xc9\x63\x97\x67\x9a\x20\x20\x65\x97\xbe\x1d\xaa\x9a\xaf\x71\x58\x6e\xd0\x1f\x0f\xa8\x0d\x60\x93\xeb\x77\xfa\xf0\x02\x7b\x1d\x82\x27\xfc\x2f\x83\xb1\x98\xc2\xd6\x7b\x34\xef\x1e\x28\x89\xb5\x95\x6b\xad\x34\x2e\x68\x8e\xd8\x26\x49\xb5\xe4\x7e\x5e\x45\x27\x42\x63\x8b\x4d\xe6\x0a\xb0\x05\x29\xe9\x63\x5f\xee\xe1\xc0\xeb\x5d\x18\x50\xd5\xc3\xf6\x46\xad\xe3\x82\xee\x2b\x18\xcc\xd9\x02\x54\x9a\x20\xba\x08\x38\x8f\x60\xb3\x36\xaf\xb6\x9b\x56\xac\xf6\x8c\x23\x19\xac\x53\xc8\x07\x39\xcf\x19\x89\x36\x68\x9d\xf4\xda\xf0\x91\xfa\x3b\x8d\x04\x04\xf8\xf6\x20\x4a\x93\x5c\x9e\x39\x8e\x5c\x53\xc8\xd5\x06\x1a\x20\x3e\x53\xa8\xe0\xa9\xf3\x0c\xf3\x3c\x34\xb7\x21\x58\xbf\xd7\x99\x5f\x13\xf0\x44\x37\xa2\x2d\xcd\x01\x34\x9a\x67\x45\xe9\x95\x85\x88\x98\x14\xde\x9e\xc1\x24\x73\x6a\x90\xcc\x62\x1e\xba\x79\x32\x1d\x74\x01\x21\xc7\x44\xf1\x71\xb2\xde\x03\xa9\x12\x5c\x59\xde\xf0\x7c\x29\xc8\x82\x0c\x9d\x29\x16\xd7\x42\xb9\xcd\xab\x1f\x01\x8f\xa5\x16\x73\x63\x44\x03\x63\xd3\xa8\x0e\xc7\x6c\x31\xe6\x86\xe3\xf6\x97\x65\x34\x5e\xe7\xbb\xa9\x49\x6c\x7f\xc9\x43\xeb\x1a\x2e\x50\x1c\x61\x1a\xb8\xa7\xb4\xe6\x5e\x62\xab\xe3\xaa\xcc\x6a\xbe\x13\xc3\x7c\x96\x29\x11\xa8\x29\xc4\x6c\x11\xbf\xf8\x24\xcb\x1b\x37\xd8\x2b\xdf\x62\x2b\xed\x3e\x6c\xd1\x1f\x84\x7f\x06\x64\xa0\xc2\x7d\xd1\x90\x73\xc4\x25\x81\xb8\xe0\x99\x1d\x31\xcc\xa4\xca\xed\x18\xf0\x04\x25\xf0\x8c\xbb\xe9\xb4\x0d\x37\x3c\xd8\x3d\x3d\x45\x6e\xdf\xe1\x36\x73\xc0\xba\x6d\x05\x4c\x0a\x54\x36\x18\xfc\x2e\xfa\x65\x0a\x36\x48\xc1\x86\x11\xb1\x4b\x22\x6e\xe8\x89\x54\xd5\x69\xb8\x1a\xb3\x31\x76\xaf\xed\x2f\x85\xa2\x70\x6d\xea\x0e\xb6\xac\xa2\x73\x30\x8e\x33\xb7\x1a\xe3\xe7\x3e\xcc\xd8\x6e\x41\xab\xf3\x06\xa6\x38\x37\xc5\x76\xb4\x51\x66\x06\xa3\x53\x73\x86\x12\x91\x24\x74\xcc\x37\x50\x03\x12\x41\x33\x0e\xf9\x19\xc3\xcf\xc4\x9e\xb5\xec\xdf\x53\xcc\x14\x9f\x47\x1c\x33\x12\x4f\x23\xb0\xb3\xc4\x93\x2c\xe1\x11\xd6\x05\x81\x8d\x59\x98\x23\x28\xc7\x70\xc5\x78\xe0\x72\xe5\x2b\x46\x6d\x16\xcb\x7e\x68\xd7\x8b\xa2\x02\x0d\x90\xf1\x52\x91\x31\xf7\x18\xe6\x01\x81\x6e\x41\xbc\x71\xee\x01\xec\x47\xe4\x4d\x53\x85\x49\x6e\x07\xe5\x76\xac\xcc\x19\x0c\x3d\xc1\x3c\x4a\xf2\x08\xcc\xb4\x21\xb6\x97\x48\x9b\x21\xdf\x3d\x9b\xdb\x8a\x6a\x0d\x99\xe1\xd2\x62\xb1\x72\x8d\x58\x89\x84\x80\x15\x0b\x93\x14\x58\xf2\x51\x22\x95\x90\xd5\x0d\x98\xc5\x5c\xdc\x03\x2a\x3e\x7e\x64\x8c\xd7\xcc\xab\x96\x1d\x01\xf5\xc2\xda\x84\x4d\x99\x90\x7b\x4b\xfb\x10\xe6\x40\x72\x2c\xc0\x1d\x62\x3a\xb5\xbc\x73\x8f\x97\x43\xbb\x20\x3c\x95\xbc\x72\x00\x2c\xd5\x29\x44\x80\xb9\xcd\xa5\xc2\x8a\x63\x0c\x89\x89\x88\x43\xc5\x39\x3c\x1a\xb6\x4f\xbb\xdd\x93\xf3\xa3\x7e\xab\xd3\x3b\x79\x87\x29\xb0\xea\x63\xb8\x1d\x5d\x12\x41\xbc\x88\xeb\x63\xa5\x92\x73\x74\x7a\x84\xee\xe6\xe6\x1e\x88\x2e\x58\x60\xe7\x93\xde\xa1\x3c\xa8\xcc\x67\x04\x65\x29\x17\xca\x04\x2d\x58\x6f\x18\xb4\x6e\xc7\x90\xac\x30\x3f\xae\x16\x51\xe8\x9d\x6c\xac\xac\x8d\xe7\xb3\x85\xb9\xa7\x9e\xd9\x3f\x50\x9f\x58\xa1\x15\x56\x1d\xf0\xda\x67\xa1\x2c\x66\x40\x1c\xa8\xf9\x5f\x72\xf8\xfa\x83\x14\xce\x31\x2e\x25\x50\x80\x2e\x9b\xa1\xd7\x39\xc6\xcc\xc2\x8a\x94\x39\x87\x76\xd5\x97\x5d\xcf\x24\x49\x1a\x5b\xa8\x99\xe2\xb0\x8e\x3a\xaf\x1c\x5a\x2f\x27\x79\xa8\x70\x73\xd9\xb9\x44\xad\xe4\xd7\x44\xa5\xde\xac\x18\x13\x4a\x7a\xb9\x63\x57\x6a\x09\x49\xcd\x0b\x6b\x34\xca\x87\xf0\xe7\x59\xb3\x5c\x92\x8a\x1c\x86\x69\xc4\xf5\x4f\x45\xdb\x95\x9a\x44\xa6\x95\x4e\x25\xe1\xe2\x3a\x01\xf3\x1d\x4d\x5c\xb1\xa6\x63\x3f\x5e\x00\x9e\x8e\x31\xfd\xff\xf7\x9f\x7f\xff\xb7\x58\xd6\xf6\x33\xf8\x8d\xff\x2b\x42\xdb\xec\xd5\xa5\xf3\x6b\x3e\xd9\xb7\xcd\xc4\x67\x67\x7a\x53\x17\x1c\x33\xce\xe3\x34\x6f\x90\x20\xe5\xe4\x04\x7d\x69\xe6\x79\xe6\x1e\x34\x9e\xef\xef\x3f\x6e\x3e\x75\x9f\x3c\x7d\xde\x30\x63\x3c\xc2\x02\x53\xa0\x0a\x88\x60\x66\xbc\x90\x17\x37\x4c\xef\x8b\x6a\xe5\x21\xd8\x18\x7c\x42\x97\x52\x79\xe4\x36\x6c\x88\x66\x4b\x8d\x97\x60\x7a\xc1\xe2\xa5\xf1\x21\xec\x62\x04\x30\x2b\x76\x7e\x97\x4d\x77\x8e\x8e\x23\x95\x93\x8a\x9b\xb8\x7a\xa9\x94\xa5\x72\xdb\x69\x73\x55\x63\x69\xe0\xdf\x38\x64\x24\x38\xfc\xd7\xb2\xad\x4e\xf3\xa0\x76\x4f\x02\xcd\xfb\x12\x68\xdc\x97\xc0\xfe\x3d\x09\x34\x9e\xdf\x97\xc0\xb3\xfb\x12\x78\x7a\x5f\x02\xbf\xfc\x41\x02\x06\x5f\xff\xc1\xb3\xf9\x7d\xb7\x73\x1f\x22\xfc\x0a\x87\x2f\x27\xd7\x23\x6c\xd2\x53\x4e\x2c\xb0\x12\x2c\xee\x45\x70\xaa\xd4\x34\xe0\x3f\x82\x1e\x56\x6a\x47\x84\x53\x87\xe9\x90\xb3\xb1\x70\x2e\x9f\xae\xc4\xa5\xe0\x37\x6f\xbb\x85\x17\xa9\xb5\x67\xe3\xdf\xdd\xfc\x7e\xb8\x5a\xb1\xb5\x74\xf5\xc5\xf5\xec\x1b\x16\x92\xc8\x72\xba\xcd\x11\x5b\x4d\x84\x87\x42\x65\x77\x91\x6f\x5a\xc3\xee\x6a\xe6\x5c\x5e\x90\x89\x90\xee\x9f\x69\xf6\x63\x11\xf6\x8f\xba\x87\xa2\x3a\x17\x85\xdb\x9b\x5f\xf9\xc2\x9b\x61\xff\x8f\xe9\xbd\x5a\xc8\xfa\x45\xf6\x1b\x9c\x11\x98\xfb\xb9\xb1\x90\x2e\x5d\x96\xc7\x62\x5a\xfe\x98\xf3\xd8\xc8\xcd\x17\x31\xa2\x87\x71\x4a\x00\xae\x1e\x35\x9a\x3f\x9a\x29\x38\x03\xa8\xd5\xbe\xcd\xfa\xcf\xd0\x97\x5e\xe5\x10\x44\x26\xbc\x15\x45\x7d\xd9\x4b\x52\x99\x3b\xe1\x7b\xae\x9c\xcb\xef\x17\x07\x55\x6c\x83\xca\x4f\xb1\x71\x16\xae\xa1\x29\xd8\xa2\xe8\x9c\xc6\x3a\xf3\xce\x62\x83\x88\x90\x67\x6f\xd0\x52\x52\x8c\xbe\xc5\x1a\xb7\xad\x66\xe9\x48\x04\x34\x9c\x99\xf3\x48\x99\xe6\x37\x43\x99\xb8\x39\x22\x62\xd9\xbb\x8e\x13\x5f\x88\xe8\x9c\x90\x2b\x36\xd6\x73\x4c\xdc\x73\xf3\x4e\x72\x8e\xf2\x79\x88\x99\x09\x52\x98\x4b\xb3\xd5\x03\xaa\xe9\x39\xbe\x3f\x9c\x63\xac\x0e\xb1\x55\x7b\xb9\x50\xad\xd3\xf6\xfb\xde\x87\xee\x6a\xb6\x27\x79\x96\x3f\x56\xf7\xd1\x9b\xde\x4e\x5c\x4b\x2e\x8d\xc8\xaf\xa5\x7b\xea\x25\x30\xf1\x7d\x12\xa6\xc8\x36\xbf\x09\x47\x3e\x06\xf1\x6e\x37\x6c\xef\x4f\xe5\x3f\x45\x04\xce\x3f\x0a\xcf\x34\xcd\x3a\xa9\x8c\xec\x97\x9f\x92\x30\xc2\xad\x3a\x5c\xdf\x84\x1f\xb3\x34\xd5\xa5\xcd\x65\x09\xed\x63\x9e\x22\xe8\x4f\x57\xa6\x7a\xb2\xf9\x49\x68\xf3\x21\xf7\x51\xdd\x5f\x04\x62\x7c\x07\xfa\xf9\x56\xe2\x73\x43\x78\xa3\x25\x4a\x59\xcf\xf5\xab\x67\xdb\xea\x6e\xbd\x6e\x95\x2b\x49\x66\x4e\xd9\x1b\xc3\x5e\x7f\xd8\x93\x9e\xe6\xf4\x82\xfb\x86\x1c\x7e\x92\x86\x63\xae\x37\x3f\xb6\x55\x2b\x88\x7b\x42\x81\xa5\x2c\x95\x09\x02\x57\x7a\xf7\xd7\xfc\xd2\xb1\xef\x62\x59\x85\xf4\x68\x71\xd7\xfc\xff\x57\x74\x99\xe8\xd7\xef\xfe\x6b\xf6\xee\xbf\xfd\xf6\x4d\x2a\xfd\x80\x7f\x40\x20\x87\xa9\x5c\xa3\xdc\xee\xc9\x89\xb2\xaf\x6f\x74\x53\x21\xf9\xbc\x48\x94\xea\x41\x36\x05\x17\xa4\xb3\x33\x30\x7e\xba\x21\x45\x31\xa6\xd6\x48\x15\x09\x2c\x07\xfb\xf5\x3d\x0f\x1f\x6e\x20\xfa\x73\x63\x77\xd7\xa2\xcb\xca\xd7\x54\x1c\xe2\xd4\xfc\x62\x4d\x43\xd8\x2e\x73\xa0\x87\xd6\xa2\xca\x35\x73\x39\x83\x51\x0f\xbd\xd1\x10\x6b\xcf\x59\xcc\xe3\x13\x25\xbb\x57\x3c\xa4\xf7\x77\x4f\x2f\xec\xbc\x4c\xa3\x0a\xbd\xce\x64\xef\x19\xf6\x36\x04\xc7\xf4\x6a\x85\xc9\x78\xce\xf5\x48\xad\x36\xdf\xc9\x13\x5f\x65\x77\x43\xca\x5d\xeb\x85\x8f\xb0\x7d\x93\x1d\xdd\xb8\x3f\xaa\x75\x08\x04\x9f\x60\x2d\xea\xd2\x13\x45\xed\x11\x7c\x5e\x39\xe2\x6b\xd2\x90\xea\x2f\xbe\xae\xfb\x98\xde\xe6\x8d\xf6\x37\xad\x67\x02\xfd\xff\x01\x00\x00\xff\xff\xee\x24\x01\xec\xb9\x25\x00\x00")

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

	info := bindataFileInfo{name: "scripts/main.bash", size: 9657, mode: os.FileMode(420), modTime: time.Unix(1467133300, 0)}
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

